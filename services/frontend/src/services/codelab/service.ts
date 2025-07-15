import { SessionInterceptor } from '@/services/auth/interceptor';
import { Service } from '@/services/shared/service';
import type {
  CreateExercisePayload,
  GetExercisePayload,
  UpdateExercisePayload,
  DeleteExercisePayload,
  CreateTestPayload,
  GetTestsByExercisePayload,
  UpdateTestPayload,
  DeleteTestPayload,
  GetExerciseForStudentPayload,
  CreateAttemptPayload,
  GetAttemptsByUserAndExercisePayload,
  GetAnswerByUserAndExercisePayload,
} from '@/types/codelab/payload';
import {
  CreateExercisePayloadSchema,
  UpdateExercisePayloadSchema,
  DeleteExercisePayloadSchema,
  CreateTestPayloadSchema,
  UpdateTestPayloadSchema,
  DeleteTestPayloadSchema,
  CreateAttemptPayloadSchema,
} from '@/types/codelab/schemas/payload';
import type { AsyncResult } from '@/types/shared/services/result';
import type { ReadonlyRequestCookies } from 'next/dist/server/web/spec-extension/adapters/request-cookies';
import type { SimpleResponse } from '@/services/shared/response';
import type {
  Exercise,
  Test,
  Answer,
  Attempt,
  ExerciseForStudents,
  ExerciseForStudentsListView,
} from '@/types/codelab/models';

class CodelabService extends Service {
  constructor() {
    super('codelab');
  }

  async createExercise(
    payload: CreateExercisePayload,
  ): AsyncResult<SimpleResponse> {
    return this.post<SimpleResponse>({
      endpoint: 'exercises',
      payload: {
        schema: CreateExercisePayloadSchema,
        data: payload,
      },
    });
  }

  async getExercise(payload: GetExercisePayload): AsyncResult<Exercise> {
    return this.get<Exercise>({
      endpoint: ['exercises', payload.id],
    });
  }

  async listExercises(): AsyncResult<Exercise[]> {
    return this.get<Exercise[]>({
      endpoint: 'exercises',
    });
  }

  async updateExercise(
    payload: UpdateExercisePayload,
  ): AsyncResult<SimpleResponse> {
    return this.put<SimpleResponse>({
      endpoint: ['exercises', payload.id],
      payload: {
        schema: UpdateExercisePayloadSchema,
        data: payload,
      },
    });
  }

  async deleteExercise(
    payload: DeleteExercisePayload,
  ): AsyncResult<SimpleResponse> {
    return this.delete<SimpleResponse>({
      endpoint: ['exercises', payload.id],
      payload: {
        schema: DeleteExercisePayloadSchema,
        data: payload,
      },
    });
  }

  // ========================================
  // TEST CRUD METHODS (for professors)
  // ========================================

  async createTest(payload: CreateTestPayload): AsyncResult<SimpleResponse> {
    return this.post<SimpleResponse>({
      endpoint: 'tests',
      payload: {
        schema: CreateTestPayloadSchema,
        data: payload,
      },
    });
  }

  async getTestsByExercise(
    payload: GetTestsByExercisePayload,
  ): AsyncResult<Test[]> {
    return this.get<Test[]>({
      endpoint: ['exercises', payload.exercise_id, 'tests'],
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

  async deleteTest(payload: DeleteTestPayload): AsyncResult<SimpleResponse> {
    return this.delete<SimpleResponse>({
      endpoint: ['tests', payload.id],
      payload: {
        schema: DeleteTestPayloadSchema,
        data: payload,
      },
    });
  }

  // ========================================
  // STUDENT METHODS (read exercises, submit attempts)
  // ========================================

  async getExerciseForStudent(
    payload: GetExerciseForStudentPayload,
  ): AsyncResult<ExerciseForStudents> {
    return this.get<ExerciseForStudents>({
      endpoint: ['student', 'exercises', payload.id],
    });
  }

  async listExercisesForStudents(): AsyncResult<ExerciseForStudentsListView[]> {
    return this.get<ExerciseForStudentsListView[]>({
      endpoint: ['student', 'exercises'],
    });
  }

  async createAttempt(
    payload: CreateAttemptPayload,
  ): AsyncResult<SimpleResponse> {
    return this.post<SimpleResponse>({
      endpoint: 'attempts',
      payload: {
        schema: CreateAttemptPayloadSchema,
        data: payload,
      },
    });
  }

  async getAttemptsByUserAndExercise(
    payload: GetAttemptsByUserAndExercisePayload,
  ): AsyncResult<Attempt[]> {
    return this.get<Attempt[]>({
      endpoint: [
        'student',
        'users',
        payload.user_id,
        'exercises',
        payload.exercise_id,
        'attempts',
      ],
    });
  }

  // ========================================
  // ANSWER MANAGEMENT METHODS
  // ========================================

  async getAnswerByUserAndExercise(
    payload: GetAnswerByUserAndExercisePayload,
  ): AsyncResult<Answer> {
    return this.get<Answer>({
      endpoint: [
        'answers',
        'user',
        payload.user_id,
        'exercise',
        payload.exercise_id,
      ],
    });
  }
}

export const codelabService = async (
  cookies: Promise<ReadonlyRequestCookies>,
) => {
  const service = new CodelabService();
  service.addInterceptor(new SessionInterceptor(await cookies));
  return service;
};
