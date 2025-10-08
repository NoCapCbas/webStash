globalThis.process ??= {}; globalThis.process.env ??= {};
import { b as getBookmarks, e as createBookmark } from '../../chunks/db_Cs_0ya7l.mjs';
export { renderers } from '../../renderers.mjs';

const GET = async ({ locals, url }) => {
  if (!locals.user) {
    return new Response(JSON.stringify({ error: "Unauthorized" }), {
      status: 401,
      headers: { "Content-Type": "application/json" }
    });
  }
  const search = url.searchParams.get("search") || void 0;
  const bookmarks = await getBookmarks(locals.runtime.env.DB, locals.user.id, search);
  return new Response(JSON.stringify(bookmarks), {
    status: 200,
    headers: { "Content-Type": "application/json" }
  });
};
const POST = async ({ locals, request }) => {
  if (!locals.user) {
    return new Response(JSON.stringify({ error: "Unauthorized" }), {
      status: 401,
      headers: { "Content-Type": "application/json" }
    });
  }
  try {
    const data = await request.json();
    if (!data.url || !data.title) {
      return new Response(JSON.stringify({ error: "URL and title are required" }), {
        status: 400,
        headers: { "Content-Type": "application/json" }
      });
    }
    const bookmark = await createBookmark(locals.runtime.env.DB, locals.user.id, data);
    return new Response(JSON.stringify(bookmark), {
      status: 201,
      headers: { "Content-Type": "application/json" }
    });
  } catch (error) {
    return new Response(JSON.stringify({ error: "Invalid request" }), {
      status: 400,
      headers: { "Content-Type": "application/json" }
    });
  }
};

const _page = /*#__PURE__*/Object.freeze(/*#__PURE__*/Object.defineProperty({
  __proto__: null,
  GET,
  POST
}, Symbol.toStringTag, { value: 'Module' }));

const page = () => _page;

export { page };
