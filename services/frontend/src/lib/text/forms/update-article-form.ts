import { UpdateArticlePayloadSchema } from '@/types/text/schemas/payload';
import type { Field } from '@/types/shared/field';
import type z from 'zod';

export const updateArticleFormSchema = UpdateArticlePayloadSchema;

export const updateArticleFormFields: Field<
  keyof z.infer<typeof updateArticleFormSchema>
>[] = [
  {
    name: 'title',
    label: 'Article Title',
    type: 'text',
    placeholder: 'e.g. Updated Article Title',
    description: 'Enter the article title (3-100 characters).',
  },
  {
    name: 'content',
    label: 'Article Content',
    type: 'html_editor',
    placeholder: 'e.g. Updated article content...',
    description: 'Enter the content of the article (minimum 10 characters).',
  },
];
