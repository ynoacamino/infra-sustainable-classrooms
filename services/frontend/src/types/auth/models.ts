import type {
  DeviceInfoSchema,
  TOTPSecretSchema,
  UserSchema,
} from '@/types/auth/schemas/models';
import type z from 'zod';

export type TOTPSecret = z.infer<typeof TOTPSecretSchema>;

export type DeviceInfo = z.infer<typeof DeviceInfoSchema>;

export type User = z.infer<typeof UserSchema>;
