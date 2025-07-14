import { codelabService } from '@/services/codelab/service';
import { cookies } from 'next/headers';
import Link from 'next/link';
import { Button } from '@/ui/button';
import { ArrowLeft, Edit, TestTube } from 'lucide-react';
import { notFound } from 'next/navigation';
import { DeleteExerciseButton } from '@/components/codelab/delete-exercise-button';

interface ExercisePageProps {
  params: Promise<{ exerciseId: string }>;
}

export default async function ExercisePage({ params }: ExercisePageProps) {
  const { exerciseId } = await params;
  const codelab = await codelabService(cookies());
  
  const exercise = await codelab.getExercise({
    id: parseInt(exerciseId),
  });

  if (!exercise.success) {
    console.error('Error fetching exercise:', exercise.error);
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

  return (
    <div className="container mx-auto px-4 py-8">
      <div className="flex items-center gap-4 mb-8">
        <Button variant="outline" size="sm" asChild>
          <Link href="/teacher/codelab/exercises">
            <ArrowLeft className="h-4 w-4 mr-2" />
            Back to Exercises
          </Link>
        </Button>
        <h1 className="text-3xl font-bold">{exercise.data.title}</h1>
        <div className="ml-auto flex gap-2">
          <Button asChild variant="outline">
            <Link href={`/teacher/codelab/exercises/${exerciseId}/edit`}>
              <Edit className="h-4 w-4 mr-2" />
              Edit
            </Link>
          </Button>
          <Button asChild variant="outline">
            <Link href={`/teacher/codelab/exercises/${exerciseId}/tests`}>
              <TestTube className="h-4 w-4 mr-2" />
              Test Cases
            </Link>
          </Button>
        </div>
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-2 gap-8">
        <div className="space-y-6">
          <div>
            <h2 className="text-xl font-semibold mb-3">Description</h2>
            <div className="bg-gray-50 rounded-lg p-4">
              <p className="whitespace-pre-wrap">{exercise.data.description}</p>
            </div>
          </div>

          <div>
            <h2 className="text-xl font-semibold mb-3">Difficulty</h2>
            <span 
              className={`px-3 py-1 rounded-full text-sm font-medium ${
                exercise.data.difficulty === 'easy' ? 'bg-green-100 text-green-800' :
                exercise.data.difficulty === 'medium' ? 'bg-yellow-100 text-yellow-800' : 
                'bg-red-100 text-red-800'
              }`}
            >
              {exercise.data.difficulty.charAt(0).toUpperCase() + exercise.data.difficulty.slice(1)}
            </span>
          </div>
        </div>

        <div className="space-y-6">
          <div>
            <h2 className="text-xl font-semibold mb-3">Initial Code Template</h2>
            <div className="bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto">
              <pre className="text-sm">
                <code>{exercise.data.initial_code}</code>
              </pre>
            </div>
          </div>

          <div>
            <h2 className="text-xl font-semibold mb-3">Solution</h2>
            <div className="bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto">
              <pre className="text-sm">
                <code>{exercise.data.solution}</code>
              </pre>
            </div>
          </div>
        </div>
      </div>

      <div className="mt-8 pt-8 border-t">
        <div className="flex justify-between items-center mb-4">
          <h2 className="text-xl font-semibold">Exercise Information</h2>
          <DeleteExerciseButton 
            exerciseId={parseInt(exerciseId)} 
            exerciseTitle={exercise.data.title}
          />
        </div>
        <div className="grid grid-cols-1 md:grid-cols-3 gap-4 text-sm text-gray-600">
          <div>
            <span className="font-medium">Created:</span>{' '}
            {new Date(exercise.data.created_at).toLocaleDateString()}
          </div>
          <div>
            <span className="font-medium">Last Updated:</span>{' '}
            {new Date(exercise.data.updated_at).toLocaleDateString()}
          </div>
          <div>
            <span className="font-medium">Created by:</span> User #{exercise.data.created_by}
          </div>
        </div>
      </div>
    </div>
  );
}
