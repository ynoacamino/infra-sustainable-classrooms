import {
  getAllCategoriesAction,
  getAllTagsAction,
  getCategoryAction,
  getCommentsAction,
  getOwnVideosAction,
  getRecommendationsAction,
  getSimilarVideosAction,
  getTagAction,
  getTagsByVideoAction,
  getVideoAction,
  getVideosByCategoryAction,
  getVideosGroupedByCategory,
  searchVideosAction,
} from '@/actions/video_learning/actions';
import { formatSWRResponse } from '@/lib/shared/swr/utils';
import type {
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
} from '@/types/video_learning/payload';
import useSWR from 'swr';

export const useGetComments = (payload: GetCommentsPayload) => {
  const response = useSWR(['get-comments', payload], ([, p]) =>
    getCommentsAction(p),
  );
  return formatSWRResponse(response);
};

export const useGetVideo = (payload: GetVideoPayload) => {
  const response = useSWR(['get-video', payload], ([, p]) => getVideoAction(p));
  return formatSWRResponse(response);
};

export const useSearchVideos = (payload: SearchVideosPayload) => {
  const response = useSWR(['search-videos', payload], ([, p]) =>
    searchVideosAction(p),
  );
  return formatSWRResponse(response);
};

export const useGetRecommendations = (payload: GetRecommendationsPayload) => {
  const response = useSWR(['get-recommendations', payload], ([, p]) =>
    getRecommendationsAction(p),
  );
  return formatSWRResponse(response);
};

export const useGetOwnVideos = (payload: GetOwnVideosPayload) => {
  const response = useSWR(['get-own-videos', payload], ([, p]) =>
    getOwnVideosAction(p),
  );
  return formatSWRResponse(response);
};

export const useGetVideosByCategory = (payload: GetVideosByCategoryPayload) => {
  const response = useSWR(['get-videos-by-category', payload], ([, p]) =>
    getVideosByCategoryAction(p),
  );
  return formatSWRResponse(response);
};

export const useGetSimilarVideos = (payload: GetSimilarVideosPayload) => {
  const response = useSWR(['get-similar-videos', payload], ([, p]) =>
    getSimilarVideosAction(p),
  );
  return formatSWRResponse(response);
};

export const useGetAllCategories = () => {
  const response = useSWR(['get-all-categories'], () =>
    getAllCategoriesAction(),
  );
  return formatSWRResponse(response);
};

export const useGetCategory = (payload: GetCategoryPayload) => {
  const response = useSWR(['get-category', payload], ([, p]) =>
    getCategoryAction(p),
  );
  return formatSWRResponse(response);
};

export const useGetAllTags = () => {
  const response = useSWR(['get-all-tags'], () => getAllTagsAction());
  return formatSWRResponse(response);
};

export const useGetTag = (payload: GetTagPayload) => {
  const response = useSWR(['get-tag', payload], () => getTagAction(payload));
  return formatSWRResponse(response);
};

export const useGetTagsByVideo = (payload: GetTagsByVideoPayload) => {
  const response = useSWR(['get-tags-by-video', payload], () =>
    getTagsByVideoAction(payload),
  );
  return formatSWRResponse(response);
};

export const useGetVideosGroupedByCategory = (
  payload: GetVideosGroupedByCategoryPayload,
) => {
  const response = useSWR(['get-videos-grouped-by-category', payload], () =>
    getVideosGroupedByCategory(payload),
  );
  return formatSWRResponse(response);
};
