import type {
  CreateExercisePayloadSchema,
  GetExercisePayloadSchema,
  UpdateExercisePayloadSchema,
  DeleteExercisePayloadSchema,
  CreateTestPayloadSchema,
  GetTestsByExercisePayloadSchema,
  UpdateTestPayloadSchema,
  DeleteTestPayloadSchema,
  GetExerciseForStudentPayloadSchema,
  CreateAttemptPayloadSchema,
  GetAttemptsByUserAndExercisePayloadSchema,
  GetAnswerByUserAndExercisePayloadSchema,
} from '@/types/codelab/schemas/payload';
import type z from 'zod';

// Exercise payload types
export type CreateExercisePayload = z.infer<typeof CreateExercisePayloadSchema>;
export type GetExercisePayload = z.infer<typeof GetExercisePayloadSchema>;
export type UpdateExercisePayload = z.infer<typeof UpdateExercisePayloadSchema>;
export type DeleteExercisePayload = z.infer<typeof DeleteExercisePayloadSchema>;

// Test payload types
export type CreateTestPayload = z.infer<typeof CreateTestPayloadSchema>;
export type GetTestsByExercisePayload = z.infer<
  typeof GetTestsByExercisePayloadSchema
>;
export type UpdateTestPayload = z.infer<typeof UpdateTestPayloadSchema>;
export type DeleteTestPayload = z.infer<typeof DeleteTestPayloadSchema>;

// Student exercise payload types
export type GetExerciseForStudentPayload = z.infer<
  typeof GetExerciseForStudentPayloadSchema
>;

// Attempt payload types
export type CreateAttemptPayload = z.infer<typeof CreateAttemptPayloadSchema>;
export type GetAttemptsByUserAndExercisePayload = z.infer<
  typeof GetAttemptsByUserAndExercisePayloadSchema
>;

// Answer payload types
export type GetAnswerByUserAndExercisePayload = z.infer<
  typeof GetAnswerByUserAndExercisePayloadSchema
>;
