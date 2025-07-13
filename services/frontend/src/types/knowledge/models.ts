import type {
  SimpleResponseSchema,
  TestSchema,
  QuestionSchema,
  QuestionFormSchema,
  AnswerSchema,
  SubmissionSchema,
  QuestionResultSchema,
  SubmissionResultSchema,
  TestsResponseSchema,
  QuestionsResponseSchema,
  FormResponseSchema,
  SubmissionsResponseSchema,
  SubmitResponseSchema,
  UserAccessResponseSchema,
  TestPreviewResponseSchema,
  BulkQuestionInputSchema,
  BulkQuestionResponseSchema,
} from '@/types/knowledge/schemas/models';
import type z from 'zod';

// Basic types
export type SimpleResponse = z.infer<typeof SimpleResponseSchema>;

export type Test = z.infer<typeof TestSchema>;

export type Question = z.infer<typeof QuestionSchema>;

export type QuestionForm = z.infer<typeof QuestionFormSchema>;

export type Answer = z.infer<typeof AnswerSchema>;

export type Submission = z.infer<typeof SubmissionSchema>;

export type QuestionResult = z.infer<typeof QuestionResultSchema>;

export type SubmissionResult = z.infer<typeof SubmissionResultSchema>;

// Response types
export type TestsResponse = z.infer<typeof TestsResponseSchema>;

export type QuestionsResponse = z.infer<typeof QuestionsResponseSchema>;

export type FormResponse = z.infer<typeof FormResponseSchema>;

export type SubmissionsResponse = z.infer<typeof SubmissionsResponseSchema>;

export type SubmitResponse = z.infer<typeof SubmitResponseSchema>;

export type UserAccessResponse = z.infer<typeof UserAccessResponseSchema>;

// Enhanced types
export type TestPreviewResponse = z.infer<typeof TestPreviewResponseSchema>;

export type BulkQuestionInput = z.infer<typeof BulkQuestionInputSchema>;

export type BulkQuestionResponse = z.infer<typeof BulkQuestionResponseSchema>;
