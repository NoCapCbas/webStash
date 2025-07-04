# WebStash - Bookmark Management System

A minimalist bookmark management system that allows users to organize, store, and access their favorite web links. Users can authenticate via magic links, manage their bookmarks, and organize their web resources efficiently.

## Features

- **Magic Link Authentication**: Secure, passwordless authentication via email
- **Bookmark Management**: Create, read, update, and delete bookmarks
- **User Sessions**: Secure session management with automatic expiration
- **Public/Private Bookmarks**: Control bookmark visibility
- **Click Tracking**: Track bookmark usage statistics
- **Responsive Design**: Clean, mobile-friendly interface
- **SQLite Database**: Lightweight, file-based database storage

## Technology Stack

### Backend

- **Python 3.11**: Modern Python runtime
- **FastAPI**: High-performance, modern web framework
- **SQLAlchemy**: Powerful SQL toolkit and ORM
- **SQLite**: Lightweight database engine
- **Uvicorn**: Lightning-fast ASGI server
- **Pydantic**: Data validation using Python type hints

### Frontend

- **Jinja2**: Template engine for dynamic HTML
- **Alpine.js**: Minimal JavaScript framework
- **Tailwind CSS**: Utility-first CSS framework

### Infrastructure

- **Docker**: Containerized deployment
- **Docker Compose**: Multi-container orchestration
- **Traefik**: Reverse proxy and load balancer
- **GitHub Actions**: CI/CD workflows
- **Watchtower**: Automated container updates

## Project Structure

```
webStash/
├── main.py                 # FastAPI application entry point
├── database.py             # Database configuration and initialization
├── models.py               # SQLAlchemy database models
├── schemas.py              # Pydantic data validation schemas
├── services/               # Business logic layer
│   ├── auth.py            # Authentication service
│   └── bookmark.py        # Bookmark management service
├── templates/              # Jinja2 HTML templates
├── static/                 # CSS, JS, and image assets
├── requirements.txt        # Python dependencies
├── dockerfile.dev          # Development Docker configuration
├── dockerfile.prod         # Production Docker configuration
├── docker-compose.dev.yml  # Development compose file
├── docker-compose.prod.yml # Production compose file
└── Makefile               # Development commands
```

## Development Setup

### Prerequisites

- Docker and Docker Compose
- Node.js and npm (for frontend development)

### Quick Start

1. **Clone the repository**

   ```bash
   git clone <repository-url>
   cd webStash
   ```

2. **Install frontend dependencies**

   ```bash
   npm install
   ```

3. **Build and run with Docker**

   ```bash
   make dev
   ```

4. **Access the application**
   - Open <http://localhost:8080> in your browser

### Frontend Development

**Install dependencies:**

```bash
npm install
```

**Watch and compile Tailwind CSS:**

```bash
npx tailwindcss -i ./static/css/input.css -o ./static/css/output.css --watch
```

### Backend Development

**Run development server with auto-reload:**

```bash
make test
```

**Or run directly with uvicorn:**

```bash
uvicorn main:app --host 0.0.0.0 --port 8080 --reload
```

## API Endpoints

### Authentication

- `POST /api/login` - Request magic link authentication
- `GET /verify` - Verify magic link token

### Bookmarks

- `POST /api/v1/bookmarks/create` - Create new bookmark
- `PUT /api/v1/bookmarks/update` - Update existing bookmark
- `DELETE /api/v1/bookmarks/delete` - Delete bookmark
- `POST /api/v1/bookmarks/read` - Get bookmark details

### Pages

- `GET /` - Home page
- `GET /view/bookmarks` - Bookmark dashboard
- `GET /policies` - Privacy/Terms pages
- `GET /404` - Not found page

## Database Schema

### Users

- `id`: Primary key
- `email`: User email (unique)
- `premium`: Premium status flag
- `created_at`, `updated_at`: Timestamps

### Bookmarks

- `id`: Primary key
- `user_id`: Foreign key to users
- `url`: Bookmark URL
- `title`: Bookmark title
- `description`: Optional description
- `public`: Visibility flag
- `click_count`: Usage counter
- `created_at`, `updated_at`: Timestamps

### Sessions

- `id`: Primary key
- `user_id`: Foreign key to users
- `token`: Session token (unique)
- `expires_at`: Expiration timestamp
- `created_at`: Creation timestamp

### Magic Links

- `id`: Primary key
- `email`: Target email
- `token`: Magic link token (unique)
- `expires_at`: Expiration timestamp
- `used`: Usage flag
- `created_at`: Creation timestamp

## Deployment

### Development

```bash
make dev
```

### Production

```bash
make prod
```

### Manual Docker Commands

```bash
# Development
docker build -f dockerfile.dev -t webstash-dev .
docker run --rm -p 8080:8080 webstash-dev

# Production
docker build -f dockerfile.prod -t webstash-prod .
docker run --rm -p 8080:8080 webstash-prod
```

### Environment Variables

```bash
export DATABASE_URL=sqlite:///./webstash.sqlite3
export DEV_EMAIL=example@domain.com
export HOST_DOMAIN=example.com
```

## Production Infrastructure

- **Traefik**: Reverse proxy and SSL termination
- **GitHub Actions**: Automated CI/CD pipeline
- **Watchtower**: Automatic container updates
- **Docker Compose**: Service orchestration

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## License

This project is licensed under the MIT License.

## Resources

- **Icons**: [Heroicons](https://heroicons.com/)
- **FastAPI Documentation**: [fastapi.tiangolo.com](https://fastapi.tiangolo.com/)
- **SQLAlchemy Documentation**: [sqlalchemy.org](https://www.sqlalchemy.org/)

