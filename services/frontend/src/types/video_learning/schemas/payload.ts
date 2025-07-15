import { PAGINATION_DEFAULT_SIZE } from '@/config/shared/const';
import {
  CommentSchema,
  VideoCategorySchema,
  VideoDetailsSchema,
  VideoSchema,
  VideoTagSchema,
} from '@/types/video_learning/schemas/models';
import z from 'zod';

const PaginationSchema = z.object({
  page: z
    .number()
    .int('Page must be an integer')
    .positive('Page must be greater than 0')
    .default(1),
  page_size: z
    .number()
    .int('Page size must be an integer')
    .positive('Page size must be greater than 0')
    .max(100, 'Page size cannot exceed 100')
    .default(PAGINATION_DEFAULT_SIZE)
    .optional(),
});

const AmountSchema = z.object({
  amount: z
    .number()
    .int('Amount must be an integer')
    .min(1, 'Amount must be at least 1')
    .max(100, 'Amount cannot exceed 100')
    .default(20)
    .optional(),
});

const UploadSchema = z.object({
  file: z
    .instanceof(File)
    .refine((file) => file.size > 0, 'File data is required'),
  filename: z
    .string()
    .min(1, 'Filename is required')
    .max(200, 'Filename cannot exceed 200 characters'),
});

export const SearchVideosPayloadSchema = PaginationSchema.extend({
  query: z
    .string()
    .min(1, 'Query is required')
    .max(200, 'Query cannot exceed 200 characters'),
  category_id: VideoCategorySchema.shape.id.optional(),
});

export const GetRecommendationsPayloadSchema = AmountSchema;

export const GetVideoPayloadSchema = VideoDetailsSchema.pick({
  id: true,
});

export const GetCommentsPayloadSchema = CommentSchema.pick({
  video_id: true,
}).merge(PaginationSchema);

export const CreateCommentPayloadSchema = CommentSchema.omit({
  id: true,
  date: true,
  author: true,
});

export const GetOwnVideosPayloadSchema = PaginationSchema;

export const UploadVideoPayloadSchema = UploadSchema;

export const CompleteUploadPayloadSchema = VideoDetailsSchema.omit({
  id: true,
  author: true,
  likes: true,
  thumbnail_url: true,
  upload_date: true,
  video_url: true,
  views: true,
}).extend({
  thumbnail_object_name: z
    .string()
    .min(1, 'Thumbnail object name is required')
    .max(200, 'Thumbnail object name cannot exceed 200 characters'),
  video_object_name: z
    .string()
    .min(1, 'Video object name is required')
    .max(200, 'Video object name cannot exceed 200 characters'),
});

export const UploadThumbnailPayloadSchema = UploadSchema;

export const ToggleVideoLikePayloadSchema = VideoSchema.pick({
  id: true,
});

export const DeleteVideoPayloadSchema = VideoSchema.pick({
  id: true,
});

export const GetVideosByCategoryPayloadSchema = VideoCategorySchema.pick({
  id: true,
}).merge(AmountSchema);

export const GetSimilarVideosPayloadSchema = VideoSchema.pick({
  id: true,
}).merge(AmountSchema);

export const DeleteCommentPayloadSchema = CommentSchema.pick({
  id: true,
});

export const CreateCategoryPayloadSchema = VideoCategorySchema.omit({
  id: true,
});

export const GetCategoryPayloadSchema = VideoCategorySchema.pick({
  id: true,
});

export const CreateTagPayloadSchema = VideoTagSchema.omit({
  id: true,
});

export const GetTagPayloadSchema = VideoTagSchema.pick({
  id: true,
});

export const GetTagsByVideoPayloadSchema = VideoSchema.pick({
  id: true,
});

export const GetVideosGroupedByCategoryPayloadSchema = AmountSchema;
