import type { ErrorResultType } from '@/types/shared/services/errors';

export type SuccessResult<T> = {
  success: true;
  data: T;
};

export type NotSuccessResult = {
  success: false;
  error: ErrorResultType;
};

export type Result<T> = SuccessResult<T> | NotSuccessResult;

export type AsyncResult<T> = Promise<Result<T>>;
