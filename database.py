from sqlalchemy import create_engine
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker
import os
import logging

logger = logging.getLogger(__name__)

# Database configuration
DATABASE_URL = os.getenv("DATABASE_URL", "sqlite:///./webstash.sqlite3")

engine = create_engine(
    DATABASE_URL, 
    connect_args={
        "check_same_thread": False,
        "timeout": 20,  # 20 second timeout for database operations
    },
    pool_pre_ping=True,  # Verify connections before use
    echo=False  # Set to True for SQL debugging
)
SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine)

Base = declarative_base()

def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()

def init_db():
    """Initialize the database and create tables"""
    logger.info("Initializing database...")
    
    # Create database file if it doesn't exist with proper permissions
    if DATABASE_URL.startswith("sqlite:///"):
        db_path = DATABASE_URL.replace("sqlite:///", "")
        db_dir = os.path.dirname(os.path.abspath(db_path)) if os.path.dirname(db_path) else os.getcwd()
        
        # Ensure directory exists and is writable
        os.makedirs(db_dir, exist_ok=True)
        
        if not os.path.exists(db_path):
            # Create the database file with write permissions
            with open(db_path, 'w') as f:
                pass  # Create empty file
            # Set proper permissions (read/write for owner)
            os.chmod(db_path, 0o644)
            logger.info(f"Created new SQLite database file at: {db_path}")
        else:
            # Ensure existing file has write permissions
            os.chmod(db_path, 0o644)
            logger.info(f"Using existing SQLite database file at: {db_path}")
    
    # Import models to ensure they're registered
    from models import User, UserSession, Bookmark, MagicLink
    
    # Create all tables
    Base.metadata.create_all(bind=engine)
    logger.info("Database initialized successfully")