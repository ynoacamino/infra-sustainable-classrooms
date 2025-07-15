import type {
  CourseSchema,
  SectionSchema,
  ArticleSchema,
} from '@/types/text/schemas/models';
import type z from 'zod';

export type Course = z.infer<typeof CourseSchema>;

export type Section = z.infer<typeof SectionSchema>;

export type Article = z.infer<typeof ArticleSchema>;
