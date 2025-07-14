import {
  getTestAction,
  getTestQuestionsAction,
} from '@/actions/knowledge/actions';
import Link from 'next/link';
import { Button } from '@/ui/button';
import { ArrowLeft, Plus, Edit, FileText } from 'lucide-react';
import { notFound } from 'next/navigation';
import { DeleteQuestionButton } from '@/components/knowledge/delete-question-button';

interface QuestionsPageProps {
  params: Promise<{ testId: string }>;
}

export default async function QuestionsPage({ params }: QuestionsPageProps) {
  const resolvedParams = await params;
  const testId = parseInt(resolvedParams.testId);

  if (isNaN(testId)) {
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
      <div className="container mx-auto px-4 py-8">
        <div className="text-center">
          <h1 className="text-2xl font-bold mb-4 text-red-600">Error</h1>
          <p>Error loading test: {testResult.error.message}</p>
        </div>
      </div>
    );
  }

  if (!questionsResult.success) {
    return (
      <div className="container mx-auto px-4 py-8">
        <div className="text-center">
          <h1 className="text-2xl font-bold mb-4 text-red-600">Error</h1>
          <p>Error loading questions: {questionsResult.error.message}</p>
        </div>
      </div>
    );
  }

  const test = testResult.data.test;
  const questions = questionsResult.data.questions;

  return (
    <div className="container mx-auto px-4 py-8">
      <div className="flex items-center justify-between mb-8">
        <div className="flex items-center gap-4">
          <Button variant="outline" size="sm" asChild>
            <Link href={`/teacher/tests/${testId}`}>
              <ArrowLeft className="h-4 w-4 mr-2" />
              Back to Test
            </Link>
          </Button>
          <div>
            <h1 className="text-3xl font-bold">
              Questions for &ldquo;{test.title}&rdquo;
            </h1>
            <p className="text-gray-600 mt-2">Manage test questions</p>
          </div>
        </div>

        <Button asChild>
          <Link href={`/teacher/tests/${testId}/questions/new`}>
            <Plus className="h-4 w-4 mr-2" />
            Add Question
          </Link>
        </Button>
      </div>

      {questions.length === 0 ? (
        <div className="text-center py-12">
          <FileText className="h-16 w-16 mx-auto text-gray-400 mb-4" />
          <h2 className="text-xl font-semibold text-gray-600 mb-2">
            No questions added yet
          </h2>
          <p className="text-gray-500 mb-4">
            Start by adding your first question to this test
          </p>
          <Button asChild>
            <Link href={`/teacher/tests/${testId}/questions/new`}>
              <Plus className="h-4 w-4 mr-2" />
              Add First Question
            </Link>
          </Button>
        </div>
      ) : (
        <div className="space-y-6">
          <div className="bg-blue-50 border border-blue-200 rounded-lg p-4 mb-6">
            <div className="flex items-center justify-between">
              <div>
                <h3 className="font-medium text-blue-900">Test Summary</h3>
                <p className="text-blue-700">
                  Total Questions: {questions.length}
                </p>
              </div>
              <div className="text-right">
                <p className="text-sm text-blue-600">
                  {questions.length > 0
                    ? 'Test is ready for students'
                    : 'Add questions to activate test'}
                </p>
              </div>
            </div>
          </div>

          <div className="space-y-4">
            {questions.map((question, index) => (
              <div
                key={question.id}
                className="bg-white border border-gray-200 rounded-lg p-6 hover:shadow-sm transition-shadow"
              >
                <div className="flex items-start justify-between mb-4">
                  <div className="flex-1">
                    <h3 className="text-lg font-semibold mb-2">
                      Question {index + 1}
                    </h3>
                    <p className="text-gray-800 mb-4">
                      {question.question_text}
                    </p>
                  </div>
                  <div className="flex gap-2 ml-4">
                    <Button variant="outline" size="sm" asChild>
                      <Link
                        href={`/teacher/tests/${testId}/questions/${question.id}/edit`}
                      >
                        <Edit className="h-4 w-4 mr-1" />
                        Edit
                      </Link>
                    </Button>
                    <DeleteQuestionButton
                      testId={testId}
                      questionId={question.id}
                      questionText={question.question_text}
                    />
                  </div>
                </div>

                <div className="grid grid-cols-1 md:grid-cols-2 gap-3 mb-4">
                  <div
                    className={`p-3 rounded-lg border ${
                      question.correct_answer === 0
                        ? 'bg-green-50 border-green-200'
                        : 'bg-gray-50 border-gray-200'
                    }`}
                  >
                    <span className="font-medium">A) </span>
                    <span
                      className={
                        question.correct_answer === 0
                          ? 'text-green-800'
                          : 'text-gray-700'
                      }
                    >
                      {question.option_a}
                    </span>
                    {question.correct_answer === 0 && (
                      <span className="ml-2 text-green-600 text-sm font-medium">
                        ✓ Correct
                      </span>
                    )}
                  </div>

                  <div
                    className={`p-3 rounded-lg border ${
                      question.correct_answer === 1
                        ? 'bg-green-50 border-green-200'
                        : 'bg-gray-50 border-gray-200'
                    }`}
                  >
                    <span className="font-medium">B) </span>
                    <span
                      className={
                        question.correct_answer === 1
                          ? 'text-green-800'
                          : 'text-gray-700'
                      }
                    >
                      {question.option_b}
                    </span>
                    {question.correct_answer === 1 && (
                      <span className="ml-2 text-green-600 text-sm font-medium">
                        ✓ Correct
                      </span>
                    )}
                  </div>

                  <div
                    className={`p-3 rounded-lg border ${
                      question.correct_answer === 2
                        ? 'bg-green-50 border-green-200'
                        : 'bg-gray-50 border-gray-200'
                    }`}
                  >
                    <span className="font-medium">C) </span>
                    <span
                      className={
                        question.correct_answer === 2
                          ? 'text-green-800'
                          : 'text-gray-700'
                      }
                    >
                      {question.option_c}
                    </span>
                    {question.correct_answer === 2 && (
                      <span className="ml-2 text-green-600 text-sm font-medium">
                        ✓ Correct
                      </span>
                    )}
                  </div>

                  <div
                    className={`p-3 rounded-lg border ${
                      question.correct_answer === 3
                        ? 'bg-green-50 border-green-200'
                        : 'bg-gray-50 border-gray-200'
                    }`}
                  >
                    <span className="font-medium">D) </span>
                    <span
                      className={
                        question.correct_answer === 3
                          ? 'text-green-800'
                          : 'text-gray-700'
                      }
                    >
                      {question.option_d}
                    </span>
                    {question.correct_answer === 3 && (
                      <span className="ml-2 text-green-600 text-sm font-medium">
                        ✓ Correct
                      </span>
                    )}
                  </div>
                </div>

                <div className="text-sm text-gray-500">
                  Order: {question.question_order}
                </div>
              </div>
            ))}
          </div>
        </div>
      )}
    </div>
  );
}
