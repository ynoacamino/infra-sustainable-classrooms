import { DeviceInfoSchema, UserIdentifier } from '@/types/auth/schemas/models';
import z from 'zod';

// /totp/generate
export const GenerateSecretPayloadSchema = z.object({
  identifier: UserIdentifier,
});

// /totp/verify
export const VerifyTOTPPayloadSchema = z.object({
  identifier: UserIdentifier,
  totp_code: z.string().regex(/^\d{6}$/, 'TOTP code must be a 6-digit number'),
  device_info: DeviceInfoSchema.optional(),
});

// /backup/verify
export const VerifyBackupCodePayloadSchema = z.object({
  identifier: UserIdentifier,
  backup_code: z.string().length(8, 'Backup code must be exactly 8 characters'),
  device_info: DeviceInfoSchema.optional(),
});
