'use server';

import { codelabService } from '@/services/codelab/service';
import type {
  CreateAttemptPayload,
  CreateExercisePayload,
  CreateTestPayload,
  DeleteExercisePayload,
  DeleteTestPayload,
  GetExercisePayload,
  GetTestsByExercisePayload,
  UpdateExercisePayload,
  UpdateTestPayload,
} from '@/types/codelab/payload';
import { cookies } from 'next/headers';

// === EXERCISE ACTIONS ===
export async function getExercise(payload: GetExercisePayload) {
  const codelab = await codelabService(cookies());
  return codelab.getExercise(payload);
}

export async function createExerciseAction(payload: CreateExercisePayload) {
  const codelab = await codelabService(cookies());
  return codelab.createExercise(payload);
}

export async function updateExerciseAction(payload: UpdateExercisePayload) {
  const codelab = await codelabService(cookies());
  return codelab.updateExercise(payload);
}

export async function deleteExerciseAction(payload: DeleteExercisePayload) {
  const codelab = await codelabService(cookies());
  return codelab.deleteExercise(payload);
}

// === TEST ACTIONS ===
export async function getTestsByExercise(payload: GetTestsByExercisePayload) {
  const codelab = await codelabService(cookies());
  return codelab.getTestsByExercise(payload);
}

export async function createTestAction(payload: CreateTestPayload) {
  const codelab = await codelabService(cookies());
  return codelab.createTest(payload);
}

export async function updateTestAction(payload: UpdateTestPayload) {
  const codelab = await codelabService(cookies());
  return codelab.updateTest(payload);
}

export async function deleteTestAction(payload: DeleteTestPayload) {
  const codelab = await codelabService(cookies());
  return codelab.deleteTest(payload);
}

// === ATTEMPT ACTIONS ===
export async function createAttemptAction(payload: CreateAttemptPayload) {
  const codelab = await codelabService(cookies());

  // revalidatePath(`/dashboard/codelab/exercises/${payload.exercise_id}`);

  return codelab.createAttempt(payload);
}
