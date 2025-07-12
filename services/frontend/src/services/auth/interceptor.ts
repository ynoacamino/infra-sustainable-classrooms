import type { Interceptor } from '@/types/shared/services/interceptors';

export class AuthInterceptor implements Interceptor {
  constructor(private token?: string) {}
  async onRequest(url: string, init: RequestInit): Promise<RequestInit> {
    if (this.token) {
      const headers = new Headers(init.headers);
      headers.set('Cookie', `session_token=${this.token}`);
      return { ...init, headers };
    }
    return init;
  }
}
