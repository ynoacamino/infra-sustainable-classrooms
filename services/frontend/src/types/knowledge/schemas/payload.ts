import z from 'zod';

// Common session token schema (used in all requests but handled via cookie)
const SessionTokenSchema = z.object({
  session_token: z
    .string()
    .optional()
    .describe('Session token (handled via cookie)'),
});

// === TEACHER PAYLOAD SCHEMAS ===

// Test payload schemas
export const CreateTestPayloadSchema = z.object({
  // session_token is handled via cookie
  title: z
    .string()
    .min(3, 'Title must be at least 3 characters')
    .max(200, 'Title cannot exceed 200 characters')
    .describe('Test title'),
});

export const GetMyTestsPayloadSchema = z.object({
  // session_token is handled via cookie
});

export const UpdateTestPayloadSchema = z.object({
  // session_token is handled via cookie
  test_id: z
    .number()
    .int('Test ID must be an integer')
    .describe('Test unique identifier'),
  title: z
    .string()
    .min(3, 'Title must be at least 3 characters')
    .max(200, 'Title cannot exceed 200 characters')
    .describe('New test title'),
});

export const DeleteTestPayloadSchema = z.object({
  // session_token is handled via cookie
  test_id: z
    .number()
    .int('Test ID must be an integer')
    .describe('Test unique identifier'),
});

export const GetTestQuestionsPayloadSchema = z.object({
  // session_token is handled via cookie
  test_id: z
    .number()
    .int('Test ID must be an integer')
    .describe('Test unique identifier'),
});

// Question payload schemas
export const AddQuestionPayloadSchema = z.object({
  // session_token is handled via cookie
  test_id: z
    .number()
    .int('Test ID must be an integer')
    .describe('Test unique identifier'),
  question_text: z
    .string()
    .min(5, 'Question text must be at least 5 characters')
    .max(500, 'Question text cannot exceed 500 characters')
    .describe('Question text'),
  option_a: z
    .string()
    .min(1, 'Option A is required')
    .max(200, 'Option A cannot exceed 200 characters')
    .describe('Answer option A'),
  option_b: z
    .string()
    .min(1, 'Option B is required')
    .max(200, 'Option B cannot exceed 200 characters')
    .describe('Answer option B'),
  option_c: z
    .string()
    .min(1, 'Option C is required')
    .max(200, 'Option C cannot exceed 200 characters')
    .describe('Answer option C'),
  option_d: z
    .string()
    .min(1, 'Option D is required')
    .max(200, 'Option D cannot exceed 200 characters')
    .describe('Answer option D'),
  correct_answer: z
    .number()
    .int('Correct answer must be an integer')
    .min(0, 'Correct answer must be between 0-3 (0=A, 1=B, 2=C, 3=D)')
    .max(3, 'Correct answer must be between 0-3 (0=A, 1=B, 2=C, 3=D)')
    .describe('Correct answer index (0=A, 1=B, 2=C, 3=D)'),
});

export const UpdateQuestionPayloadSchema = z.object({
  // session_token is handled via cookie
  test_id: z
    .number()
    .int('Test ID must be an integer')
    .describe('Test unique identifier'),
  question_id: z
    .number()
    .int('Question ID must be an integer')
    .describe('Question unique identifier'),
  question_text: z
    .string()
    .min(5, 'Question text must be at least 5 characters')
    .max(500, 'Question text cannot exceed 500 characters')
    .describe('Question text'),
  option_a: z
    .string()
    .min(1, 'Option A is required')
    .max(200, 'Option A cannot exceed 200 characters')
    .describe('Answer option A'),
  option_b: z
    .string()
    .min(1, 'Option B is required')
    .max(200, 'Option B cannot exceed 200 characters')
    .describe('Answer option B'),
  option_c: z
    .string()
    .min(1, 'Option C is required')
    .max(200, 'Option C cannot exceed 200 characters')
    .describe('Answer option C'),
  option_d: z
    .string()
    .min(1, 'Option D is required')
    .max(200, 'Option D cannot exceed 200 characters')
    .describe('Answer option D'),
  correct_answer: z
    .number()
    .int('Correct answer must be an integer')
    .min(0, 'Correct answer must be between 0-3 (0=A, 1=B, 2=C, 3=D)')
    .max(3, 'Correct answer must be between 0-3 (0=A, 1=B, 2=C, 3=D)')
    .describe('Correct answer index (0=A, 1=B, 2=C, 3=D)'),
});

export const DeleteQuestionPayloadSchema = z.object({
  // session_token is handled via cookie
  test_id: z
    .number()
    .int('Test ID must be an integer')
    .describe('Test unique identifier'),
  question_id: z
    .number()
    .int('Question ID must be an integer')
    .describe('Question unique identifier'),
});

// === STUDENT PAYLOAD SCHEMAS ===

export const GetAvailableTestsPayloadSchema = z.object({
  // session_token is handled via cookie
});

export const GetTestFormPayloadSchema = z.object({
  // session_token is handled via cookie
  test_id: z
    .number()
    .int('Test ID must be an integer')
    .describe('Test unique identifier'),
});

export const SubmitTestPayloadSchema = z.object({
  // session_token is handled via cookie
  test_id: z
    .number()
    .int('Test ID must be an integer')
    .describe('Test unique identifier'),
  answers: z
    .array(
      z.object({
        question_id: z
          .number()
          .int('Question ID must be an integer')
          .describe('Question unique identifier'),
        selected_answer: z
          .number()
          .int('Selected answer must be an integer')
          .min(0, 'Selected answer must be between 0-3 (0=A, 1=B, 2=C, 3=D)')
          .max(3, 'Selected answer must be between 0-3 (0=A, 1=B, 2=C, 3=D)')
          .describe('Selected answer index (0=A, 1=B, 2=C, 3=D)'),
      }),
    )
    .min(1, 'At least one answer is required')
    .describe('List of answer submissions'),
});

export const GetMySubmissionsPayloadSchema = z.object({
  // session_token is handled via cookie
});

export const GetSubmissionResultPayloadSchema = z.object({
  // session_token is handled via cookie
  submission_id: z
    .number()
    .int('Submission ID must be an integer')
    .describe('Submission unique identifier'),
});

// === ENHANCED PAYLOAD SCHEMAS ===

export const GetTestPreviewPayloadSchema = z.object({
  // session_token is handled via cookie
  test_id: z
    .number()
    .int('Test ID must be an integer')
    .describe('Test unique identifier'),
});

export const BulkAddQuestionsPayloadSchema = z.object({
  // session_token is handled via cookie
  test_id: z
    .number()
    .int('Test ID must be an integer')
    .describe('Test unique identifier'),
  questions: z
    .array(
      z.object({
        question_text: z
          .string()
          .min(5, 'Question text must be at least 5 characters')
          .max(500, 'Question text cannot exceed 500 characters')
          .describe('Question text'),
        options: z
          .array(
            z
              .string()
              .min(1, 'Option cannot be empty')
              .max(200, 'Option cannot exceed 200 characters'),
          )
          .min(2, 'At least 2 options are required')
          .max(4, 'Maximum 4 options allowed')
          .describe('Multiple choice options'),
        correct_answer: z
          .number()
          .int('Correct answer must be an integer')
          .min(0, 'Correct answer index must be non-negative')
          .describe('Index of correct answer (0-based)'),
        explanation: z
          .string()
          .max(500, 'Explanation cannot exceed 500 characters')
          .optional()
          .describe('Explanation for the correct answer'),
        points: z
          .number()
          .int('Points must be an integer')
          .min(1, 'Points must be at least 1')
          .default(1)
          .describe('Points awarded for correct answer'),
        order: z
          .number()
          .int('Order must be an integer')
          .min(1, 'Order must be at least 1')
          .optional()
          .describe('Question order in test (auto-assigned if not provided)'),
      }),
    )
    .min(1, 'At least one question is required')
    .describe('List of questions to add'),
});

// List payloads (for consistency, though most don't have parameters)
export const ListTestsPayloadSchema = GetMyTestsPayloadSchema;
export const ListQuestionsPayloadSchema = GetTestQuestionsPayloadSchema;
export const ListSubmissionsPayloadSchema = GetMySubmissionsPayloadSchema;
