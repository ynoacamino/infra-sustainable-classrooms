import { CreateCategoryPayloadSchema } from '@/types/video_learning/schemas/payload';
import type { Field } from '@/types/shared/field';
import type z from 'zod';

export const createCategoryFormSchema = CreateCategoryPayloadSchema;

export const createCategoryFormFields: Field<
  keyof z.infer<typeof createCategoryFormSchema>
>[] = [
  {
    name: 'name',
    label: 'Category Name',
    type: 'text',
    placeholder: 'e.g. Science, Mathematics, History',
    description: 'Enter the category name for organizing videos.',
  },
];
