'use server';

import { codelabService } from '@/services/codelab/service';
import type {
  CreateExercisePayloadSchema,
  UpdateExercisePayloadSchema,
  DeleteExercisePayloadSchema,
  CreateTestPayloadSchema,
  UpdateTestPayloadSchema,
  DeleteTestPayloadSchema,
  CreateAttemptPayloadSchema,
} from '@/types/codelab/schemas/payload';
import { revalidatePath } from 'next/cache';
import { cookies } from 'next/headers';
import type z from 'zod';

// === EXERCISE ACTIONS ===
export async function getExercise(exerciseId: number) {
  const codelab = await codelabService(cookies());
  return codelab.getExercise({
    id: exerciseId,
  });
}

export async function createExerciseAction(
  payload: z.infer<typeof CreateExercisePayloadSchema>,
) {
  const codelab = await codelabService(cookies());
  return codelab.createExercise(payload);
}

export async function updateExerciseAction(
  exerciseId: number,
  exercise: z.infer<typeof UpdateExercisePayloadSchema>['exercise'],
) {
  const codelab = await codelabService(cookies());
  return codelab.updateExercise({
    id: exerciseId,
    exercise,
  });
}

export async function deleteExerciseAction(exerciseId: number) {
  const codelab = await codelabService(cookies());
  return codelab.deleteExercise({
    id: exerciseId,
  });
}

// === TEST ACTIONS ===
export async function getTestsByExercise(exerciseId: number) {
  const codelab = await codelabService(cookies());
  return codelab.getTestsByExercise({
    exercise_id: exerciseId,
  });
}

export async function createTestAction(
  payload: z.infer<typeof CreateTestPayloadSchema>,
) {
  const codelab = await codelabService(cookies());
  return codelab.createTest(payload);
}

export async function updateTestAction(
  testId: number,
  test: z.infer<typeof UpdateTestPayloadSchema>['test'],
) {
  const codelab = await codelabService(cookies());
  return codelab.updateTest({
    id: testId,
    test,
  });
}

export async function deleteTestAction(testId: number) {
  const codelab = await codelabService(cookies());
  return codelab.deleteTest({
    id: testId,
  });
}

// === ATTEMPT ACTIONS ===
export async function createAttemptAction(
  payload: z.infer<typeof CreateAttemptPayloadSchema>,
) {
  const codelab = await codelabService(cookies());

  revalidatePath(`/dashboard/codelab/exercises/${payload.exercise_id}`);

  return codelab.createAttempt(payload);
}
