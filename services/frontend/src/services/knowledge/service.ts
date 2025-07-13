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
} from '@/types/knowledge/payload';
import type {
  CreateTestResponse,
  GetMyTestsResponse,
  UpdateTestResponse,
  DeleteTestResponse,
  GetTestQuestionsResponse,
  AddQuestionResponse,
  UpdateQuestionResponse,
  DeleteQuestionResponse,
  GetAvailableTestsResponse,
  GetTestFormResponse,
  SubmitTestResponse,
  GetMySubmissionsResponse,
  GetSubmissionResultResponse,
} from '@/types/knowledge/responses';
import {
  CreateTestPayloadSchema,
  GetMyTestsPayloadSchema,
  UpdateTestPayloadSchema,
  GetTestQuestionsPayloadSchema,
  AddQuestionPayloadSchema,
  UpdateQuestionPayloadSchema,
  GetAvailableTestsPayloadSchema,
  GetTestFormPayloadSchema,
  SubmitTestPayloadSchema,
  GetMySubmissionsPayloadSchema,
  GetSubmissionResultPayloadSchema,
} from '@/types/knowledge/schemas/payload';
import type { ReadonlyRequestCookies } from 'next/dist/server/web/spec-extension/adapters/request-cookies';
import type { AsyncResult } from '@/types/shared/services/result';

class KnowledgeService extends Service {
  constructor() {
    super('knowledge');
  }

  // === TEACHER METHODS ===

  /**
   * Create a new test
   */
  async createTest(
    payload: CreateTestPayload,
  ): AsyncResult<CreateTestResponse> {
    return this.post<CreateTestResponse>({
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
      payload: {
        schema: GetMyTestsPayloadSchema,
        data: {},
      },
    });
  }

  /**
   * Update test title
   */
  async updateTest(
    payload: UpdateTestPayload,
  ): AsyncResult<UpdateTestResponse> {
    return this.put<UpdateTestResponse>({
      endpoint: ['tests', payload.test_id],
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
    return this.delete<DeleteTestResponse>([
      'tests',
      payload.test_id.toString(),
    ]);
  }

  /**
   * Get questions for a test
   */
  async getTestQuestions(
    payload: GetTestQuestionsPayload,
  ): AsyncResult<GetTestQuestionsResponse> {
    return this.get<GetTestQuestionsResponse>({
      endpoint: ['tests', payload.test_id, 'questions'],
      payload: {
        schema: GetTestQuestionsPayloadSchema,
        data: payload,
      },
    });
  }

  /**
   * Add a question to a test
   */
  async addQuestion(
    payload: AddQuestionPayload,
  ): AsyncResult<AddQuestionResponse> {
    return this.post<AddQuestionResponse>({
      endpoint: ['tests', payload.test_id, 'questions'],
      payload: {
        schema: AddQuestionPayloadSchema,
        data: payload,
      },
    });
  }

  /**
   * Update a question
   */
  async updateQuestion(
    payload: UpdateQuestionPayload,
  ): AsyncResult<UpdateQuestionResponse> {
    return this.put<UpdateQuestionResponse>({
      endpoint: ['tests', payload.test_id, 'questions', payload.question_id],
      payload: {
        schema: UpdateQuestionPayloadSchema,
        data: payload,
      },
    });
  }

  /**
   * Delete a question
   */
  async deleteQuestion(
    payload: DeleteQuestionPayload,
  ): AsyncResult<DeleteQuestionResponse> {
    return this.delete<DeleteQuestionResponse>([
      'tests',
      payload.test_id.toString(),
      'questions',
      payload.question_id.toString(),
    ]);
  }

  // === STUDENT METHODS ===

  /**
   * Get available tests for students
   */
  async getAvailableTests(): AsyncResult<GetAvailableTestsResponse> {
    return this.get<GetAvailableTestsResponse>({
      endpoint: 'tests/available',
      payload: {
        schema: GetAvailableTestsPayloadSchema,
        data: {},
      },
    });
  }

  /**
   * Get test form for taking
   */
  async getTestForm(
    payload: GetTestFormPayload,
  ): AsyncResult<GetTestFormResponse> {
    return this.get<GetTestFormResponse>({
      endpoint: ['tests', payload.test_id, 'form'],
      payload: {
        schema: GetTestFormPayloadSchema,
        data: payload,
      },
    });
  }

  /**
   * Submit test answers
   */
  async submitTest(
    payload: SubmitTestPayload,
  ): AsyncResult<SubmitTestResponse> {
    return this.post<SubmitTestResponse>({
      endpoint: ['tests', payload.test_id, 'submit'],
      payload: {
        schema: SubmitTestPayloadSchema,
        data: payload,
      },
    });
  }

  /**
   * Get my test submissions
   */
  async getMySubmissions(): AsyncResult<GetMySubmissionsResponse> {
    return this.get<GetMySubmissionsResponse>({
      endpoint: 'submissions/my',
      payload: {
        schema: GetMySubmissionsPayloadSchema,
        data: {},
      },
    });
  }

  /**
   * Get detailed submission result
   */
  async getSubmissionResult(
    payload: GetSubmissionResultPayload,
  ): AsyncResult<GetSubmissionResultResponse> {
    return this.get<GetSubmissionResultResponse>({
      endpoint: ['submissions', payload.submission_id, 'result'],
      payload: {
        schema: GetSubmissionResultPayloadSchema,
        data: payload,
      },
    });
  }

  // === CONVENIENCE METHODS ===

  /**
   * Get a specific test by ID (for teachers)
   * This uses getMyTests and filters by test ID since there's no direct getTest endpoint
   */
  async getTest(testId: number): AsyncResult<Test> {
    const result = await this.getMyTests();
    if (!result.success) {
      return result;
    }

    const test = result.data.tests.find((t) => t.id === testId);
    if (!test) {
      return {
        success: false,
        error: {
          message: 'Test not found',
          status: 404,
          reason: 'NOT_FOUND',
          extend: {},
        },
      };
    }

    return {
      success: true,
      data: test,
    };
  }

  /**
   * Get a specific question by ID
   * This uses getTestQuestions and filters by question ID since there's no direct endpoint
   */
  async getQuestion(testId: number, questionId: number): AsyncResult<Question> {
    try {
      const result = await this.getTestQuestions({ test_id: testId });
      if (!result.success) {
        return result;
      }

      const question = result.data.questions.find((q) => q.id === questionId);
      if (!question) {
        return {
          success: false,
          error: {
            message: 'Question not found',
            status: 404,
            reason: 'NOT_FOUND',
            extend: {},
          },
        };
      }

      return {
        success: true,
        data: question,
      };
    } catch (error) {
      return {
        success: false,
        error: {
          message: 'Failed to fetch question',
          status: 500,
          reason: 'INTERNAL_ERROR',
          extend: {},
        },
      };
    }
  }

  /**
   * Get a specific submission by ID
   * This uses getSubmissionResult and extracts the submission data
   */
  async getSubmission(submissionId: number): AsyncResult<Submission> {
    try {
      const result = await this.getSubmissionResult({
        submission_id: submissionId,
      });
      if (!result.success) {
        return result;
      }

      return {
        success: true,
        data: result.data.submission,
      };
    } catch (error) {
      return {
        success: false,
        error: {
          message: 'Failed to fetch submission',
          status: 500,
          reason: 'INTERNAL_ERROR',
          extend: {},
        },
      };
    }
  }
}

export const knowledgeService = async (
  cookies: Promise<ReadonlyRequestCookies>,
) => {
  const service = new KnowledgeService();
  service.addInterceptor(new SessionInterceptor(await cookies));
  return service;
};
