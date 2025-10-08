globalThis.process ??= {}; globalThis.process.env ??= {};
import { d as defineMiddleware, s as sequence } from './chunks/index_D5DjRt9R.mjs';
import { g as getSession } from './chunks/auth_BoIW-EfV.mjs';
import './chunks/astro-designed-error-pages_C-x_UzC0.mjs';

class MockD1PreparedStatement {
  constructor(sql, params = []) {
    this.sql = sql;
    this.params = params;
  }
  bind(...values) {
    this.params = values;
    return this;
  }
  async first() {
    const result = await this.all();
    return result.results[0] || null;
  }
  async all() {
    const db = globalThis.__MOCK_DB__;
    if (!db) {
      return { results: [], success: true, meta: { changes: 0 } };
    }
    if (this.sql.includes("SELECT * FROM bookmarks WHERE user_id")) {
      const userId = this.params[0];
      const bookmarks = Array.from(db.values()).filter(
        (item) => item._type === "bookmark" && item.user_id === userId
      );
      return { results: bookmarks, success: true, meta: { changes: 0 } };
    }
    if (this.sql.includes("SELECT * FROM bookmarks WHERE id")) {
      const id = this.params[0];
      const bookmark = db.get(`bookmark:${id}`);
      return { results: bookmark ? [bookmark] : [], success: true, meta: { changes: 0 } };
    }
    if (this.sql.includes("SELECT * FROM users WHERE email")) {
      const email = this.params[0];
      const user = Array.from(db.values()).find(
        (item) => item._type === "user" && item.email === email
      );
      return { results: user ? [user] : [], success: true, meta: { changes: 0 } };
    }
    return { results: [], success: true, meta: { changes: 0 } };
  }
  async run() {
    const db = globalThis.__MOCK_DB__;
    if (!db) {
      return { results: [], success: true, meta: { changes: 0 } };
    }
    if (this.sql.includes("INSERT INTO users")) {
      const [id, email, password_hash, created_at] = this.params;
      db.set(`user:${id}`, { _type: "user", id, email, password_hash, created_at });
      return { results: [], success: true, meta: { changes: 1 } };
    }
    if (this.sql.includes("INSERT INTO bookmarks")) {
      const [id, user_id, url, title, description, tags, created_at, updated_at] = this.params;
      db.set(`bookmark:${id}`, {
        _type: "bookmark",
        id,
        user_id,
        url,
        title,
        description,
        tags,
        created_at,
        updated_at
      });
      return { results: [], success: true, meta: { changes: 1 } };
    }
    if (this.sql.includes("UPDATE bookmarks")) {
      const id = this.params[this.params.length - 2];
      const bookmark = db.get(`bookmark:${id}`);
      if (bookmark) {
        db.set(`bookmark:${id}`, { ...bookmark, updated_at: Math.floor(Date.now() / 1e3) });
        return { results: [], success: true, meta: { changes: 1 } };
      }
      return { results: [], success: true, meta: { changes: 0 } };
    }
    if (this.sql.includes("DELETE FROM bookmarks")) {
      const id = this.params[0];
      const deleted = db.delete(`bookmark:${id}`);
      return { results: [], success: true, meta: { changes: deleted ? 1 : 0 } };
    }
    return { results: [], success: true, meta: { changes: 0 } };
  }
}
class MockD1Database {
  prepare(sql) {
    return new MockD1PreparedStatement(sql);
  }
}
class MockKVNamespace {
  store = /* @__PURE__ */ new Map();
  async get(key) {
    return this.store.get(key) || null;
  }
  async put(key, value, options) {
    this.store.set(key, value);
  }
  async delete(key) {
    this.store.delete(key);
  }
}
if (typeof globalThis !== "undefined" && !globalThis.__MOCK_DB__) {
  globalThis.__MOCK_DB__ = /* @__PURE__ */ new Map();
}
function createMockRuntime() {
  return {
    env: {
      DB: new MockD1Database(),
      SESSIONS: new MockKVNamespace()
    }
  };
}

const onRequest$2 = defineMiddleware(async ({ locals, cookies, url }, next) => {
  if (!locals.runtime?.env) {
    locals.runtime = createMockRuntime();
  }
  const sessionId = cookies.get("sessionId")?.value;
  if (sessionId && locals.runtime?.env?.SESSIONS) {
    const session = await getSession(locals.runtime.env.SESSIONS, sessionId);
    if (session) {
      locals.user = {
        id: session.userId,
        email: session.email
      };
    }
  }
  const protectedRoutes = ["/bookmarks", "/api/bookmarks"];
  const isProtectedRoute = protectedRoutes.some((route) => url.pathname.startsWith(route));
  if (isProtectedRoute && !locals.user && !url.pathname.startsWith("/api/")) {
    return Response.redirect(new URL("/login", url));
  }
  return next();
});

const When = {
                	Client: 'client',
                	Server: 'server',
                	Prerender: 'prerender',
                	StaticBuild: 'staticBuild',
                	DevServer: 'devServer',
              	};
            	
              const isBuildContext = Symbol.for('astro:when/buildContext');
              const whenAmI = globalThis[isBuildContext] ? When.Prerender : When.Server;

const middlewares = {
  [When.Client]: () => {
    throw new Error("Client should not run a middleware!");
  },
  [When.DevServer]: (_, next) => next(),
  [When.Server]: (_, next) => next(),
  [When.Prerender]: (ctx, next) => {
    if (ctx.locals.runtime === void 0) {
      ctx.locals.runtime = {
        env: process.env
      };
    }
    return next();
  },
  [When.StaticBuild]: (_, next) => next()
};
const onRequest$1 = middlewares[whenAmI];

const onRequest = sequence(
	onRequest$1,
	onRequest$2
	
);

export { onRequest };
