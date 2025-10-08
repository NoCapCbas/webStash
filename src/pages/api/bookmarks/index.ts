import type { APIRoute } from 'astro';
import { getBookmarks, createBookmark } from '../../../lib/db';

export const GET: APIRoute = async ({ locals, url }) => {
  if (!locals.user) {
    return new Response(JSON.stringify({ error: 'Unauthorized' }), {
      status: 401,
      headers: { 'Content-Type': 'application/json' }
    });
  }

  const search = url.searchParams.get('search') || undefined;
  const filter = (url.searchParams.get('filter') as 'all' | 'title' | 'url' | 'description' | 'tags') || 'all';
  const bookmarks = await getBookmarks(locals.runtime.env.DB, locals.user.id, search, filter);

  return new Response(JSON.stringify(bookmarks), {
    status: 200,
    headers: { 'Content-Type': 'application/json' }
  });
};

export const POST: APIRoute = async ({ locals, request }) => {
  if (!locals.user) {
    return new Response(JSON.stringify({ error: 'Unauthorized' }), {
      status: 401,
      headers: { 'Content-Type': 'application/json' }
    });
  }

  try {
    const data = await request.json() as { url: string; title: string; description?: string; tags?: string[] };

    if (!data.url || !data.title) {
      return new Response(JSON.stringify({ error: 'URL and title are required' }), {
        status: 400,
        headers: { 'Content-Type': 'application/json' }
      });
    }

    const bookmark = await createBookmark(locals.runtime.env.DB, locals.user.id, data);

    return new Response(JSON.stringify(bookmark), {
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
