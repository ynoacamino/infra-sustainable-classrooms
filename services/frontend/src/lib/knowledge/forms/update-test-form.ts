import { UpdateTestPayloadSchema } from '@/types/knowledge/schemas/payload';
import type { Field } from '@/types/shared/field';
import type z from 'zod';

export const updateTestFormSchema = UpdateTestPayloadSchema;

export const updateTestFormFields: Field<
  keyof z.infer<typeof updateTestFormSchema>
>[] = [
  {
    name: 'title',
    label: 'Test Title',
    type: 'text',
    placeholder: 'e.g. JavaScript Basics Quiz',
    description: 'Enter the test title (3-200 characters).',
  },
];
