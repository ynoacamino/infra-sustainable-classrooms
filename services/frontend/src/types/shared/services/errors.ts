export type ErrorResultType<
  T extends Record<string, unknown> = Record<string, unknown>,
> = {
  message: string;
  status: number;
  reason: string;
  extend: T;
};
