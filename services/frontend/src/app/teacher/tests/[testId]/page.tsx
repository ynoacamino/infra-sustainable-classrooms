import {
  getTestAction,
  getTestQuestionsAction,
} from '@/actions/knowledge/actions';
import Link from 'next/link';
import { Button } from '@/ui/button';
import { ArrowLeft, Edit, Plus, FileText, Users } from 'lucide-react';
import { notFound } from 'next/navigation';

interface TestPageProps {
  params: Promise<{ testId: string }>;
}

export default async function TestPage({ params }: TestPageProps) {
  const resolvedParams = await params;
  const testId = parseInt(resolvedParams.testId);

  if (isNaN(testId)) {
    notFound();
  }

  // Get test details and questions in parallel
  const [testResult, questionsResult] = await Promise.all([
    getTestAction(testId),
    getTestQuestionsAction(testId),
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

  const test = testResult.data;
  const questions = questionsResult.data.questions;

  return (
    <div className="container mx-auto px-4 py-8">
      <div className="flex items-center gap-4 mb-8">
        <Button variant="outline" size="sm" asChild>
          <Link href="/teacher/tests">
            <ArrowLeft className="h-4 w-4 mr-2" />
            Back to Tests
          </Link>
        </Button>
        <div className="flex-1">
          <h1 className="text-3xl font-bold">{test.title}</h1>
          <p className="text-gray-600 mt-2">Test overview and management</p>
        </div>
        <div className="flex gap-2">
          <Button variant="outline" asChild>
            <Link href={`/teacher/tests/${testId}/edit`}>
              <Edit className="h-4 w-4 mr-2" />
              Edit Test
            </Link>
          </Button>
          <Button asChild>
            <Link href={`/teacher/tests/${testId}/questions`}>
              <FileText className="h-4 w-4 mr-2" />
              Manage Questions
            </Link>
          </Button>
        </div>
      </div>

      <div className="grid gap-6 lg:grid-cols-3">
        {/* Test Information */}
        <div className="lg:col-span-2">
          <div className="bg-white rounded-lg border border-gray-200 p-6 mb-6">
            <h2 className="text-xl font-semibold mb-4">Test Information</h2>
            <div className="grid gap-4 sm:grid-cols-2">
              <div>
                <label className="text-sm font-medium text-gray-600">
                  Title
                </label>
                <p className="text-lg">{test.title}</p>
              </div>
              <div>
                <label className="text-sm font-medium text-gray-600">
                  Created
                </label>
                <p className="text-lg">
                  {new Date(test.created_at * 1000).toLocaleDateString()}
                </p>
              </div>
              <div>
                <label className="text-sm font-medium text-gray-600">
                  Total Questions
                </label>
                <p className="text-lg">{questions.length}</p>
              </div>
              <div>
                <label className="text-sm font-medium text-gray-600">
                  Creator ID
                </label>
                <p className="text-lg">{test.created_by}</p>
              </div>
            </div>
          </div>

          {/* Questions Preview */}
          <div className="bg-white rounded-lg border border-gray-200 p-6">
            <div className="flex items-center justify-between mb-4">
              <h2 className="text-xl font-semibold">Questions Preview</h2>
              <Button size="sm" asChild>
                <Link href={`/teacher/tests/${testId}/questions/new`}>
                  <Plus className="h-4 w-4 mr-2" />
                  Add Question
                </Link>
              </Button>
            </div>

            {questions.length === 0 ? (
              <div className="text-center py-8">
                <FileText className="h-12 w-12 mx-auto text-gray-400 mb-3" />
                <p className="text-gray-600 mb-4">No questions added yet</p>
                <Button asChild>
                  <Link href={`/teacher/tests/${testId}/questions/new`}>
                    <Plus className="h-4 w-4 mr-2" />
                    Add First Question
                  </Link>
                </Button>
              </div>
            ) : (
              <div className="space-y-4">
                {questions.slice(0, 3).map((question, index) => (
                  <div
                    key={question.id}
                    className="border border-gray-100 rounded-lg p-4"
                  >
                    <h4 className="font-medium mb-2">
                      Question {index + 1}: {question.question_text}
                    </h4>
                    <div className="grid grid-cols-2 gap-2 text-sm text-gray-600">
                      <p>A) {question.option_a}</p>
                      <p>B) {question.option_b}</p>
                      <p>C) {question.option_c}</p>
                      <p>D) {question.option_d}</p>
                    </div>
                    <p className="text-sm text-green-600 mt-2">
                      Correct: {['A', 'B', 'C', 'D'][question.correct_answer]}
                    </p>
                  </div>
                ))}

                {questions.length > 3 && (
                  <div className="text-center pt-4">
                    <Button variant="outline" asChild>
                      <Link href={`/teacher/tests/${testId}/questions`}>
                        View All {questions.length} Questions
                      </Link>
                    </Button>
                  </div>
                )}
              </div>
            )}
          </div>
        </div>

        {/* Quick Stats */}
        <div className="space-y-6">
          <div className="bg-white rounded-lg border border-gray-200 p-6">
            <h3 className="text-lg font-semibold mb-4">Quick Stats</h3>
            <div className="space-y-3">
              <div className="flex items-center justify-between">
                <span className="text-gray-600">Total Questions</span>
                <span className="font-semibold">{questions.length}</span>
              </div>
              <div className="flex items-center justify-between">
                <span className="text-gray-600">Test Status</span>
                <span className="font-semibold text-green-600">
                  {questions.length > 0 ? 'Ready' : 'Draft'}
                </span>
              </div>
            </div>
          </div>

          <div className="bg-white rounded-lg border border-gray-200 p-6">
            <h3 className="text-lg font-semibold mb-4">Quick Actions</h3>
            <div className="space-y-2">
              <Button className="w-full" variant="outline" asChild>
                <Link href={`/teacher/tests/${testId}/questions`}>
                  <FileText className="h-4 w-4 mr-2" />
                  Manage Questions
                </Link>
              </Button>
              <Button className="w-full" variant="outline" asChild>
                <Link href={`/teacher/tests/${testId}/edit`}>
                  <Edit className="h-4 w-4 mr-2" />
                  Edit Test
                </Link>
              </Button>
              {questions.length > 0 && (
                <Button className="w-full" asChild>
                  <Link href={`/student/tests/${testId}/preview`}>
                    <Users className="h-4 w-4 mr-2" />
                    Preview as Student
                  </Link>
                </Button>
              )}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
