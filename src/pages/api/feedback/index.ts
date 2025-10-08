import type { APIRoute } from 'astro';

export const POST: APIRoute = async ({ locals, request }) => {
  if (!locals.user) {
    return new Response(JSON.stringify({ error: 'Unauthorized' }), {
      status: 401,
      headers: { 'Content-Type': 'application/json' }
    });
  }

  try {
    const data = await request.json();

    if (!data.type || !data.message) {
      return new Response(JSON.stringify({ error: 'Type and message are required' }), {
        status: 400,
        headers: { 'Content-Type': 'application/json' }
      });
    }

    if (!['bug', 'feature'].includes(data.type)) {
      return new Response(JSON.stringify({ error: 'Type must be either "bug" or "feature"' }), {
        status: 400,
        headers: { 'Content-Type': 'application/json' }
      });
    }

    const id = crypto.randomUUID();
    const now = Math.floor(Date.now() / 1000);

    await locals.runtime.env.DB
      .prepare('INSERT INTO feedback (id, user_id, user_email, type, message, created_at) VALUES (?, ?, ?, ?, ?, ?)')
      .bind(id, locals.user.id, locals.user.email, data.type, data.message, now)
      .run();

    return new Response(JSON.stringify({ success: true, id }), {
      status: 201,
      headers: { 'Content-Type': 'application/json' }
    });
  } catch (error) {
    return new Response(JSON.stringify({ error: 'Invalid request' }), {
      status: 400,
      headers: { 'Content-Type': 'application/json' }
    });
  }
};
