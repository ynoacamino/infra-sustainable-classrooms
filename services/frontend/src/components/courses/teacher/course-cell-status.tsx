import { CourseStates } from '@/lib/courses/enums/courses';
import { cn } from '@/lib/shared/utils';
import type { CourseStatus } from '@/types/courses/courses';

interface CoursesCellStatusProps {
  status: CourseStatus;
}

function CoursesCellStatus({ status }: CoursesCellStatusProps) {
  return (
    <div
      className={cn(
        'rounded-l-full rounded-r-full inline-flex w-full justify-center font-medium py-1',
        {
          'bg-red-300/80': status === CourseStates.INACTIVE,
          'bg-green-300/80': status === CourseStates.ACTIVE,
        },
      )}
    >
      {status}
    </div>
  );
}

export { CoursesCellStatus };
