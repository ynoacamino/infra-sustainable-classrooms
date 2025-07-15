import { getSubmissionResultAction } from '@/actions/knowledge/actions';
import Link from 'next/link';
import { Button } from '@/ui/button';
import { ArrowLeft, CheckCircle, XCircle, Clock } from 'lucide-react';
import { notFound } from 'next/navigation';

interface SubmissionResultPageProps {
  params: Promise<{ submissionId: string }>;
}

export default async function SubmissionResultPage({
  params,
}: SubmissionResultPageProps) {
  const resolvedParams = await params;
  const submissionId = parseInt(resolvedParams.submissionId);

  if (isNaN(submissionId)) {
    notFound();
  }

  const submissionResult = await getSubmissionResultAction({
    id: submissionId,
  });

  if (!submissionResult.success) {
    if (submissionResult.error.status === 404) {
      notFound();
    }
    return (
      <div className="container mx-auto px-4 py-8">
        <div className="text-center">
          <h1 className="text-2xl font-bold mb-4 text-red-600">Error</h1>
          <p>Error loading submission: {submissionResult.error.message}</p>
        </div>
      </div>
    );
  }

  const { submission, questions } = submissionResult.data;

  return (
    <div className="container mx-auto px-4 py-8">
      <div className="flex items-center gap-4 mb-8">
        <Button variant="outline" size="sm" asChild>
          <Link href="/dashboard/tests/results">
            <ArrowLeft className="h-4 w-4 mr-2" />
            Back to Results
          </Link>
        </Button>
        <div>
          <h1 className="text-3xl font-bold">Test Result</h1>
          <p className="text-gray-600 mt-2">Submission #{submission.id}</p>
        </div>
      </div>

      {/* Summary Card */}
      <div className="bg-white border border-gray-200 rounded-lg p-6 mb-8">
        <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
          <div className="text-center">
            <div
              className={`inline-flex items-center justify-center w-16 h-16 rounded-full mb-2 ${
                submission.score >= 70 ? 'bg-green-100' : 'bg-red-100'
              }`}
            >
              {submission.score >= 70 ? (
                <CheckCircle className="h-8 w-8 text-green-600" />
              ) : (
                <XCircle className="h-8 w-8 text-red-600" />
              )}
            </div>
            <div
              className={`text-3xl font-bold ${
                submission.score >= 70 ? 'text-green-600' : 'text-red-600'
              }`}
            >
              {submission.score}%
            </div>
            <p className="text-gray-600">Your Score</p>
          </div>

          <div className="text-center">
            <div className="inline-flex items-center justify-center w-16 h-16 rounded-full bg-blue-100 mb-2">
              <CheckCircle className="h-8 w-8 text-blue-600" />
            </div>
            <div className="text-3xl font-bold text-blue-600">
              {questions.filter((q) => q.is_correct).length}
            </div>
            <p className="text-gray-600">Correct Answers</p>
          </div>

          <div className="text-center">
            <div className="inline-flex items-center justify-center w-16 h-16 rounded-full bg-gray-100 mb-2">
              <Clock className="h-8 w-8 text-gray-600" />
            </div>
            <div className="text-3xl font-bold text-gray-600">
              {questions.length}
            </div>
            <p className="text-gray-600">Total Questions</p>
          </div>
        </div>

        <div className="mt-6 pt-6 border-t border-gray-200">
          <p className="text-sm text-gray-600">
            Submitted on:{' '}
            {new Date(submission.submitted_at * 1000).toLocaleString()}
          </p>
        </div>
      </div>

      {/* Questions Review */}
      <div className="space-y-6">
        <h2 className="text-2xl font-bold">Question Review</h2>

        {questions.map((questionResult, index) => (
          <div
            key={questionResult.question.id}
            className="bg-white border border-gray-200 rounded-lg p-6"
          >
            <div className="flex items-start justify-between mb-4">
              <h3 className="text-lg font-semibold">Question {index + 1}</h3>
              <div
                className={`flex items-center ${
                  questionResult.is_correct ? 'text-green-600' : 'text-red-600'
                }`}
              >
                {questionResult.is_correct ? (
                  <CheckCircle className="h-5 w-5 mr-1" />
                ) : (
                  <XCircle className="h-5 w-5 mr-1" />
                )}
                <span className="font-medium">
                  {questionResult.is_correct ? 'Correct' : 'Incorrect'}
                </span>
              </div>
            </div>

            <p className="text-gray-700 mb-4">
              {questionResult.question.question_text}
            </p>

            <div className="grid grid-cols-1 sm:grid-cols-2 gap-3">
              {[
                {
                  index: 0,
                  label: 'A',
                  text: questionResult.question.option_a,
                },
                {
                  index: 1,
                  label: 'B',
                  text: questionResult.question.option_b,
                },
                {
                  index: 2,
                  label: 'C',
                  text: questionResult.question.option_c,
                },
                {
                  index: 3,
                  label: 'D',
                  text: questionResult.question.option_d,
                },
              ].map((option) => {
                const isSelected =
                  questionResult.selected_answer === option.index;
                const isCorrect =
                  questionResult.question.correct_answer === option.index;

                return (
                  <div
                    key={option.index}
                    className={`p-3 border rounded-lg ${
                      isCorrect && isSelected
                        ? 'border-green-500 bg-green-50'
                        : isCorrect
                          ? 'border-green-500 bg-green-50'
                          : isSelected && !isCorrect
                            ? 'border-red-500 bg-red-50'
                            : 'border-gray-200'
                    }`}
                  >
                    <div className="flex items-center">
                      <span className="font-medium mr-2">{option.label}.</span>
                      <span className="flex-1">{option.text}</span>
                      {isCorrect && (
                        <CheckCircle className="h-4 w-4 text-green-600 ml-2" />
                      )}
                      {isSelected && !isCorrect && (
                        <XCircle className="h-4 w-4 text-red-600 ml-2" />
                      )}
                    </div>

                    {isSelected && (
                      <div className="mt-1 text-xs text-gray-600">
                        Your answer
                      </div>
                    )}

                    {isCorrect && !isSelected && (
                      <div className="mt-1 text-xs text-green-600">
                        Correct answer
                      </div>
                    )}
                  </div>
                );
              })}
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}
