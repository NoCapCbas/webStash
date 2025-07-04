from sqlalchemy.orm import Session
from sqlalchemy import and_
from datetime import datetime
from typing import List, Optional
import logging

from models import Bookmark
from schemas import BookmarkCreate, BookmarkCreateInternal, BookmarkUpdate

logger = logging.getLogger(__name__)

class BookmarkService:
    def create_bookmark(self, bookmark_data: BookmarkCreateInternal, db: Session) -> Bookmark:
        """Create a new bookmark"""
        bookmark = Bookmark(
            user_id=bookmark_data.user_id,
            url=bookmark_data.url,
            title=bookmark_data.title,
            description=bookmark_data.description,
            public=bookmark_data.public
        )
        
        db.add(bookmark)
        db.commit()
        db.refresh(bookmark)
        
        logger.info(f"Created bookmark: {bookmark.title} for user {bookmark.user_id}")
        return bookmark
    
    def create_bookmark_for_user(self, bookmark_data: BookmarkCreate, user_id: int, db: Session) -> Bookmark:
        """Create a new bookmark for a specific user"""
        internal_data = BookmarkCreateInternal(
            user_id=user_id,
            url=bookmark_data.url,
            title=bookmark_data.title,
            description=bookmark_data.description,
            public=bookmark_data.public
        )
        return self.create_bookmark(internal_data, db)
    
    def get_bookmark_by_id(self, bookmark_id: int, db: Session) -> Optional[Bookmark]:
        """Get bookmark by ID"""
        return db.query(Bookmark).filter(Bookmark.id == bookmark_id).first()
    
    def get_bookmarks_by_user_id(self, user_id: int, db: Session) -> List[Bookmark]:
        """Get all bookmarks for a user"""
        return db.query(Bookmark).filter(Bookmark.user_id == user_id).order_by(Bookmark.created_at.desc()).all()
    
    def update_bookmark(self, bookmark_data: BookmarkUpdate, db: Session) -> Optional[Bookmark]:
        """Update a bookmark"""
        bookmark = db.query(Bookmark).filter(
            and_(
                Bookmark.id == bookmark_data.id,
                Bookmark.user_id == bookmark_data.user_id
            )
        ).first()
        
        if not bookmark:
            return None
        
        bookmark.url = bookmark_data.url
        bookmark.title = bookmark_data.title
        bookmark.description = bookmark_data.description
        bookmark.public = bookmark_data.public
        bookmark.updated_at = datetime.utcnow()
        
        db.commit()
        db.refresh(bookmark)
        
        logger.info(f"Updated bookmark: {bookmark.title}")
        return bookmark
    
    def delete_bookmark(self, bookmark_id: int, user_id: int, db: Session) -> bool:
        """Delete a bookmark"""
        bookmark = db.query(Bookmark).filter(
            and_(
                Bookmark.id == bookmark_id,
                Bookmark.user_id == user_id
            )
        ).first()
        
        if not bookmark:
            return False
        
        db.delete(bookmark)
        db.commit()
        
        logger.info(f"Deleted bookmark with ID {bookmark_id} for user {user_id}")
        return True
    
    def increment_click_count(self, bookmark_id: int, db: Session) -> bool:
        """Increment click count for a bookmark"""
        bookmark = db.query(Bookmark).filter(Bookmark.id == bookmark_id).first()
        
        if not bookmark:
            return False
        
        bookmark.click_count += 1
        db.commit()
        
        return True