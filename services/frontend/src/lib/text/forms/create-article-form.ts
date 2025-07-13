import { CreateArticlePayloadSchema } from '@/types/text/schemas/payload';
import type { Field } from '@/types/shared/field';
import type z from 'zod';

export const createArticleFormSchema = CreateArticlePayloadSchema;

export const createArticleFormFields: Field<
  keyof z.infer<typeof createArticleFormSchema>
>[] = [
  {
    name: 'title',
    label: 'Article Title',
    type: 'text',
    placeholder: 'e.g. What is JavaScript?',
    description: 'Enter the article title (3-100 characters).',
  },
  {
    name: 'content',
    label: 'Article Content',
    type: 'textarea',
    placeholder: 'e.g. JavaScript is a programming language...',
    description: 'Enter the content of the article (minimum 10 characters).',
  },
];
