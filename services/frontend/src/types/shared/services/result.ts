import type { ErrorResult } from '@/services/shared/errors/base';

export type SuccessResult<T> = {
  success: true;
  data: T;
};

export type NotSuccessResult = {
  success: false;
  error: ErrorResult;
};

export type Result<T> = SuccessResult<T> | NotSuccessResult;

export type AsyncResult<T> = Promise<Result<T>>;
