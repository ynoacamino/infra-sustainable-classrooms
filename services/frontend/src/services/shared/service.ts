import { API_BASE_URL } from '@/config/shared/env';
import type { ServiceError } from '@/services/shared/errors/base';
import {
  InvalidEndpointError,
  ValidationError,
} from '@/services/shared/errors/client';
import {
  createHttpError,
  createNetworkError,
} from '@/services/shared/errors/utils';
import type { Interceptor } from '@/types/shared/services/interceptors';
import type { ServiceRequest } from '@/types/shared/services/request';
import type { AsyncResult, Result } from '@/types/shared/services/result';
import type { ZodSchema } from 'zod';
import FormData from 'form-data';

export abstract class Service {
  private apiBaseUrl: string;
  private interceptors: Interceptor[];
  private endpointPrefix: string;

  // The services can have a custom API URL, but if not provided, it will use the default API base URL.
  protected constructor(endpointPrefix: string, apiUrl?: string) {
    this.apiBaseUrl = apiUrl ?? API_BASE_URL;
    this.interceptors = [];
    this.endpointPrefix = this.normalizeEndpoint(endpointPrefix) || '';
  }

  public addInterceptor(interceptor: Interceptor): void {
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

  private async request<T, B extends ZodSchema = ZodSchema>({
    endpoint,
    payload,
    multipart,
    query,
    options,
  }: ServiceRequest<B>): AsyncResult<T> {
    // Validate the endpoint type
    const safeEndpoint = this.normalizeEndpoint(endpoint);

    if (!safeEndpoint) return this.error(new InvalidEndpointError());
    const safeEndpointPrefix = this.endpointPrefix
      ? `${this.endpointPrefix}/`
      : '';
    // Payload validation
    if (payload) {
      const parsed = payload.schema.safeParse(payload.data);
      if (!parsed.success) {
        return this.error(
          ValidationError.fromFieldErrors(parsed.error.flatten().fieldErrors),
        );
      }
    }
    // Handle query parameters
    const queryParams = new URLSearchParams();
    if (query && Array.isArray(query) && query.length > 0) {
      for (const key of query) {
        if (payload?.data && key in payload.data) {
          queryParams.append(key.toString(), String(payload.data[key]));
        }
      }
    }

    const queryString = queryParams.toString();

    // Handle body payload
    const jsonbodyUnsafe: Record<string, unknown> = {};
    const formDataUnsafe: FormData = new FormData();
    if (multipart) {
      const arrayBuffer = await payload?.data?.file?.arrayBuffer();
      const buffer = Buffer.from(arrayBuffer || '');
      formDataUnsafe.append('file', buffer, payload?.data.filename);
      formDataUnsafe.append('filename', payload?.data.filename);
    } else {
      Object.keys(payload?.data || {}).forEach((key) => {
        if (
          !queryParams.has(key) &&
          payload?.data &&
          payload.data[key] !== undefined
        ) {
          jsonbodyUnsafe[key] = payload.data[key];
        }
      });
    }
    const bodySafe = multipart
      ? formDataUnsafe
      : JSON.stringify(jsonbodyUnsafe);
    // Make the request
    try {
      const url = `${this.apiBaseUrl}/${safeEndpointPrefix}${safeEndpoint}${queryString ? `?${queryString}` : ''}`;

      // Ensure the method is uppercase and defaults to GET
      const method = options?.method?.toUpperCase() || 'GET';

      const reqInit: RequestInit = {
        ...options,
        method,
        headers: multipart
          ? {
              ...formDataUnsafe.getHeaders(),
              ...(options?.headers || {}),
            }
          : {
              'Content-Type': 'application/json',
              ...(options?.headers || {}),
            },
        body: method !== 'GET' ? bodySafe : undefined,
      };
      const finalReqInit = await this.applyRequestInterceptors(url, reqInit);
      console.log('finalReqInit', finalReqInit);
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
        const errorMessage =
          typeof body === 'object' && 'message' in body
            ? body.message
            : typeof body === 'string'
              ? body
              : 'Unknown error from remote service';
        const extendedError =
          typeof body === 'object' ? body : { raw: String(body) };
        const error = createHttpError(
          finalResponse.status,
          errorMessage,
          extendedError,
        );
        return this.error(error);
      }

      return this.result<T>(body);
    } catch (err) {
      if (err instanceof Error) {
        return this.error(createNetworkError(err));
      }
      return this.error(createNetworkError(new Error('Unexpected error')));
    }
  }

  protected result<T>(data: T): Result<T> {
    return { success: true, data };
  }

  protected error(error: ServiceError): Result<never> {
    return { success: false, error: error.toJSON() };
  }

  private normalizeEndpoint(
    endpoint: (string | number) | (string | number)[],
  ): string | undefined {
    const raw = Array.isArray(endpoint) ? endpoint.join('/') : endpoint;
    const cleaned = raw
      .toString()
      .trim()
      .replace(/\/+/g, '/')
      .replace(/^\/|\/$/g, '');

    if (!cleaned || /\s/.test(cleaned)) {
      return undefined;
    }

    return cleaned;
  }

  // TODO: make payload validation for get requests
  protected async get<T, B extends ZodSchema = ZodSchema>({
    endpoint,
    payload,
    query,
    options,
  }: ServiceRequest<B>): AsyncResult<T> {
    return this.request<T, B>({
      endpoint,
      payload,
      query,
      options: { ...options, method: 'GET' },
    });
  }

  protected async post<T, B extends ZodSchema = ZodSchema>({
    endpoint,
    payload,
    multipart,
    query,
    options,
  }: ServiceRequest<B>): AsyncResult<T> {
    return this.request<T>({
      endpoint,
      payload,
      query,
      multipart,
      options: {
        ...options,
        method: 'POST',
      },
    });
  }

  protected async put<T, B extends ZodSchema = ZodSchema>({
    endpoint,
    payload,
    query,
    multipart,
    options,
  }: ServiceRequest<B>): AsyncResult<T> {
    return this.request<T>({
      endpoint,
      payload,
      query,
      multipart,
      options: {
        ...options,
        method: 'PUT',
      },
    });
  }

  protected async patch<T, B extends ZodSchema = ZodSchema>({
    endpoint,
    payload,
    query,
    options,
  }: ServiceRequest<B>): AsyncResult<T> {
    return this.request<T>({
      endpoint,
      payload,
      query,
      options: {
        ...options,
        method: 'PATCH',
      },
    });
  }

  // TODO: make payload validation for delete requests
  protected async delete<T, B extends ZodSchema = ZodSchema>({
    endpoint,
    payload,
    query,
    options,
  }: ServiceRequest<B>): AsyncResult<T> {
    return this.request<T>({
      endpoint,
      payload,
      query,
      options: { ...options, method: 'DELETE' },
    });
  }
}
