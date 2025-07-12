import z from 'zod';

export const UserIdentifier = z
  .string()
  .min(3, 'Identifier must be at least 3 characters long')
  .max(100, 'Identifier must be at most 100 characters long')
  .describe('User identifier (username/email)');

export const TOTPSecretSchema = z.object({
  totp_url: z.string().url(),
  backup_codes: z.array(z.string()),
  issuer: z.string(),
});

export const UserSchema = z.object({
  id: z.number().int(),
  identifier: UserIdentifier,
  created_at: z.number().int(),
  last_login: z.number().int(),
  is_verified: z.boolean(),
  metadata: z.record(z.string()),
});

export const DeviceInfoSchema = z.object({
  user_agent: z.string().optional(),
  ip_address: z.string().optional(),
  device_id: z.string().optional(),
  platform: z.string().optional(),
});
