'use server';

import { textService } from '@/services/text/service';
import type {
  CreateArticlePayload,
  CreateCoursePayload,
  CreateSectionPayload,
  DeleteArticlePayload,
  DeleteCoursePayload,
  DeleteSectionPayload,
  UpdateArticlePayload,
  UpdateCoursePayload,
  UpdateSectionPayload,
} from '@/types/text/payload';
import { cookies } from 'next/headers';

// === COURSE ACTIONS ===
export async function createCourseAction(payload: CreateCoursePayload) {
  const text = await textService(cookies());
  return text.createCourse(payload);
}

export async function updateCourseAction(payload: UpdateCoursePayload) {
  const text = await textService(cookies());
  return text.updateCourse(payload);
}

export async function deleteCourseAction(payload: DeleteCoursePayload) {
  const text = await textService(cookies());
  return text.deleteCourse(payload);
}

// === SECTION ACTIONS ===
export async function createSectionAction(payload: CreateSectionPayload) {
  const text = await textService(cookies());
  return text.createSection(payload);
}

export async function updateSectionAction(payload: UpdateSectionPayload) {
  const text = await textService(cookies());
  return text.updateSection(payload);
}

export async function deleteSectionAction(payload: DeleteSectionPayload) {
  const text = await textService(cookies());
  return text.deleteSection(payload);
}

// === ARTICLE ACTIONS ===
export async function createArticleAction(payload: CreateArticlePayload) {
  const text = await textService(cookies());
  return text.createArticle(payload);
}

export async function updateArticleAction(payload: UpdateArticlePayload) {
  const text = await textService(cookies());
  return text.updateArticle(payload);
}

export async function deleteArticleAction(payload: DeleteArticlePayload) {
  const text = await textService(cookies());
  return text.deleteArticle(payload);
}
