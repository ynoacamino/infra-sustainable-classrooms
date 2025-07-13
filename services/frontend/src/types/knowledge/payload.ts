import type {
  CreateTestPayloadSchema,
  GetMyTestsPayloadSchema,
  UpdateTestPayloadSchema,
  DeleteTestPayloadSchema,
  GetTestQuestionsPayloadSchema,
  AddQuestionPayloadSchema,
  UpdateQuestionPayloadSchema,
  DeleteQuestionPayloadSchema,
  GetAvailableTestsPayloadSchema,
  GetTestFormPayloadSchema,
  SubmitTestPayloadSchema,
  GetMySubmissionsPayloadSchema,
  GetSubmissionResultPayloadSchema,
  GetTestPreviewPayloadSchema,
  BulkAddQuestionsPayloadSchema,
  ListTestsPayloadSchema,
  ListQuestionsPayloadSchema,
  ListSubmissionsPayloadSchema,
} from '@/types/knowledge/schemas/payload';
import type z from 'zod';

// === TEACHER PAYLOAD TYPES ===

// Test payload types
export type CreateTestPayload = z.infer<typeof CreateTestPayloadSchema>;
export type GetMyTestsPayload = z.infer<typeof GetMyTestsPayloadSchema>;
export type UpdateTestPayload = z.infer<typeof UpdateTestPayloadSchema>;
export type DeleteTestPayload = z.infer<typeof DeleteTestPayloadSchema>;
export type GetTestQuestionsPayload = z.infer<
  typeof GetTestQuestionsPayloadSchema
>;

// Question payload types
export type AddQuestionPayload = z.infer<typeof AddQuestionPayloadSchema>;
export type UpdateQuestionPayload = z.infer<typeof UpdateQuestionPayloadSchema>;
export type DeleteQuestionPayload = z.infer<typeof DeleteQuestionPayloadSchema>;

// === STUDENT PAYLOAD TYPES ===
export type GetAvailableTestsPayload = z.infer<
  typeof GetAvailableTestsPayloadSchema
>;
export type GetTestFormPayload = z.infer<typeof GetTestFormPayloadSchema>;
export type SubmitTestPayload = z.infer<typeof SubmitTestPayloadSchema>;
export type GetMySubmissionsPayload = z.infer<
  typeof GetMySubmissionsPayloadSchema
>;
export type GetSubmissionResultPayload = z.infer<
  typeof GetSubmissionResultPayloadSchema
>;

// === ENHANCED PAYLOAD TYPES ===
export type GetTestPreviewPayload = z.infer<typeof GetTestPreviewPayloadSchema>;
export type BulkAddQuestionsPayload = z.infer<
  typeof BulkAddQuestionsPayloadSchema
>;

// === LIST PAYLOAD TYPES ===
export type ListTestsPayload = z.infer<typeof ListTestsPayloadSchema>;
export type ListQuestionsPayload = z.infer<typeof ListQuestionsPayloadSchema>;
export type ListSubmissionsPayload = z.infer<
  typeof ListSubmissionsPayloadSchema
>;
