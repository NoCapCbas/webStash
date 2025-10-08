globalThis.process ??= {}; globalThis.process.env ??= {};
import { a as getBookmark, u as updateBookmark, d as deleteBookmark } from '../../../chunks/db_SJnkQxxZ.mjs';
export { renderers } from '../../../renderers.mjs';

const GET = async ({ locals, params }) => {
  if (!locals.user) {
    return new Response(JSON.stringify({ error: "Unauthorized" }), {
      status: 401,
      headers: { "Content-Type": "application/json" }
    });
  }
  const bookmark = await getBookmark(locals.runtime.env.DB, params.id, locals.user.id);
  if (!bookmark) {
    return new Response(JSON.stringify({ error: "Bookmark not found" }), {
      status: 404,
      headers: { "Content-Type": "application/json" }
    });
  }
  return new Response(JSON.stringify(bookmark), {
    status: 200,
    headers: { "Content-Type": "application/json" }
  });
};
const PUT = async ({ locals, params, request }) => {
  if (!locals.user) {
    return new Response(JSON.stringify({ error: "Unauthorized" }), {
      status: 401,
      headers: { "Content-Type": "application/json" }
    });
  }
  try {
    const data = await request.json();
    const updated = await updateBookmark(locals.runtime.env.DB, params.id, locals.user.id, data);
    if (!updated) {
      return new Response(JSON.stringify({ error: "Bookmark not found" }), {
        status: 404,
        headers: { "Content-Type": "application/json" }
      });
    }
    return new Response(JSON.stringify({ success: true }), {
      status: 200,
      headers: { "Content-Type": "application/json" }
    });
  } catch (error) {
    return new Response(JSON.stringify({ error: "Invalid request" }), {
      status: 400,
      headers: { "Content-Type": "application/json" }
    });
  }
};
const DELETE = async ({ locals, params }) => {
  if (!locals.user) {
    return new Response(JSON.stringify({ error: "Unauthorized" }), {
      status: 401,
      headers: { "Content-Type": "application/json" }
    });
  }
  const deleted = await deleteBookmark(locals.runtime.env.DB, params.id, locals.user.id);
  if (!deleted) {
    return new Response(JSON.stringify({ error: "Bookmark not found" }), {
      status: 404,
      headers: { "Content-Type": "application/json" }
    });
  }
  return new Response(JSON.stringify({ success: true }), {
    status: 200,
    headers: { "Content-Type": "application/json" }
  });
};

const _page = /*#__PURE__*/Object.freeze(/*#__PURE__*/Object.defineProperty({
  __proto__: null,
  DELETE,
  GET,
  PUT
}, Symbol.toStringTag, { value: 'Module' }));

const page = () => _page;

export { page };
