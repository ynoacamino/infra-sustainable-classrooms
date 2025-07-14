import type {
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
import type z from 'zod';

export type SearchVideosPayload = z.infer<typeof SearchVideosPayloadSchema>;

export type GetRecommendationsPayload = z.infer<
  typeof GetRecommendationsPayloadSchema
>;

export type GetVideoPayload = z.infer<typeof GetVideoPayloadSchema>;

export type GetCommentsPayload = z.infer<typeof GetCommentsPayloadSchema>;

export type CreateCommentPayload = z.infer<typeof CreateCommentPayloadSchema>;

export type GetOwnVideosPayload = z.infer<typeof GetOwnVideosPayloadSchema>;

export type UploadVideoPayload = z.infer<typeof UploadVideoPayloadSchema>;

export type CompleteUploadPayload = z.infer<typeof CompleteUploadPayloadSchema>;

export type UploadThumbnailPayload = z.infer<
  typeof UploadThumbnailPayloadSchema
>;

export type ToggleVideoLikePayload = z.infer<
  typeof ToggleVideoLikePayloadSchema
>;

export type DeleteVideoPaylod = z.infer<typeof DeleteVideoPayloadSchema>;

export type GetVideosByCategoryPayload = z.infer<
  typeof GetVideosByCategoryPayloadSchema
>;

export type GetSimilarVideosPayload = z.infer<
  typeof GetSimilarVideosPayloadSchema
>;

export type DeleteCommentPayload = z.infer<typeof DeleteCommentPayloadSchema>;

export type CreateCategoryPayload = z.infer<typeof CreateCategoryPayloadSchema>;

export type GetCategoryPayload = z.infer<typeof GetCategoryPayloadSchema>;

export type CreateTagPayload = z.infer<typeof CreateTagPayloadSchema>;

export type GetTagPayload = z.infer<typeof GetTagPayloadSchema>;
