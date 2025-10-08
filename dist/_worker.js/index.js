globalThis.process ??= {}; globalThis.process.env ??= {};
import { renderers } from './renderers.mjs';
import { createExports } from './_@astrojs-ssr-adapter.mjs';
import { manifest } from './manifest_D8-22dQN.mjs';

const _page0 = () => import('./pages/_image.astro.mjs');
const _page1 = () => import('./pages/api/auth/login.astro.mjs');
const _page2 = () => import('./pages/api/auth/logout.astro.mjs');
const _page3 = () => import('./pages/api/auth/register.astro.mjs');
const _page4 = () => import('./pages/api/bookmarks/_id_.astro.mjs');
const _page5 = () => import('./pages/api/bookmarks.astro.mjs');
const _page6 = () => import('./pages/api/feedback.astro.mjs');
const _page7 = () => import('./pages/bookmarks.astro.mjs');
const _page8 = () => import('./pages/feedback.astro.mjs');
const _page9 = () => import('./pages/login.astro.mjs');
const _page10 = () => import('./pages/logout.astro.mjs');
const _page11 = () => import('./pages/register.astro.mjs');
const _page12 = () => import('./pages/index.astro.mjs');

const pageMap = new Map([
    ["node_modules/@astrojs/cloudflare/dist/entrypoints/image-endpoint.js", _page0],
    ["src/pages/api/auth/login.ts", _page1],
    ["src/pages/api/auth/logout.ts", _page2],
    ["src/pages/api/auth/register.ts", _page3],
    ["src/pages/api/bookmarks/[id].ts", _page4],
    ["src/pages/api/bookmarks/index.ts", _page5],
    ["src/pages/api/feedback/index.ts", _page6],
    ["src/pages/bookmarks.astro", _page7],
    ["src/pages/feedback.astro", _page8],
    ["src/pages/login.astro", _page9],
    ["src/pages/logout.astro", _page10],
    ["src/pages/register.astro", _page11],
    ["src/pages/index.astro", _page12]
]);
const serverIslandMap = new Map();
const _manifest = Object.assign(manifest, {
    pageMap,
    serverIslandMap,
    renderers,
    middleware: () => import('./_astro-internal_middleware.mjs')
});
const _exports = createExports(_manifest);
const __astrojsSsrVirtualEntry = _exports.default;

export { __astrojsSsrVirtualEntry as default, pageMap };
