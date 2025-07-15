import { ServiceError } from './base';

export class BadRequestError extends ServiceError {
  constructor(message = 'Bad request', extend: Record<string, unknown> = {}) {
    super({
      message,
      reason: 'BadRequest',
      status: 400,
      extend,
    });
  }
}

export class UnauthorizedError extends ServiceError {
  constructor(message = 'Unauthorized', extend: Record<string, unknown> = {}) {
    super({
      message,
      reason: 'Unauthorized',
      status: 401,
      extend,
    });
  }
}

export class ForbiddenError extends ServiceError {
  constructor(message = 'Forbidden', extend: Record<string, unknown> = {}) {
    super({
      message,
      reason: 'Forbidden',
      status: 403,
      extend,
    });
  }
}

export class NotFoundError extends ServiceError {
  constructor(message = 'Not found', extend: Record<string, unknown> = {}) {
    super({
      message,
      reason: 'NotFound',
      status: 404,
      extend,
    });
  }
}

export class MethodNotAllowedError extends ServiceError {
  constructor(
    message = 'Method not allowed',
    extend: Record<string, unknown> = {},
  ) {
    super({
      message,
      reason: 'MethodNotAllowed',
      status: 405,
      extend,
    });
  }
}

export class ConflictError extends ServiceError {
  constructor(message = 'Conflict', extend: Record<string, unknown> = {}) {
    super({
      message,
      reason: 'Conflict',
      status: 409,
      extend,
    });
  }
}

export class UnprocessableEntityError extends ServiceError {
  constructor(
    message = 'Unprocessable entity',
    extend: Record<string, unknown> = {},
  ) {
    super({
      message,
      reason: 'UnprocessableEntity',
      status: 422,
      extend,
    });
  }
}

export class TooManyRequestsError extends ServiceError {
  constructor(
    message = 'Too many requests',
    extend: Record<string, unknown> = {},
  ) {
    super({
      message,
      reason: 'TooManyRequests',
      status: 429,
      extend,
    });
  }
}

// 5xx Server Errors
export class InternalServerError extends ServiceError {
  constructor(
    message = 'Internal server error',
    extend: Record<string, unknown> = {},
  ) {
    super({
      message,
      reason: 'InternalServerError',
      status: 500,
      extend,
    });
  }
}

export class NotImplementedError extends ServiceError {
  constructor(
    message = 'Not implemented',
    extend: Record<string, unknown> = {},
  ) {
    super({
      message,
      reason: 'NotImplemented',
      status: 501,
      extend,
    });
  }
}

export class BadGatewayError extends ServiceError {
  constructor(message = 'Bad gateway', extend: Record<string, unknown> = {}) {
    super({
      message,
      reason: 'BadGateway',
      status: 502,
      extend,
    });
  }
}

export class ServiceUnavailableError extends ServiceError {
  constructor(
    message = 'Service unavailable',
    extend: Record<string, unknown> = {},
  ) {
    super({
      message,
      reason: 'ServiceUnavailable',
      status: 503,
      extend,
    });
  }
}

export class GatewayTimeoutError extends ServiceError {
  constructor(
    message = 'Gateway timeout',
    extend: Record<string, unknown> = {},
  ) {
    super({
      message,
      reason: 'GatewayTimeout',
      status: 504,
      extend,
    });
  }
}
