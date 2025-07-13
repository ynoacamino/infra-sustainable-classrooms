import {
  ArticleSchema,
  CourseSchema,
  SectionSchema,
} from '@/types/text/schemas/models';

// Course payload schemas
export const CreateCoursePayloadSchema = CourseSchema.omit({
  id: true,
  created_at: true,
  updated_at: true,
});

export const GetCoursePayloadSchema = CourseSchema.pick({
  id: true,
});

export const DeleteCoursePayloadSchema = CourseSchema.pick({
  id: true,
});

export const UpdateCoursePayloadSchema = CourseSchema.omit({
  created_at: true,
  updated_at: true,
}).partial({
  description: true,
  imageUrl: true,
  title: true,
});

// Section payload schemas
export const CreateSectionPayloadSchema = SectionSchema.omit({
  id: true,
  created_at: true,
  updated_at: true,
});

export const GetSectionPayloadSchema = SectionSchema.pick({
  id: true,
});

export const ListSectionsPayloadSchema = SectionSchema.pick({
  course_id: true,
});

export const UpdateSectionPayloadSchema = SectionSchema.omit({
  course_id: true,
  created_at: true,
  updated_at: true,
}).partial({
  title: true,
  description: true,
  order: true,
});

export const DeleteSectionPayloadSchema = SectionSchema.pick({
  id: true,
});

// Article payload schemas
export const CreateArticlePayloadSchema = ArticleSchema.omit({
  id: true,
  created_at: true,
  updated_at: true,
});

export const GetArticlePayloadSchema = ArticleSchema.pick({
  id: true,
});

export const ListArticlesPayloadSchema = ArticleSchema.pick({
  section_id: true,
});

export const UpdateArticlePayloadSchema = ArticleSchema.omit({
  section_id: true,
  created_at: true,
  updated_at: true,
}).partial({
  title: true,
  content: true,
});

export const DeleteArticlePayloadSchema = ArticleSchema.pick({
  id: true,
});
