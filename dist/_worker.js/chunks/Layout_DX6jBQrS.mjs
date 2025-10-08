globalThis.process ??= {}; globalThis.process.env ??= {};
import { c as createComponent, a as createAstro, b as addAttribute, r as renderHead, d as renderSlot, e as renderTemplate } from './astro/server_CrZeW1zu.mjs';
/* empty css                             */

const $$Astro = createAstro();
const $$Layout = createComponent(($$result, $$props, $$slots) => {
  const Astro2 = $$result.createAstro($$Astro, $$props, $$slots);
  Astro2.self = $$Layout;
  const { title } = Astro2.props;
  return renderTemplate`<html lang="en"> <head><meta charset="UTF-8"><meta name="description" content="webStash - Save your favorite websites"><meta name="viewport" content="width=device-width"><link rel="icon" type="image/svg+xml" href="/favicon.svg"><meta name="generator"${addAttribute(Astro2.generator, "content")}><title>${title}</title>${renderHead()}</head> <body> <nav class="navbar"> <div class="container"> <h1>📚 webStash</h1> <div class="nav-links"> <a href="/">Home</a> <a href="/bookmarks">Bookmarks</a> <a href="/logout">Logout</a> </div> </div> </nav> <main class="container"> ${renderSlot($$result, $$slots["default"])} </main>  </body> </html>`;
}, "/Users/cbas/Documents/code-projects/webStash/src/layouts/Layout.astro", void 0);

export { $$Layout as $ };
