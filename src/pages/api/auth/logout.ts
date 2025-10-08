import type { APIRoute } from 'astro';
import { deleteSession, clearSessionCookie } from '../../../lib/auth';

export const POST: APIRoute = async ({ locals, cookies }) => {
  const sessionId = cookies.get('sessionId')?.value;

  if (sessionId) {
    await deleteSession(locals.runtime.env.SESSIONS, sessionId);
  }

  return new Response(JSON.stringify({ success: true }), {
    status: 200,
    headers: {
      'Content-Type': 'application/json',
      'Set-Cookie': clearSessionCookie()
    }
  });
};
