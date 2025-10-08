import type { KVNamespace } from '@cloudflare/workers-types';

export interface Session {
  userId: string;
  email: string;
  createdAt: number;
}

export async function hashPassword(password: string): Promise<string> {
  const encoder = new TextEncoder();
  const data = encoder.encode(password);
  const hash = await crypto.subtle.digest('SHA-256', data);
  return Array.from(new Uint8Array(hash))
    .map(b => b.toString(16).padStart(2, '0'))
    .join('');
}

export async function verifyPassword(password: string, hash: string): Promise<boolean> {
  const passwordHash = await hashPassword(password);
  return passwordHash === hash;
}

export function generateSessionId(): string {
  return crypto.randomUUID();
}

export async function createSession(
  kv: KVNamespace,
  userId: string,
  email: string
): Promise<string> {
  const sessionId = generateSessionId();
  const session: Session = {
    userId,
    email,
    createdAt: Date.now()
  };

  // Session expires in 7 days
  await kv.put(`session:${sessionId}`, JSON.stringify(session), {
    expirationTtl: 60 * 60 * 24 * 7
  });

  return sessionId;
}

export async function getSession(kv: KVNamespace, sessionId: string): Promise<Session | null> {
  const sessionData = await kv.get(`session:${sessionId}`);
  if (!sessionData) return null;
  return JSON.parse(sessionData);
}

export async function deleteSession(kv: KVNamespace, sessionId: string): Promise<void> {
  await kv.delete(`session:${sessionId}`);
}

export function createSessionCookie(sessionId: string, maxAge: number = 60 * 60 * 24 * 7): string {
  const isDev = import.meta.env.DEV;
  const secure = isDev ? '' : 'Secure; ';
  return `sessionId=${sessionId}; Path=/; HttpOnly; ${secure}SameSite=Lax; Max-Age=${maxAge}`;
}

export function clearSessionCookie(): string {
  const isDev = import.meta.env.DEV;
  const secure = isDev ? '' : 'Secure; ';
  return `sessionId=; Path=/; HttpOnly; ${secure}SameSite=Lax; Max-Age=0`;
}
