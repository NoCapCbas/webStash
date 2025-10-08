globalThis.process ??= {}; globalThis.process.env ??= {};
/* empty css                                     */
import { c as createComponent, a as createAstro, m as maybeRenderHead, f as addAttribute, r as renderTemplate, d as renderComponent } from '../chunks/astro/server_COAJwCcf.mjs';
import { $ as $$Layout } from '../chunks/Layout_y2Z-hqE8.mjs';
import { b as getBookmarks } from '../chunks/db_SJnkQxxZ.mjs';
export { renderers } from '../renderers.mjs';

function getFaviconUrl(url) {
  try {
    const domain = new URL(url).hostname;
    return `https://www.google.com/s2/favicons?domain=${domain}&sz=64`;
  } catch {
    return "";
  }
}
function getDomain(url) {
  try {
    return new URL(url).hostname;
  } catch {
    return url;
  }
}

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
  const faviconUrl = getFaviconUrl(url);
  const domain = getDomain(url);
  return renderTemplate`${maybeRenderHead()}<div class="bg-white rounded-lg border border-slate-200 p-5 hover:shadow-md transition-shadow"${addAttribute(id, "data-id")}> <div class="flex justify-between items-start mb-2"> <div class="flex items-start gap-3 flex-1"> <img${addAttribute(faviconUrl, "src")} alt="" class="w-6 h-6 mt-1 rounded flex-shrink-0" onerror="this.style.display='none'"> <div class="flex-1 min-w-0"> <h3 class="text-lg font-semibold text-slate-900"> <a${addAttribute(url, "href")} target="_blank" rel="noopener noreferrer" class="hover:text-blue-600 transition-colors"> ${title} </a> </h3> <p class="text-sm text-slate-500 mt-1">${domain}</p> </div> </div> <div class="flex gap-2 ml-4"> <button class="edit-btn text-sm px-3 py-1 text-green-600 hover:text-green-700 hover:bg-green-50 rounded transition-colors"${addAttribute(id, "data-id")}>
Edit
</button> <button class="delete-btn text-sm px-3 py-1 text-red-600 hover:text-red-700 hover:bg-red-50 rounded transition-colors"${addAttribute(id, "data-id")}>
Delete
</button> </div> </div> ${description && renderTemplate`<p class="text-slate-700 mb-3 leading-relaxed ml-9">${description}</p>`} <div class="flex items-center justify-between flex-wrap gap-2 ml-9"> ${tags && tags.length > 0 && renderTemplate`<div class="flex gap-2 flex-wrap"> ${tags.map((tag) => renderTemplate`<span class="px-2 py-1 text-xs bg-blue-50 text-blue-700 rounded-full"> ${tag} </span>`)} </div>`} <span class="text-sm text-slate-500">${formatDate(created_at)}</span> </div> </div>`;
}, "/Users/cbas/Documents/code-projects/webStash/src/components/BookmarkCard.astro", void 0);

const $$Astro = createAstro();
const $$Bookmarks = createComponent(async ($$result, $$props, $$slots) => {
  const Astro2 = $$result.createAstro($$Astro, $$props, $$slots);
  Astro2.self = $$Bookmarks;
  if (!Astro2.locals.user) {
    return Astro2.redirect("/login");
  }
  const search = Astro2.url.searchParams.get("search") || void 0;
  const filter = Astro2.url.searchParams.get("filter") || "all";
  const bookmarks = await getBookmarks(Astro2.locals.runtime.env.DB, Astro2.locals.user.id, search, filter);
  return renderTemplate`${renderComponent($$result, "Layout", $$Layout, { "title": "Bookmarks - webStash" }, { "default": async ($$result2) => renderTemplate` ${maybeRenderHead()}<div class="flex justify-between items-center mb-6"> <h2 class="text-3xl font-bold text-slate-900">My Bookmarks</h2> <button id="addBookmarkBtn" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors font-medium">
+ Add Bookmark
</button> </div> <form id="searchForm" class="mb-6 space-y-3"> <div class="flex gap-3"> <div class="flex-1"> <input type="text" id="searchInput" name="search" placeholder="Search bookmarks..."${addAttribute(search || "", "value")} class="w-full px-4 py-2 border border-slate-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"> </div> <select id="filterSelect" name="filter" class="px-4 py-2 border border-slate-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent bg-white"> <option value="all"${addAttribute(filter === "all", "selected")}>All Fields</option> <option value="title"${addAttribute(filter === "title", "selected")}>Title</option> <option value="url"${addAttribute(filter === "url", "selected")}>URL</option> <option value="description"${addAttribute(filter === "description", "selected")}>Description</option> <option value="tags"${addAttribute(filter === "tags", "selected")}>Tags</option> </select> </div> <button type="submit" class="w-full px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors font-medium">
Search
</button> </form> <div id="bookmarksList" class="space-y-4"> ${bookmarks.length === 0 ? renderTemplate`<div class="text-center py-16 text-slate-500"> <p class="text-lg">No bookmarks yet. Add your first bookmark to get started!</p> </div>` : bookmarks.map((bookmark) => renderTemplate`${renderComponent($$result2, "BookmarkCard", $$BookmarkCard, { ...bookmark })}`)} </div>  <div id="bookmarkModal" class="hidden fixed inset-0 bg-black bg-opacity-50 z-50 flex items-center justify-center p-4"> <div class="bg-white rounded-lg shadow-xl w-full max-w-md p-6 relative"> <button class="close absolute top-4 right-4 text-slate-400 hover:text-slate-600 text-2xl font-bold">
&times;
</button> <h3 id="modalTitle" class="text-2xl font-bold text-slate-900 mb-6">Add Bookmark</h3> <form id="bookmarkForm" class="space-y-4"> <input type="hidden" id="bookmarkId"> <div> <label for="url" class="block text-sm font-medium text-slate-700 mb-1">
URL
</label> <input type="url" id="url" name="url" required class="w-full px-3 py-2 border border-slate-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"> </div> <div> <label for="title" class="block text-sm font-medium text-slate-700 mb-1">
Title
</label> <input type="text" id="title" name="title" required class="w-full px-3 py-2 border border-slate-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"> </div> <div> <label for="description" class="block text-sm font-medium text-slate-700 mb-1">
Description
</label> <textarea id="description" name="description" rows="3" class="w-full px-3 py-2 border border-slate-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"></textarea> </div> <div> <label for="tags" class="block text-sm font-medium text-slate-700 mb-1">
Tags (comma-separated)
</label> <input type="text" id="tags" name="tags" placeholder="javascript, tutorial, web" class="w-full px-3 py-2 border border-slate-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"> </div> <div id="modalError" class="hidden text-sm text-red-600 bg-red-50 border border-red-200 rounded-lg p-3"></div> <button type="submit" class="w-full px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors font-medium">
Save Bookmark
</button> </form> </div> </div> ` })} `;
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
