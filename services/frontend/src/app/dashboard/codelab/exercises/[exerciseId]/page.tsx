import { codelabService } from '@/services/codelab/service';
import { cookies } from 'next/headers';
import Link from 'next/link';
import { Button } from '@/ui/button';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/ui/card';
import { Badge } from '@/ui/badge';
import { Separator } from '@/ui/separator';
import { ArrowLeft, Play, Code, CheckCircle, XCircle, Calendar } from 'lucide-react';
import { notFound } from 'next/navigation';
import { CodeEditor } from '@/components/codelab/code-editor';

interface ExercisePageProps {
  params: {
    exerciseId: string;
  };
}

export default async function ExercisePage({ params }: ExercisePageProps) {
  const exerciseId = parseInt(params.exerciseId);

  if (isNaN(exerciseId)) {
    notFound();
  }

  const codelab = await codelabService(cookies());

  const exercise = await codelab.getExerciseForStudent({
    id: exerciseId,
  });


  if (!exercise.success) {
    if (exercise.error.message.includes('not found')) {
      notFound();
    }
    return (
      <div className="flex flex-col items-center justify-center w-full h-full">
        <h1 className="text-2xl font-bold mb-4 text-red-600">Error</h1>
        <p>Error loading exercise: {exercise.error.message}</p>
      </div>
    );
  }
  const tests = exercise.data.tests
  const attempts = exercise.data.attempts || [];

  console.log('Exercise Tests:', tests);

  return (
    <div className="container mx-auto px-4 py-8">
      <div className="flex items-center gap-4 mb-8">
        <Button variant="outline" size="sm" asChild>
          <Link href="/dashboard/codelab/exercises">
            <ArrowLeft className="h-4 w-4 mr-2" />
            Back to Exercises
          </Link>
        </Button>
        <div className="flex-1">
          <div className="flex items-center gap-3 mb-2">
            <h1 className="text-3xl font-bold">{exercise.data.title}</h1>
            <Badge
              variant={
                exercise.data.difficulty === 'easy' ? 'default' :
                  exercise.data.difficulty === 'medium' ? 'secondary' : 'destructive'
              }
            >
              {exercise.data.difficulty}
            </Badge>
          </div>
          <p className="text-muted-foreground">
            Solve this coding challenge
          </p>
        </div>
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-2 gap-8">
        {/* Problem Description */}
        <div className="space-y-6">
          <Card>
            <CardHeader>
              <CardTitle className="flex items-center gap-2">
                <Code className="h-5 w-5" />
                Problem Description
              </CardTitle>
            </CardHeader>
            <CardContent>
              <div className="prose prose-sm max-w-none">
                <p className="whitespace-pre-wrap">{exercise.data.description}</p>
              </div>
            </CardContent>
          </Card>

          {/* Sample Test Cases */}
          <Card>
            <CardHeader>
              <CardTitle className="flex items-center gap-2">
                <CheckCircle className="h-5 w-5 text-green-600" />
                Example Test Cases
              </CardTitle>
              <CardDescription>
                These are sample inputs and expected outputs
              </CardDescription>
            </CardHeader>
            <CardContent>
              <div className="space-y-4">
                {tests.length > 0 ? (
                  <div className="space-y-3">
                    {tests.map((test, index) => (
                      <div key={test.id} className="border rounded-lg p-3">
                        <div className="text-sm font-medium mb-2">Example {index + 1}:</div>
                        <div className="grid grid-cols-2 gap-3 text-sm">
                          <div>
                            <div className="font-medium text-muted-foreground mb-1">Input:</div>
                            <pre className="bg-muted p-2 rounded text-xs whitespace-pre-wrap">"{test.input}"</pre>
                          </div>
                          <div>
                            <div className="font-medium text-muted-foreground mb-1">Output:</div>
                            <pre className="bg-muted p-2 rounded text-xs whitespace-pre-wrap">"{test.output}"</pre>
                          </div>
                        </div>
                      </div>
                    ))}
                  </div>
                ) : (
                  <div className="text-center py-4">
                    <div className="text-sm text-muted-foreground">
                      No public test cases available.
                      Your solution will be tested against hidden test cases.
                    </div>
                  </div>
                )}

                <div className="text-sm text-muted-foreground border-t pt-3">
                  Note: There may be additional hidden test cases that will be used to evaluate your solution.
                </div>
              </div>
            </CardContent>
          </Card>

          <div className="space-y-4">
            <h2 className="text-xl font-semibold">Recent Attempts</h2>

            {attempts.length === 0 ? (
              <Card>
                <CardContent className="py-12">
                  <div className="text-center">
                    <Code className="h-12 w-12 text-muted-foreground mx-auto mb-4" />
                    <h3 className="text-lg font-medium mb-2">No attempts yet</h3>
                    <p className="text-muted-foreground mb-4">
                      Start solving coding challenges to see your progress here
                    </p>
                  </div>
                </CardContent>
              </Card>
            ) : (
              <div className="space-y-4">
                {attempts.map((attempt) => (
                  <Card key={attempt.id}>
                    <CardHeader>
                      <div className="flex items-center justify-between">
                        <div className="flex items-center gap-3">
                          <CardTitle className="text-lg">{exercise.data.title}</CardTitle>
                          <Badge
                            variant={
                              exercise.data.difficulty === 'easy' ? 'default' :
                                exercise.data.difficulty === 'medium' ? 'secondary' : 'destructive'
                            }
                          >
                            {exercise.data.difficulty}
                          </Badge>
                          <Badge
                            variant={attempt.success ? 'default' : 'secondary'}
                            className={attempt.success ? 'bg-green-600' : 'bg-gray-500'}
                          >
                            {attempt.success ? (
                              <CheckCircle className="w-3 h-3 mr-1" />
                            ) : (
                              <XCircle className="w-3 h-3 mr-1" />
                            )}
                            {attempt.success ? 'Passed' : 'Failed'}
                          </Badge>
                        </div>
                        <div className="flex items-center gap-2 text-sm text-muted-foreground">
                          <Calendar className="w-4 h-4" />
                          {new Date(attempt.created_at).toLocaleString()}
                        </div>
                      </div>
                    </CardHeader>

                    <CardContent>
                      <div className="space-y-4">
                        <div>
                          <h4 className="font-medium mb-2">Submitted Code:</h4>
                          <pre className="bg-muted p-3 rounded-md text-sm overflow-x-auto border">
                            {attempt.code}
                          </pre>
                        </div>

                        <div className="flex gap-2">
                          {!attempt.success && (
                            <Button asChild size="sm">
                              <Link href={`/dashboard/codelab/exercises/${exercise.data.id}`}>
                                Continue Working
                              </Link>
                            </Button>
                          )}
                        </div>
                      </div>
                    </CardContent>
                  </Card>
                ))}
              </div>
            )}
          </div>
        </div>

        {/* Code Editor */}
        <div className="space-y-6">
          <Card>
            <CardHeader>
              <CardTitle className="flex items-center gap-2">
                <Play className="h-5 w-5" />
                Your Solution
              </CardTitle>
              <CardDescription>
                Write your code below and test it
              </CardDescription>
            </CardHeader>
            <CardContent>
              <CodeEditor
                exerciseId={exerciseId}
                initialCode={exercise.data.initial_code}
                testCases={tests}
              />
            </CardContent>
          </Card>

          {/* Exercise Info */}
          <Card>
            <CardHeader>
              <CardTitle>Exercise Information</CardTitle>
            </CardHeader>
            <CardContent>
              <div className="space-y-2 text-sm">
                <div className="flex justify-between">
                  <span className="text-muted-foreground">Difficulty:</span>
                  <Badge
                    variant={
                      exercise.data.difficulty === 'easy' ? 'default' :
                        exercise.data.difficulty === 'medium' ? 'secondary' : 'destructive'
                    }
                  >
                    {exercise.data.difficulty}
                  </Badge>
                </div>
                <Separator />
                <div className="flex justify-between">
                  <span className="text-muted-foreground">Created:</span>
                  <span>{new Date(exercise.data.created_at).toLocaleDateString()}</span>
                </div>
                <div className="flex justify-between">
                  <span className="text-muted-foreground">Last Updated:</span>
                  <span>{new Date(exercise.data.updated_at).toLocaleDateString()}</span>
                </div>
              </div>
            </CardContent>
          </Card>
        </div>
      </div>
    </div>
  );
}
