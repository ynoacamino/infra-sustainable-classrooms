import type { Interceptor } from '@/types/shared/services/interceptors';
import type { ReadonlyRequestCookies } from 'next/dist/server/web/spec-extension/adapters/request-cookies';

export class AuthInterceptor implements Interceptor {
  private token: string;
  constructor(private cookies: ReadonlyRequestCookies) {
    this.token = this.cookies.get('session')?.value || '';
  }
  async onRequest(url: string, init: RequestInit): Promise<RequestInit> {
    if (this.token) {
      const headers = new Headers(init.headers);
      headers.set('Cookie', `session=${this.token}`);
      return { ...init, headers };
    }
    return init;
  }
  async onResponse(url: string, response: Response): Promise<Response> {
    // set cookie from response headers if available
    const setCookie = response.headers.get('set-cookie');
    if (!setCookie) {
      return response; // no cookie to set
    }
    const parsedCookie = setCookie.split('=')?.at(1); // get the cookie value part
    if (!parsedCookie) {
      return response; // no valid cookie value
    }
    const cookies = this.cookies; // cast to allow setting cookies
    cookies.set('session', parsedCookie, { httpOnly: true, secure: true });
    return response;
  }
}
