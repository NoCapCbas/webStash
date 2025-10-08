globalThis.process ??= {}; globalThis.process.env ??= {};
import { g as getUserByEmail } from '../../../chunks/db_Cs_0ya7l.mjs';
import { v as verifyPassword, c as createSession } from '../../../chunks/auth_BoIW-EfV.mjs';
export { renderers } from '../../../renderers.mjs';

const POST = async ({ locals, request }) => {
  try {
    const data = await request.json();
    if (!data.email || !data.password) {
      return new Response(JSON.stringify({ error: "Email and password are required" }), {
        status: 400,
        headers: { "Content-Type": "application/json" }
      });
    }
    const user = await getUserByEmail(locals.runtime.env.DB, data.email);
    if (!user) {
      return new Response(JSON.stringify({ error: "Invalid credentials" }), {
        status: 401,
        headers: { "Content-Type": "application/json" }
      });
    }
    const validPassword = await verifyPassword(data.password, user.password_hash);
    if (!validPassword) {
      return new Response(JSON.stringify({ error: "Invalid credentials" }), {
        status: 401,
        headers: { "Content-Type": "application/json" }
      });
    }
    const sessionId = await createSession(locals.runtime.env.SESSIONS, user.id, user.email);
    return new Response(JSON.stringify({ success: true }), {
      status: 200,
      headers: {
        "Content-Type": "application/json",
        "Set-Cookie": `sessionId=${sessionId}; Path=/; HttpOnly; Secure; SameSite=Strict; Max-Age=${60 * 60 * 24 * 7}`
      }
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
  POST
}, Symbol.toStringTag, { value: 'Module' }));

const page = () => _page;

export { page };
