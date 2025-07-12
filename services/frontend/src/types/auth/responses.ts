import type { User } from '@/types/auth/models';

export interface AuthResponse {
  success: boolean;
  message: string;
  user: User;
  // This is a cookie
  // session: string;
}

export interface BackupCodeResponse {
  success: boolean;
  message: string;
  user: User;
  remaining_codes: number;
  // This is a cookie
  // session: string;
}
