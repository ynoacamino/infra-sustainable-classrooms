import z from 'zod';
import { ExerciseSchema, TestSchema } from '@/types/codelab/schemas/models';

// Exercise payload schemas
export const CreateExercisePayloadSchema = ExerciseSchema.omit({
  id: true,
  created_at: true,
  updated_at: true,
});

export const GetExercisePayloadSchema = z.object({
  id: z.number().int('Exercise ID must be an integer'),
});


export const UpdateExercisePayloadSchema = z.object({
  id: z.number().int('Exercise ID must be an integer'),
  exercise: ExerciseSchema.omit({
    id: true,
    created_at: true,
    updated_at: true,
    created_by: true,
  }),
});

export const DeleteExercisePayloadSchema = z.object({
  id: z.number().int('Exercise ID must be an integer'),
});

// Test payload schemas
export const CreateTestPayloadSchema = TestSchema.omit({
  id: true,
  created_at: true,
  updated_at: true,
});

export const GetTestsByExercisePayloadSchema = z.object({
  exercise_id: z.number().int('Exercise ID must be an integer'),
});

export const UpdateTestPayloadSchema = z.object({
  id: z.number().int('Test ID must be an integer'),
  test: TestSchema.omit({
    id: true,
    created_at: true,
    updated_at: true,
    exercise_id: true,
  }),
});

export const DeleteTestPayloadSchema = z.object({
  id: z.number().int('Test ID must be an integer'),
});

// Student exercise payload schemas
export const GetExerciseForStudentPayloadSchema = z.object({
  id: z.number().int('Exercise ID must be an integer'),
});

// Attempt payload schemas
export const CreateAttemptPayloadSchema = z.object({
  exercise_id: z.number().int('Exercise ID must be an integer'),
  code: z.string().min(1, 'Code is required').describe('Submitted code'),
  success: z.boolean().describe('Whether the attempt was successful'),
});

export const GetAttemptsByUserAndExercisePayloadSchema = z.object({
  user_id: z.number().int('User ID must be an integer'),
  exercise_id: z.number().int('Exercise ID must be an integer'),
});

// Answer payload schemas
export const GetAnswerByUserAndExercisePayloadSchema = z.object({
  user_id: z.number().int('User ID must be an integer'),
  exercise_id: z.number().int('Exercise ID must be an integer'),
});
