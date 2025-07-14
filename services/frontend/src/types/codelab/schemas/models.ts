import z from 'zod';

export const SimpleResponseSchema = z.object({
  success: z.boolean().describe('Operation success status'),
  message: z.string().describe('Response message'),
});

export const ExerciseSchema = z.object({
  id: z
    .number()
    .int('Exercise ID must be an integer')
    .describe('Unique identifier for the exercise'),
  title: z
    .string()
    .min(1, 'Title is required')
    .max(200, 'Title must be at most 200 characters')
    .describe('Exercise title'),
  description: z
    .string()
    .min(1, 'Description is required')
    .describe('Exercise description'),
  initial_code: z
    .string()
    .min(1, 'Initial code is required')
    .describe('Initial code template'),
  solution: z
    .string()
    .min(1, 'Solution is required')
    .describe('Exercise solution'),
  difficulty: z
    .enum(['easy', 'medium', 'hard'])
    .describe('Exercise difficulty level'),
  created_by: z
    .number()
    .int('Creator user ID must be an integer')
    .describe('ID of user who created the exercise'),
  created_at: z
    .number()
    .int('Creation timestamp must be an integer')
    .describe(
      'Timestamp when the exercise was created (milliseconds since epoch)',
    ),
  updated_at: z
    .number()
    .int('Last update timestamp must be an integer')
    .describe(
      'Timestamp when the exercise was last updated (milliseconds since epoch)',
    ),
});

export const TestSchema = z.object({
  id: z
    .number()
    .int('Test ID must be an integer')
    .describe('Unique identifier for the test'),
  input: z.string().min(1, 'Input is required').describe('Test input'),
  output: z.string().min(1, 'Output is required').describe('Expected output'),
  public: z.boolean().describe('Whether test is visible to students'),
  exercise_id: z
    .number()
    .int('Exercise ID must be an integer')
    .describe('Associated exercise ID'),
  created_at: z
    .number()
    .int('Creation timestamp must be an integer')
    .describe('Timestamp when the test was created (milliseconds since epoch)'),
  updated_at: z
    .number()
    .int('Last update timestamp must be an integer')
    .describe(
      'Timestamp when the test was last updated (milliseconds since epoch)',
    ),
});

export const AnswerSchema = z.object({
  id: z
    .number()
    .int('Answer ID must be an integer')
    .describe('Unique identifier for the answer'),
  exercise_id: z
    .number()
    .int('Exercise ID must be an integer')
    .describe('Associated exercise ID'),
  user_id: z
    .number()
    .int('User ID must be an integer')
    .describe('Student user ID'),
  completed: z.boolean().describe('Whether the exercise is completed'),
  created_at: z
    .number()
    .int('Creation timestamp must be an integer')
    .describe(
      'Timestamp when the answer was created (milliseconds since epoch)',
    ),
  updated_at: z
    .number()
    .int('Last update timestamp must be an integer')
    .describe(
      'Timestamp when the answer was last updated (milliseconds since epoch)',
    ),
});

export const AttemptSchema = z.object({
  id: z
    .number()
    .int('Attempt ID must be an integer')
    .describe('Unique identifier for the attempt'),
  answer_id: z
    .number()
    .int('Answer ID must be an integer')
    .describe('Associated answer ID'),
  code: z.string().min(1, 'Code is required').describe('Submitted code'),
  success: z.boolean().describe('Whether the attempt was successful'),
  created_at: z
    .number()
    .int('Creation timestamp must be an integer')
    .describe(
      'Timestamp when the attempt was created (milliseconds since epoch)',
    ),
});

export const ExerciseForStudentsSchema = z.object({
  id: z
    .number()
    .int('Exercise ID must be an integer')
    .describe('Unique identifier for the exercise'),
  title: z
    .string()
    .min(1, 'Title is required')
    .max(200, 'Title must be at most 200 characters')
    .describe('Exercise title'),
  description: z
    .string()
    .min(1, 'Description is required')
    .describe('Exercise description'),
  initial_code: z
    .string()
    .min(1, 'Initial code is required')
    .describe('Initial code template'),
  difficulty: z
    .enum(['easy', 'medium', 'hard'])
    .describe('Exercise difficulty level'),
  tests: z.array(TestSchema).describe('Associated tests for the exercise'),
  attempts: z
    .array(AttemptSchema)
    .describe('Associated attempts for the exercise'),
  answer: AnswerSchema.describe(
    "Student's answer/participation in the exercise",
  ),
  created_by: z
    .number()
    .int('Creator user ID must be an integer')
    .describe('ID of user who created the exercise'),
  created_at: z
    .number()
    .int('Creation timestamp must be an integer')
    .describe(
      'Timestamp when the exercise was created (milliseconds since epoch)',
    ),
  updated_at: z
    .number()
    .int('Last update timestamp must be an integer')
    .describe(
      'Timestamp when the exercise was last updated (milliseconds since epoch)',
    ),
});

export const ExerciseForStudentsListViewSchema = z.object({
  id: z
    .number()
    .int('Exercise ID must be an integer')
    .describe('Unique identifier for the exercise'),
  title: z
    .string()
    .min(1, 'Title is required')
    .max(200, 'Title must be at most 200 characters')
    .describe('Exercise title'),
  description: z
    .string()
    .min(1, 'Description is required')
    .describe('Exercise description'),
  difficulty: z
    .enum(['easy', 'medium', 'hard'])
    .describe('Exercise difficulty level'),
  completed: z
    .boolean()
    .optional()
    .describe('Whether the exercise is completed by the student'),
  created_by: z
    .number()
    .int('Creator user ID must be an integer')
    .describe('ID of user who created the exercise'),
  created_at: z
    .number()
    .int('Creation timestamp must be an integer')
    .describe(
      'Timestamp when the exercise was created (milliseconds since epoch)',
    ),
  updated_at: z
    .number()
    .int('Last update timestamp must be an integer')
    .describe(
      'Timestamp when the exercise was last updated (milliseconds since epoch)',
    ),
});
