import type { APIRoute } from 'astro';
import { createUser, getUserByEmail } from '../../../lib/db';
import { hashPassword, createSession, createSessionCookie } from '../../../lib/auth';

export const POST: APIRoute = async ({ locals, request }) => {
  try {
    const data = await request.json();

    if (!data.email || !data.password) {
      return new Response(JSON.stringify({ error: 'Email and password are required' }), {
        status: 400,
        headers: { 'Content-Type': 'application/json' }
      });
    }

    // Check if user already exists
    const existingUser = await getUserByEmail(locals.runtime.env.DB, data.email);
    if (existingUser) {
      return new Response(JSON.stringify({ error: 'User already exists' }), {
        status: 400,
        headers: { 'Content-Type': 'application/json' }
      });
    }

    // Create user
    const passwordHash = await hashPassword(data.password);
    const user = await createUser(locals.runtime.env.DB, data.email, passwordHash);

    // Create session
    const sessionId = await createSession(locals.runtime.env.SESSIONS, user.id, user.email);

    return new Response(JSON.stringify({ success: true }), {
      status: 201,
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
