import { columns } from '@/components/courses/teacher/columns-modules';
import type { CourseModule } from '@/types/courses/courses';
import { DataTable } from '@/ui/data-table';

interface ModulesTableProps {
  modules: CourseModule[];
}

function ModulesTable({ modules }: ModulesTableProps) {
  return <DataTable data={modules} columns={columns} />;
}

export { ModulesTable };
