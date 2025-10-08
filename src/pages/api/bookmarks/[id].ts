import type { APIRoute } from 'astro';
import { getBookmark, updateBookmark, deleteBookmark } from '../../../lib/db';

export const GET: APIRoute = async ({ locals, params }) => {
  if (!locals.user) {
    return new Response(JSON.stringify({ error: 'Unauthorized' }), {
      status: 401,
      headers: { 'Content-Type': 'application/json' }
    });
  }

  const bookmark = await getBookmark(locals.runtime.env.DB, params.id!, locals.user.id);

  if (!bookmark) {
    return new Response(JSON.stringify({ error: 'Bookmark not found' }), {
      status: 404,
      headers: { 'Content-Type': 'application/json' }
    });
  }

  return new Response(JSON.stringify(bookmark), {
    status: 200,
    headers: { 'Content-Type': 'application/json' }
  });
};

export const PUT: APIRoute = async ({ locals, params, request }) => {
  if (!locals.user) {
    return new Response(JSON.stringify({ error: 'Unauthorized' }), {
      status: 401,
      headers: { 'Content-Type': 'application/json' }
    });
  }

  try {
    const data = await request.json() as { url?: string; title?: string; description?: string; tags?: string[] };
    const updated = await updateBookmark(locals.runtime.env.DB, params.id!, locals.user.id, data);

    if (!updated) {
      return new Response(JSON.stringify({ error: 'Bookmark not found' }), {
        status: 404,
        headers: { 'Content-Type': 'application/json' }
      });
    }

    return new Response(JSON.stringify({ success: true }), {
      status: 200,
      headers: { 'Content-Type': 'application/json' }
    });
  } catch (error) {
    return new Response(JSON.stringify({ error: 'Invalid request' }), {
      status: 400,
      headers: { 'Content-Type': 'application/json' }
    });
  }
};

export const DELETE: APIRoute = async ({ locals, params }) => {
  if (!locals.user) {
    return new Response(JSON.stringify({ error: 'Unauthorized' }), {
      status: 401,
      headers: { 'Content-Type': 'application/json' }
    });
  }

  const deleted = await deleteBookmark(locals.runtime.env.DB, params.id!, locals.user.id);

  if (!deleted) {
    return new Response(JSON.stringify({ error: 'Bookmark not found' }), {
      status: 404,
      headers: { 'Content-Type': 'application/json' }
    });
  }

  return new Response(JSON.stringify({ success: true }), {
    status: 200,
    headers: { 'Content-Type': 'application/json' }
  });
};
