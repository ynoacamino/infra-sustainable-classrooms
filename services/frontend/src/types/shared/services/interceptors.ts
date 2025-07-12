export interface Interceptor {
  onRequest?(
    url: string,
    init: RequestInit,
  ): Promise<RequestInit> | RequestInit;
  onResponse?(url: string, response: Response): Promise<Response> | Response;
}
