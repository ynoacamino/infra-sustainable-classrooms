import z from 'zod';

export const SimpleResponseSchema = z.object({
  success: z.boolean().describe('Operation success status'),
  message: z.string().describe('Response message'),
});

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

export const QuestionResultSchema = z.object({
  question: QuestionSchema,
  selected_answer: z
    .number()
    .int('Selected answer must be an integer')
    .min(0, 'Selected answer must be between 0-3')
    .max(3, 'Selected answer must be between 0-3')
    .describe('Answer selected by the user'),
  is_correct: z.boolean().describe('Whether the selected answer was correct'),
});

export const SubmissionResultSchema = z.object({
  submission: SubmissionSchema,
  questions: z
    .array(QuestionResultSchema)
    .describe('Detailed results for each question'),
});

// Response schemas
export const TestsResponseSchema = z.object({
  tests: z.array(TestSchema).describe('List of tests'),
});

export const QuestionsResponseSchema = z.object({
  questions: z.array(QuestionSchema).describe('List of questions'),
});

export const FormResponseSchema = z.object({
  test: TestSchema,
  questions: z
    .array(QuestionFormSchema)
    .describe('List of questions for the form (without correct answers)'),
});

export const SubmissionsResponseSchema = z.object({
  submissions: z.array(SubmissionSchema).describe('List of submissions'),
});

export const SubmitResponseSchema = z.object({
  success: z.boolean().describe('Operation success status'),
  message: z.string().describe('Response message'),
  submission_id: z
    .number()
    .int('Submission ID must be an integer')
    .describe('ID of the created submission'),
  score: z
    .number()
    .min(0, 'Score cannot be negative')
    .max(100, 'Score cannot exceed 100')
    .describe('Score percentage (0-100)'),
});

// User access response schema
export const UserAccessResponseSchema = z.object({
  user_id: z
    .number()
    .int('User ID must be an integer')
    .describe('User identifier'),
  username: z.string().min(1, 'Username is required').describe('Username'),
  email: z.string().email('Invalid email format').describe('User email'),
  permissions: z.array(z.string()).describe('User permissions'),
  roles: z.array(z.string()).describe('User roles'),
  is_active: z.boolean().describe('Whether user is active'),
  last_login: z.string().optional().describe('Last login timestamp'),
  session_valid: z.boolean().describe('Whether session is valid'),
});

// Enhanced types
export const TestPreviewResponseSchema = z.object({
  test_id: z
    .number()
    .int('Test ID must be an integer')
    .describe('Test identifier'),
  title: z.string().min(1, 'Title is required').describe('Test title'),
  description: z.string().optional().describe('Test description'),
  difficulty_level: z
    .string()
    .min(1, 'Difficulty level is required')
    .describe('Difficulty level'),
  duration_minutes: z
    .number()
    .int('Duration must be an integer')
    .min(1, 'Duration must be at least 1 minute')
    .describe('Duration in minutes'),
  total_questions: z
    .number()
    .int('Total questions must be an integer')
    .min(0, 'Total questions cannot be negative')
    .describe('Total number of questions'),
  expires_at: z
    .number()
    .int('Expiration timestamp must be an integer')
    .optional()
    .describe('Test expiration timestamp (milliseconds since epoch)'),
  created_by: z
    .number()
    .int('Creator ID must be an integer')
    .describe('Creator ID'),
});

export const BulkQuestionInputSchema = z.object({
  question_text: z
    .string()
    .min(1, 'Question text is required')
    .describe('Question text'),
  options: z
    .array(z.string().min(1, 'Option cannot be empty'))
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
    .default(1)
    .describe('Question order in test'),
});

export const BulkQuestionResponseSchema = z.object({
  test_id: z
    .number()
    .int('Test ID must be an integer')
    .describe('Test identifier'),
  questions_added: z
    .number()
    .int('Questions added must be an integer')
    .min(0, 'Questions added cannot be negative')
    .describe('Number of questions successfully added'),
  questions_failed: z
    .number()
    .int('Questions failed must be an integer')
    .min(0, 'Questions failed cannot be negative')
    .describe('Number of questions that failed to add'),
  question_ids: z
    .array(z.number().int())
    .describe('List of created question IDs'),
  errors: z
    .array(z.string())
    .optional()
    .describe('List of errors for failed questions'),
  total_questions: z
    .number()
    .int('Total questions must be an integer')
    .min(0, 'Total questions cannot be negative')
    .describe('Total questions in test after bulk add'),
});
