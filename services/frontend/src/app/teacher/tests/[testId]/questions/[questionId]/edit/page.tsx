import { notFound } from 'next/navigation';
import { UpdateQuestionForm } from '@/components/knowledge/forms';
import {
  getTestAction,
  getTestQuestionsAction,
} from '@/actions/knowledge/actions';
import type { Question } from '@/types/knowledge/models';
import Link from 'next/link';
import { Button } from '@/ui/button';
import { ArrowLeft } from 'lucide-react';

interface EditQuestionPageProps {
  params: Promise<{
    testId: string;
    questionId: string;
  }>;
}

export default async function EditQuestionPage({
  params,
}: EditQuestionPageProps) {
  const resolvedParams = await params;
  const testId = parseInt(resolvedParams.testId);
  const questionId = parseInt(resolvedParams.questionId);

  if (isNaN(testId) || isNaN(questionId)) {
    notFound();
  }

  // Get test details and questions in parallel
  const [testResult, questionsResult] = await Promise.all([
    getTestAction({ id: testId }),
    getTestQuestionsAction({ id: testId }),
  ]);

  if (!testResult.success) {
    if (testResult.error.status === 404) {
      notFound();
    }
    return (
      <div className="container mx-auto p-6">
        <div className="text-center">
          <h1 className="text-2xl font-bold mb-4 text-red-600">Error</h1>
          <p>Error loading test: {testResult.error.message}</p>
        </div>
      </div>
    );
  }

  if (!questionsResult.success) {
    return (
      <div className="container mx-auto p-6">
        <div className="text-center">
          <h1 className="text-2xl font-bold mb-4 text-red-600">Error</h1>
          <p>Error loading questions: {questionsResult.error.message}</p>
        </div>
      </div>
    );
  }

  const question = questionsResult.data.questions.find(
    (q: Question) => q.id === questionId,
  );

  if (!question) {
    notFound();
  }

  return (
    <div className="container mx-auto p-6">
      <div className="mb-6">
        <div className="flex items-center gap-4 mb-4">
          <Button variant="outline" size="sm" asChild>
            <Link href={`/teacher/tests/${testId}/questions`}>
              <ArrowLeft className="h-4 w-4 mr-2" />
              Back to Questions
            </Link>
          </Button>
          <div>
            <h1 className="text-3xl font-bold">Edit Question</h1>
            <p className="text-muted-foreground">
              Test: {testResult.data.test.title}
            </p>
          </div>
        </div>
      </div>

      <UpdateQuestionForm question={question} />
    </div>
  );
}
