import { ErrorResult } from '@/services/shared/errors';
import type { ErrorResultType } from '@/types/shared/services/errors';

export abstract class ServiceError<
  // eslint-disable-next-line @typescript-eslint/no-empty-object-type
  T extends Record<string, unknown> = {},
> extends ErrorResult<T> {
  constructor(params: ErrorResultType<T>) {
    super(params);
    Object.setPrototypeOf(this, new.target.prototype); // fix instanceof
  }
}

export class UnknownError extends ServiceError {
  constructor(message: string) {
    super({
      message,
      reason: 'Unknown',
      status: 500,
      extend: {},
    });
  }
}

export class InternalServerError extends ServiceError {
  constructor(
    message = 'Internal server error',
    extend: Record<string, unknown> = {},
  ) {
    super({
      message,
      reason: 'Internal',
      status: 500,
      extend,
    });
  }
}

export class InvalidEndpointError extends ServiceError {
  constructor() {
    super({
      message: 'The endpoint provided is invalid',
      reason: 'InvalidEndpoint',
      status: 400,
      extend: {},
    });
  }
}

export class NetworkError extends ServiceError {
  constructor(message = 'Network request failed') {
    super({
      message,
      reason: 'NetworkError',
      status: 503,
      extend: {},
    });
  }
}

export class RemoteServiceError<
  T extends Record<string, unknown>,
> extends ServiceError<T> {
  constructor({
    message = 'Remote service error',
    reason = 'RemoteError',
    status = 500,
    extend = {} as T,
  }: Partial<ErrorResultType<T>> = {}) {
    super({ message, reason, status, extend });
  }
}
