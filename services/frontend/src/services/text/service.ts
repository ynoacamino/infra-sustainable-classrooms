import { SessionInterceptor } from '@/services/auth/interceptor';
import { Service } from '@/services/shared/service';
import type {
  Course,
  Section,
  Article,
  SimpleResponse,
} from '@/types/text/models';
import type {
  CreateCoursePayload,
  GetCoursePayload,
  UpdateCoursePayload,
  DeleteCoursePayload,
  ListCoursesPayload,
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
import type {
  CreateCourseResponse,
  GetCourseResponse,
  ListCoursesResponse,
  UpdateCourseResponse,
  DeleteCourseResponse,
  CreateSectionResponse,
  GetSectionResponse,
  ListSectionsResponse,
  UpdateSectionResponse,
  DeleteSectionResponse,
  CreateArticleResponse,
  GetArticleResponse,
  ListArticlesResponse,
  UpdateArticleResponse,
  DeleteArticleResponse,
} from '@/types/text/responses';
import {
  CreateCoursePayloadSchema,
  GetCoursePayloadSchema,
  UpdateCoursePayloadSchema,
  ListCoursesPayloadSchema,
  CreateSectionPayloadSchema,
  GetSectionPayloadSchema,
  ListSectionsPayloadSchema,
  UpdateSectionPayloadSchema,
  CreateArticlePayloadSchema,
  GetArticlePayloadSchema,
  ListArticlesPayloadSchema,
  UpdateArticlePayloadSchema,
} from '@/types/text/schemas/payload';
import type { AsyncResult } from '@/types/shared/services/result';
import type { ReadonlyRequestCookies } from 'next/dist/server/web/spec-extension/adapters/request-cookies';

class TextService extends Service {
  constructor() {
    super('text');
  }

  // === COURSE METHODS ===
  async createCourse(
    payload: CreateCoursePayload,
  ): AsyncResult<CreateCourseResponse> {
    return this.post<CreateCourseResponse>({
      endpoint: 'courses',
      payload: {
        schema: CreateCoursePayloadSchema,
        data: payload,
      },
    });
  }

  async getCourse(payload: GetCoursePayload): AsyncResult<GetCourseResponse> {
    return this.get<GetCourseResponse>({
      endpoint: ['courses', payload.course_id],
      payload: {
        schema: GetCoursePayloadSchema,
        data: payload,
      },
    });
  }

  async listCourses(
    payload: ListCoursesPayload,
  ): AsyncResult<ListCoursesResponse> {
    return this.get<ListCoursesResponse>({
      endpoint: 'courses',
      payload: {
        schema: ListCoursesPayloadSchema,
        data: payload,
      },
    });
  }

  async updateCourse(
    payload: UpdateCoursePayload,
  ): AsyncResult<UpdateCourseResponse> {
    return this.patch<UpdateCourseResponse>({
      endpoint: ['courses', payload.course_id],
      payload: {
        schema: UpdateCoursePayloadSchema,
        data: payload,
      },
    });
  }

  async deleteCourse(
    payload: DeleteCoursePayload,
  ): AsyncResult<DeleteCourseResponse> {
    return this.delete<DeleteCourseResponse>([
      'courses',
      payload.course_id.toString(),
    ]);
  }

  // === SECTION METHODS ===
  async createSection(
    payload: CreateSectionPayload,
  ): AsyncResult<CreateSectionResponse> {
    return this.post<CreateSectionResponse>({
      endpoint: ['courses', payload.course_id, 'sections'],
      payload: {
        schema: CreateSectionPayloadSchema,
        data: payload,
      },
    });
  }

  async getSection(
    payload: GetSectionPayload,
  ): AsyncResult<GetSectionResponse> {
    return this.get<GetSectionResponse>({
      endpoint: ['sections', payload.section_id],
      payload: {
        schema: GetSectionPayloadSchema,
        data: payload,
      },
    });
  }

  async listSections(
    payload: ListSectionsPayload,
  ): AsyncResult<ListSectionsResponse> {
    return this.get<ListSectionsResponse>({
      endpoint: ['courses', payload.course_id, 'sections'],
      payload: {
        schema: ListSectionsPayloadSchema,
        data: payload,
      },
    });
  }

  async updateSection(
    payload: UpdateSectionPayload,
  ): AsyncResult<UpdateSectionResponse> {
    return this.patch<UpdateSectionResponse>({
      endpoint: ['sections', payload.section_id],
      payload: {
        schema: UpdateSectionPayloadSchema,
        data: payload,
      },
    });
  }

  async deleteSection(
    payload: DeleteSectionPayload,
  ): AsyncResult<DeleteSectionResponse> {
    return this.delete<DeleteSectionResponse>([
      'sections',
      payload.section_id.toString(),
    ]);
  }

  // === ARTICLE METHODS ===
  async createArticle(
    payload: CreateArticlePayload,
  ): AsyncResult<CreateArticleResponse> {
    return this.post<CreateArticleResponse>({
      endpoint: ['sections', payload.section_id, 'articles'],
      payload: {
        schema: CreateArticlePayloadSchema,
        data: payload,
      },
    });
  }

  async getArticle(
    payload: GetArticlePayload,
  ): AsyncResult<GetArticleResponse> {
    return this.get<GetArticleResponse>({
      endpoint: ['articles', payload.article_id],
      payload: {
        schema: GetArticlePayloadSchema,
        data: payload,
      },
    });
  }

  async listArticles(
    payload: ListArticlesPayload,
  ): AsyncResult<ListArticlesResponse> {
    return this.get<ListArticlesResponse>({
      endpoint: ['sections', payload.section_id, 'articles'],
      payload: {
        schema: ListArticlesPayloadSchema,
        data: payload,
      },
    });
  }

  async updateArticle(
    payload: UpdateArticlePayload,
  ): AsyncResult<UpdateArticleResponse> {
    return this.patch<UpdateArticleResponse>({
      endpoint: ['articles', payload.article_id],
      payload: {
        schema: UpdateArticlePayloadSchema,
        data: payload,
      },
    });
  }

  async deleteArticle(
    payload: DeleteArticlePayload,
  ): AsyncResult<DeleteArticleResponse> {
    return this.delete<DeleteArticleResponse>([
      'articles',
      payload.article_id.toString(),
    ]);
  }

  // === UTILITY METHODS ===
  /**
   * Helper method to check if a response is a SimpleResponse
   */
  isSimpleResponse(response: any): response is SimpleResponse {
    return (
      response &&
      typeof response.success === 'boolean' &&
      typeof response.message === 'string'
    );
  }

  /**
   * Helper method to check if a response is a Course
   */
  isCourse(response: any): response is Course {
    return (
      response &&
      typeof response.id === 'number' &&
      typeof response.title === 'string'
    );
  }

  /**
   * Helper method to check if a response is a Section
   */
  isSection(response: any): response is Section {
    return (
      response &&
      typeof response.id === 'number' &&
      typeof response.course_id === 'number'
    );
  }

  /**
   * Helper method to check if a response is an Article
   */
  isArticle(response: any): response is Article {
    return (
      response &&
      typeof response.id === 'number' &&
      typeof response.section_id === 'number'
    );
  }
}

export const textService = async (cookies: Promise<ReadonlyRequestCookies>) => {
  const service = new TextService();
  service.addInterceptor(new SessionInterceptor(await cookies));
  return service;
};
