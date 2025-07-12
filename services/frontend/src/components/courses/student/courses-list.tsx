'use client';

import { useCourses } from '@/hooks/courses/use-courses';
import { type Course as CourseProps } from '@/types/courses/courses';
import { Link } from '@/ui/link';
import Image from 'next/image';

function CourseSkeleton() {
  return (
    <div className="grid grid-cols-2 gap-4 animate-pulse">
      <div className="flex flex-col gap-2">
        <div className="h-6 bg-gray-300 rounded w-3/4"></div>
        <div className="h-4 bg-gray-300 rounded w-full"></div>
        <div className="h-4 bg-gray-300 rounded w-1/2"></div>
      </div>
      <div className="bg-gray-300 aspect-video max-w-sm w-full rounded-[12px] justify-self-end"></div>
    </div>
  );
}

function Course({
  description,
  imageUrl,
  title,
  id,
}: Omit<CourseProps, 'status' | 'date' | 'modules'>) {
  return (
    <div className="grid grid-cols-2">
      <div>
        <h3 className="font-bold">{title}</h3>
        <p className="text-sm">{description}</p>
        <Link
          href={`/dashboard/courses/${id}`}
          className="mt-3"
          variant={'secondary'}
        >
          View Course
        </Link>
      </div>
      <Image
        src={imageUrl}
        alt={title}
        width={500}
        height={300}
        className="w-full aspect-video object-cover rounded-[12px] max-w-sm justify-self-end bg-gray-300"
      />
    </div>
  );
}
export function Courses() {
  const { courses, isLoading } = useCourses();

  if (isLoading || !courses) {
    return (
      <div className="flex flex-col gap-4">
        <CourseSkeleton />
        <CourseSkeleton />
        <CourseSkeleton />
        <CourseSkeleton />
      </div>
    );
  }

  return (
    <div className="flex flex-col gap-10">
      {courses.map((course) => (
        <Course
          key={course.id}
          id={course.id}
          title={course.title}
          description={course.description}
          imageUrl={course.imageUrl}
        />
      ))}
    </div>
  );
}
