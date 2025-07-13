import { getTestFormAction } from '@/actions/knowledge/actions';
import { SubmitTestForm } from '@/components/knowledge/forms';
import Link from 'next/link';
import { Button } from '@/ui/button';
import { ArrowLeft, Clock, FileText } from 'lucide-react';
import { notFound } from 'next/navigation';

interface TakeTestPageProps {
  params: Promise<{ testId: string }>;
}

export default async function TakeTestPage({ params }: TakeTestPageProps) {
  const resolvedParams = await params;
  const testId = parseInt(resolvedParams.testId);

  if (isNaN(testId)) {
    notFound();
  }

  const testFormResult = await getTestFormAction(testId);

  if (!testFormResult.success) {
    if (testFormResult.error.status === 404) {
      notFound();
    }
    return (
      <div className="container mx-auto px-4 py-8">
        <div className="text-center">
          <h1 className="text-2xl font-bold mb-4 text-red-600">Error</h1>
          <p>Error loading test: {testFormResult.error.message}</p>
        </div>
      </div>
    );
  }

  const { test, questions } = testFormResult.data;

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
          <h1 className="text-3xl font-bold">{test.title}</h1>
          <p className="text-gray-600 mt-2">Take your time and answer all questions</p>
        </div>
      </div>

      <div className="mb-6 p-4 bg-blue-50 border border-blue-200 rounded-lg">
        <div className="flex items-center gap-4 mb-2">
          <div className="flex items-center">
            <FileText className="h-5 w-5 mr-2 text-blue-600" />
            <span className="text-blue-800 font-medium">
              {questions.length} Questions
            </span>
          </div>
          <div className="flex items-center">
            <Clock className="h-5 w-5 mr-2 text-blue-600" />
            <span className="text-blue-800 font-medium">
              No time limit
            </span>
          </div>
        </div>
        <p className="text-blue-700 text-sm mb-2">
          Read each question carefully and select the best answer. You can review and change your answers before submitting.
        </p>
        <details className="text-blue-700 text-sm">
          <summary className="cursor-pointer hover:text-blue-800 font-medium">
            Keyboard shortcuts
          </summary>
          <div className="mt-2 space-y-1 text-xs bg-blue-100 p-2 rounded">
            <div><kbd className="px-1 py-0.5 bg-white rounded border">←</kbd> Previous question</div>
            <div><kbd className="px-1 py-0.5 bg-white rounded border">→</kbd> Next question</div>
            <div><kbd className="px-1 py-0.5 bg-white rounded border">1-4</kbd> Select answer A-D</div>
          </div>
        </details>
      </div>

      <div className="max-w-4xl mx-auto">
        <SubmitTestForm test={test} questions={questions} />
      </div>
    </div>
  );
}
