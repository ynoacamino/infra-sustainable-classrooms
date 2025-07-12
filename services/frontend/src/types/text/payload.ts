import type {
  CreateCoursePayloadSchema,
  GetCoursePayloadSchema,
  UpdateCoursePayloadSchema,
  DeleteCoursePayloadSchema,
  ListCoursesPayloadSchema,
  CreateSectionPayloadSchema,
  GetSectionPayloadSchema,
  ListSectionsPayloadSchema,
  UpdateSectionPayloadSchema,
  DeleteSectionPayloadSchema,
  CreateArticlePayloadSchema,
  GetArticlePayloadSchema,
  ListArticlesPayloadSchema,
  UpdateArticlePayloadSchema,
  DeleteArticlePayloadSchema,
} from '@/types/text/schemas/payload';
import type z from 'zod';

// Course payload types
export type CreateCoursePayload = z.infer<typeof CreateCoursePayloadSchema>;
export type GetCoursePayload = z.infer<typeof GetCoursePayloadSchema>;
export type UpdateCoursePayload = z.infer<typeof UpdateCoursePayloadSchema>;
export type DeleteCoursePayload = z.infer<typeof DeleteCoursePayloadSchema>;
export type ListCoursesPayload = z.infer<typeof ListCoursesPayloadSchema>;

// Section payload types
export type CreateSectionPayload = z.infer<typeof CreateSectionPayloadSchema>;
export type GetSectionPayload = z.infer<typeof GetSectionPayloadSchema>;
export type ListSectionsPayload = z.infer<typeof ListSectionsPayloadSchema>;
export type UpdateSectionPayload = z.infer<typeof UpdateSectionPayloadSchema>;
export type DeleteSectionPayload = z.infer<typeof DeleteSectionPayloadSchema>;

// Article payload types
export type CreateArticlePayload = z.infer<typeof CreateArticlePayloadSchema>;
export type GetArticlePayload = z.infer<typeof GetArticlePayloadSchema>;
export type ListArticlesPayload = z.infer<typeof ListArticlesPayloadSchema>;
export type UpdateArticlePayload = z.infer<typeof UpdateArticlePayloadSchema>;
export type DeleteArticlePayload = z.infer<typeof DeleteArticlePayloadSchema>;