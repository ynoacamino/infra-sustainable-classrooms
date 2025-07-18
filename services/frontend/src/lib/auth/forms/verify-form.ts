import { SupportedFields } from '@/lib/shared/enums/field';
import { VerifyTOTPPayloadSchema } from '@/types/auth/schemas/payload';
import type { Field } from '@/types/shared/field';
import { REGEXP_ONLY_DIGITS } from 'input-otp';
import type z from 'zod';

export const verifyFormSchema = VerifyTOTPPayloadSchema;
export const verifyFormFields: Field<keyof z.infer<typeof verifyFormSchema>>[] =
  [
    {
      name: 'identifier',
      label: 'Username',
      type: SupportedFields.TEXT,
      placeholder: 'e.g. example12',
      description: 'Enter your username associated with your account.',
    },
    {
      name: 'totp_code',
      label: 'TOTP Code',
      type: SupportedFields.OTP,
      maxLength: 6,
      pattern: REGEXP_ONLY_DIGITS,
      description:
        'Enter the 6-digit TOTP code generated by your authenticator app.',
    },
  ];
