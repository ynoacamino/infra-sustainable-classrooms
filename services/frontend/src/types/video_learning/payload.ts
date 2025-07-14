import type {
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
import type z from 'zod';

export type SearchVideosPayload = z.infer<typeof SearchVideosPayloadSchema>;

export type GetRecommendationsPayload = z.infer<
  typeof GetRecommendationsPayloadSchema
>;

export type GetVideoDetailsPayload = z.infer<
  typeof GetVideoDetailsPayloadSchema
>;

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

export type GetOrCreateCategoryPayload = z.infer<
  typeof GetOrCreateCategoryPayloadSchema
>;

export type GetOrCreateTagPayload = z.infer<typeof GetOrCreateTagPayloadSchema>;
