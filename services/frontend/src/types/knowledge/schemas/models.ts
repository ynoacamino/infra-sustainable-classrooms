import z from 'zod';

export const TestSchema = z.object({
  id: z
    .number()
    .int('Test ID must be an integer')
    .describe('Unique identifier for the test'),
  title: z.string().min(1, 'Title is required').describe('Title of the test'),
  created_by: z
    .number()
    .int('Creator ID must be an integer')
    .describe('ID of the user who created the test'),
  created_at: z
    .number()
    .int('Creation timestamp must be an integer')
    .describe('Timestamp when the test was created (milliseconds since epoch)'),
  question_count: z
    .number()
    .int('Question count must be an integer')
    .min(0, 'Question count cannot be negative')
    .optional()
    .describe('Number of questions in the test'),
});

export const QuestionSchema = z.object({
  id: z
    .number()
    .int('Question ID must be an integer')
    .describe('Unique identifier for the question'),
  test_id: z
    .number()
    .int('Test ID must be an integer')
    .describe('ID of the parent test'),
  question_text: z
    .string()
    .min(1, 'Question text is required')
    .describe('Text of the question'),
  option_a: z
    .string()
    .min(1, 'Option A is required')
    .describe('Answer option A'),
  option_b: z
    .string()
    .min(1, 'Option B is required')
    .describe('Answer option B'),
  option_c: z
    .string()
    .min(1, 'Option C is required')
    .describe('Answer option C'),
  option_d: z
    .string()
    .min(1, 'Option D is required')
    .describe('Answer option D'),
  correct_answer: z
    .number()
    .int('Correct answer must be an integer')
    .min(0, 'Correct answer must be between 0-3')
    .max(3, 'Correct answer must be between 0-3')
    .describe('Correct answer index (0=A, 1=B, 2=C, 3=D)'),
  question_order: z
    .number()
    .int('Question order must be an integer')
    .min(1, 'Question order must be at least 1')
    .describe('Order of the question in the test'),
});

export const QuestionFormSchema = z.object({
  id: z
    .number()
    .int('Question ID must be an integer')
    .describe('Unique identifier for the question'),
  question_text: z
    .string()
    .min(1, 'Question text is required')
    .describe('Text of the question'),
  option_a: z
    .string()
    .min(1, 'Option A is required')
    .describe('Answer option A'),
  option_b: z
    .string()
    .min(1, 'Option B is required')
    .describe('Answer option B'),
  option_c: z
    .string()
    .min(1, 'Option C is required')
    .describe('Answer option C'),
  option_d: z
    .string()
    .min(1, 'Option D is required')
    .describe('Answer option D'),
  question_order: z
    .number()
    .int('Question order must be an integer')
    .min(1, 'Question order must be at least 1')
    .describe('Order of the question in the test'),
});

export const AnswerSchema = z.object({
  question_id: z
    .number()
    .int('Question ID must be an integer')
    .describe('ID of the question being answered'),
  selected_answer: z
    .number()
    .int('Selected answer must be an integer')
    .min(0, 'Selected answer must be between 0-3')
    .max(3, 'Selected answer must be between 0-3')
    .describe('Selected answer index (0=A, 1=B, 2=C, 3=D)'),
});

export const SubmissionSchema = z.object({
  id: z
    .number()
    .int('Submission ID must be an integer')
    .describe('Unique identifier for the submission'),
  test_id: z
    .number()
    .int('Test ID must be an integer')
    .describe('ID of the test that was submitted'),
  test_title: z
    .string()
    .min(1, 'Test title is required')
    .describe('Title of the test'),
  score: z
    .number()
    .min(0, 'Score cannot be negative')
    .max(100, 'Score cannot exceed 100')
    .describe('Score percentage (0-100)'),
  submitted_at: z
    .number()
    .int('Submission timestamp must be an integer')
    .describe(
      'Timestamp when the test was submitted (milliseconds since epoch)',
    ),
});
