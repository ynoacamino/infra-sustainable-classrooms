import { CreateTestPayloadSchema } from '@/types/knowledge/schemas/payload';
import type { Field } from '@/types/shared/field';
import type z from 'zod';

export const createTestFormSchema = CreateTestPayloadSchema;

export const createTestFormFields: Field<
  keyof z.infer<typeof createTestFormSchema>
>[] = [
  {
    name: 'title',
    label: 'Test Title',
    type: 'text',
    placeholder: 'e.g. JavaScript Basics Quiz',
    description: 'Enter the test title (3-200 characters).',
  },
];
