import type { ErrorResultType } from '@/types/shared/services/errors';
import type { Result } from '@/types/shared/services/result';
import type { KeyedMutator } from 'swr';

export type SWRFormattedResponse<B> =
  | {
      isLoading: true;
      data: null;
      error: null;
      mutate: () => void;
    }
  | {
      isLoading: false;
      data: null;
      error: ErrorResultType;
      mutate: () => void;
    }
  | {
      isLoading: false;
      data: B;
      error: null;
      mutate: KeyedMutator<Result<B>>;
    }
  | {
      isLoading: false;
      data: null;
      error: null;
      mutate: () => void;
    };

export type SWRAllResult<T extends unknown[]> = {
  isLoading: boolean;
  errors: ErrorResultType[];
  data: {
    [K in keyof T]: T[K] extends SWRFormattedResponse<infer U> ? U : never;
  };
  mutateAll: () => void;
};
