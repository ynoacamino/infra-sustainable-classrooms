import { getAvailableTestsAction, getMySubmissionsAction } from '@/actions/knowledge/actions';
import Link from 'next/link';
import { Button } from '@/ui/button';
import { FileText, Users, CheckCircle, Eye } from 'lucide-react';

export default async function StudentsTestsPage() {
  const [testsResult, submissionsResult] = await Promise.all([
    getAvailableTestsAction(),
    getMySubmissionsAction()
  ]);

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
  const submissions = submissionsResult.success ? submissionsResult.data.submissions : [];
  
  // Create a map of submissions by test_id for quick lookup
  const submissionsByTestId = submissions.reduce((acc, submission) => {
    acc[submission.test_id] = submission;
    return acc;
  }, {} as Record<number, any>);

  return (
    <div className="container mx-auto px-4 py-8">
      <div className="flex items-center justify-between mb-8">
        <div>
          <h1 className="text-3xl font-bold">Tests</h1>
          <p className="text-gray-600 mt-2">Take tests and view your results</p>
        </div>
        {submissions.length > 0 && (
          <Button variant="outline" asChild>
            <Link href="/dashboard/tests/results">
              View All Results
            </Link>
          </Button>
        )}
      </div>

      {tests.length === 0 ? (
        <div className="text-center py-12">
          <FileText className="h-16 w-16 mx-auto text-gray-400 mb-4" />
          <h2 className="text-xl font-semibold text-gray-600 mb-2">
            No tests available
          </h2>
          <p className="text-gray-500">
            Check back later for new tests from your teachers
          </p>
        </div>
      ) : (
        <div className="grid gap-6 md:grid-cols-2 lg:grid-cols-3">
          {tests.map((test) => {
            const submission = submissionsByTestId[test.id];
            const hasSubmission = !!submission;
            
            return (
              <div
                key={test.id}
                className="border border-gray-200 rounded-lg p-6 hover:shadow-md transition-shadow"
              >
                <div className="flex items-start justify-between mb-4">
                  <h3 className="text-lg font-semibold line-clamp-2">
                    {test.title}
                  </h3>
                  {hasSubmission && (
                    <div className="ml-2">
                      <CheckCircle className="h-5 w-5 text-green-500" />
                    </div>
                  )}
                </div>

                <div className="space-y-2 text-sm text-gray-600 mb-4">
                  <div className="flex items-center">
                    <FileText className="h-4 w-4 mr-2" />
                    <span>Questions: {test.question_count || 0}</span>
                  </div>
                  <div className="flex items-center">
                    <Users className="h-4 w-4 mr-2" />
                    <span>Created: {new Date(test.created_at * 1000).toLocaleDateString()}</span>
                  </div>
                  {hasSubmission && (
                    <div className="flex items-center">
                      <CheckCircle className="h-4 w-4 mr-2 text-green-500" />
                      <span className="text-green-600 font-medium">Score: {submission.score}%</span>
                    </div>
                  )}
                </div>

                <div className="flex gap-2">
                  {hasSubmission ? (
                    <Button variant="outline" size="sm" asChild className="flex-1">
                      <Link href={`/dashboard/tests/results/${submission.id}`}>
                        <Eye className="h-4 w-4 mr-1" />
                        View Result
                      </Link>
                    </Button>
                  ) : (
                    <Button size="sm" asChild className="flex-1">
                      <Link href={`/dashboard/tests/${test.id}/take`}>
                        Take Test
                      </Link>
                    </Button>
                  )}
                </div>

                {hasSubmission && (
                  <div className="mt-3 text-xs text-gray-500">
                    Completed on: {new Date(submission.submitted_at * 1000).toLocaleDateString()}
                  </div>
                )}
              </div>
            );
          })}
        </div>
      )}
    </div>
  );
}
