import type {
  CreateTestPayloadSchema,
  UpdateTestPayloadSchema,
  DeleteTestPayloadSchema,
  GetTestQuestionsPayloadSchema,
  AddQuestionPayloadSchema,
  UpdateQuestionPayloadSchema,
  DeleteQuestionPayloadSchema,
  GetTestFormPayloadSchema,
  SubmitTestPayloadSchema,
  GetSubmissionResultPayloadSchema,
  GetTestPayloadSchema,
  GetQuestionPayloadSchema,
  GetSubmissionPayloadSchema,
} from '@/types/knowledge/schemas/payload';
import type z from 'zod';

// === TEACHER PAYLOAD TYPES ===

// Test payload types
export type CreateTestPayload = z.infer<typeof CreateTestPayloadSchema>;
export type GetTestPayload = z.infer<typeof GetTestPayloadSchema>;
export type UpdateTestPayload = z.infer<typeof UpdateTestPayloadSchema>;
export type DeleteTestPayload = z.infer<typeof DeleteTestPayloadSchema>;

export type GetTestQuestionsPayload = z.infer<
  typeof GetTestQuestionsPayloadSchema
>;

// Question payload types
export type AddQuestionPayload = z.infer<typeof AddQuestionPayloadSchema>;
export type GetQuestionPayload = z.infer<typeof GetQuestionPayloadSchema>;
export type UpdateQuestionPayload = z.infer<typeof UpdateQuestionPayloadSchema>;
export type DeleteQuestionPayload = z.infer<typeof DeleteQuestionPayloadSchema>;

export type GetTestFormPayload = z.infer<typeof GetTestFormPayloadSchema>;
export type SubmitTestPayload = z.infer<typeof SubmitTestPayloadSchema>;
export type GetSubmissionPayload = z.infer<typeof GetSubmissionPayloadSchema>;

export type GetSubmissionResultPayload = z.infer<
  typeof GetSubmissionResultPayloadSchema
>;
