import z from 'zod';
import {
  AttemptSchema,
  ExerciseForStudentsSchema,
  ExerciseSchema,
  TestSchema,
} from '@/types/codelab/schemas/models';

// Exercise payload schemas
export const CreateExercisePayloadSchema = ExerciseSchema.omit({
  id: true,
  created_at: true,
  updated_at: true,
});

export const GetExercisePayloadSchema = ExerciseSchema.pick({
  id: true,
});

export const UpdateExercisePayloadSchema = z.object({
  id: ExerciseSchema.shape.id,
  exercise: ExerciseSchema.omit({
    id: true,
    created_at: true,
    updated_at: true,
    created_by: true,
  }),
});

export const DeleteExercisePayloadSchema = ExerciseSchema.pick({
  id: true,
});

// Test payload schemas
export const CreateTestPayloadSchema = TestSchema.omit({
  id: true,
  created_at: true,
  updated_at: true,
});

export const GetTestsByExercisePayloadSchema = TestSchema.pick({
  exercise_id: true,
});

export const UpdateTestPayloadSchema = z.object({
  id: TestSchema.shape.id,
  test: TestSchema.omit({
    id: true,
    created_at: true,
    updated_at: true,
    exercise_id: true,
  }),
});

export const DeleteTestPayloadSchema = TestSchema.pick({
  id: true,
});

// Student exercise payload schemas
export const GetExerciseForStudentPayloadSchema =
  ExerciseForStudentsSchema.pick({
    id: true,
  });

// Attempt payload schemas
export const CreateAttemptPayloadSchema = AttemptSchema.omit({
  id: true,
  answer_id: true,
  created_at: true,
}).extend({
  exercise_id: ExerciseSchema.shape.id,
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
