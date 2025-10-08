globalThis.process ??= {}; globalThis.process.env ??= {};
/* empty css                                     */
import { c as createComponent, a as createAstro } from '../chunks/astro/server_COAJwCcf.mjs';
import { d as deleteSession } from '../chunks/auth_DxkeBhQK.mjs';
export { renderers } from '../renderers.mjs';

const $$Astro = createAstro();
const $$Logout = createComponent(async ($$result, $$props, $$slots) => {
  const Astro2 = $$result.createAstro($$Astro, $$props, $$slots);
  Astro2.self = $$Logout;
  const sessionId = Astro2.cookies.get("sessionId")?.value;
  if (sessionId && Astro2.locals.runtime?.env?.SESSIONS) {
    await deleteSession(Astro2.locals.runtime.env.SESSIONS, sessionId);
  }
  Astro2.cookies.delete("sessionId", {
    path: "/"
  });
  return Astro2.redirect("/");
}, "/Users/cbas/Documents/code-projects/webStash/src/pages/logout.astro", void 0);

const $$file = "/Users/cbas/Documents/code-projects/webStash/src/pages/logout.astro";
const $$url = "/logout";

const _page = /*#__PURE__*/Object.freeze(/*#__PURE__*/Object.defineProperty({
  __proto__: null,
  default: $$Logout,
  file: $$file,
  url: $$url
}, Symbol.toStringTag, { value: 'Module' }));

const page = () => _page;

export { page };
