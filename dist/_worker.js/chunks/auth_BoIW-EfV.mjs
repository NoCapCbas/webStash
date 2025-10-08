globalThis.process ??= {}; globalThis.process.env ??= {};
async function hashPassword(password) {
  const encoder = new TextEncoder();
  const data = encoder.encode(password);
  const hash = await crypto.subtle.digest("SHA-256", data);
  return Array.from(new Uint8Array(hash)).map((b) => b.toString(16).padStart(2, "0")).join("");
}
async function verifyPassword(password, hash) {
  const passwordHash = await hashPassword(password);
  return passwordHash === hash;
}
function generateSessionId() {
  return crypto.randomUUID();
}
async function createSession(kv, userId, email) {
  const sessionId = generateSessionId();
  const session = {
    userId,
    email,
    createdAt: Date.now()
  };
  await kv.put(`session:${sessionId}`, JSON.stringify(session), {
    expirationTtl: 60 * 60 * 24 * 7
  });
  return sessionId;
}
async function getSession(kv, sessionId) {
  const sessionData = await kv.get(`session:${sessionId}`);
  if (!sessionData) return null;
  return JSON.parse(sessionData);
}
async function deleteSession(kv, sessionId) {
  await kv.delete(`session:${sessionId}`);
}

export { createSession as c, deleteSession as d, getSession as g, hashPassword as h, verifyPassword as v };
