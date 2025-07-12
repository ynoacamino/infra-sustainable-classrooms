import type { ErrorResultType } from '@/types/shared/services/errors';

export class ErrorResult<
  T extends Record<string, unknown> = Record<string, unknown>,
> {
  private readonly _message: string;
  private readonly _status: number;
  private readonly _reason: string;
  private readonly _extend: T;

  constructor({
    message,
    status,
    reason,
    extend,
  }: Partial<ErrorResultType<T>>) {
    this._message = message ?? 'An error has occurred';
    this._status = status ?? 500;
    this._reason = reason ?? 'Unknown';
    this._extend = extend ?? ({} as T);
  }

  get message(): string {
    return this._message;
  }

  get status(): number {
    return this._status;
  }

  get reason(): string {
    return this._reason;
  }

  get extend(): T {
    return this._extend;
  }

  toJSON(): ErrorResultType<T> {
    return {
      message: this._message,
      reason: this._reason,
      status: this._status,
      extend: this._extend,
    };
  }
}

export abstract class ServiceError<
  T extends Record<string, unknown> = Record<string, unknown>,
> extends ErrorResult<T> {
  constructor(params: ErrorResultType<T>) {
    super(params);
    Object.setPrototypeOf(this, new.target.prototype); // fix instanceof
  }
}
