import Link from 'next/link';
import { Button } from '@/ui/button';

export default function TeacherCodelabPage() {
  return (
    <div className="container mx-auto px-4 py-8">
      <div className="flex justify-between items-center mb-8">
        <h1 className="text-3xl font-bold">Codelab Management</h1>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div className="p-6 border rounded-lg hover:shadow-md transition-shadow">
          <h2 className="text-xl font-semibold mb-4">Exercises</h2>
          <p className="text-gray-600 mb-4">
            Create and manage coding exercises for your students.
          </p>
          <Button asChild className="w-full">
            <Link href="/teacher/codelab/exercises">Manage Exercises</Link>
          </Button>
        </div>

        <div className="p-6 border rounded-lg hover:shadow-md transition-shadow">
          <h2 className="text-xl font-semibold mb-4">Test Cases</h2>
          <p className="text-gray-600 mb-4">
            Create test cases to validate student solutions automatically.
          </p>
          <Button asChild variant="outline" className="w-full">
            <Link href="/teacher/codelab/exercises">Manage Test Cases</Link>
          </Button>
        </div>
      </div>
    </div>
  );
}
