import type { SimpleResponse } from '@/services/shared/response';
import type { User } from '@/types/auth/models';

export type AuthResponse = SimpleResponse & {
  user: User;
  // This is a cookie
  // session: string;
};

export type BackupCodeResponse = AuthResponse & {
  remaining_codes: number;
  // This is a cookie
  // session: string;
};
