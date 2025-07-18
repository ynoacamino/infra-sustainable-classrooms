import { SessionInterceptor } from '@/services/auth/interceptor';
import { Service } from '@/services/shared/service';
import type {
  CreateTestPayload,
  UpdateTestPayload,
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
  QuestionResponse,
  QuestionsResponse,
  SubmissionResponse,
  SubmissionsResponse,
  SubmitTestResponse,
  TestResponse,
  TestsResponse,
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

  async getMyTests(): AsyncResult<TestsResponse> {
    return this.get<TestsResponse>({
      endpoint: ['tests', 'my'],
    });
  }

  async getTest(payload: GetTestPayload): AsyncResult<TestResponse> {
    return this.get<TestResponse>({
      endpoint: ['tests', payload.id],
      payload: {
        schema: GetTestPayloadSchema,
        data: payload,
      },
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

  async deleteTest(payload: GetTestPayload): AsyncResult<SimpleResponse> {
    return this.delete<SimpleResponse>({
      endpoint: ['tests', payload.id],
      payload: {
        schema: GetTestPayloadSchema,
        data: payload,
      },
    });
  }

  async getTestQuestions(
    payload: GetTestQuestionsPayload,
  ): AsyncResult<QuestionsResponse> {
    return this.get<QuestionsResponse>({
      endpoint: ['tests', payload.id, 'questions'],
      payload: {
        schema: GetTestQuestionsPayloadSchema,
        data: payload,
      },
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

  async getQuestion(
    payload: GetQuestionPayload,
  ): AsyncResult<QuestionResponse> {
    return this.get<QuestionResponse>({
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
  ): AsyncResult<SimpleResponse> {
    return this.delete<SimpleResponse>({
      endpoint: ['tests', payload.test_id, 'questions', payload.id],
      payload: {
        schema: DeleteQuestionPayloadSchema,
        data: payload,
      },
    });
  }

  // === STUDENT METHODS ===
  async getAvailableTests(): AsyncResult<TestsResponse> {
    return this.get<TestsResponse>({
      endpoint: ['tests', 'available'],
    });
  }

  async getTestForm(
    payload: GetTestFormPayload,
  ): AsyncResult<GetTestFormResponse> {
    return this.get<GetTestFormResponse>({
      endpoint: ['tests', payload.id, 'form'],
      payload: {
        schema: GetTestFormPayloadSchema,
        data: payload,
      },
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

  async getMySubmissions(): AsyncResult<SubmissionsResponse> {
    return this.get<SubmissionsResponse>({
      endpoint: 'submissions/my',
    });
  }

  async getSubmission(
    payload: GetSubmissionPayload,
  ): AsyncResult<SubmissionResponse> {
    return this.get<SubmissionResponse>({
      endpoint: ['submissions', payload.id],
      payload: {
        schema: GetSubmissionPayloadSchema,
        data: payload,
      },
    });
  }

  async getSubmissionResult(
    payload: GetSubmissionResultPayload,
  ): AsyncResult<GetSubmissionResultResponse> {
    return this.get<GetSubmissionResultResponse>({
      endpoint: ['submissions', payload.id, 'result'],
      payload: {
        schema: GetSubmissionResultPayloadSchema,
        data: payload,
      },
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
