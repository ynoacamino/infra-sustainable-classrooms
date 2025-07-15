import type { Video, VideoCategory } from '@/types/video_learning/models';

export type UploadResponse = {
  object_name: string;
  presigned_url?: string;
};

export type GetVideosGroupedByCategoryResponse = {
  category: VideoCategory;
  videos: Video[];
}[];
