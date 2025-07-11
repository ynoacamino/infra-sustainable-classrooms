import type { ErrorResult } from '@/services/shared/errors';

export type SuccessResultType<T> = {
  success: true;
  data: T;
};

export type NotSuccessResultType = {
  success: false;
  error: ErrorResult;
};

export type ResultType<T> = SuccessResultType<T> | NotSuccessResultType;
