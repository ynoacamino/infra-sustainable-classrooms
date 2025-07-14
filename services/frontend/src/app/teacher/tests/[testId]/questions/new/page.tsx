import { AddQuestionForm } from '@/components/knowledge/forms';
import { getTestAction } from '@/actions/knowledge/actions';
import Link from 'next/link';
import { Button } from '@/ui/button';
import { ArrowLeft } from 'lucide-react';
import { notFound } from 'next/navigation';

interface NewQuestionPageProps {
  params: Promise<{ testId: string }>;
}

export default async function NewQuestionPage({
  params,
}: NewQuestionPageProps) {
  const resolvedParams = await params;
  const testId = parseInt(resolvedParams.testId);

  if (isNaN(testId)) {
    notFound();
  }

  const testResult = await getTestAction({ id: testId });

  if (!testResult.success) {
    if (testResult.error.status === 404) {
      notFound();
    }
    return (
      <div className="container mx-auto px-4 py-8">
        <div className="text-center">
          <h1 className="text-2xl font-bold mb-4 text-red-600">Error</h1>
          <p>Error loading test: {testResult.error.message}</p>
        </div>
      </div>
    );
  }

  const test = testResult.data;

  return (
    <div className="container mx-auto px-4 py-8">
      <div className="flex items-center gap-4 mb-8">
        <Button variant="outline" size="sm" asChild>
          <Link href={`/teacher/tests/${testId}/questions`}>
            <ArrowLeft className="h-4 w-4 mr-2" />
            Back to Questions
          </Link>
        </Button>
        <div>
          <h1 className="text-3xl font-bold">Add Question</h1>
          <p className="text-gray-600 mt-2">
            Add a new question to &ldquo;{test.test.title}&rdquo;
          </p>
        </div>
      </div>

      <div className="max-w-3xl mx-auto">
        <div className="bg-white rounded-lg border border-gray-200 p-6">
          <div className="mb-6 p-4 bg-blue-50 border border-blue-200 rounded-lg">
            <h3 className="font-medium text-blue-900 mb-2">
              Question Guidelines
            </h3>
            <ul className="text-blue-700 text-sm space-y-1">
              <li>• Write clear and concise questions</li>
              <li>• Ensure all answer options are plausible</li>
              <li>• Make sure only one answer is clearly correct</li>
              <li>• Use simple language appropriate for your students</li>
            </ul>
          </div>

          <AddQuestionForm testId={testId} />
        </div>
      </div>
    </div>
  );
}
