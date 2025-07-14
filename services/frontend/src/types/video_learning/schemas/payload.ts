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
    .default(20),
});

const UploadSchema = z.object({
  file: z
    .instanceof(Buffer)
    .refine((file) => file.length > 0, 'File data is required'),
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
  category_id: z
    .number()
    .int('Category ID must be an integer')
    .positive('Category ID must be greater than 0')
    .optional(),
});

export const GetRecommendationsPayloadSchema = z.object({
  amount: z
    .number()
    .int('Amount must be an integer')
    .min(1, 'Amount must be at least 1')
    .max(100, 'Amount cannot exceed 100')
    .default(20)
    .optional(),
});

export const GetVideoDetailsPayloadSchema = VideoDetailsSchema.pick({
  id: true,
});

export const GetCommentsPayloadSchema = VideoSchema.pick({
  id: true,
}).merge(PaginationSchema);

export const CreateCommentPayloadSchema = CommentSchema.pick({
  title: true,
  body: true,
}).extend({
  video_id: z
    .number()
    .int('Video ID must be an integer')
    .positive('Video ID must be greater than 0'),
});

export const GetOwnVideosPayloadSchema = PaginationSchema;

export const UploadVideoPayloadSchema = UploadSchema;

export const CompleteUploadPayloadSchema = VideoDetailsSchema.pick({
  title: true,
  description: true,
  tags: true,
}).extend({
  category_id: z.number().int('Id must be an integer'),
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

export const GetVideosByCategoryPayloadSchema = z.object({
  category_id: z.number().int('Id must be an integer'),
  amount: z
    .number()
    .int('Amount must be an integer')
    .min(1, 'Amount must be at least 1')
    .max(100, 'Amount cannot exceed 100')
    .default(20),
});

export const GetSimilarVideosPayloadSchema = VideoSchema.pick({
  id: true,
}).extend({
  amount: z
    .number()
    .int('Amount must be an integer')
    .min(1, 'Amount must be at least 1')
    .max(100, 'Amount cannot exceed 100')
    .default(20),
});

export const DeleteCommentPayloadSchema = CommentSchema.pick({
  id: true,
});

export const GetOrCreateCategoryPayloadSchema = VideoCategorySchema.omit({
  id: true,
});

export const GetOrCreateTagPayloadSchema = VideoTagSchema.omit({
  id: true,
});
