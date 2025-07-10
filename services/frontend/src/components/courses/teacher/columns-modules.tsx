import { CourseModuleType } from '@/components/courses/teacher/course-module-type';
import type {
  CourseModule,
  CourseModuleType as ModuleType,
} from '@/types/courses/courses';
import { Button } from '@/ui/button';
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuTrigger,
} from '@/ui/dropdown-menu';
import { Link } from '@/ui/link';
import type { ColumnDef } from '@tanstack/react-table';
import { MoreHorizontal } from 'lucide-react';

export const columns: ColumnDef<CourseModule>[] = [
  {
    accessorKey: 'title',
    header: 'Module Name',
  },
  {
    accessorKey: 'type',
    header: 'Type',
    cell: ({ row }) => {
      const type = row.getValue<ModuleType>('type');
      return <CourseModuleType type={type} />;
    },
  },
  {
    id: 'actions',
    cell: ({ row }) => {
      const moduleData = row.original;
      return (
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <Button variant="ghost" size="icon">
              <span className="sr-only">Open menu</span>
              <MoreHorizontal className="size-4" />
            </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent align="end">
            <DropdownMenuLabel>Actions</DropdownMenuLabel>
            <DropdownMenuItem asChild>
              <Link
                variant="ghost"
                href={`/teacher/courses/${moduleData.idCourse}/modules/${moduleData.id}`}
                className="justify-start"
              >
                View
              </Link>
            </DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu>
      );
    },
  },
];
