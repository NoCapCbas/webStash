import uuid

from pydantic import EmailStr
from sqlmodel import Field, Relationship, SQLModel


# Shared properties
class UserBase(SQLModel):
    email: EmailStr = Field(unique=True, index=True, max_length=255)
    is_active: bool = True
    is_superuser: bool = False
    full_name: str | None = Field(default=None, max_length=255)


# Properties to receive via API on creation
class UserCreate(UserBase):
    password: str = Field(min_length=8, max_length=40)


class UserRegister(SQLModel):
    email: EmailStr = Field(max_length=255)
    password: str = Field(min_length=8, max_length=40)
    full_name: str | None = Field(default=None, max_length=255)


# Properties to receive via API on update, all are optional
class UserUpdate(UserBase):
    email: EmailStr | None = Field(default=None, max_length=255)  # type: ignore
    password: str | None = Field(default=None, min_length=8, max_length=40)


class UserUpdateMe(SQLModel):
    full_name: str | None = Field(default=None, max_length=255)
    email: EmailStr | None = Field(default=None, max_length=255)


class UpdatePassword(SQLModel):
    current_password: str = Field(min_length=8, max_length=40)
    new_password: str = Field(min_length=8, max_length=40)


# Database model, database table inferred from class name
class User(UserBase, table=True):
    id: uuid.UUID = Field(default_factory=uuid.uuid4, primary_key=True)
    hashed_password: str
    items: list["Item"] = Relationship(back_populates="owner", cascade_delete=True)
    tags: list["Tag"] = Relationship(back_populates="owner", cascade_delete=True)
    bookmarks: list["Bookmark"] = Relationship(back_populates="owner", cascade_delete=True)


# Properties to return via API, id is always required
class UserPublic(UserBase):
    id: uuid.UUID


class UsersPublic(SQLModel):
    data: list[UserPublic]
    count: int


# Shared properties
class ItemBase(SQLModel):
    title: str = Field(min_length=1, max_length=255)
    description: str | None = Field(default=None, max_length=255)


# Properties to receive on item creation
class ItemCreate(ItemBase):
    pass


# Properties to receive on item update
class ItemUpdate(ItemBase):
    title: str | None = Field(default=None, min_length=1, max_length=255)  # type: ignore


# Database model, database table inferred from class name
class Item(ItemBase, table=True):
    id: uuid.UUID = Field(default_factory=uuid.uuid4, primary_key=True)
    title: str = Field(max_length=255)
    owner_id: uuid.UUID = Field(
        foreign_key="user.id", nullable=False, ondelete="CASCADE"
    )
    owner: User | None = Relationship(back_populates="items")


# Properties to return via API, id is always required
class ItemPublic(ItemBase):
    id: uuid.UUID
    owner_id: uuid.UUID


class ItemsPublic(SQLModel):
    data: list[ItemPublic]
    count: int


# Generic message
class Message(SQLModel):
    message: str


# JSON payload containing access token
class Token(SQLModel):
    access_token: str
    token_type: str = "bearer"


# Contents of JWT token
class TokenPayload(SQLModel):
    sub: str | None = None


class NewPassword(SQLModel):
    token: str
    new_password: str = Field(min_length=8, max_length=40)

# Bookmark Specific Models
"""
Domain models for bookmarks
"""
# Shared properties
class BookmarkBase(SQLModel):
    url: str = Field(max_length=1000)
    title: str | None = Field(default=None, max_length=50)
    description: str | None = Field(default=None, max_length=1000)

# Properties to receive on bookmark creation
class BookmarkCreate(BookmarkBase):
    pass

# Properties to receive on bookmark update
class BookmarkUpdate(BookmarkBase):
    url: str | None = Field(default=None, max_length=1000)
    title: str | None = Field(default=None, max_length=50)
    description: str | None = Field(default=None, max_length=1000)

# Database model, database table inferred from class name
class Bookmark(BookmarkBase, table=True):
    id: uuid.UUID = Field(default_factory=uuid.uuid4, primary_key=True)
    url: str = Field(max_length=1000)
    title: str | None = Field(default=None, max_length=50)
    description: str | None = Field(default=None, max_length=1000)
    owner_id: uuid.UUID = Field(foreign_key="user.id", nullable=False, ondelete="CASCADE")
    owner: User | None = Relationship(back_populates="bookmarks")
    tags: list["Tag"] = Relationship(back_populates="bookmark", cascade_delete=True)

# Properties to return via API, id is always required
class BookmarkPublic(BookmarkBase):
    id: uuid.UUID
    owner_id: uuid.UUID

class BookmarksPublic(SQLModel):
    data: list[BookmarkPublic]
    count: int

"""
Used to categorize bookmarks
"""
class TagBase(SQLModel):
    name: str = Field(max_length=255)

class TagCreate(TagBase):
    pass

class TagUpdate(TagBase):
    name: str | None = Field(default=None, max_length=255)

class Tag(TagBase, table=True):
    id: uuid.UUID = Field(default_factory=uuid.uuid4, primary_key=True)
    name: str = Field(max_length=255)
    owner_id: uuid.UUID = Field(foreign_key="user.id", nullable=False, ondelete="CASCADE")
    owner: User | None = Relationship(back_populates="tags")
    bookmark_id: uuid.UUID = Field(foreign_key="bookmark.id", nullable=False, ondelete="CASCADE")
    bookmark: Bookmark | None = Relationship(back_populates="tags")

class TagPublic(TagBase):
    id: uuid.UUID
    owner_id: uuid.UUID

class TagsPublic(SQLModel):
    data: list[TagPublic]
    count: int

