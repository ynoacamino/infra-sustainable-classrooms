import type { Role } from '@/types/auth/role';

export interface User {
  id: number;
  name: string;
  email: string;
  photo: string;
  role: Role;
  mfaEnabled: boolean;
}

export type UserRequired = {
  user: User;
};
