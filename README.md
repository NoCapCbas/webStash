# webStash

A bookmark management application built with Astro and Cloudflare Workers.

## Features

- Save and organize your favorite websites
- Full-text search across bookmarks
- Tags and categories
- User authentication
- Deployed on Cloudflare Workers

## Quick Start

```bash
# Install dependencies
npm install

# Start local development server
npm run dev
```

Visit `http://localhost:4321` and create an account to start saving bookmarks!

### Local Development

**Note for macOS 12.x users:** The app uses an in-memory mock runtime for local development on older macOS versions. Your data will reset when you restart the dev server, but all features work normally. When deployed to Cloudflare, it uses the real D1 database.

For macOS 13.5+ users, you can enable the Cloudflare Workers runtime by setting `platformProxy.enabled: true` in `astro.config.mjs`.

### Production Deployment

See [DEPLOYMENT.md](./DEPLOYMENT.md) for detailed deployment instructions.

```bash
# Build for production
npm run build

# Deploy to Cloudflare
npm run deploy
```

## Tech Stack

- **Frontend**: Astro
- **Backend**: Cloudflare Workers
- **Database**: Cloudflare D1 (SQLite)
- **Authentication**: Cloudflare Workers KV
