import { SessionInterceptor } from '@/services/auth/interceptor';
import type { SimpleResponse } from '@/services/shared/response';
import { Service } from '@/services/shared/service';
import type { AsyncResult } from '@/types/shared/services/result';
import type {
  Comment,
  OwnVideo,
  Video,
  VideoCategory,
  VideoDetails,
  VideoTag,
} from '@/types/video_learning/models';
import type {
  CompleteUploadPayload,
  CreateCategoryPayload,
  CreateCommentPayload,
  CreateTagPayload,
  DeleteCommentPayload,
  DeleteVideoPaylod,
  GetCategoryPayload,
  GetCommentsPayload,
  GetOwnVideosPayload,
  GetRecommendationsPayload,
  GetSimilarVideosPayload,
  GetTagPayload,
  GetTagsByVideoPayload,
  GetVideoPayload,
  GetVideosByCategoryPayload,
  GetVideosGroupedByCategoryPayload,
  SearchVideosPayload,
  ToggleVideoLikePayload,
  UploadThumbnailPayload,
  UploadVideoPayload,
} from '@/types/video_learning/payload';
import type {
  GetVideosGroupedByCategoryResponse,
  UploadResponse,
} from '@/types/video_learning/responses';
import {
  CompleteUploadPayloadSchema,
  CreateCategoryPayloadSchema,
  CreateCommentPayloadSchema,
  CreateTagPayloadSchema,
  DeleteCommentPayloadSchema,
  DeleteVideoPayloadSchema,
  GetCategoryPayloadSchema,
  GetCommentsPayloadSchema,
  GetOwnVideosPayloadSchema,
  GetRecommendationsPayloadSchema,
  GetSimilarVideosPayloadSchema,
  GetTagPayloadSchema,
  GetVideoPayloadSchema,
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
      query: ['amount'],
    });
  }

  async getVideo(payload: GetVideoPayload): AsyncResult<VideoDetails> {
    return this.get<VideoDetails>({
      endpoint: ['video', payload.id],
      payload: {
        schema: GetVideoPayloadSchema,
        data: payload,
      },
    });
  }

  // Doubt in backend, it must be the video id in the url as a parameter
  async getComments(payload: GetCommentsPayload): AsyncResult<Comment[]> {
    return this.get<Comment[], typeof GetCommentsPayloadSchema>({
      endpoint: ['comments', payload.video_id],
      payload: {
        schema: GetCommentsPayloadSchema,
        data: payload,
      },
      query: ['page', 'page_size'],
    });
  }

  async createComment(
    payload: CreateCommentPayload,
  ): AsyncResult<SimpleResponse> {
    return this.post<SimpleResponse>({
      endpoint: ['comments', payload.video_id, 'create'],
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
      endpoint: ['upload', 'video', encodeURIComponent(payload.filename)],
      payload: {
        schema: UploadVideoPayloadSchema,
        data: payload,
      },
      multipart: true,
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
      endpoint: ['upload', 'thumbnail', encodeURIComponent(payload.filename)],
      payload: {
        schema: UploadThumbnailPayloadSchema,
        data: payload,
      },
      multipart: true,
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

  async toggleVideoLike(
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
      endpoint: ['category', payload.id],
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

  async createCategory(
    payload: CreateCategoryPayload,
  ): AsyncResult<VideoCategory> {
    return this.post<VideoCategory>({
      endpoint: 'category',
      payload: {
        schema: CreateCategoryPayloadSchema,
        data: payload,
      },
    });
  }

  async getCategory(payload: GetCategoryPayload): AsyncResult<VideoCategory> {
    return this.get<VideoCategory>({
      endpoint: ['category', payload.id],
      payload: {
        schema: GetCategoryPayloadSchema,
        data: payload,
      },
    });
  }

  async createTag(payload: CreateTagPayload): AsyncResult<VideoTag> {
    return this.post<VideoTag>({
      endpoint: 'tag',
      payload: {
        schema: CreateTagPayloadSchema,
        data: payload,
      },
    });
  }

  async getTag(payload: GetTagPayload): AsyncResult<VideoTag> {
    return this.get<VideoTag>({
      endpoint: ['tag', payload.id],
      payload: {
        schema: GetTagPayloadSchema,
        data: payload,
      },
    });
  }

  async getTagsByVideo(
    payload: GetTagsByVideoPayload,
  ): AsyncResult<VideoTag[]> {
    const resTags = await this.getAllTags();
    if (!resTags.success) {
      return resTags;
    }
    const resVideo = await this.getVideo(payload);
    if (!resVideo.success) {
      return resVideo;
    }
    const videoTags = resTags.data.filter((tag) =>
      resVideo.data.tags_ids.includes(tag.id),
    );
    return this.result(videoTags);
  }

  async getVideosGroupedByCategory(
    payload: GetVideosGroupedByCategoryPayload,
  ): AsyncResult<GetVideosGroupedByCategoryResponse> {
    const resCategories = await this.getAllCategories();
    if (!resCategories.success) {
      return resCategories;
    }

    const categories = resCategories.data;
    const promises = categories.map(async (category) => {
      const resVideos = await this.getVideosByCategory({
        id: category.id,
        amount: payload.amount || 10,
      });
      if (!resVideos.success) {
        return { category, videos: [] };
      }
      return { category, videos: resVideos.data };
    });
    const videosByCategory = await Promise.all(promises);

    return this.result(videosByCategory);
  }
}

export const videoLearningService = async (
  cookies: Promise<ReadonlyRequestCookies>,
) => {
  const service = new VideoLearningService();
  service.addInterceptor(new SessionInterceptor(await cookies));
  return service;
};
