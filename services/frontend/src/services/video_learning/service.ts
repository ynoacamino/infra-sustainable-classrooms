import { SessionInterceptor } from '@/services/auth/interceptor';
import type { SimpleResponse } from '@/services/shared/response';
import { Service } from '@/services/shared/service';
import type { AsyncResult } from '@/types/shared/services/result';
import type {
  OwnVideo,
  Video,
  VideoCategory,
  VideoDetails,
  VideoTag,
} from '@/types/video_learning/models';
import type {
  CompleteUploadPayload,
  CreateCommentPayload,
  DeleteCommentPayload,
  DeleteVideoPaylod,
  GetCommentsPayload,
  GetOrCreateCategoryPayload,
  GetOrCreateTagPayload,
  GetOwnVideosPayload,
  GetRecommendationsPayload,
  GetSimilarVideosPayload,
  GetVideoDetailsPayload,
  GetVideosByCategoryPayload,
  SearchVideosPayload,
  ToggleVideoLikePayload,
  UploadThumbnailPayload,
  UploadVideoPayload,
} from '@/types/video_learning/payload';
import type { UploadResponse } from '@/types/video_learning/responses';
import {
  CompleteUploadPayloadSchema,
  CreateCommentPayloadSchema,
  DeleteCommentPayloadSchema,
  DeleteVideoPayloadSchema,
  GetCommentsPayloadSchema,
  GetOrCreateCategoryPayloadSchema,
  GetOrCreateTagPayloadSchema,
  GetOwnVideosPayloadSchema,
  GetRecommendationsPayloadSchema,
  GetSimilarVideosPayloadSchema,
  GetVideoDetailsPayloadSchema,
  GetVideosByCategoryPayloadSchema,
  SearchVideosPayloadSchema,
  ToggleVideoLikePayloadSchema,
  UploadThumbnailPayloadSchema,
  UploadVideoPayloadSchema,
} from '@/types/video_learning/schemas/payload';
import type { ReadonlyRequestCookies } from 'next/dist/server/web/spec-extension/adapters/request-cookies';

class VideoLearningService extends Service {
  constructor() {
    super('video_learning');
  }

  async searchVideos(payload: SearchVideosPayload): AsyncResult<Video[]> {
    return this.get<Video[], typeof SearchVideosPayloadSchema>({
      endpoint: 'search',
      payload: {
        schema: SearchVideosPayloadSchema,
        data: payload,
      },
      query: ['query', 'category_id', 'page_size', 'page'],
    });
  }

  async getRecommendations(
    payload: GetRecommendationsPayload,
  ): AsyncResult<Video[]> {
    return this.get<Video[], typeof GetRecommendationsPayloadSchema>({
      endpoint: 'recommendations',
      payload: {
        schema: GetRecommendationsPayloadSchema,
        data: payload,
      },
      query: ['ammount'],
    });
  }

  async getVideoDetails(
    payload: GetVideoDetailsPayload,
  ): AsyncResult<VideoDetails> {
    return this.get<VideoDetails>({
      endpoint: ['video', payload.id],
      payload: {
        schema: GetVideoDetailsPayloadSchema,
        data: payload,
      },
    });
  }

  // Doubt in backend, it must be the video id in the url as a parameter
  async getComments(payload: GetCommentsPayload): AsyncResult<Comment[]> {
    return this.get<Comment[], typeof GetCommentsPayloadSchema>({
      endpoint: ['comments'],
      payload: {
        schema: GetCommentsPayloadSchema,
        data: payload,
      },
      query: ['id', 'page', 'page_size'],
    });
  }

  async createComment(
    payload: CreateCommentPayload,
  ): AsyncResult<SimpleResponse> {
    return this.post<SimpleResponse>({
      endpoint: ['comment', 'create'],
      payload: {
        schema: CreateCommentPayloadSchema,
        data: payload,
      },
    });
  }

  async getOwnVideos(payload: GetOwnVideosPayload): AsyncResult<OwnVideo[]> {
    return this.get<OwnVideo[], typeof GetOwnVideosPayloadSchema>({
      endpoint: 'my-videos',
      payload: {
        schema: GetOwnVideosPayloadSchema,
        data: payload,
      },
      query: ['page', 'page_size'],
    });
  }

  // TODO: Implement video upload with form data if needed
  async initialUpload(
    payload: UploadVideoPayload,
  ): AsyncResult<UploadResponse> {
    return this.post<UploadResponse>({
      endpoint: ['upload', 'video'],
      payload: {
        schema: UploadVideoPayloadSchema,
        data: payload,
      },
    });
  }

  async completeUpload(
    payload: CompleteUploadPayload,
  ): AsyncResult<SimpleResponse> {
    return this.post<SimpleResponse>({
      endpoint: ['upload', 'complete'],
      payload: {
        schema: CompleteUploadPayloadSchema,
        data: payload,
      },
    });
  }

  async uploadThumbnail(
    payload: UploadThumbnailPayload,
  ): AsyncResult<UploadResponse> {
    return this.post<UploadResponse>({
      endpoint: ['upload', 'thumbnail'],
      payload: {
        schema: UploadThumbnailPayloadSchema,
        data: payload,
      },
    });
  }

  async getAllCategories(): AsyncResult<VideoCategory[]> {
    return this.get<VideoCategory[]>({
      endpoint: 'categories',
    });
  }

  async getAllTags(): AsyncResult<VideoTag[]> {
    return this.get<VideoTag[]>({
      endpoint: 'tags',
    });
  }

  async toogleVideoLike(
    payload: ToggleVideoLikePayload,
  ): AsyncResult<SimpleResponse> {
    return this.post<SimpleResponse>({
      endpoint: ['video', payload.id, 'like'],
      payload: {
        schema: ToggleVideoLikePayloadSchema,
        data: payload,
      },
    });
  }

  async deleteVideo(payload: DeleteVideoPaylod): AsyncResult<SimpleResponse> {
    return this.delete<SimpleResponse>({
      endpoint: ['video', payload.id],
      payload: {
        schema: DeleteVideoPayloadSchema,
        data: payload,
      },
    });
  }

  async getVideosByCategory(
    payload: GetVideosByCategoryPayload,
  ): AsyncResult<Video[]> {
    return this.get<Video[], typeof GetVideosByCategoryPayloadSchema>({
      endpoint: ['category', payload.category_id],
      payload: {
        schema: GetVideosByCategoryPayloadSchema,
        data: payload,
      },
      query: ['amount'],
    });
  }

  async getSimilarVideos(
    payload: GetSimilarVideosPayload,
  ): AsyncResult<Video[]> {
    return this.get<Video[], typeof GetSimilarVideosPayloadSchema>({
      endpoint: ['video', payload.id, 'similar'],
      payload: {
        schema: GetSimilarVideosPayloadSchema,
        data: payload,
      },
      query: ['amount'],
    });
  }

  async deleteComment(
    payload: DeleteCommentPayload,
  ): AsyncResult<SimpleResponse> {
    return this.delete<SimpleResponse>({
      endpoint: ['comment', payload.id],
      payload: {
        schema: DeleteCommentPayloadSchema,
        data: payload,
      },
    });
  }

  async getOrCreateCategory(
    payload: GetOrCreateCategoryPayload,
  ): AsyncResult<VideoCategory> {
    return this.post<VideoCategory>({
      endpoint: 'category',
      payload: {
        schema: GetOrCreateCategoryPayloadSchema,
        data: payload,
      },
    });
  }

  async getOrCreateTag(payload: GetOrCreateTagPayload): AsyncResult<VideoTag> {
    return this.post<VideoTag>({
      endpoint: 'tag',
      payload: {
        schema: GetOrCreateTagPayloadSchema,
        data: payload,
      },
    });
  }
}

export const videoLearningService = async (
  cookies: Promise<ReadonlyRequestCookies>,
) => {
  const service = new VideoLearningService();
  service.addInterceptor(new SessionInterceptor(await cookies));
  return service;
};
