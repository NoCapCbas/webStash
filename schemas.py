from pydantic import BaseModel, EmailStr
from datetime import datetime
from typing import Optional

class LoginRequest(BaseModel):
    email: EmailStr

class BookmarkCreate(BaseModel):
    url: str
    title: str
    description: Optional[str] = ""
    public: bool = False

class BookmarkCreateInternal(BaseModel):
    user_id: int
    url: str
    title: str
    description: Optional[str] = ""
    public: bool = False

class BookmarkUpdate(BaseModel):
    id: int
    user_id: int
    url: str
    title: str
    description: Optional[str] = ""
    public: bool = False

class BookmarkDelete(BaseModel):
    id: int
    user_id: int

class BookmarkRead(BaseModel):
    id: int

class BookmarkResponse(BaseModel):
    id: int
    user_id: int
    url: str
    title: str
    description: Optional[str]
    public: bool
    click_count: int
    created_at: datetime
    updated_at: Optional[datetime]
    
    class Config:
        from_attributes = True

class UserResponse(BaseModel):
    id: int
    email: str
    created_at: datetime
    updated_at: Optional[datetime]
    
    class Config:
        from_attributes = True