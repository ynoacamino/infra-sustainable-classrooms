import { CreateAttemptPayloadSchema } from '@/types/codelab/schemas/payload';
import type { Field } from '@/types/shared/field';
import type z from 'zod';

export const createAttemptFormSchema = CreateAttemptPayloadSchema;

export const createAttemptFormFields: Field<
  keyof z.infer<typeof createAttemptFormSchema>
>[] = [
  {
    name: 'code',
    label: 'Your Code Solution',
    type: 'textarea',
    placeholder: 'def sum_two_numbers(a, b):\n    return a + b',
    description: 'Write your code solution here.',
  },
];
