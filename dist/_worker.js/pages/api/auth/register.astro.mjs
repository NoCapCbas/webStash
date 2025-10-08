globalThis.process ??= {}; globalThis.process.env ??= {};
import { g as getUserByEmail, c as createUser } from '../../../chunks/db_Cs_0ya7l.mjs';
import { h as hashPassword, c as createSession } from '../../../chunks/auth_BoIW-EfV.mjs';
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
    const existingUser = await getUserByEmail(locals.runtime.env.DB, data.email);
    if (existingUser) {
      return new Response(JSON.stringify({ error: "User already exists" }), {
        status: 400,
        headers: { "Content-Type": "application/json" }
      });
    }
    const passwordHash = await hashPassword(data.password);
    const user = await createUser(locals.runtime.env.DB, data.email, passwordHash);
    const sessionId = await createSession(locals.runtime.env.SESSIONS, user.id, user.email);
    return new Response(JSON.stringify({ success: true }), {
      status: 201,
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
