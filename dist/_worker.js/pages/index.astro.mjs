globalThis.process ??= {}; globalThis.process.env ??= {};
/* empty css                                     */
import { c as createComponent, a as createAstro, d as renderComponent, r as renderTemplate, m as maybeRenderHead } from '../chunks/astro/server_COAJwCcf.mjs';
import { $ as $$Layout } from '../chunks/Layout_y2Z-hqE8.mjs';
export { renderers } from '../renderers.mjs';

const $$Astro = createAstro();
const $$Index = createComponent(($$result, $$props, $$slots) => {
  const Astro2 = $$result.createAstro($$Astro, $$props, $$slots);
  Astro2.self = $$Index;
  const user = Astro2.locals.user;
  return renderTemplate`${renderComponent($$result, "Layout", $$Layout, { "title": "webStash - Save Your Favorite Websites" }, { "default": ($$result2) => renderTemplate` ${maybeRenderHead()}<div class="max-w-4xl mx-auto"> <!-- Hero Section --> <div class="text-center py-16"> <h1 class="text-5xl font-bold text-slate-900 mb-4">
Welcome to <span class="text-blue-600">webStash</span> </h1> <p class="text-xl text-slate-600 mb-8">
Your personal bookmark manager. Access your favorite sites from anywhere
</p> ${user ? renderTemplate`<div> <a href="/bookmarks" class="inline-block px-6 py-3 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors font-medium">
Go to Bookmarks →
</a> </div>` : renderTemplate`<div class="flex justify-center gap-4"> <a href="/register" class="px-6 py-3 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors font-medium">
Get Started
</a> <a href="/login" class="px-6 py-3 bg-white text-slate-700 border border-slate-300 rounded-lg hover:bg-slate-50 transition-colors font-medium">
Sign In
</a> </div>`} </div> <!-- Features --> <div class="grid md:grid-cols-3 gap-8 py-12"> <div class="text-center"> <div class="text-4xl mb-4">🔖</div> <h3 class="text-lg font-semibold text-slate-900 mb-2">Save Bookmarks</h3> <p class="text-slate-600">
Quickly save your favorite websites with title, description, and tags
</p> </div> <div class="text-center"> <div class="text-4xl mb-4">🔍</div> <h3 class="text-lg font-semibold text-slate-900 mb-2">Search</h3> <p class="text-slate-600">
Find your bookmarks instantly with full-text search
</p> </div> <div class="text-center"> <div class="text-4xl mb-4">🏷️</div> <h3 class="text-lg font-semibold text-slate-900 mb-2">Organize with Tags</h3> <p class="text-slate-600">
Categorize and filter your bookmarks with custom tags
</p> </div> </div> </div> ` })}`;
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
