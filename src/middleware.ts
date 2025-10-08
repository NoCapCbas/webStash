import { defineMiddleware } from 'astro:middleware';
import { getSession } from './lib/auth';
import { createMockRuntime } from './lib/mock-runtime';

export const onRequest = defineMiddleware(async ({ locals, cookies, url }, next) => {
  // Initialize runtime (use mock for local development)
  if (!locals.runtime?.env) {
    locals.runtime = createMockRuntime() as any;
  }

  const sessionId = cookies.get('sessionId')?.value;

  if (sessionId && locals.runtime?.env?.SESSIONS) {
    const session = await getSession(locals.runtime.env.SESSIONS, sessionId);
    if (session) {
      locals.user = {
        id: session.userId,
        email: session.email
      };
    }
  }

  // Redirect to login if not authenticated and accessing protected routes
  const protectedRoutes = ['/bookmarks', '/api/bookmarks', '/feedback', '/api/feedback'];
  const isProtectedRoute = protectedRoutes.some(route => url.pathname.startsWith(route));

  if (isProtectedRoute && !locals.user && !url.pathname.startsWith('/api/')) {
    return Response.redirect(new URL('/login', url));
  }

  return next();
});
