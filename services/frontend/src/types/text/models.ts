import type {
  CourseSchema,
  SectionSchema,
  ArticleSchema,
  SimpleResponseSchema,
} from '@/types/text/schemas/models';
import type z from 'zod';

export type Course = z.infer<typeof CourseSchema>;

export type Section = z.infer<typeof SectionSchema>;

export type Article = z.infer<typeof ArticleSchema>;

export type SimpleResponse = z.infer<typeof SimpleResponseSchema>;