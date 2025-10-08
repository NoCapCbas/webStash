import type { APIRoute } from 'astro';
import { getUserByEmail } from '../../../lib/db';
import { verifyPassword, createSession, createSessionCookie } from '../../../lib/auth';

export const POST: APIRoute = async ({ locals, request }) => {
  try {
    const data = await request.json() as { email: string; password: string };

    if (!data.email || !data.password) {
      return new Response(JSON.stringify({ error: 'Email and password are required' }), {
        status: 400,
        headers: { 'Content-Type': 'application/json' }
      });
    }

    // Get user
    const user = await getUserByEmail(locals.runtime.env.DB, data.email);
    if (!user) {
      return new Response(JSON.stringify({ error: 'Invalid credentials' }), {
        status: 401,
        headers: { 'Content-Type': 'application/json' }
      });
    }

    // Verify password
    const validPassword = await verifyPassword(data.password, user.password_hash);
    if (!validPassword) {
      return new Response(JSON.stringify({ error: 'Invalid credentials' }), {
        status: 401,
        headers: { 'Content-Type': 'application/json' }
      });
    }

    // Create session
    const sessionId = await createSession(locals.runtime.env.SESSIONS, user.id, user.email);

    return new Response(JSON.stringify({ success: true }), {
      status: 200,
      headers: {
        'Content-Type': 'application/json',
        'Set-Cookie': createSessionCookie(sessionId)
      }
    });
  } catch (error) {
    return new Response(JSON.stringify({ error: 'Invalid request' }), {
      status: 400,
      headers: { 'Content-Type': 'application/json' }
    });
  }
};
