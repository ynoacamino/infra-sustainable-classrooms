import type {
  CommentSchema,
  OwnVideoSchema,
  VideoCategorySchema,
  VideoDetailsSchema,
  VideoSchema,
  VideoTagSchema,
} from '@/types/video_learning/schemas/models';
import type z from 'zod';

export type Video = z.infer<typeof VideoSchema>;

export type VideoDetails = z.infer<typeof VideoDetailsSchema>;

export type OwnVideo = z.infer<typeof OwnVideoSchema>;

export type Comment = z.infer<typeof CommentSchema>;

export type VideoCategory = z.infer<typeof VideoCategorySchema>;

export type VideoTag = z.infer<typeof VideoTagSchema>;
