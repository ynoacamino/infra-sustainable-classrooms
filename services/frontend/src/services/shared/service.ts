import { API_BASE_URL } from '@/config/shared/env';
import { Result } from '@/services/shared/result';
import {
  InvalidEndpointError,
  NetworkError,
  RemoteServiceError,
  type ServiceError,
} from '@/services/shared/service-errors';
import type { Interceptor } from '@/types/shared/services/interceptors';

export abstract class Service {
  private apiBaseUrl: string;
  private interceptors: Interceptor[];

  // The services can have a custom API URL, but if not provided, it will use the default API base URL.
  protected constructor(apiUrl?: string) {
    this.apiBaseUrl = apiUrl ?? API_BASE_URL;
    this.interceptors = [];
  }

  protected addInterceptor(interceptor: Interceptor): void {
    this.interceptors.push(interceptor);
  }

  private async applyRequestInterceptors(
    url: string,
    options?: RequestInit,
  ): Promise<RequestInit> {
    let requestOptions = options || {};

    for (const interceptor of this.interceptors) {
      if (interceptor.onRequest) {
        requestOptions = await interceptor.onRequest(url, requestOptions);
      }
    }

    return requestOptions;
  }

  private async applyResponseInterceptors(
    url: string,
    response: Response,
  ): Promise<Response> {
    let res = response;

    for (const interceptor of this.interceptors) {
      if (interceptor.onResponse) {
        res = await interceptor.onResponse(url, res);
      }
    }

    return res;
  }

  private async request<T>(
    endpoint: string | string[],
    options?: RequestInit,
  ): Promise<Result<T>> {
    // Validate the endpoint type
    const safeEndpoint = this.normalizeEndpoint(endpoint);

    if (!safeEndpoint) return this.error(new InvalidEndpointError());
    // Make the request
    try {
      const url = `${this.apiBaseUrl}/${safeEndpoint}`;
      const reqInit: RequestInit = {
        ...options,
        headers: {
          'Content-Type': 'application/json',
          ...(options?.headers || {}),
        },
      };
      const finalReqInit = await this.applyRequestInterceptors(url, reqInit);
      const response = await fetch(url, finalReqInit);
      const finalResponse = await this.applyResponseInterceptors(url, response);

      // Handle 204 No Content
      if (finalResponse.status === 204) {
        return this.result<T>(undefined as T);
      }

      // Handle other responses
      const contentType = finalResponse.headers.get('content-type') || '';

      const isJson = contentType.includes('application/json');
      const body = isJson
        ? await finalResponse.json()
        : await finalResponse.text();

      if (!finalResponse.ok) {
        const error = new RemoteServiceError({
          message:
            typeof body === 'object' && 'message' in body
              ? body.message
              : typeof body === 'string'
                ? body
                : 'Unknown error from remote service',
          status: finalResponse.status,
          reason:
            typeof body === 'object' && 'reason' in body
              ? body.reason
              : finalResponse.statusText || 'RemoteServiceError',
          extend: typeof body === 'object' ? body : { raw: String(body) },
        });

        return this.error(error);
      }

      return this.result<T>(body);
    } catch (err) {
      if (err instanceof Error) {
        return this.error(new NetworkError(err.message));
      }
      return this.error(new NetworkError('Unexpected error'));
    }
  }

  private result<T>(data: T): Result<T> {
    return Result.ok(data);
  }

  private error(error: ServiceError): Result<never> {
    return Result.fail(error);
  }

  private normalizeEndpoint(endpoint: string | string[]): string | undefined {
    const raw = Array.isArray(endpoint) ? endpoint.join('/') : endpoint;
    const cleaned = raw
      .trim()
      .replace(/\/+/g, '/')
      .replace(/^\/|\/$/g, '');

    if (!cleaned || /\s/.test(cleaned)) {
      return undefined;
    }

    return cleaned;
  }

  protected async get<T>(
    endpoint: string | string[],
    options?: RequestInit,
  ): Promise<Result<T>> {
    return this.request<T>(endpoint, { ...options, method: 'GET' });
  }

  protected async post<T>(
    endpoint: string | string[],
    body?: Record<string, unknown>,
    options?: RequestInit,
  ): Promise<Result<T>> {
    return this.request<T>(endpoint, {
      ...options,
      method: 'POST',
      body: body ? JSON.stringify(body) : undefined,
    });
  }

  protected async put<T>(
    endpoint: string | string[],
    body?: Record<string, unknown>,
    options?: RequestInit,
  ): Promise<Result<T>> {
    return this.request<T>(endpoint, {
      ...options,
      method: 'PUT',
      body: body ? JSON.stringify(body) : undefined,
    });
  }

  protected async delete<T>(
    endpoint: string | string[],
    options?: RequestInit,
  ): Promise<Result<T>> {
    return this.request<T>(endpoint, { ...options, method: 'DELETE' });
  }
}
