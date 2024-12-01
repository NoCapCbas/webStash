import uuid
from typing import Any

from fastapi import APIRouter, HTTPException
from sqlmodel import func, select

from app.api.deps import CurrentUser, SessionDep
from app.models import Bookmark, BookmarkCreate, BookmarkPublic, BookmarksPublic, BookmarkUpdate, Message

router = APIRouter()


@router.get("/", response_model=BookmarksPublic)
def read_bookmarks(
    session: SessionDep, current_user: CurrentUser, skip: int = 0, limit: int = 100
) -> Any:
    """
    Retrieve bookmarks.
    """

    if current_user.is_superuser:
        count_statement = select(func.count()).select_from(Bookmark)
        count = session.exec(count_statement).one()
        statement = select(Bookmark).offset(skip).limit(limit)
        bookmarks = session.exec(statement).all()
    else:
        count_statement = (
            select(func.count())
            .select_from(Bookmark)
            .where(Bookmark.owner_id == current_user.id)
        )
        count = session.exec(count_statement).one()
        statement = (
            select(Bookmark)
            .where(Bookmark.owner_id == current_user.id)
            .offset(skip)
            .limit(limit)
        )
        bookmarks = session.exec(statement).all()

    return BookmarksPublic(data=bookmarks, count=count)


@router.get("/{id}", response_model=BookmarkPublic)
def read_bookmark(session: SessionDep, current_user: CurrentUser, id: uuid.UUID) -> Any:
    """
    Get bookmark by ID.
    """
    bookmark = session.get(Bookmark, id)
    if not bookmark:
        raise HTTPException(status_code=404, detail="Bookmark not found")
    if not current_user.is_superuser and (bookmark.owner_id != current_user.id):
        raise HTTPException(status_code=400, detail="Not enough permissions")
    return bookmark


@router.post("/", response_model=BookmarkPublic)
def create_bookmark(
    *, session: SessionDep, current_user: CurrentUser, bookmark_in: BookmarkCreate
) -> Any:
    """
    Create new bookmark.
    """
    bookmark = Bookmark.model_validate(bookmark_in, update={"owner_id": current_user.id})
    session.add(bookmark)
    session.commit()
    session.refresh(bookmark)
    return bookmark


@router.put("/{id}", response_model=BookmarkPublic)
def update_bookmark(
    *,
    session: SessionDep,
    current_user: CurrentUser,
    id: uuid.UUID,
    bookmark_in: BookmarkUpdate,
) -> Any:
    """
    Update a bookmark.
    """
    bookmark = session.get(Bookmark, id)
    if not bookmark:
        raise HTTPException(status_code=404, detail="Bookmark not found")
    if not current_user.is_superuser and (bookmark.owner_id != current_user.id):
        raise HTTPException(status_code=400, detail="Not enough permissions")
    update_dict = bookmark_in.model_dump(exclude_unset=True)
    bookmark.sqlmodel_update(update_dict)
    session.add(bookmark)
    session.commit()
    session.refresh(bookmark)
    return bookmark


@router.delete("/{id}")
def delete_bookmark(
    session: SessionDep, current_user: CurrentUser, id: uuid.UUID
) -> Message:
    """
    Delete a bookmark.
    """
    bookmark = session.get(Bookmark, id)
    if not bookmark:
        raise HTTPException(status_code=404, detail="Bookmark not found")
    if not current_user.is_superuser and (bookmark.owner_id != current_user.id):
        raise HTTPException(status_code=400, detail="Not enough permissions")
    session.delete(bookmark)
    session.commit()
    return Message(message="Bookmark deleted successfully")
