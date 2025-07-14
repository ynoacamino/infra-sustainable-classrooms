import { SessionInterceptor } from '@/services/auth/interceptor';
import { Service } from '@/services/shared/service';
import type { Test, Question, Submission } from '@/types/knowledge/models';
import type {
  CreateTestPayload,
  UpdateTestPayload,
  DeleteTestPayload,
  GetTestQuestionsPayload,
  AddQuestionPayload,
  UpdateQuestionPayload,
  DeleteQuestionPayload,
  GetTestFormPayload,
  SubmitTestPayload,
  GetSubmissionResultPayload,
  GetTestPayload,
  GetQuestionPayload,
  GetSubmissionPayload,
} from '@/types/knowledge/payload';
import {
  CreateTestPayloadSchema,
  UpdateTestPayloadSchema,
  GetTestQuestionsPayloadSchema,
  AddQuestionPayloadSchema,
  UpdateQuestionPayloadSchema,
  GetTestFormPayloadSchema,
  SubmitTestPayloadSchema,
  GetSubmissionResultPayloadSchema,
  GetTestPayloadSchema,
  DeleteTestPayloadSchema,
  GetQuestionPayloadSchema,
  DeleteQuestionPayloadSchema,
  GetSubmissionPayloadSchema,
} from '@/types/knowledge/schemas/payload';
import type { ReadonlyRequestCookies } from 'next/dist/server/web/spec-extension/adapters/request-cookies';
import type { AsyncResult } from '@/types/shared/services/result';
import type { SimpleResponse } from '@/services/shared/response';
import type {
  GetSubmissionResultResponse,
  GetTestFormResponse,
  SubmitTestResponse,
} from '@/types/knowledge/responses';

class KnowledgeService extends Service {
  constructor() {
    super('knowledge');
  }

  async createTest(payload: CreateTestPayload): AsyncResult<SimpleResponse> {
    return this.post<SimpleResponse>({
      endpoint: 'tests',
      payload: {
        schema: CreateTestPayloadSchema,
        data: payload,
      },
    });
  }

  /**
   * Get my created tests
   */
  async getMyTests(): AsyncResult<GetMyTestsResponse> {
    return this.get<GetMyTestsResponse>({
      endpoint: 'tests/my',
    });
  }

  async updateTest(payload: UpdateTestPayload): AsyncResult<SimpleResponse> {
    return this.put<SimpleResponse>({
      endpoint: ['tests', payload.id],
      payload: {
        schema: UpdateTestPayloadSchema,
        data: payload,
      },
    });
  }

  /**
   * Delete a test
   */
  async deleteTest(
    payload: DeleteTestPayload,
  ): AsyncResult<DeleteTestResponse> {
    return this.delete<DeleteTestResponse>({
      endpoint: ['tests', payload.test_id.toString()]
    });
  }

  async getTestQuestions(
    payload: GetTestQuestionsPayload,
  ): AsyncResult<GetTestQuestionsResponse> {
    return this.get<GetTestQuestionsResponse>({
      endpoint: ['tests', payload.test_id, 'questions'],
    });
  }

  async addQuestion(payload: AddQuestionPayload): AsyncResult<SimpleResponse> {
    return this.post<SimpleResponse>({
      endpoint: ['tests', payload.test_id, 'questions'],
      payload: {
        schema: AddQuestionPayloadSchema,
        data: payload,
      },
    });
  }

  async getQuestion(payload: GetQuestionPayload): AsyncResult<Question> {
    return this.get<Question>({
      endpoint: ['tests', payload.test_id, 'questions', payload.id],
      payload: {
        schema: GetQuestionPayloadSchema,
        data: payload,
      },
    });
  }

  async updateQuestion(
    payload: UpdateQuestionPayload,
  ): AsyncResult<SimpleResponse> {
    return this.put<SimpleResponse>({
      endpoint: ['tests', payload.test_id, 'questions', payload.id],
      payload: {
        schema: UpdateQuestionPayloadSchema,
        data: payload,
      },
    });
  }

  async deleteQuestion(
    payload: DeleteQuestionPayload,
  ): AsyncResult<DeleteQuestionResponse> {
    return this.delete<DeleteQuestionResponse>({
      endpoint: ['tests', payload.test_id.toString(), 'questions', payload.question_id.toString()]
    });
  }

  // === STUDENT METHODS ===
  /**
   * Get available tests for students
   */
  async getAvailableTests(): AsyncResult<GetAvailableTestsResponse> {
    return this.get<GetAvailableTestsResponse>({
      endpoint: 'tests/available',
    });
  }

  async getTestForm(
    payload: GetTestFormPayload,
  ): AsyncResult<GetTestFormResponse> {
    return this.get<GetTestFormResponse>({
      endpoint: ['tests', payload.test_id, 'form'],
    });
  }

  async submitTest(
    payload: SubmitTestPayload,
  ): AsyncResult<SubmitTestResponse> {
    return this.post<SubmitTestResponse>({
      endpoint: ['tests', payload.id, 'submit'],
      payload: {
        schema: SubmitTestPayloadSchema,
        data: payload,
      },
    });
  }

  async getMySubmissions(): AsyncResult<Submission[]> {
    return this.get<Submission[]>({
      endpoint: 'submissions/my',
    });
  }

  async getSubmissionResult(
    payload: GetSubmissionResultPayload,
  ): AsyncResult<GetSubmissionResultResponse> {
    return this.get<GetSubmissionResultResponse>({
      endpoint: ['submissions', payload.submission_id, 'result'],
    });
  }
}

export const knowledgeService = async (
  cookies: Promise<ReadonlyRequestCookies>,
) => {
  const service = new KnowledgeService();
  service.addInterceptor(new SessionInterceptor(await cookies));
  return service;
};
