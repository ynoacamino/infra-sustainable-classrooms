import { COURSES } from '@/lib/courses/mock';
import type { Course } from '@/types/courses/courses';
import useSWR from 'swr';

export const getOneCourse = async ({
  courseId,
}: {
  courseId: string;
}): Promise<Course> => {
  const course = COURSES.find((c) => c.id === courseId);
  return new Promise<Course>((resolve, reject) => {
    setTimeout(() => {
      if (course) {
        resolve(course);
      }
      reject('The course did not found');
    }, 1000);
  });
};

export const getMyCourses = async (): Promise<Course[]> => {
  return new Promise((resolve) => {
    setTimeout(() => {
      resolve(COURSES);
    }, 1000);
  });
};

export const useCourses = () => {
  // const { data, isLoading, mutate } = useSWR('courses', getMyCourses);
  const { data, isLoading, mutate } = useSWR('courses', getMyCourses);

  return {
    courses: data,
    isLoading,
    mutate,
  };
};
