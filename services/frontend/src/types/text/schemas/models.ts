import z from 'zod';

export const CourseSchema = z.object({
  id: z
    .number()
    .int('Course ID must be an integer')
    .describe('Unique identifier for the course'),
  title: z.string().min(1, 'Title is required').describe('Title of the course'),
  description: z
    .string()
    .min(1, 'Description is required')
    .describe('Description of the course'),
  imageUrl: z
    .string()
    .url('Invalid URL')
    .optional()
    .describe('URL of the course image'),
  created_at: z
    .number()
    .int('Creation timestamp must be an integer')
    .describe(
      'Timestamp when the course was created (milliseconds since epoch)',
    ),
  updated_at: z
    .number()
    .int('Last update timestamp must be an integer')
    .describe(
      'Timestamp when the course was last updated (milliseconds since epoch)',
    ),
});

export const SectionSchema = z.object({
  id: z
    .number()
    .int('Section ID must be an integer')
    .describe('Unique identifier for the section'),
  course_id: z
    .number()
    .int('Course ID must be an integer')
    .describe('ID of the parent course'),
  title: z
    .string()
    .min(1, 'Title is required')
    .describe('Title of the section'),
  description: z
    .string()
    .min(1, 'Description is required')
    .describe('Description of the section'),
  order: z
    .number()
    .int('Order must be an integer')
    .min(1, 'Order must be at least 1')
    .describe(
      'Order of the section in the course (autonumbered for frontend rendering)',
    ),
  created_at: z
    .number()
    .int('Creation timestamp must be an integer')
    .describe(
      'Timestamp when the section was created (milliseconds since epoch)',
    ),
  updated_at: z
    .number()
    .int('Last update timestamp must be an integer')
    .describe(
      'Timestamp when the section was last updated (milliseconds since epoch)',
    ),
});

export const ArticleSchema = z.object({
  id: z
    .number()
    .int('Article ID must be an integer')
    .describe('Unique identifier for the article'),
  section_id: z
    .number()
    .int('Section ID must be an integer')
    .describe('ID of the parent section'),
  title: z
    .string()
    .min(1, 'Title is required')
    .describe('Title of the article'),
  content: z
    .string()
    .min(1, 'Content is required')
    .describe('Content of the article'),
  created_at: z
    .number()
    .int('Creation timestamp must be an integer')
    .describe(
      'Timestamp when the article was created (milliseconds since epoch)',
    ),
  updated_at: z
    .number()
    .int('Last update timestamp must be an integer')
    .describe(
      'Timestamp when the article was last updated (milliseconds since epoch)',
    ),
});

export const SimpleResponseSchema = z.object({
  success: z.boolean().describe('Operation success status'),
  message: z.string().describe('Response message'),
});
