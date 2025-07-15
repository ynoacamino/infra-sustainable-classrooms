'use server';

import { videoLearningService } from '@/services/video_learning/service';
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
import { cookies } from 'next/headers';

export async function searchVideosAction(payload: SearchVideosPayload) {
  const videoLearning = await videoLearningService(cookies());
  return videoLearning.searchVideos(payload);
}

export async function getRecommendationsAction(
  payload: GetRecommendationsPayload,
) {
  const videoLearning = await videoLearningService(cookies());
  return videoLearning.getRecommendations(payload);
}

export async function getVideoAction(payload: GetVideoPayload) {
  const videoLearning = await videoLearningService(cookies());
  return videoLearning.getVideo(payload);
}

export async function getCommentsAction(payload: GetCommentsPayload) {
  const videoLearning = await videoLearningService(cookies());
  return videoLearning.getComments(payload);
}

export async function createCommentAction(payload: CreateCommentPayload) {
  const videoLearning = await videoLearningService(cookies());
  return videoLearning.createComment(payload);
}

export async function deleteCommentAction(payload: DeleteCommentPayload) {
  const videoLearning = await videoLearningService(cookies());
  return videoLearning.deleteComment(payload);
}

export async function getOwnVideosAction(payload: GetOwnVideosPayload) {
  const videoLearning = await videoLearningService(cookies());
  return videoLearning.getOwnVideos(payload);
}

export async function initialUploadAction(payload: UploadVideoPayload) {
  const videoLearning = await videoLearningService(cookies());
  return videoLearning.initialUpload(payload);
}

export async function completeUploadAction(payload: CompleteUploadPayload) {
  const videoLearning = await videoLearningService(cookies());
  return videoLearning.completeUpload(payload);
}

export async function uploadThumbnailAction(payload: UploadThumbnailPayload) {
  const videoLearning = await videoLearningService(cookies());
  return videoLearning.uploadThumbnail(payload);
}

export async function toggleVideoLikeAction(payload: ToggleVideoLikePayload) {
  const videoLearning = await videoLearningService(cookies());
  return videoLearning.toggleVideoLike(payload);
}

export async function deleteVideoAction(payload: DeleteVideoPaylod) {
  const videoLearning = await videoLearningService(cookies());
  return videoLearning.deleteVideo(payload);
}

export async function getVideosByCategoryAction(
  payload: GetVideosByCategoryPayload,
) {
  const videoLearning = await videoLearningService(cookies());
  return videoLearning.getVideosByCategory(payload);
}

export async function getSimilarVideosAction(payload: GetSimilarVideosPayload) {
  const videoLearning = await videoLearningService(cookies());
  return videoLearning.getSimilarVideos(payload);
}

export async function getAllCategoriesAction() {
  const videoLearning = await videoLearningService(cookies());
  return videoLearning.getAllCategories();
}

export async function getCategoryAction(payload: GetCategoryPayload) {
  const videoLearning = await videoLearningService(cookies());
  return videoLearning.getCategory(payload);
}

export async function createCategoryAction(payload: CreateCategoryPayload) {
  const videoLearning = await videoLearningService(cookies());
  return videoLearning.createCategory(payload);
}

export async function getAllTagsAction() {
  const videoLearning = await videoLearningService(cookies());
  return videoLearning.getAllTags();
}

export async function getTagAction(payload: GetTagPayload) {
  const videoLearning = await videoLearningService(cookies());
  return videoLearning.getTag(payload);
}

export async function createTagAction(payload: CreateTagPayload) {
  const videoLearning = await videoLearningService(cookies());
  return videoLearning.createTag(payload);
}

export async function getTagsByVideoAction(payload: GetTagsByVideoPayload) {
  const videoLearning = await videoLearningService(cookies());
  return videoLearning.getTagsByVideo(payload);
}

export async function getVideosGroupedByCategory(
  payload: GetVideosGroupedByCategoryPayload,
) {
  const videoLearning = await videoLearningService(cookies());
  return videoLearning.getVideosGroupedByCategory(payload);
}
