globalThis.process ??= {}; globalThis.process.env ??= {};
import { c as createComponent, a as createAstro, f as renderComponent, e as renderTemplate, m as maybeRenderHead } from '../chunks/astro/server_CrZeW1zu.mjs';
import { $ as $$Layout } from '../chunks/Layout_DX6jBQrS.mjs';
/* empty css                                 */
export { renderers } from '../renderers.mjs';

const $$Astro = createAstro();
const $$Login = createComponent(async ($$result, $$props, $$slots) => {
  const Astro2 = $$result.createAstro($$Astro, $$props, $$slots);
  Astro2.self = $$Login;
  const user = Astro2.locals.user;
  if (user) {
    return Astro2.redirect("/bookmarks");
  }
  return renderTemplate`${renderComponent($$result, "Layout", $$Layout, { "title": "Login - webStash", "data-astro-cid-sgpqyurt": true }, { "default": async ($$result2) => renderTemplate` ${maybeRenderHead()}<div class="auth-container" data-astro-cid-sgpqyurt> <div class="card" data-astro-cid-sgpqyurt> <h2 data-astro-cid-sgpqyurt>Login to webStash</h2> <form id="loginForm" data-astro-cid-sgpqyurt> <div class="form-group" data-astro-cid-sgpqyurt> <label for="email" data-astro-cid-sgpqyurt>Email</label> <input type="email" id="email" name="email" required data-astro-cid-sgpqyurt> </div> <div class="form-group" data-astro-cid-sgpqyurt> <label for="password" data-astro-cid-sgpqyurt>Password</label> <input type="password" id="password" name="password" required data-astro-cid-sgpqyurt> </div> <div id="error" class="error" data-astro-cid-sgpqyurt></div> <button type="submit" data-astro-cid-sgpqyurt>Login</button> </form> <p style="margin-top: 1rem; text-align: center;" data-astro-cid-sgpqyurt>
Don't have an account? <a href="/register" data-astro-cid-sgpqyurt>Register</a> </p> </div> </div> ` })}  `;
}, "/Users/cbas/Documents/code-projects/webStash/src/pages/login.astro", void 0);

const $$file = "/Users/cbas/Documents/code-projects/webStash/src/pages/login.astro";
const $$url = "/login";

const _page = /*#__PURE__*/Object.freeze(/*#__PURE__*/Object.defineProperty({
  __proto__: null,
  default: $$Login,
  file: $$file,
  url: $$url
}, Symbol.toStringTag, { value: 'Module' }));

const page = () => _page;

export { page };
