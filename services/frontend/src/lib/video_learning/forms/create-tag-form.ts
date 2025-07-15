import { CreateTagPayloadSchema } from '@/types/video_learning/schemas/payload';
import type { Field } from '@/types/shared/field';
import type z from 'zod';

export const createTagFormSchema = CreateTagPayloadSchema;

export const createTagFormFields: Field<
  keyof z.infer<typeof createTagFormSchema>
>[] = [
  {
    name: 'name',
    label: 'Tag Name',
    type: 'text',
    placeholder: 'e.g. beginner, advanced, tutorial',
    description: 'Enter the tag name for classifying videos.',
  },
];
