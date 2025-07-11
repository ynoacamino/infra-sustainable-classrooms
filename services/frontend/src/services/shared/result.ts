import type { ErrorResult } from '@/services/shared/errors';
import type {
  NotSuccessResultType,
  ResultType,
} from '@/types/shared/services/result';

export class Result<T> {
  private readonly _success: boolean;
  private readonly _data?: T;
  private readonly _error?: ErrorResult;

  private constructor(result: ResultType<T>) {
    this._success = result.success ?? false;
    if (result.success) {
      this._data = result.data;
      this._error = undefined;
    } else {
      this._data = undefined;
      this._error = result.error;
    }
  }

  static ok<T>(data: T): Result<T> {
    return new Result<T>({ success: true, data });
  }

  static fail(error: ErrorResult): Result<never> {
    return new Result({ success: false, error });
  }

  isOk(): this is Result<T> {
    return this._success;
  }

  isError(): this is NotSuccessResultType {
    return !this._success;
  }

  unwrap(): T {
    if (!this._success) throw new Error(this._error?.message);
    return this._data!;
  }

  unwrapError(): ErrorResult {
    if (this._success) throw new Error('No error available');
    return this._error!;
  }
}
