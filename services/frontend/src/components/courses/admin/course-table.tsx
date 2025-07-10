import { columns } from '@/components/courses/admin/columns-courses';
import { useCourses } from '@/hooks/courses/use-courses';
import type { Course } from '@/types/courses/courses';
import { DataTable } from '@/ui/data-table';
import { Skeleton } from '@/ui/skeleton';
import { useMemo } from 'react';

function CoursesTable() {
  const { courses, isLoading } = useCourses();
  const data = useMemo<Course[]>(
    () =>
      isLoading || !courses
        ? Array.from({ length: 10 }, () => ({}) as Course)
        : courses,
    [isLoading, courses],
  );
  const tableColumns = useMemo(
    () =>
      isLoading
        ? columns.map((column) => ({
            ...column,
            cell: () => <Skeleton className="h-6" />,
          }))
        : columns,
    [isLoading],
  );

  return <DataTable data={data} columns={tableColumns} />;
}

export { CoursesTable };
