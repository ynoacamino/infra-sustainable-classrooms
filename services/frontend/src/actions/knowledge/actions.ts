'use server';

import { knowledgeService } from '@/services/knowledge/service';
import type {
  AddQuestionPayload,
  CreateTestPayload,
  DeleteQuestionPayload,
  DeleteTestPayload,
  GetSubmissionResultPayload,
  GetTestFormPayload,
  GetTestPayload,
  GetTestQuestionsPayload,
  SubmitTestPayload,
  UpdateQuestionPayload,
  UpdateTestPayload,
} from '@/types/knowledge/payload';
import { cookies } from 'next/headers';

// === TEST ACTIONS (Teacher) ===
export async function createTestAction(payload: CreateTestPayload) {
  const knowledge = await knowledgeService(cookies());
  return knowledge.createTest(payload);
}

export async function updateTestAction(payload: UpdateTestPayload) {
  const knowledge = await knowledgeService(cookies());
  return knowledge.updateTest(payload);
}

export async function deleteTestAction(payload: DeleteTestPayload) {
  const knowledge = await knowledgeService(cookies());
  return knowledge.deleteTest(payload);
}

export async function getMyTestsAction() {
  const knowledge = await knowledgeService(cookies());
  return knowledge.getMyTests();
}

export async function getTestAction(payload: GetTestPayload) {
  const knowledge = await knowledgeService(cookies());
  return knowledge.getTest(payload);
}

// === QUESTION ACTIONS (Teacher) ===
export async function addQuestionAction(payload: AddQuestionPayload) {
  const knowledge = await knowledgeService(cookies());
  return knowledge.addQuestion(payload);
}

export async function updateQuestionAction(payload: UpdateQuestionPayload) {
  const knowledge = await knowledgeService(cookies());
  return knowledge.updateQuestion(payload);
}

export async function deleteQuestionAction(payload: DeleteQuestionPayload) {
  const knowledge = await knowledgeService(cookies());
  return knowledge.deleteQuestion(payload);
}

export async function getTestQuestionsAction(payload: GetTestQuestionsPayload) {
  const knowledge = await knowledgeService(cookies());
  return knowledge.getTestQuestions(payload);
}

// === STUDENT ACTIONS ===
export async function getAvailableTestsAction() {
  const knowledge = await knowledgeService(cookies());
  return knowledge.getAvailableTests();
}

export async function getTestFormAction(payload: GetTestFormPayload) {
  const knowledge = await knowledgeService(cookies());
  return knowledge.getTestForm(payload);
}

export async function submitTestAction(payload: SubmitTestPayload) {
  const knowledge = await knowledgeService(cookies());
  return knowledge.submitTest(payload);
}

export async function getMySubmissionsAction() {
  const knowledge = await knowledgeService(cookies());
  return knowledge.getMySubmissions();
}

export async function getSubmissionResultAction(
  payload: GetSubmissionResultPayload,
) {
  const knowledge = await knowledgeService(cookies());
  return knowledge.getSubmissionResult(payload);
}
