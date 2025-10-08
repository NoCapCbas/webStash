globalThis.process ??= {}; globalThis.process.env ??= {};
async function getBookmarks(db, userId, search) {
  let query = "SELECT * FROM bookmarks WHERE user_id = ?";
  const params = [userId];
  if (search) {
    query += " AND (title LIKE ? OR description LIKE ? OR url LIKE ?)";
    const searchPattern = `%${search}%`;
    params.push(searchPattern, searchPattern, searchPattern);
  }
  query += " ORDER BY created_at DESC";
  const result = await db.prepare(query).bind(...params).all();
  return result.results.map((row) => ({
    ...row,
    tags: row.tags ? JSON.parse(row.tags) : []
  }));
}
async function getBookmark(db, id, userId) {
  const result = await db.prepare("SELECT * FROM bookmarks WHERE id = ? AND user_id = ?").bind(id, userId).first();
  if (!result) return null;
  return {
    ...result,
    tags: result.tags ? JSON.parse(result.tags) : []
  };
}
async function createBookmark(db, userId, data) {
  const id = crypto.randomUUID();
  const now = Math.floor(Date.now() / 1e3);
  const tags = data.tags ? JSON.stringify(data.tags) : null;
  await db.prepare(
    "INSERT INTO bookmarks (id, user_id, url, title, description, tags, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
  ).bind(id, userId, data.url, data.title, data.description || null, tags, now, now).run();
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
async function updateBookmark(db, id, userId, data) {
  const now = Math.floor(Date.now() / 1e3);
  const updates = [];
  const params = [];
  if (data.url) {
    updates.push("url = ?");
    params.push(data.url);
  }
  if (data.title) {
    updates.push("title = ?");
    params.push(data.title);
  }
  if (data.description !== void 0) {
    updates.push("description = ?");
    params.push(data.description);
  }
  if (data.tags !== void 0) {
    updates.push("tags = ?");
    params.push(JSON.stringify(data.tags));
  }
  updates.push("updated_at = ?");
  params.push(now);
  params.push(id, userId);
  const result = await db.prepare(`UPDATE bookmarks SET ${updates.join(", ")} WHERE id = ? AND user_id = ?`).bind(...params).run();
  return result.meta.changes > 0;
}
async function deleteBookmark(db, id, userId) {
  const result = await db.prepare("DELETE FROM bookmarks WHERE id = ? AND user_id = ?").bind(id, userId).run();
  return result.meta.changes > 0;
}
async function getUserByEmail(db, email) {
  const result = await db.prepare("SELECT * FROM users WHERE email = ?").bind(email).first();
  return result;
}
async function createUser(db, email, passwordHash) {
  const id = crypto.randomUUID();
  const now = Math.floor(Date.now() / 1e3);
  await db.prepare("INSERT INTO users (id, email, password_hash, created_at) VALUES (?, ?, ?, ?)").bind(id, email, passwordHash, now).run();
  return {
    id,
    email,
    password_hash: passwordHash,
    created_at: now
  };
}

export { getBookmark as a, getBookmarks as b, createUser as c, deleteBookmark as d, createBookmark as e, getUserByEmail as g, updateBookmark as u };
