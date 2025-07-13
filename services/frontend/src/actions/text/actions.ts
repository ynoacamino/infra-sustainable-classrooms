'use server';

import { textService } from '@/services/text/service';
import type {
  CreateCoursePayloadSchema,
  UpdateCoursePayloadSchema,
  DeleteCoursePayloadSchema,
  CreateSectionPayloadSchema,
  UpdateSectionPayloadSchema,
  DeleteSectionPayloadSchema,
  CreateArticlePayloadSchema,
  UpdateArticlePayloadSchema,
  DeleteArticlePayloadSchema,
} from '@/types/text/schemas/payload';
import { cookies } from 'next/headers';
import type z from 'zod';

// === COURSE ACTIONS ===
export async function createCourseAction(
  payload: z.infer<typeof CreateCoursePayloadSchema>,
) {
  const text = await textService(cookies());
  return text.createCourse(payload);
}

export async function updateCourseAction(
  payload: z.infer<typeof UpdateCoursePayloadSchema>,
) {
  const text = await textService(cookies());
  return text.updateCourse(payload);
}

export async function deleteCourseAction(courseId: number) {
  const text = await textService(cookies());
  return text.deleteCourse({ course_id: courseId });
}

// === SECTION ACTIONS ===
export async function createSectionAction(
  payload: z.infer<typeof CreateSectionPayloadSchema>,
) {
  const text = await textService(cookies());
  return text.createSection(payload);
}

export async function updateSectionAction(
  payload: z.infer<typeof UpdateSectionPayloadSchema>,
) {
  const text = await textService(cookies());
  return text.updateSection(payload);
}

export async function deleteSectionAction(sectionId: number) {
  const text = await textService(cookies());
  return text.deleteSection({ section_id: sectionId });
}

// === ARTICLE ACTIONS ===
export async function createArticleAction(
  payload: z.infer<typeof CreateArticlePayloadSchema>,
) {
  const text = await textService(cookies());
  return text.createArticle(payload);
}

export async function updateArticleAction(
  payload: z.infer<typeof UpdateArticlePayloadSchema>,
) {
  const text = await textService(cookies());
  return text.updateArticle(payload);
}

export async function deleteArticleAction(articleId: number) {
  const text = await textService(cookies());
  return text.deleteArticle({ article_id: articleId });
}
