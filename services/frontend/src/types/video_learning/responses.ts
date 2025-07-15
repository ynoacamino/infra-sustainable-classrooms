import type { Video, VideoCategory } from '@/types/video_learning/models';

export type UploadResponse = {
  object_name: string;
  presigned_url?: string;
};

export type GetVideosGroupedByCategoryResponse = {
  category: VideoCategory;
  videos: Video[];
}[];

export type SearchVideosResponse = {
  videos: Video[];
};

export type GetRecommendationsResponse = {
  videos: Video[];
};

export type GetVideosByCategoryResponse = {
  videos: Video[];
};
