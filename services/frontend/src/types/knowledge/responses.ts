import type {
  Test,
  Question,
  QuestionForm,
  Submission,
  SubmissionResult,
  SimpleResponse,
  TestsResponse,
  QuestionsResponse,
  FormResponse,
  SubmissionsResponse,
  SubmitResponse,
  UserAccessResponse,
  TestPreviewResponse,
  BulkQuestionResponse,
} from '@/types/knowledge/models';

// === TEACHER RESPONSE TYPES ===

// Test response types
export type CreateTestResponse = SimpleResponse;
export type GetMyTestsResponse = TestsResponse;
export type UpdateTestResponse = SimpleResponse;
export type DeleteTestResponse = SimpleResponse;
export type GetTestQuestionsResponse = QuestionsResponse;

// Question response types
export type AddQuestionResponse = SimpleResponse;
export type UpdateQuestionResponse = SimpleResponse;
export type DeleteQuestionResponse = SimpleResponse;

// === STUDENT RESPONSE TYPES ===
export type GetAvailableTestsResponse = TestsResponse;
export type GetTestFormResponse = FormResponse;
export type SubmitTestResponse = SubmitResponse;
export type GetMySubmissionsResponse = SubmissionsResponse;
export type GetSubmissionResultResponse = SubmissionResult;

// === AUTH RESPONSE TYPES ===
export type UserAccessResponseType = UserAccessResponse;

// === ENHANCED RESPONSE TYPES ===
export type GetTestPreviewResponseType = TestPreviewResponse;
export type BulkAddQuestionsResponse = BulkQuestionResponse;

// === LIST RESPONSE TYPES ===
export type ListTestsResponse = TestsResponse;
export type ListQuestionsResponse = QuestionsResponse;
export type ListSubmissionsResponse = SubmissionsResponse;

// === INDIVIDUAL ENTITY RESPONSE TYPES ===
export type GetTestResponse = Test;
export type GetQuestionResponse = Question;
export type GetSubmissionResponse = Submission;
