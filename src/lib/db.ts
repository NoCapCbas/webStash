import type { D1Database } from '@cloudflare/workers-types';

export interface Bookmark {
  id: string;
  user_id: string;
  url: string;
  title: string;
  description?: string;
  favicon_url?: string;
  tags?: string[];
  created_at: number;
  updated_at: number;
}

export interface User {
  id: string;
  email: string;
  password_hash: string;
  created_at: number;
}

export async function getBookmarks(
  db: D1Database,
  userId: string,
  search?: string,
  filter?: 'all' | 'title' | 'url' | 'description' | 'tags'
): Promise<Bookmark[]> {
  let query = 'SELECT * FROM bookmarks WHERE user_id = ?';
  const params: any[] = [userId];

  if (search) {
    const searchPattern = `%${search}%`;

    if (!filter || filter === 'all') {
      query += ' AND (title LIKE ? OR description LIKE ? OR url LIKE ? OR tags LIKE ?)';
      params.push(searchPattern, searchPattern, searchPattern, searchPattern);
    } else if (filter === 'title') {
      query += ' AND title LIKE ?';
      params.push(searchPattern);
    } else if (filter === 'url') {
      query += ' AND url LIKE ?';
      params.push(searchPattern);
    } else if (filter === 'description') {
      query += ' AND description LIKE ?';
      params.push(searchPattern);
    } else if (filter === 'tags') {
      query += ' AND tags LIKE ?';
      params.push(searchPattern);
    }
  }

  query += ' ORDER BY created_at DESC';

  const result = await db.prepare(query).bind(...params).all();

  return (result.results as any[]).map(row => ({
    ...row,
    tags: row.tags ? JSON.parse(row.tags) : []
  }));
}

export async function getBookmark(db: D1Database, id: string, userId: string): Promise<Bookmark | null> {
  const result = await db
    .prepare('SELECT * FROM bookmarks WHERE id = ? AND user_id = ?')
    .bind(id, userId)
    .first();

  if (!result) return null;

  return {
    ...(result as any),
    tags: result.tags ? JSON.parse(result.tags as string) : []
  };
}

export async function createBookmark(
  db: D1Database,
  userId: string,
  data: { url: string; title: string; description?: string; tags?: string[] }
): Promise<Bookmark> {
  const id = crypto.randomUUID();
  const now = Math.floor(Date.now() / 1000);
  const tags = data.tags ? JSON.stringify(data.tags) : null;

  await db
    .prepare(
      'INSERT INTO bookmarks (id, user_id, url, title, description, tags, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)'
    )
    .bind(id, userId, data.url, data.title, data.description || null, tags, now, now)
    .run();

  return {
    id,
    user_id: userId,
    url: data.url,
    title: data.title,
    description: data.description,
    tags: data.tags || [],
    created_at: now,
    updated_at: now
  };
}

export async function updateBookmark(
  db: D1Database,
  id: string,
  userId: string,
  data: { url?: string; title?: string; description?: string; tags?: string[] }
): Promise<boolean> {
  const now = Math.floor(Date.now() / 1000);
  const updates: string[] = [];
  const params: any[] = [];

  if (data.url) {
    updates.push('url = ?');
    params.push(data.url);
  }
  if (data.title) {
    updates.push('title = ?');
    params.push(data.title);
  }
  if (data.description !== undefined) {
    updates.push('description = ?');
    params.push(data.description);
  }
  if (data.tags !== undefined) {
    updates.push('tags = ?');
    params.push(JSON.stringify(data.tags));
  }

  updates.push('updated_at = ?');
  params.push(now);

  params.push(id, userId);

  const result = await db
    .prepare(`UPDATE bookmarks SET ${updates.join(', ')} WHERE id = ? AND user_id = ?`)
    .bind(...params)
    .run();

  return result.meta.changes > 0;
}

export async function deleteBookmark(db: D1Database, id: string, userId: string): Promise<boolean> {
  const result = await db
    .prepare('DELETE FROM bookmarks WHERE id = ? AND user_id = ?')
    .bind(id, userId)
    .run();

  return result.meta.changes > 0;
}

export async function getUserByEmail(db: D1Database, email: string): Promise<User | null> {
  const result = await db
    .prepare('SELECT * FROM users WHERE email = ?')
    .bind(email)
    .first();

  return result as User | null;
}

export async function createUser(db: D1Database, email: string, passwordHash: string): Promise<User> {
  const id = crypto.randomUUID();
  const now = Math.floor(Date.now() / 1000);

  await db
    .prepare('INSERT INTO users (id, email, password_hash, created_at) VALUES (?, ?, ?, ?)')
    .bind(id, email, passwordHash, now)
    .run();

  return {
    id,
    email,
    password_hash: passwordHash,
    created_at: now
  };
}
