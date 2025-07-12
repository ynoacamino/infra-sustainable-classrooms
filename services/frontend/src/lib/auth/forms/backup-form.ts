import { SupportedFields } from '@/lib/shared/enums/field';
import { VerifyBackupCodePayloadSchema } from '@/types/auth/schemas/payload';
import type { Field } from '@/types/shared/field';
import { REGEXP_ONLY_DIGITS_AND_CHARS } from 'input-otp';
import type z from 'zod';

export const backupFormSchema = VerifyBackupCodePayloadSchema;
export const backupFormFields: Field<keyof z.infer<typeof backupFormSchema>>[] =
  [
    {
      name: 'identifier',
      label: 'Email',
      type: SupportedFields.EMAIL,
      placeholder: 'e.g. example@example.com',
      description: 'Enter the email address associated with your account.',
    },
    {
      name: 'backup_code',
      label: 'Backup Code',
      type: SupportedFields.OTP,
      maxLength: 8,
      pattern: REGEXP_ONLY_DIGITS_AND_CHARS,
      placeholder: 'Enter your 8-digit backup code',
      description:
        'Enter one of your 8-digit backup codes that you saved when setting up two-factor authentication.',
    },
  ];
