globalThis.process ??= {}; globalThis.process.env ??= {};
import { c as createComponent, a as createAstro } from '../chunks/astro/server_CrZeW1zu.mjs';
export { renderers } from '../renderers.mjs';

const $$Astro = createAstro();
const $$Logout = createComponent(async ($$result, $$props, $$slots) => {
  const Astro2 = $$result.createAstro($$Astro, $$props, $$slots);
  Astro2.self = $$Logout;
  await fetch(new URL("/api/auth/logout", Astro2.url), {
    method: "POST"
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
