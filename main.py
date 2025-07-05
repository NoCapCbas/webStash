from fastapi import FastAPI, HTTPException, Request, Response, Cookie, Depends, Form
from fastapi.responses import HTMLResponse, RedirectResponse, JSONResponse
from fastapi.staticfiles import StaticFiles
from fastapi.templating import Jinja2Templates
from sqlalchemy.orm import Session
import os
from datetime import datetime, timedelta
import logging

from database import get_db, init_db
from models import User, UserSession, Bookmark, MagicLink
from services.auth import AuthService
from services.bookmark import BookmarkService
from schemas import LoginRequest, BookmarkCreate, BookmarkCreateInternal, BookmarkUpdate, BookmarkDelete, BookmarkRead

# Setup logging
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

app = FastAPI(title="WebStash", version="1.0.0")

# Mount static files
app.mount("/static", StaticFiles(directory="static"), name="static")

# Setup templates
templates = Jinja2Templates(directory="templates")

# Initialize database
init_db()

# Initialize services
auth_service = AuthService()
bookmark_service = BookmarkService()

class PageData:
    def __init__(self, title: str = "", message: str = ""):
        self.title = title
        self.message = message

@app.get("/", response_class=HTMLResponse)
async def index(request: Request):
    data = PageData(title="WebStash", message="Welcome to WebStash")
    return templates.TemplateResponse("index.html", {"request": request, "data": data})

@app.post("/api/login")
async def login(request: LoginRequest, db: Session = Depends(get_db)):
    logger.info(f"Generating magic link for {request.email}")
    
    try:
        # Generate magic link
        magic_link = auth_service.generate_magic_link(request.email, db)
        
        # Create user if doesn't exist
        auth_service.create_user(request.email, db)
        
        # Generate the full magic link URL
        HOST_DOMAIN = os.getenv('HOST_DOMAIN', 'localhost:8080')
        magic_link_url = f"http://{HOST_DOMAIN}/verify?token={magic_link.token}"
        
        logger.info("Magic link sent to email")
        return JSONResponse({
            "message": "Magic link displayed due to email service being offline",
            "magic_link": magic_link_url
        })
    except Exception as e:
        logger.error(f"Error in login: {e}")
        raise HTTPException(status_code=500, detail="Error generating magic link")

@app.get("/verify")
async def verify(token: str, response: Response, db: Session = Depends(get_db)):
    logger.info("Verifying magic link")
    
    if not token:
        raise HTTPException(status_code=400, detail="Missing token")
    
    try:
        # Validate the magic link token
        email = auth_service.validate_magic_link(token, db)
        
        # Generate a new session token
        session_token = auth_service.generate_session_token(email, db)
        
        # Set session cookie
        response = RedirectResponse(url="/view/bookmarks", status_code=303)
        response.set_cookie(
            key="session_token",
            value=session_token,
            path="/",
            httponly=True,
            max_age=3600,  # 1 hour
            samesite="lax"
        )
        
        return response
    except Exception as e:
        logger.error(f"Invalid magic link token: {e}")
        raise HTTPException(status_code=401, detail=str(e))

@app.post("/api/v1/bookmarks/add-default")
async def add_default_bookmarks(request: Request, db: Session = Depends(get_db)):
    logger.info("Adding default bookmarks")
    
    session_token = request.cookies.get("session_token")
    if not session_token:
        return RedirectResponse(url="/404", status_code=303)
    
    try:
        email = auth_service.validate_session(session_token, db)
        user = auth_service.get_user_by_email(email, db)
        
        # Create default bookmarks
        google_bookmark = bookmark_service.create_bookmark(BookmarkCreateInternal(
            user_id=user.id,
            title="Google",
            url="https://google.com",
            description="Google's homepage",
            public=False
        ), db)
        logger.info(f"Created default bookmark: {google_bookmark.title}")
        
        github_bookmark = bookmark_service.create_bookmark(BookmarkCreateInternal(
            user_id=user.id,
            title="Github",
            url="https://github.com",
            description="Github's homepage",
            public=False
        ), db)
        logger.info(f"Created default bookmark: {github_bookmark.title}")
        
        return RedirectResponse(url="/view/bookmarks", status_code=303)
    except Exception as e:
        logger.error(f"Error adding default bookmarks: {e}")
        return RedirectResponse(url="/404", status_code=303)

@app.get("/view/bookmarks", response_class=HTMLResponse)
async def bookmark_view(request: Request, db: Session = Depends(get_db)):
    logger.info("Bookmark view handler")
    
    session_token = request.cookies.get("session_token")
    if not session_token:
        return RedirectResponse(url="/404", status_code=303)
    
    try:
        email = auth_service.validate_session(session_token, db)
        user = auth_service.get_user_by_email(email, db)
        bookmarks = bookmark_service.get_bookmarks_by_user_id(user.id, db)
        
        logger.info(f"Retrieved {len(bookmarks)} bookmarks for user {email}")
        
        # Debug: Log the bookmarks data
        for bookmark in bookmarks:
            logger.info(f"Bookmark: {bookmark.title} - {bookmark.url}")
        
        data = {
            "email": email,
            "bookmarks": bookmarks
        }
        
        return templates.TemplateResponse("bookmarks/index.html", {"request": request, **data})
    except Exception as e:
        logger.error(f"Error in bookmark view: {e}")
        return RedirectResponse(url="/404", status_code=303)

@app.get("/policies", response_class=HTMLResponse)
async def policies(request: Request):
    return templates.TemplateResponse("policies/main.html", {"request": request})

@app.post("/api/logout")
async def logout(request: Request, db: Session = Depends(get_db)):
    """Logout user by clearing session cookie and invalidating session"""
    session_token = request.cookies.get("session_token")
    
    # Invalidate session in database if token exists
    if session_token:
        auth_service.invalidate_session(session_token, db)
    
    # Clear session cookie and redirect
    response = RedirectResponse(url="/", status_code=303)
    response.delete_cookie(key="session_token", path="/")
    return response

@app.get("/404", response_class=HTMLResponse)
async def not_found(request: Request):
    return templates.TemplateResponse("404.html", {"request": request})

@app.post("/api/v1/bookmarks/create")
async def bookmark_create(request: Request, bookmark: BookmarkCreate, db: Session = Depends(get_db)):
    logger.info("Bookmark create handler")
    
    # Get session token from cookies
    session_token = request.cookies.get("session_token")
    if not session_token:
        raise HTTPException(status_code=401, detail="No session token")
    
    try:
        # Validate session and get user
        email = auth_service.validate_session(session_token, db)
        user = auth_service.get_user_by_email(email, db)
        
        # Create bookmark for authenticated user
        created_bookmark = bookmark_service.create_bookmark_for_user(bookmark, user.id, db)
        logger.info(f"Successfully created bookmark: {created_bookmark.title} for user {user.email}")
        return created_bookmark
    except Exception as e:
        logger.error(f"Error creating bookmark: {e}")
        raise HTTPException(status_code=500, detail="Error creating bookmark")

@app.put("/api/v1/bookmarks/update")
async def bookmark_update(bookmark: BookmarkUpdate, db: Session = Depends(get_db)):
    logger.info("Bookmark update handler")
    try:
        updated_bookmark = bookmark_service.update_bookmark(bookmark, db)
        return updated_bookmark
    except Exception as e:
        logger.error(f"Error updating bookmark: {e}")
        raise HTTPException(status_code=500, detail="Error updating bookmark")

@app.delete("/api/v1/bookmarks/delete")
async def bookmark_delete(bookmark: BookmarkDelete, db: Session = Depends(get_db)):
    logger.info("Bookmark delete handler")
    try:
        logger.info(f"Deleting bookmark with ID {bookmark.id} for user {bookmark.user_id}")
        bookmark_service.delete_bookmark(bookmark.id, bookmark.user_id, db)
        return {"message": "Bookmark deleted successfully"}
    except Exception as e:
        logger.error(f"Error deleting bookmark: {e}")
        raise HTTPException(status_code=500, detail="Error deleting bookmark")

@app.post("/api/v1/bookmarks/read")
async def bookmark_read(req: BookmarkRead, db: Session = Depends(get_db)):
    logger.info("Bookmark read handler")
    try:
        bookmark = bookmark_service.get_bookmark_by_id(req.id, db)
        if not bookmark:
            raise HTTPException(status_code=404, detail="Bookmark not found")
        return bookmark
    except Exception as e:
        logger.error(f"Error reading bookmark: {e}")
        raise HTTPException(status_code=500, detail="Error reading bookmark")

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=8080)
