globalThis.process ??= {}; globalThis.process.env ??= {};
import { c as createComponent, a as createAstro, f as renderComponent, e as renderTemplate, m as maybeRenderHead } from '../chunks/astro/server_CrZeW1zu.mjs';
import { $ as $$Layout } from '../chunks/Layout_DX6jBQrS.mjs';
/* empty css                                 */
export { renderers } from '../renderers.mjs';

const $$Astro = createAstro();
const $$Index = createComponent(($$result, $$props, $$slots) => {
  const Astro2 = $$result.createAstro($$Astro, $$props, $$slots);
  Astro2.self = $$Index;
  const user = Astro2.locals.user;
  return renderTemplate`${renderComponent($$result, "Layout", $$Layout, { "title": "webStash - Save Your Favorite Websites", "data-astro-cid-j7pv25f6": true }, { "default": ($$result2) => renderTemplate` ${maybeRenderHead()}<div class="hero" data-astro-cid-j7pv25f6> <h1 data-astro-cid-j7pv25f6>📚 Welcome to webStash</h1> <p data-astro-cid-j7pv25f6>Your personal bookmark manager in the cloud</p> ${user ? renderTemplate`<div class="actions" data-astro-cid-j7pv25f6> <a href="/bookmarks" class="button" data-astro-cid-j7pv25f6>Go to Bookmarks</a> </div>` : renderTemplate`<div class="actions" data-astro-cid-j7pv25f6> <a href="/login" class="button" data-astro-cid-j7pv25f6>Login</a> <a href="/register" class="button secondary" data-astro-cid-j7pv25f6>Register</a> </div>`} </div> <div class="features" data-astro-cid-j7pv25f6> <div class="feature" data-astro-cid-j7pv25f6> <h3 data-astro-cid-j7pv25f6>🔖 Save Bookmarks</h3> <p data-astro-cid-j7pv25f6>Quickly save your favorite websites with title, description, and tags</p> </div> <div class="feature" data-astro-cid-j7pv25f6> <h3 data-astro-cid-j7pv25f6>🔍 Search</h3> <p data-astro-cid-j7pv25f6>Find your bookmarks instantly with full-text search</p> </div> <div class="feature" data-astro-cid-j7pv25f6> <h3 data-astro-cid-j7pv25f6>☁️ Cloud Sync</h3> <p data-astro-cid-j7pv25f6>Access your bookmarks from anywhere with Cloudflare Workers</p> </div> </div> ` })} `;
}, "/Users/cbas/Documents/code-projects/webStash/src/pages/index.astro", void 0);

const $$file = "/Users/cbas/Documents/code-projects/webStash/src/pages/index.astro";
const $$url = "";

const _page = /*#__PURE__*/Object.freeze(/*#__PURE__*/Object.defineProperty({
  __proto__: null,
  default: $$Index,
  file: $$file,
  url: $$url
}, Symbol.toStringTag, { value: 'Module' }));

const page = () => _page;

export { page };
