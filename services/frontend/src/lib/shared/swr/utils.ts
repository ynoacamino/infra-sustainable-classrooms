import type { ErrorResult } from '@/services/shared/errors/base';
import type { Result } from '@/types/shared/services/result';
import type { SWRAllResult, SWRFormattedResponse } from '@/types/shared/swr';
import type { SWRResponse } from 'swr';

export function formatSWRResponse<B>({
  isLoading,
  data,
  error,
  mutate,
}: SWRResponse<Result<B>, ErrorResult>): SWRFormattedResponse<B> {
  if (isLoading) {
    return {
      isLoading: true,
      data: null,
      error: null,
      mutate: () => {},
    };
  }
  if (error) {
    return {
      isLoading: false,
      data: null,
      error,
      mutate: () => {},
    };
  }
  if (data) {
    if (data.success) {
      return {
        isLoading: false,
        data: data.data,
        error: null,
        mutate,
      };
    } else {
      return {
        isLoading: false,
        data: null,
        error: data.error,
        mutate: () => {},
      };
    }
  }
  return {
    isLoading: false,
    data: null,
    error: null,
    mutate: () => {},
  };
}

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export function useSWRAll<T extends SWRFormattedResponse<any>[]>(
  hooks: [...T],
): SWRAllResult<T> {
  const isLoading = hooks.some((h) => h.isLoading);
  const errors = hooks.map((h) => h.error).filter((e) => e !== null);
  const data = hooks.map((h) => h.data) as SWRAllResult<T>['data'];

  const mutateAll = () => {
    hooks.forEach((h) => h.mutate());
  };

  return {
    isLoading,
    errors,
    data,
    mutateAll,
  };
}
