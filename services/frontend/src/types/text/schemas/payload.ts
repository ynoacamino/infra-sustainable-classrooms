import { CourseSchema, SectionSchema, ArticleSchema } from '@/types/text/schemas/models';
import z from 'zod';

// Course payload schemas
export const CreateCoursePayloadSchema = z.object({
  // session_token is handled via cookie
  title: z
    .string()
    .min(3, 'Title must be at least 3 characters')
    .max(150, 'Title cannot exceed 150 characters')
    .describe('Course title'),
  description: z
    .string()
    .min(10, 'Description must be at least 10 characters')
    .max(300, 'Description cannot exceed 300 characters')
    .describe('Course description'),
  imageUrl: z
    .string()
    .url('Invalid URL')
    .min(5, 'Image URL must be at least 5 characters')
    .max(500, 'Image URL cannot exceed 500 characters')
    .optional()
    .describe('Course image URL'),
});

export const GetCoursePayloadSchema = z.object({
  // session_token is handled via cookie
  course_id: z
    .number()
    .int('Course ID must be an integer')
    .describe('Course unique identifier'),
});

export const UpdateCoursePayloadSchema = z.object({
  // session_token is handled via cookie
  course_id: z
    .number()
    .int('Course ID must be an integer')
    .describe('Course unique identifier'),
  title: z
    .string()
    .min(3, 'Title must be at least 3 characters')
    .max(150, 'Title cannot exceed 150 characters')
    .optional()
    .describe('Course title'),
  description: z
    .string()
    .min(10, 'Description must be at least 10 characters')
    .max(300, 'Description cannot exceed 300 characters')
    .optional()
    .describe('Course description'),
  imageUrl: z
    .string()
    .url('Invalid URL')
    .min(5, 'Image URL must be at least 5 characters')
    .max(500, 'Image URL cannot exceed 500 characters')
    .optional()
    .describe('Course image URL'),
});

export const DeleteCoursePayloadSchema = z.object({
  // session_token is handled via cookie
  course_id: z
    .number()
    .int('Course ID must be an integer')
    .describe('Course unique identifier'),
});

// Section payload schemas
export const CreateSectionPayloadSchema = z.object({
  // session_token is handled via cookie
  course_id: z
    .number()
    .int('Course ID must be an integer')
    .describe('Course unique identifier'),
  title: z
    .string()
    .min(3, 'Title must be at least 3 characters')
    .max(100, 'Title cannot exceed 100 characters')
    .describe('Section title'),
  description: z
    .string()
    .min(5, 'Description must be at least 5 characters')
    .max(200, 'Description cannot exceed 200 characters')
    .describe('Section description'),
  order: z
    .number()
    .int('Order must be an integer')
    .min(1, 'Order must be at least 1')
    .optional()
    .describe('Order of the section in the course (optional, if not set it will be auto-numbered)'),
});

export const GetSectionPayloadSchema = z.object({
  // session_token is handled via cookie
  section_id: z
    .number()
    .int('Section ID must be an integer')
    .describe('Section unique identifier'),
});

export const ListSectionsPayloadSchema = z.object({
  // session_token is handled via cookie
  course_id: z
    .number()
    .int('Course ID must be an integer')
    .describe('Course unique identifier'),
});

export const UpdateSectionPayloadSchema = z.object({
  // session_token is handled via cookie
  section_id: z
    .number()
    .int('Section ID must be an integer')
    .describe('Section unique identifier'),
  title: z
    .string()
    .min(3, 'Title must be at least 3 characters')
    .max(100, 'Title cannot exceed 100 characters')
    .optional()
    .describe('Section title'),
  description: z
    .string()
    .min(5, 'Description must be at least 5 characters')
    .max(200, 'Description cannot exceed 200 characters')
    .optional()
    .describe('Section description'),
  order: z
    .number()
    .int('Order must be an integer')
    .min(1, 'Order must be at least 1')
    .optional()
    .describe('Order of the section in the course (optional, if set will update the order)'),
});

export const DeleteSectionPayloadSchema = z.object({
  // session_token is handled via cookie
  section_id: z
    .number()
    .int('Section ID must be an integer')
    .describe('Section unique identifier'),
});

// Article payload schemas
export const CreateArticlePayloadSchema = z.object({
  // session_token is handled via cookie
  section_id: z
    .number()
    .int('Section ID must be an integer')
    .describe('Section unique identifier'),
  title: z
    .string()
    .min(3, 'Title must be at least 3 characters')
    .max(100, 'Title cannot exceed 100 characters')
    .describe('Article title'),
  content: z
    .string()
    .min(10, 'Content must be at least 10 characters')
    .describe('Article content'),
});

export const GetArticlePayloadSchema = z.object({
  // session_token is handled via cookie
  article_id: z
    .number()
    .int('Article ID must be an integer')
    .describe('Article unique identifier'),
});

export const ListArticlesPayloadSchema = z.object({
  // session_token is handled via cookie
  section_id: z
    .number()
    .int('Section ID must be an integer')
    .describe('Section unique identifier'),
});

export const UpdateArticlePayloadSchema = z.object({
  // session_token is handled via cookie
  article_id: z
    .number()
    .int('Article ID must be an integer')
    .describe('Article unique identifier'),
  title: z
    .string()
    .min(3, 'Title must be at least 3 characters')
    .max(100, 'Title cannot exceed 100 characters')
    .optional()
    .describe('Article title'),
  content: z
    .string()
    .min(10, 'Content must be at least 10 characters')
    .optional()
    .describe('Article content'),
});

export const DeleteArticlePayloadSchema = z.object({
  // session_token is handled via cookie
  article_id: z
    .number()
    .int('Article ID must be an integer')
    .describe('Article unique identifier'),
});

// Simple payload schema for endpoints that only need authentication
export const ListCoursesPayloadSchema = z.object({
  // session_token is handled via cookie
  // No additional fields needed
});