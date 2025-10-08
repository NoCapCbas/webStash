globalThis.process ??= {}; globalThis.process.env ??= {};
import { c as createComponent, a as createAstro, m as maybeRenderHead, b as addAttribute, e as renderTemplate, f as renderComponent } from '../chunks/astro/server_CrZeW1zu.mjs';
import { $ as $$Layout } from '../chunks/Layout_DX6jBQrS.mjs';
/* empty css                                     */
import { b as getBookmarks } from '../chunks/db_Cs_0ya7l.mjs';
export { renderers } from '../renderers.mjs';

const $$Astro$1 = createAstro();
const $$BookmarkCard = createComponent(($$result, $$props, $$slots) => {
  const Astro2 = $$result.createAstro($$Astro$1, $$props, $$slots);
  Astro2.self = $$BookmarkCard;
  const { id, url, title, description, tags, created_at } = Astro2.props;
  const formatDate = (timestamp) => {
    return new Date(timestamp * 1e3).toLocaleDateString("en-US", {
      year: "numeric",
      month: "short",
      day: "numeric"
    });
  };
  const getDomain = (url2) => {
    try {
      return new URL(url2).hostname;
    } catch {
      return url2;
    }
  };
  return renderTemplate`${maybeRenderHead()}<div class="bookmark-card"${addAttribute(id, "data-id")} data-astro-cid-65t4mxqd> <div class="bookmark-header" data-astro-cid-65t4mxqd> <h3 data-astro-cid-65t4mxqd> <a${addAttribute(url, "href")} target="_blank" rel="noopener noreferrer" data-astro-cid-65t4mxqd>${title}</a> </h3> <div class="bookmark-actions" data-astro-cid-65t4mxqd> <button class="edit-btn"${addAttribute(id, "data-id")} data-astro-cid-65t4mxqd>Edit</button> <button class="delete-btn"${addAttribute(id, "data-id")} data-astro-cid-65t4mxqd>Delete</button> </div> </div> <p class="bookmark-url" data-astro-cid-65t4mxqd>${getDomain(url)}</p> ${description && renderTemplate`<p class="bookmark-description" data-astro-cid-65t4mxqd>${description}</p>`} <div class="bookmark-footer" data-astro-cid-65t4mxqd> ${tags && tags.length > 0 && renderTemplate`<div class="tags" data-astro-cid-65t4mxqd> ${tags.map((tag) => renderTemplate`<span class="tag" data-astro-cid-65t4mxqd>${tag}</span>`)} </div>`} <span class="date" data-astro-cid-65t4mxqd>${formatDate(created_at)}</span> </div> </div> `;
}, "/Users/cbas/Documents/code-projects/webStash/src/components/BookmarkCard.astro", void 0);

const $$Astro = createAstro();
const $$Bookmarks = createComponent(async ($$result, $$props, $$slots) => {
  const Astro2 = $$result.createAstro($$Astro, $$props, $$slots);
  Astro2.self = $$Bookmarks;
  if (!Astro2.locals.user) {
    return Astro2.redirect("/login");
  }
  const bookmarks = await getBookmarks(Astro2.locals.runtime.env.DB, Astro2.locals.user.id);
  return renderTemplate`${renderComponent($$result, "Layout", $$Layout, { "title": "Bookmarks - webStash", "data-astro-cid-tvjzaet4": true }, { "default": async ($$result2) => renderTemplate` ${maybeRenderHead()}<div class="bookmarks-header" data-astro-cid-tvjzaet4> <h2 data-astro-cid-tvjzaet4>My Bookmarks</h2> <button id="addBookmarkBtn" class="button" data-astro-cid-tvjzaet4>+ Add Bookmark</button> </div> <div class="search-bar" data-astro-cid-tvjzaet4> <input type="text" id="searchInput" placeholder="Search bookmarks..." data-astro-cid-tvjzaet4> </div> <div id="bookmarksList" data-astro-cid-tvjzaet4> ${bookmarks.length === 0 ? renderTemplate`<div class="empty-state" data-astro-cid-tvjzaet4> <p data-astro-cid-tvjzaet4>No bookmarks yet. Add your first bookmark to get started!</p> </div>` : bookmarks.map((bookmark) => renderTemplate`${renderComponent($$result2, "BookmarkCard", $$BookmarkCard, { ...bookmark, "data-astro-cid-tvjzaet4": true })}`)} </div>  <div id="bookmarkModal" class="modal" data-astro-cid-tvjzaet4> <div class="modal-content" data-astro-cid-tvjzaet4> <span class="close" data-astro-cid-tvjzaet4>&times;</span> <h3 id="modalTitle" data-astro-cid-tvjzaet4>Add Bookmark</h3> <form id="bookmarkForm" data-astro-cid-tvjzaet4> <input type="hidden" id="bookmarkId" data-astro-cid-tvjzaet4> <div class="form-group" data-astro-cid-tvjzaet4> <label for="url" data-astro-cid-tvjzaet4>URL</label> <input type="url" id="url" name="url" required data-astro-cid-tvjzaet4> </div> <div class="form-group" data-astro-cid-tvjzaet4> <label for="title" data-astro-cid-tvjzaet4>Title</label> <input type="text" id="title" name="title" required data-astro-cid-tvjzaet4> </div> <div class="form-group" data-astro-cid-tvjzaet4> <label for="description" data-astro-cid-tvjzaet4>Description</label> <textarea id="description" name="description" rows="3" data-astro-cid-tvjzaet4></textarea> </div> <div class="form-group" data-astro-cid-tvjzaet4> <label for="tags" data-astro-cid-tvjzaet4>Tags (comma-separated)</label> <input type="text" id="tags" name="tags" placeholder="javascript, tutorial, web" data-astro-cid-tvjzaet4> </div> <div id="modalError" class="error" data-astro-cid-tvjzaet4></div> <button type="submit" data-astro-cid-tvjzaet4>Save Bookmark</button> </form> </div> </div> ` })}  `;
}, "/Users/cbas/Documents/code-projects/webStash/src/pages/bookmarks.astro", void 0);

const $$file = "/Users/cbas/Documents/code-projects/webStash/src/pages/bookmarks.astro";
const $$url = "/bookmarks";

const _page = /*#__PURE__*/Object.freeze(/*#__PURE__*/Object.defineProperty({
  __proto__: null,
  default: $$Bookmarks,
  file: $$file,
  url: $$url
}, Symbol.toStringTag, { value: 'Module' }));

const page = () => _page;

export { page };
