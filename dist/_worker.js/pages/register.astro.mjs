globalThis.process ??= {}; globalThis.process.env ??= {};
/* empty css                                     */
import { c as createComponent, a as createAstro, d as renderComponent, r as renderTemplate, m as maybeRenderHead } from '../chunks/astro/server_COAJwCcf.mjs';
import { $ as $$Layout } from '../chunks/Layout_y2Z-hqE8.mjs';
export { renderers } from '../renderers.mjs';

const $$Astro = createAstro();
const $$Register = createComponent(async ($$result, $$props, $$slots) => {
  const Astro2 = $$result.createAstro($$Astro, $$props, $$slots);
  Astro2.self = $$Register;
  const user = Astro2.locals.user;
  if (user) {
    return Astro2.redirect("/bookmarks");
  }
  return renderTemplate`${renderComponent($$result, "Layout", $$Layout, { "title": "Register - webStash" }, { "default": async ($$result2) => renderTemplate` ${maybeRenderHead()}<div class="max-w-md mx-auto mt-16"> <div class="bg-white rounded-lg shadow-sm border border-slate-200 p-8"> <h2 class="text-2xl font-bold text-slate-900 mb-6 text-center">Create your account</h2> <form id="registerForm" class="space-y-5"> <div> <label for="email" class="block text-sm font-medium text-slate-700 mb-1">
Email
</label> <input type="email" id="email" name="email" required class="w-full px-3 py-2 border border-slate-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"> </div> <div> <label for="password" class="block text-sm font-medium text-slate-700 mb-1">
Password
</label> <input type="password" id="password" name="password" required minlength="6" class="w-full px-3 py-2 border border-slate-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"> </div> <div> <label for="confirmPassword" class="block text-sm font-medium text-slate-700 mb-1">
Confirm Password
</label> <input type="password" id="confirmPassword" name="confirmPassword" required minlength="6" class="w-full px-3 py-2 border border-slate-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"> </div> <div id="error" class="hidden text-sm text-red-600 bg-red-50 border border-red-200 rounded-lg p-3"></div> <button type="submit" class="w-full px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors font-medium">
Create account
</button> </form> <p class="mt-6 text-center text-sm text-slate-600">
Already have an account? <a href="/login" class="text-blue-600 hover:text-blue-700 font-medium">Sign in</a> </p> </div> </div> ` })} `;
}, "/Users/cbas/Documents/code-projects/webStash/src/pages/register.astro", void 0);

const $$file = "/Users/cbas/Documents/code-projects/webStash/src/pages/register.astro";
const $$url = "/register";

const _page = /*#__PURE__*/Object.freeze(/*#__PURE__*/Object.defineProperty({
  __proto__: null,
  default: $$Register,
  file: $$file,
  url: $$url
}, Symbol.toStringTag, { value: 'Module' }));

const page = () => _page;

export { page };
