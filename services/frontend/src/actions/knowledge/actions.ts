'use server';

import { knowledgeService } from '@/services/knowledge/service';
import type {
  CreateTestPayloadSchema,
  UpdateTestPayloadSchema,
  DeleteTestPayloadSchema,
  AddQuestionPayloadSchema,
  UpdateQuestionPayloadSchema,
  DeleteQuestionPayloadSchema,
  SubmitTestPayloadSchema,
} from '@/types/knowledge/schemas/payload';
import { cookies } from 'next/headers';
import type z from 'zod';

// === TEST ACTIONS (Teacher) ===
export async function createTestAction(
  payload: z.infer<typeof CreateTestPayloadSchema>,
) {
  const knowledge = await knowledgeService(cookies());
  return knowledge.createTest(payload);
}

export async function updateTestAction(
  payload: z.infer<typeof UpdateTestPayloadSchema>,
) {
  const knowledge = await knowledgeService(cookies());
  return knowledge.updateTest(payload);
}

export async function deleteTestAction(testId: number) {
  const knowledge = await knowledgeService(cookies());
  return knowledge.deleteTest({ test_id: testId });
}

export async function getMyTestsAction() {
  const knowledge = await knowledgeService(cookies());
  return knowledge.getMyTests();
}

export async function getTestAction(testId: number) {
  const knowledge = await knowledgeService(cookies());
  return knowledge.getTest(testId);
}

// === QUESTION ACTIONS (Teacher) ===
export async function addQuestionAction(
  payload: z.infer<typeof AddQuestionPayloadSchema>,
) {
  const knowledge = await knowledgeService(cookies());
  return knowledge.addQuestion(payload);
}

export async function updateQuestionAction(
  payload: z.infer<typeof UpdateQuestionPayloadSchema>,
) {
  const knowledge = await knowledgeService(cookies());
  return knowledge.updateQuestion(payload);
}

export async function deleteQuestionAction(testId: number, questionId: number) {
  const knowledge = await knowledgeService(cookies());
  return knowledge.deleteQuestion({ test_id: testId, question_id: questionId });
}

export async function getTestQuestionsAction(testId: number) {
  const knowledge = await knowledgeService(cookies());
  return knowledge.getTestQuestions({ test_id: testId });
}

// === STUDENT ACTIONS ===
export async function getAvailableTestsAction() {
  const knowledge = await knowledgeService(cookies());
  return knowledge.getAvailableTests();
}

export async function getTestFormAction(testId: number) {
  const knowledge = await knowledgeService(cookies());
  return knowledge.getTestForm({ test_id: testId });
}

export async function submitTestAction(
  payload: z.infer<typeof SubmitTestPayloadSchema>,
) {
  const knowledge = await knowledgeService(cookies());
  return knowledge.submitTest(payload);
}

export async function getMySubmissionsAction() {
  const knowledge = await knowledgeService(cookies());
  return knowledge.getMySubmissions();
}

export async function getSubmissionResultAction(submissionId: number) {
  const knowledge = await knowledgeService(cookies());
  return knowledge.getSubmissionResult({ submission_id: submissionId });
}
