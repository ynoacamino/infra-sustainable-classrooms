import { notFound } from 'next/navigation';
import { Button } from '@/ui/button';
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '@/ui/card';
import { Badge } from '@/ui/badge';
import { Separator } from '@/ui/separator';
import { Plus, Edit } from 'lucide-react';
import Link from 'next/link';
import { getExercise, getTestsByExercise } from '@/actions/codelab/actions';
import { DeleteTestButton } from '@/components/codelab/delete-test-button';

interface TestsPageProps {
  params: Promise<{
    exerciseId: string;
  }>;
}

export default async function TestsPage({ params }: TestsPageProps) {
  const exerciseId = parseInt((await params).exerciseId);

  if (isNaN(exerciseId)) {
    notFound();
  }

  const [exerciseResult, testsResult] = await Promise.all([
    getExercise({ id: exerciseId }),
    getTestsByExercise({ exercise_id: exerciseId }),
  ]);

  if (!exerciseResult?.success) {
    notFound();
  }

  const exercise = exerciseResult.data;
  const tests = testsResult?.success ? testsResult.data : [];

  return (
    <div className="container mx-auto py-8">
      <div className="max-w-6xl mx-auto">
        <div className="flex items-center justify-between mb-8">
          <div>
            <h1 className="text-3xl font-bold">Test Cases</h1>
            <p className="text-muted-foreground">
              Manage test cases for: {exercise.title}
            </p>
          </div>
          <Button asChild>
            <Link href={`/teacher/codelab/exercises/${exerciseId}/tests/new`}>
              <Plus className="w-4 h-4 mr-2" />
              New Test Case
            </Link>
          </Button>
        </div>

        <div className="mb-6">
          <Card>
            <CardHeader>
              <CardTitle>{exercise.title}</CardTitle>
              <CardDescription>{exercise.description}</CardDescription>
            </CardHeader>
            <CardContent>
              <div className="flex gap-2">
                <Badge variant="outline">{exercise.difficulty}</Badge>
                <Badge variant="outline">Programming Exercise</Badge>
              </div>
            </CardContent>
          </Card>
        </div>

        <Separator className="my-6" />

        <div className="grid gap-4">
          {tests && tests.length > 0 ? (
            tests.map((test, index) => (
              <Card key={test.id || index}>
                <CardHeader>
                  <div className="flex items-center justify-between">
                    <CardTitle className="text-lg">
                      Test Case #{index + 1}
                    </CardTitle>
                    <div className="flex gap-2">
                      <Button variant="outline" size="sm" asChild>
                        <Link
                          href={`/teacher/codelab/exercises/${exerciseId}/tests/${test.id}/edit`}
                        >
                          <Edit className="w-4 h-4 mr-2" />
                          Edit
                        </Link>
                      </Button>
                      <DeleteTestButton testId={test.id} index={index} />
                    </div>
                  </div>
                </CardHeader>
                <CardContent>
                  <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                    <div>
                      <h4 className="font-medium mb-2">Input</h4>
                      <pre className="bg-muted p-3 rounded-md text-sm overflow-x-auto">
                        {test.input || 'No input'}
                      </pre>
                    </div>
                    <div>
                      <h4 className="font-medium mb-2">Expected Output</h4>
                      <pre className="bg-muted p-3 rounded-md text-sm overflow-x-auto">
                        {test.output || 'No expected output'}
                      </pre>
                    </div>
                  </div>
                  {!test.public && (
                    <div className="mt-4">
                      <Badge variant="secondary">Hidden Test</Badge>
                    </div>
                  )}
                </CardContent>
              </Card>
            ))
          ) : (
            <Card>
              <CardContent className="py-12">
                <div className="text-center">
                  <h3 className="text-lg font-medium mb-2">
                    No test cases yet
                  </h3>
                  <p className="text-muted-foreground mb-4">
                    Create your first test case to validate student submissions
                  </p>
                  <Button asChild>
                    <Link
                      href={`/teacher/codelab/exercises/${exerciseId}/tests/new`}
                    >
                      <Plus className="w-4 h-4 mr-2" />
                      Create Test Case
                    </Link>
                  </Button>
                </div>
              </CardContent>
            </Card>
          )}
        </div>
      </div>
    </div>
  );
}
