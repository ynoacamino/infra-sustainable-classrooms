import { SessionInterceptor } from '@/services/auth/interceptor';
import { Service } from '@/services/shared/service';
import type {
  CreateCoursePayload,
  GetCoursePayload,
  UpdateCoursePayload,
  DeleteCoursePayload,
  CreateSectionPayload,
  GetSectionPayload,
  ListSectionsPayload,
  UpdateSectionPayload,
  DeleteSectionPayload,
  CreateArticlePayload,
  GetArticlePayload,
  ListArticlesPayload,
  UpdateArticlePayload,
  DeleteArticlePayload,
} from '@/types/text/payload';
import {
  CreateCoursePayloadSchema,
  GetCoursePayloadSchema,
  UpdateCoursePayloadSchema,
  CreateSectionPayloadSchema,
  GetSectionPayloadSchema,
  ListSectionsPayloadSchema,
  UpdateSectionPayloadSchema,
  CreateArticlePayloadSchema,
  GetArticlePayloadSchema,
  ListArticlesPayloadSchema,
  UpdateArticlePayloadSchema,
  DeleteCoursePayloadSchema,
  DeleteSectionPayloadSchema,
  DeleteArticlePayloadSchema,
} from '@/types/text/schemas/payload';
import type { AsyncResult } from '@/types/shared/services/result';
import type { ReadonlyRequestCookies } from 'next/dist/server/web/spec-extension/adapters/request-cookies';
import type { SimpleResponse } from '@/services/shared/response';
import type { Article, Course, Section } from '@/types/text/models';

class TextService extends Service {
  constructor() {
    super('text');
  }

  async createCourse(
    payload: CreateCoursePayload,
  ): AsyncResult<SimpleResponse> {
    return this.post<SimpleResponse>({
      endpoint: 'courses',
      payload: {
        schema: CreateCoursePayloadSchema,
        data: payload,
      },
    });
  }

  async getCourse(payload: GetCoursePayload): AsyncResult<Course> {
    return this.get<Course>({
      endpoint: ['courses', payload.id],
      payload: {
        schema: GetCoursePayloadSchema,
        data: payload,
      },
    });
  }

  async listCourses(): AsyncResult<Course[]> {
    return this.get<Course[]>({
      endpoint: 'courses',
    });
  }

  async deleteCourse(
    payload: DeleteCoursePayload,
  ): AsyncResult<SimpleResponse> {
    return this.delete<SimpleResponse>({
      endpoint: ['courses', payload.id],
      payload: {
        schema: DeleteCoursePayloadSchema,
        data: payload,
      },
    });
  }

  async updateCourse(
    payload: UpdateCoursePayload,
  ): AsyncResult<SimpleResponse> {
    return this.patch<SimpleResponse>({
      endpoint: ['courses', payload.id],
      payload: {
        schema: UpdateCoursePayloadSchema,
        data: payload,
      },
    });
  }

  // === SECTION METHODS ===
  async createSection(
    payload: CreateSectionPayload,
  ): AsyncResult<SimpleResponse> {
    return this.post<SimpleResponse>({
      endpoint: ['courses', payload.course_id, 'sections'],
      payload: {
        schema: CreateSectionPayloadSchema,
        data: payload,
      },
    });
  }

  async getSection(payload: GetSectionPayload): AsyncResult<Section> {
    return this.get<Section>({
      endpoint: ['sections', payload.id],
      payload: {
        schema: GetSectionPayloadSchema,
        data: payload,
      },
    });
  }

  async listSections(payload: ListSectionsPayload): AsyncResult<Section[]> {
    return this.get<Section[]>({
      endpoint: ['courses', payload.course_id, 'sections'],
      payload: {
        schema: ListSectionsPayloadSchema,
        data: payload,
      },
    });
  }

  async updateSection(
    payload: UpdateSectionPayload,
  ): AsyncResult<SimpleResponse> {
    return this.patch<SimpleResponse>({
      endpoint: ['sections', payload.id],
      payload: {
        schema: UpdateSectionPayloadSchema,
        data: payload,
      },
    });
  }

  async deleteSection(
    payload: DeleteSectionPayload,
  ): AsyncResult<SimpleResponse> {
    return this.delete<SimpleResponse>({
      endpoint: ['sections', payload.id],
      payload: {
        schema: DeleteSectionPayloadSchema,
        data: payload,
      },
    });
  }

  // === ARTICLE METHODS ===
  async createArticle(
    payload: CreateArticlePayload,
  ): AsyncResult<SimpleResponse> {
    return this.post<SimpleResponse>({
      endpoint: ['sections', payload.section_id, 'articles'],
      payload: {
        schema: CreateArticlePayloadSchema,
        data: payload,
      },
    });
  }

  async getArticle(payload: GetArticlePayload): AsyncResult<Article> {
    return this.get<Article>({
      endpoint: ['articles', payload.id],
      payload: {
        schema: GetArticlePayloadSchema,
        data: payload,
      },
    });
  }

  async listArticles(payload: ListArticlesPayload): AsyncResult<Article[]> {
    return this.get<Article[]>({
      endpoint: ['sections', payload.section_id, 'articles'],
      payload: {
        schema: ListArticlesPayloadSchema,
        data: payload,
      },
    });
  }

  async updateArticle(
    payload: UpdateArticlePayload,
  ): AsyncResult<SimpleResponse> {
    return this.patch<SimpleResponse>({
      endpoint: ['articles', payload.id],
      payload: {
        schema: UpdateArticlePayloadSchema,
        data: payload,
      },
    });
  }

  async deleteArticle(
    payload: DeleteArticlePayload,
  ): AsyncResult<SimpleResponse> {
    return this.delete<SimpleResponse>({
      endpoint: ['articles', payload.id],
      payload: {
        schema: DeleteArticlePayloadSchema,
        data: payload,
      },
    });
  }
}

export const textService = async (cookies: Promise<ReadonlyRequestCookies>) => {
  const service = new TextService();
  service.addInterceptor(new SessionInterceptor(await cookies));
  return service;
};
