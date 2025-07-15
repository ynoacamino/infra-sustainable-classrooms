import { UpdateSectionPayloadSchema } from '@/types/text/schemas/payload';
import type { Field } from '@/types/shared/field';
import type z from 'zod';

export const updateSectionFormSchema = UpdateSectionPayloadSchema;

export const updateSectionFormFields: Field<
  keyof z.infer<typeof updateSectionFormSchema>
>[] = [
  {
    name: 'title',
    label: 'Section Title',
    type: 'text',
    placeholder: 'e.g. Updated Section Title',
    description: 'Enter the section title (3-100 characters).',
  },
  {
    name: 'description',
    label: 'Section Description',
    type: 'textarea',
    placeholder: 'e.g. Updated section description...',
    description: 'Enter a description of the section (5-200 characters).',
  },
  {
    name: 'order',
    label: 'Section Order',
    type: 'number',
    placeholder: 'e.g. 2',
    description:
      'Enter the order of the section in the course (optional, will update the order if set).',
  },
];
