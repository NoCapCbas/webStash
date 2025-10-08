globalThis.process ??= {}; globalThis.process.env ??= {};
import { c as createComponent, a as createAstro, f as renderComponent, e as renderTemplate, m as maybeRenderHead } from '../chunks/astro/server_CrZeW1zu.mjs';
import { $ as $$Layout } from '../chunks/Layout_DX6jBQrS.mjs';
/* empty css                                    */
export { renderers } from '../renderers.mjs';

const $$Astro = createAstro();
const $$Register = createComponent(async ($$result, $$props, $$slots) => {
  const Astro2 = $$result.createAstro($$Astro, $$props, $$slots);
  Astro2.self = $$Register;
  const user = Astro2.locals.user;
  if (user) {
    return Astro2.redirect("/bookmarks");
  }
  return renderTemplate`${renderComponent($$result, "Layout", $$Layout, { "title": "Register - webStash", "data-astro-cid-qraosrxq": true }, { "default": async ($$result2) => renderTemplate` ${maybeRenderHead()}<div class="auth-container" data-astro-cid-qraosrxq> <div class="card" data-astro-cid-qraosrxq> <h2 data-astro-cid-qraosrxq>Create Account</h2> <form id="registerForm" data-astro-cid-qraosrxq> <div class="form-group" data-astro-cid-qraosrxq> <label for="email" data-astro-cid-qraosrxq>Email</label> <input type="email" id="email" name="email" required data-astro-cid-qraosrxq> </div> <div class="form-group" data-astro-cid-qraosrxq> <label for="password" data-astro-cid-qraosrxq>Password</label> <input type="password" id="password" name="password" required minlength="6" data-astro-cid-qraosrxq> </div> <div class="form-group" data-astro-cid-qraosrxq> <label for="confirmPassword" data-astro-cid-qraosrxq>Confirm Password</label> <input type="password" id="confirmPassword" name="confirmPassword" required minlength="6" data-astro-cid-qraosrxq> </div> <div id="error" class="error" data-astro-cid-qraosrxq></div> <button type="submit" data-astro-cid-qraosrxq>Register</button> </form> <p style="margin-top: 1rem; text-align: center;" data-astro-cid-qraosrxq>
Already have an account? <a href="/login" data-astro-cid-qraosrxq>Login</a> </p> </div> </div> ` })}  `;
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
