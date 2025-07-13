import { getMySubmissionsAction } from '@/actions/knowledge/actions';
import Link from 'next/link';
import { Button } from '@/ui/button';
import { ArrowLeft, FileText, CheckCircle, XCircle, Eye } from 'lucide-react';

export default async function MySubmissionsPage() {
  const submissionsResult = await getMySubmissionsAction();

  if (!submissionsResult.success) {
    return (
      <div className="container mx-auto px-4 py-8">
        <div className="text-center">
          <h1 className="text-2xl font-bold mb-4 text-red-600">Error</h1>
          <p>Error loading submissions: {submissionsResult.error.message}</p>
        </div>
      </div>
    );
  }

  const submissions = submissionsResult.data.submissions;

  return (
    <div className="container mx-auto px-4 py-8">
      <div className="flex items-center gap-4 mb-8">
        <Button variant="outline" size="sm" asChild>
          <Link href="/dashboard/tests">
            <ArrowLeft className="h-4 w-4 mr-2" />
            Back to Tests
          </Link>
        </Button>
        <div>
          <h1 className="text-3xl font-bold">My Test Results</h1>
          <p className="text-gray-600 mt-2">View your completed test submissions</p>
        </div>
      </div>

      {submissions.length === 0 ? (
        <div className="text-center py-12">
          <FileText className="h-16 w-16 mx-auto text-gray-400 mb-4" />
          <h2 className="text-xl font-semibold text-gray-600 mb-2">
            No submissions yet
          </h2>
          <p className="text-gray-500 mb-4">
            Complete some tests to see your results here
          </p>
          <Button asChild>
            <Link href="/dashboard/tests">
              Take a Test
            </Link>
          </Button>
        </div>
      ) : (
        <div className="space-y-4">
          {submissions.map((submission) => (
            <div
              key={submission.id}
              className="border border-gray-200 rounded-lg p-6 hover:shadow-md transition-shadow"
            >
              <div className="flex items-start justify-between mb-4">
                <div>
                  <h3 className="text-lg font-semibold mb-2">
                    Test #{submission.test_id}
                  </h3>
                  <div className="flex items-center gap-4 text-sm text-gray-600">
                    <span>Submitted: {new Date(submission.submitted_at * 1000).toLocaleString()}</span>
                    {submission.score !== undefined && (
                      <div className="flex items-center">
                        {submission.score >= 70 ? (
                          <CheckCircle className="h-4 w-4 mr-1 text-green-500" />
                        ) : (
                          <XCircle className="h-4 w-4 mr-1 text-red-500" />
                        )}
                        <span className={`font-medium ${
                          submission.score >= 70 ? 'text-green-600' : 'text-red-600'
                        }`}>
                          Score: {submission.score}%
                        </span>
                      </div>
                    )}
                  </div>
                </div>
              </div>

              <div className="flex gap-2">
                <Button variant="outline" size="sm" asChild>
                  <Link href={`/dashboard/tests/results/${submission.id}`}>
                    <Eye className="h-4 w-4 mr-1" />
                    View Details
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
