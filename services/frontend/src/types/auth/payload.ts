import type {
  GenerateSecretPayloadSchema,
  VerifyBackupCodePayloadSchema,
  VerifyTOTPPayloadSchema,
} from '@/types/auth/schemas/payload';
import type z from 'zod';

export type GenerateSecretPayload = z.infer<typeof GenerateSecretPayloadSchema>;

export type VerifyTOTPPayload = z.infer<typeof VerifyTOTPPayloadSchema>;

export type VerifyBackupCodePayload = z.infer<
  typeof VerifyBackupCodePayloadSchema
>;
