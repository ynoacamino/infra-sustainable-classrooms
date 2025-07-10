import type {
  CourseModuleTypes,
  CourseStates,
} from '@/lib/courses/enums/courses';

export type CourseStatus = (typeof CourseStates)[keyof typeof CourseStates];
export type CourseModuleType =
  (typeof CourseModuleTypes)[keyof typeof CourseModuleTypes];

export interface CourseModule {
  id: string;
  idCourse: string;
  title: string;
  description: string;
  type: CourseModuleType;
}

export interface Course {
  id: string;
  title: string;
  description: string;
  imageUrl: string;
  status: CourseStatus;
  modules: CourseModule[];
  date: Date;
}
