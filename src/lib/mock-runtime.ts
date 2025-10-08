// Mock runtime for local development without Cloudflare Workers
// This provides in-memory implementations of D1 and KV

interface D1Result {
  results: any[];
  success: boolean;
  meta: {
    changes: number;
    last_row_id?: number;
  };
}

class MockD1PreparedStatement {
  constructor(private sql: string, private params: any[] = []) {}

  bind(...values: any[]) {
    this.params = values;
    return this;
  }

  async first() {
    const result = await this.all();
    return result.results[0] || null;
  }

  async all(): Promise<D1Result> {
    // For development, we'll use an in-memory store
    const db = globalThis.__MOCK_DB__ as Map<string, any>;

    if (!db) {
      return { results: [], success: true, meta: { changes: 0 } };
    }

    // Simple query parsing (this is a simplified mock)
    if (this.sql.includes('SELECT * FROM bookmarks WHERE user_id')) {
      const userId = this.params[0];
      let bookmarks = Array.from(db.values()).filter(
        (item: any) => item._type === 'bookmark' && item.user_id === userId
      );

      // Handle search filtering
      if (this.params.length > 1) {
        const searchPattern = this.params[1];
        // Remove % wildcards to get the actual search term
        const searchTerm = searchPattern.replace(/%/g, '').toLowerCase();

        if (this.sql.includes('title LIKE ? OR description LIKE ? OR url LIKE ? OR tags LIKE ?')) {
          // Filter: 'all' - search across all fields
          bookmarks = bookmarks.filter((item: any) => {
            const title = (item.title || '').toLowerCase();
            const description = (item.description || '').toLowerCase();
            const url = (item.url || '').toLowerCase();
            const tags = (item.tags || '').toLowerCase();
            return title.includes(searchTerm) ||
                   description.includes(searchTerm) ||
                   url.includes(searchTerm) ||
                   tags.includes(searchTerm);
          });
        } else if (this.sql.includes('title LIKE ?')) {
          // Filter: 'title'
          bookmarks = bookmarks.filter((item: any) => {
            const title = (item.title || '').toLowerCase();
            return title.includes(searchTerm);
          });
        } else if (this.sql.includes('url LIKE ?')) {
          // Filter: 'url'
          bookmarks = bookmarks.filter((item: any) => {
            const url = (item.url || '').toLowerCase();
            return url.includes(searchTerm);
          });
        } else if (this.sql.includes('description LIKE ?')) {
          // Filter: 'description'
          bookmarks = bookmarks.filter((item: any) => {
            const description = (item.description || '').toLowerCase();
            return description.includes(searchTerm);
          });
        } else if (this.sql.includes('tags LIKE ?')) {
          // Filter: 'tags'
          bookmarks = bookmarks.filter((item: any) => {
            const tags = (item.tags || '').toLowerCase();
            return tags.includes(searchTerm);
          });
        }
      }

      return { results: bookmarks, success: true, meta: { changes: 0 } };
    }

    if (this.sql.includes('SELECT * FROM bookmarks WHERE id')) {
      const id = this.params[0];
      const bookmark = db.get(`bookmark:${id}`);
      return { results: bookmark ? [bookmark] : [], success: true, meta: { changes: 0 } };
    }

    if (this.sql.includes('SELECT * FROM users WHERE email')) {
      const email = this.params[0];
      const user = Array.from(db.values()).find(
        (item: any) => item._type === 'user' && item.email === email
      );
      return { results: user ? [user] : [], success: true, meta: { changes: 0 } };
    }

    return { results: [], success: true, meta: { changes: 0 } };
  }

  async run(): Promise<D1Result> {
    const db = globalThis.__MOCK_DB__ as Map<string, any>;

    if (!db) {
      return { results: [], success: true, meta: { changes: 0 } };
    }

    // INSERT operations
    if (this.sql.includes('INSERT INTO users')) {
      const [id, email, password_hash, created_at] = this.params;
      db.set(`user:${id}`, { _type: 'user', id, email, password_hash, created_at });
      return { results: [], success: true, meta: { changes: 1 } };
    }

    if (this.sql.includes('INSERT INTO bookmarks')) {
      const [id, user_id, url, title, description, tags, created_at, updated_at] = this.params;
      db.set(`bookmark:${id}`, {
        _type: 'bookmark',
        id,
        user_id,
        url,
        title,
        description,
        tags,
        created_at,
        updated_at
      });
      return { results: [], success: true, meta: { changes: 1 } };
    }

    if (this.sql.includes('INSERT INTO feedback')) {
      const [id, user_id, user_email, type, message, created_at] = this.params;
      db.set(`feedback:${id}`, {
        _type: 'feedback',
        id,
        user_id,
        user_email,
        type,
        message,
        created_at
      });
      return { results: [], success: true, meta: { changes: 1 } };
    }

    // UPDATE operations
    if (this.sql.includes('UPDATE bookmarks')) {
      const id = this.params[this.params.length - 2];
      const bookmark = db.get(`bookmark:${id}`);
      if (bookmark) {
        // Simple update - in real implementation would parse SET clause
        db.set(`bookmark:${id}`, { ...bookmark, updated_at: Math.floor(Date.now() / 1000) });
        return { results: [], success: true, meta: { changes: 1 } };
      }
      return { results: [], success: true, meta: { changes: 0 } };
    }

    // DELETE operations
    if (this.sql.includes('DELETE FROM bookmarks')) {
      const id = this.params[0];
      const deleted = db.delete(`bookmark:${id}`);
      return { results: [], success: true, meta: { changes: deleted ? 1 : 0 } };
    }

    return { results: [], success: true, meta: { changes: 0 } };
  }
}

class MockD1Database {
  prepare(sql: string) {
    return new MockD1PreparedStatement(sql);
  }
}

// Global KV store that persists across requests
const globalKVStore = new Map<string, string>();

class MockKVNamespace {
  async get(key: string): Promise<string | null> {
    return globalKVStore.get(key) || null;
  }

  async put(key: string, value: string, options?: any): Promise<void> {
    globalKVStore.set(key, value);
  }

  async delete(key: string): Promise<void> {
    globalKVStore.delete(key);
  }
}

// Initialize global mock storage
if (typeof globalThis !== 'undefined' && !globalThis.__MOCK_DB__) {
  globalThis.__MOCK_DB__ = new Map();
}

// Singleton mock runtime instance
let mockRuntimeInstance: any = null;

export function createMockRuntime() {
  if (!mockRuntimeInstance) {
    mockRuntimeInstance = {
      env: {
        DB: new MockD1Database(),
        SESSIONS: new MockKVNamespace()
      }
    };
  }
  return mockRuntimeInstance;
}

declare global {
  var __MOCK_DB__: Map<string, any>;
}
