import { getMyTestsAction } from '@/actions/knowledge/actions';
import Link from 'next/link';
import { Button } from '@/ui/button';
import { Plus, Edit, Trash2, Eye, FileText } from 'lucide-react';
import { notFound } from 'next/navigation';

export default async function TestsPage() {
  const testsResult = await getMyTestsAction();

  if (!testsResult.success) {
    return (
      <div className="container mx-auto px-4 py-8">
        <div className="text-center">
          <h1 className="text-2xl font-bold mb-4 text-red-600">Error</h1>
          <p>Error loading tests: {testsResult.error.message}</p>
        </div>
      </div>
    );
  }

  const tests = testsResult.data.tests;

  return (
    <div className="container mx-auto px-4 py-8">
      <div className="flex items-center justify-between mb-8">
        <div>
          <h1 className="text-3xl font-bold">My Tests</h1>
          <p className="text-gray-600 mt-2">Manage your tests and questions</p>
        </div>
        <Button asChild>
          <Link href="/teacher/tests/new">
            <Plus className="h-4 w-4 mr-2" />
            Create Test
          </Link>
        </Button>
      </div>

      {tests.length === 0 ? (
        <div className="text-center py-12">
          <FileText className="h-16 w-16 mx-auto text-gray-400 mb-4" />
          <h2 className="text-xl font-semibold text-gray-600 mb-2">
            No tests created yet
          </h2>
          <p className="text-gray-500 mb-4">
            Start by creating your first test
          </p>
          <Button asChild>
            <Link href="/teacher/tests/new">
              <Plus className="h-4 w-4 mr-2" />
              Create Your First Test
            </Link>
          </Button>
        </div>
      ) : (
        <div className="grid gap-6 md:grid-cols-2 lg:grid-cols-3">
          {tests.map((test) => (
            <div
              key={test.id}
              className="border border-gray-200 rounded-lg p-6 hover:shadow-md transition-shadow"
            >
              <div className="flex items-start justify-between mb-4">
                <h3 className="text-lg font-semibold line-clamp-2">
                  {test.title}
                </h3>
              </div>

              <div className="space-y-2 text-sm text-gray-600 mb-4">
                <p>Questions: {test.question_count || 0}</p>
                <p>Created: {new Date(test.created_at * 1000).toLocaleDateString()}</p>
              </div>

              <div className="flex gap-2">
                <Button variant="outline" size="sm" asChild>
                  <Link href={`/teacher/tests/${test.id}`}>
                    <Eye className="h-4 w-4 mr-1" />
                    View
                  </Link>
                </Button>
                <Button variant="outline" size="sm" asChild>
                  <Link href={`/teacher/tests/${test.id}/edit`}>
                    <Edit className="h-4 w-4 mr-1" />
                    Edit
                  </Link>
                </Button>
                <Button variant="outline" size="sm" asChild>
                  <Link href={`/teacher/tests/${test.id}/questions`}>
                    <FileText className="h-4 w-4 mr-1" />
                    Questions
                  </Link>
                </Button>
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
}
