globalThis.process ??= {}; globalThis.process.env ??= {};
import { d as deleteSession, b as clearSessionCookie } from '../../../chunks/auth_DxkeBhQK.mjs';
export { renderers } from '../../../renderers.mjs';

const POST = async ({ locals, cookies }) => {
  const sessionId = cookies.get("sessionId")?.value;
  if (sessionId) {
    await deleteSession(locals.runtime.env.SESSIONS, sessionId);
  }
  return new Response(JSON.stringify({ success: true }), {
    status: 200,
    headers: {
      "Content-Type": "application/json",
      "Set-Cookie": clearSessionCookie()
    }
  });
};

const _page = /*#__PURE__*/Object.freeze(/*#__PURE__*/Object.defineProperty({
  __proto__: null,
  POST
}, Symbol.toStringTag, { value: 'Module' }));

const page = () => _page;

export { page };
