import type { Interceptor } from '@/types/shared/services/interceptors';
import { cookies } from 'next/headers';

export class AuthInterceptor implements Interceptor {
  constructor() {}
  async onRequest(url: string, init: RequestInit): Promise<RequestInit> {
    const token = (await cookies()).get('session')?.value;
    if (token) {
      const headers = new Headers(init.headers);
      headers.set('Cookie', `session_token=${token}`);
      return { ...init, headers };
    }
    return init;
  }
}
