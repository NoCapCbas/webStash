# Deployment Guide

## Prerequisites

1. Node.js 18 or higher
2. A Cloudflare account
3. Wrangler CLI installed (`npm install -g wrangler`)

## Setup

### 1. Install Dependencies

```bash
npm install
```

### 2. Create D1 Database

```bash
# Create the database
wrangler d1 create webstash

# Copy the database_id from the output and update wrangler.toml
```

Update `wrangler.toml` with your database ID:

```toml
[[d1_databases]]
binding = "DB"
database_name = "webstash"
database_id = "your-database-id-here"
```

### 3. Initialize Database Schema

```bash
# Apply schema to production database
wrangler d1 execute webstash --file=./schema.sql
```

### 4. Create KV Namespace

```bash
# Create KV namespace for sessions
wrangler kv:namespace create SESSIONS

# Copy the id from the output and update wrangler.toml
```

Update `wrangler.toml` with your KV namespace ID:

```toml
[[kv_namespaces]]
binding = "SESSIONS"
id = "your-kv-id-here"
```

## Local Development

```bash
# Start local dev server with Cloudflare Workers emulation
npm run dev

# The local dev environment will use local D1 and KV bindings
```

## Production Deployment

### Build and Deploy

```bash
# Build the project
npm run build

# Deploy to Cloudflare Pages
npm run deploy
```

### Using Wrangler Pages

Alternatively, you can use Cloudflare Pages with GitHub integration:

1. Push your code to GitHub
2. Go to Cloudflare Dashboard > Pages
3. Create a new project
4. Connect your GitHub repository
5. Set build command: `npm run build`
6. Set build output directory: `dist`
7. Add D1 and KV bindings in the Pages settings

## Environment Variables

No environment variables are required for this application. All configuration is done through Cloudflare bindings (D1 and KV).

## Database Management

### Backup Database

```bash
wrangler d1 export webstash --output=backup.sql
```

### Restore Database

```bash
wrangler d1 execute webstash --file=backup.sql
```

### Query Database

```bash
wrangler d1 execute webstash --command="SELECT * FROM users"
```

## Troubleshooting

### Local Development Issues

If you encounter issues with local development:

1. Make sure you're using the latest version of Wrangler
2. Clear the `.wrangler` directory and restart dev server
3. Check that your Astro and Cloudflare adapter versions are compatible

### Deployment Issues

1. Ensure all bindings (D1 and KV) are properly configured in wrangler.toml
2. Verify that the database schema has been applied
3. Check Cloudflare Pages logs for any errors

## Security Notes

- The app uses SHA-256 for password hashing (for simplicity). In production, consider using bcrypt or Argon2.
- Sessions are stored in KV with a 7-day expiration
- All cookies are set with HttpOnly, Secure, and SameSite=Strict flags
