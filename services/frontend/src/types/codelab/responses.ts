import type {
  Exercise,
  Test,
  Answer,
  Attempt,
  ExerciseForStudents,
  ExerciseForStudentsListView,
  SimpleResponse,
} from '@/types/codelab/models';

// Exercise response types
export type CreateExerciseResponse = SimpleResponse;
export type GetExerciseResponse = Exercise;
export type ListExercisesResponse = Exercise[];
export type UpdateExerciseResponse = SimpleResponse;
export type DeleteExerciseResponse = SimpleResponse;

// Test response types
export type CreateTestResponse = SimpleResponse;
export type GetTestsByExerciseResponse = Test[];
export type UpdateTestResponse = SimpleResponse;
export type DeleteTestResponse = SimpleResponse;

// Student exercise response types
export type GetExerciseForStudentResponse = ExerciseForStudents;
export type ListExercisesForStudentsResponse = ExerciseForStudentsListView[];

// Attempt response types
export type CreateAttemptResponse = SimpleResponse;
export type GetAttemptsByUserAndExerciseResponse = Attempt[];

// Answer response types
export type GetAnswerByUserAndExerciseResponse = Answer;
