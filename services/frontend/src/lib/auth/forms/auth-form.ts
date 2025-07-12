import { SupportedFields } from '@/lib/shared/enums/field';
import { GenerateSecretPayloadSchema } from '@/types/auth/schemas/payload';
import type { Field } from '@/types/shared/field';
import { z } from 'zod';

export const authFormSchema = GenerateSecretPayloadSchema;

export const authFormFields: Field<keyof z.infer<typeof authFormSchema>>[] = [
  {
    name: 'identifier',
    label: 'Username',
    type: SupportedFields.TEXT,
    placeholder: 'e.g. example12',
    description: 'Enter your username',
  },
];
