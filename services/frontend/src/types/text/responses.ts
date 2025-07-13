import type {
  Course,
  Section,
  Article,
  SimpleResponse,
} from '@/types/text/models';

// Course response types
export type CreateCourseResponse = SimpleResponse;
export type GetCourseResponse = Course;
export type ListCoursesResponse = Course[];
export type UpdateCourseResponse = SimpleResponse;
export type DeleteCourseResponse = SimpleResponse;

// Section response types
export type CreateSectionResponse = SimpleResponse;
export type GetSectionResponse = Section;
export type ListSectionsResponse = Section[];
export type UpdateSectionResponse = SimpleResponse;
export type DeleteSectionResponse = SimpleResponse;

// Article response types
export type CreateArticleResponse = SimpleResponse;
export type GetArticleResponse = Article;
export type ListArticlesResponse = Article[];
export type UpdateArticleResponse = SimpleResponse;
export type DeleteArticleResponse = SimpleResponse;
