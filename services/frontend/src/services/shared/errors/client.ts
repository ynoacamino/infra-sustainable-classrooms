import { ServiceError } from './base';

// Client-side errors (not HTTP-related)
export class NetworkError extends ServiceError {
  constructor(
    message = 'Network request failed',
    extend: Record<string, unknown> = {},
  ) {
    super({
      message,
      reason: 'NetworkError',
      status: 503,
      extend,
    });
  }
}

export class TimeoutError extends ServiceError {
  constructor(
    message = 'Request timeout',
    extend: Record<string, unknown> = {},
  ) {
    super({
      message,
      reason: 'TimeoutError',
      status: 408,
      extend,
    });
  }
}

export class ConnectionError extends ServiceError {
  constructor(
    message = 'Connection failed',
    extend: Record<string, unknown> = {},
  ) {
    super({
      message,
      reason: 'ConnectionError',
      status: 503,
      extend,
    });
  }
}

export class AbortError extends ServiceError {
  constructor(
    message = 'Request aborted',
    extend: Record<string, unknown> = {},
  ) {
    super({
      message,
      reason: 'AbortError',
      status: 499,
      extend,
    });
  }
}

// Application-specific errors
export class ValidationError extends ServiceError {
  constructor(
    messageOrFieldErrors:
      | string
      | Record<string, unknown> = 'Validation failed',
    extend: Record<string, unknown> = {},
  ) {
    // Handle field errors passed as first parameter
    if (
      typeof messageOrFieldErrors === 'object' &&
      messageOrFieldErrors !== null
    ) {
      super({
        message: 'Validation failed',
        reason: 'ValidationError',
        status: 400,
        extend: {
          fieldErrors: messageOrFieldErrors,
          ...extend,
        },
      });
    } else {
      super({
        message: messageOrFieldErrors,
        reason: 'ValidationError',
        status: 400,
        extend,
      });
    }
  }

  // Static method to create validation error from field errors
  static fromFieldErrors(
    fieldErrors: Record<string, unknown>,
    message = 'Validation failed',
  ): ValidationError {
    return new ValidationError(message, { fieldErrors });
  }

  // Get field errors from extend data
  get fieldErrors(): Record<string, unknown> | undefined {
    return this.extend.fieldErrors as Record<string, unknown> | undefined;
  }
}

export class InvalidEndpointError extends ServiceError {
  constructor(
    message = 'Invalid endpoint provided',
    extend: Record<string, unknown> = {},
  ) {
    super({
      message,
      reason: 'InvalidEndpoint',
      status: 400,
      extend,
    });
  }
}

export class ConfigurationError extends ServiceError {
  constructor(
    message = 'Configuration error',
    extend: Record<string, unknown> = {},
  ) {
    super({
      message,
      reason: 'ConfigurationError',
      status: 500,
      extend,
    });
  }
}

export class SerializationError extends ServiceError {
  constructor(
    message = 'Serialization failed',
    extend: Record<string, unknown> = {},
  ) {
    super({
      message,
      reason: 'SerializationError',
      status: 500,
      extend,
    });
  }
}

export class DeserializationError extends ServiceError {
  constructor(
    message = 'Deserialization failed',
    extend: Record<string, unknown> = {},
  ) {
    super({
      message,
      reason: 'DeserializationError',
      status: 500,
      extend,
    });
  }
}

export class UnknownError extends ServiceError {
  constructor(
    message = 'Unknown error occurred',
    extend: Record<string, unknown> = {},
  ) {
    super({
      message,
      reason: 'UnknownError',
      status: 500,
      extend,
    });
  }
}

export class RemoteServiceError<
  T extends Record<string, unknown> = Record<string, unknown>,
> extends ServiceError<T> {
  constructor({
    message = 'Remote service error',
    reason = 'RemoteServiceError',
    status = 500,
    extend = {} as T,
  }: {
    message?: string;
    reason?: string;
    status?: number;
    extend?: T;
  } = {}) {
    super({ message, reason, status, extend });
  }
}
