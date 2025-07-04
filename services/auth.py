import secrets
import base64
from datetime import datetime, timedelta
from sqlalchemy.orm import Session
from sqlalchemy import and_
import logging

from models import User, UserSession, MagicLink

logger = logging.getLogger(__name__)

class AuthService:
    def generate_magic_link(self, email: str, db: Session) -> MagicLink:
        """Generate a magic link for email authentication"""
        token = self._generate_random_token()
        expires_at = datetime.utcnow() + timedelta(minutes=60)
        
        # Create magic link
        magic_link = MagicLink(
            email=email,
            token=token,
            expires_at=expires_at
        )
        
        db.add(magic_link)
        db.commit()
        db.refresh(magic_link)
        
        logger.info(f"Generated magic link for {email}")
        return magic_link
    
    def validate_magic_link(self, token: str, db: Session) -> str:
        """Validate magic link token and return email"""
        magic_link = db.query(MagicLink).filter(
            and_(
                MagicLink.token == token,
                MagicLink.expires_at > datetime.utcnow(),
                MagicLink.used == False
            )
        ).first()
        
        if not magic_link:
            raise Exception("Invalid token, please request a new one")
        
        # Mark as used
        magic_link.used = True
        db.commit()
        
        # Clean up expired tokens
        self._delete_expired_magic_links(db)
        
        return magic_link.email
    
    def create_user(self, email: str, db: Session) -> User:
        """Create user if not exists"""
        user = db.query(User).filter(User.email == email).first()
        
        if not user:
            user = User(email=email)
            db.add(user)
            db.commit()
            db.refresh(user)
            logger.info(f"Created new user: {email}")
        
        return user
    
    def get_user_by_email(self, email: str, db: Session) -> User:
        """Get user by email"""
        user = db.query(User).filter(User.email == email).first()
        if not user:
            raise Exception("User not found")
        return user
    
    def create_session(self, user_id: int, db: Session) -> UserSession:
        """Create a new session for user"""
        token = self._generate_random_token()
        expires_at = datetime.utcnow() + timedelta(hours=24)
        
        session = UserSession(
            user_id=user_id,
            token=token,
            expires_at=expires_at
        )
        
        db.add(session)
        db.commit()
        db.refresh(session)
        
        return session
    
    def validate_session(self, token: str, db: Session) -> str:
        """Validate session token and return user email"""
        session = db.query(UserSession).filter(
            and_(
                UserSession.token == token,
                UserSession.expires_at > datetime.utcnow()
            )
        ).first()
        
        if not session:
            raise Exception("Invalid session")
        
        user = db.query(User).filter(User.id == session.user_id).first()
        if not user:
            raise Exception("User not found")
        
        return user.email
    
    def generate_session_token(self, email: str, db: Session) -> str:
        """Generate session token for user"""
        user = self.get_user_by_email(email, db)
        session = self.create_session(user.id, db)
        return session.token
    
    def _generate_random_token(self) -> str:
        """Generate a random token"""
        return base64.urlsafe_b64encode(secrets.token_bytes(32)).decode('utf-8')
    
    def invalidate_session(self, token: str, db: Session) -> bool:
        """Invalidate a specific session by token"""
        try:
            session = db.query(UserSession).filter(UserSession.token == token).first()
            if session:
                db.delete(session)
                db.commit()
                logger.info(f"Invalidated session for user ID: {session.user_id}")
                return True
            return False
        except Exception as e:
            logger.error(f"Error invalidating session: {e}")
            return False
    
    def _delete_expired_magic_links(self, db: Session):
        """Delete expired magic links"""
        db.query(MagicLink).filter(
            MagicLink.expires_at < datetime.utcnow()
        ).delete()
        db.commit()