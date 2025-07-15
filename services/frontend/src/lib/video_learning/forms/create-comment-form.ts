import { CreateCommentPayloadSchema } from '@/types/video_learning/schemas/payload';
import type { Field } from '@/types/shared/field';
import type z from 'zod';

export const createCommentFormSchema = CreateCommentPayloadSchema;

export const createCommentFormFields: Field<
  keyof z.infer<typeof createCommentFormSchema>
>[] = [
  {
    name: 'title',
    label: 'Comment Title',
    type: 'text',
    placeholder: 'e.g. Great explanation!',
    description: 'Enter a title for your comment.',
  },
  {
    name: 'body',
    label: 'Comment',
    type: 'textarea',
    placeholder: 'e.g. This video helped me understand the concept better...',
    description: 'Write your comment about the video.',
  },
];
