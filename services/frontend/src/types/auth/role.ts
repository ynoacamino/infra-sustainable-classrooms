import type { Roles } from '@/lib/auth/enums/roles';

export type Role = (typeof Roles)[keyof typeof Roles];
