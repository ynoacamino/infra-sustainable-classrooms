import type { ServiceError } from '@/services/shared/errors/base';
import {
  BadRequestError,
  UnauthorizedError,
  ForbiddenError,
  NotFoundError,
  MethodNotAllowedError,
  ConflictError,
  UnprocessableEntityError,
  TooManyRequestsError,
  InternalServerError,
  NotImplementedError,
  BadGatewayError,
  ServiceUnavailableError,
  GatewayTimeoutError,
} from '@/services/shared/errors/http';
import {
  NetworkError,
  TimeoutError,
  ConnectionError,
  AbortError,
  RemoteServiceError,
} from '@/services/shared/errors/client';

// HTTP Status Code to Error Class mapping
const HTTP_ERROR_MAP = {
  400: BadRequestError,
  401: UnauthorizedError,
  403: ForbiddenError,
  404: NotFoundError,
  405: MethodNotAllowedError,
  408: TimeoutError,
  409: ConflictError,
  422: UnprocessableEntityError,
  429: TooManyRequestsError,
  500: InternalServerError,
  501: NotImplementedError,
  502: BadGatewayError,
  503: ServiceUnavailableError,
  504: GatewayTimeoutError,
} as const;

// Type for HTTP status codes that have specific error classes
export type KnownHttpStatusCode = keyof typeof HTTP_ERROR_MAP;

/**
 * Creates an appropriate error instance based on HTTP status code
 * @param status - HTTP status code
 * @param message - Optional custom error message
 * @param extend - Optional additional error data
 * @returns ServiceError instance
 */
export function createHttpError(
  status: number,
  message?: string,
  extend?: Record<string, unknown>,
): ServiceError {
  const ErrorClass = HTTP_ERROR_MAP[status as KnownHttpStatusCode];

  if (ErrorClass) {
    return new ErrorClass(message, extend);
  }

  // For unknown status codes, use RemoteServiceError
  return new RemoteServiceError({
    message: message || `HTTP Error ${status}`,
    reason: `HttpError${status}`,
    status,
    extend,
  });
}

/**
 * Creates an error based on fetch API error types
 * @param error - Error from fetch API
 * @param extend - Optional additional error data
 * @returns ServiceError instance
 */
export function createNetworkError(
  error: Error,
  extend?: Record<string, unknown>,
): ServiceError {
  if (error.name === 'AbortError') {
    return new AbortError(error.message, extend);
  }

  if (error.name === 'TimeoutError') {
    return new TimeoutError(error.message, extend);
  }

  if (error.message.toLowerCase().includes('network')) {
    return new NetworkError(error.message, extend);
  }

  if (error.message.toLowerCase().includes('connection')) {
    return new ConnectionError(error.message, extend);
  }

  // Default to NetworkError for unknown fetch errors
  return new NetworkError(error.message, extend);
}
