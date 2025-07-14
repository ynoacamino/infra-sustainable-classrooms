import type {
  ExerciseSchema,
  TestSchema,
  AnswerSchema,
  AttemptSchema,
  ExerciseForStudentsSchema,
  ExerciseForStudentsListViewSchema,
  SimpleResponseSchema,
} from '@/types/codelab/schemas/models';
import type z from 'zod';

export type Exercise = z.infer<typeof ExerciseSchema>;

export type Test = z.infer<typeof TestSchema>;

export type Answer = z.infer<typeof AnswerSchema>;

export type Attempt = z.infer<typeof AttemptSchema>;

export type ExerciseForStudents = z.infer<typeof ExerciseForStudentsSchema>;

export type ExerciseForStudentsListView = z.infer<
  typeof ExerciseForStudentsListViewSchema
>;

export type SimpleResponse = z.infer<typeof SimpleResponseSchema>;
