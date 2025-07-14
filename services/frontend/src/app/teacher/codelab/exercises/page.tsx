import { codelabService } from '@/services/codelab/service';
import { cookies } from 'next/headers';
import Link from 'next/link';
import { Button } from '@/ui/button';
import { CompactDeleteExerciseButton } from '@/components/codelab/compact-delete-exercise-button';

export default async function ExercisesPage() {
  const codelab = await codelabService(cookies());
  const exercises = await codelab.listExercises();

  if (!exercises.success) {
    return (
      <div className="flex flex-col items-center justify-center w-full h-full">
        <h1 className="text-2xl font-bold mb-4 text-red-600">Error</h1>
        <p>Error loading exercises: {exercises.error.message}</p>
      </div>
    );
  }

  return (
    <div className="container mx-auto px-4 py-8">
      <div className="flex justify-between items-center mb-8">
        <h1 className="text-3xl font-bold">Coding Exercises</h1>
        <Button asChild>
          <Link href="/teacher/codelab/exercises/new">Create New Exercise</Link>
        </Button>
      </div>

      {exercises.data.length === 0 ? (
        <div className="text-center py-12">
          <h2 className="text-xl font-semibold mb-4">No exercises yet</h2>
          <p className="text-gray-600 mb-6">
            Create your first coding exercise to get started.
          </p>
          <Button asChild>
            <Link href="/teacher/codelab/exercises/new">Create Exercise</Link>
          </Button>
        </div>
      ) : (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {exercises.data.map((exercise) => (
            <div 
              key={exercise.id} 
              className="border rounded-lg p-6 hover:shadow-md transition-shadow bg-white"
            >
              <div className="flex justify-between items-start mb-4">
                <h3 className="text-lg font-semibold">{exercise.title}</h3>
                <div className="flex items-center gap-2">
                  <span 
                    className={`px-2 py-1 rounded text-sm font-medium ${
                      exercise.difficulty === 'easy' ? 'bg-green-100 text-green-800' :
                      exercise.difficulty === 'medium' ? 'bg-yellow-100 text-yellow-800' : 
                      'bg-red-100 text-red-800'
                    }`}
                  >
                    {exercise.difficulty}
                  </span>
                  <CompactDeleteExerciseButton 
                    exerciseId={exercise.id} 
                    exerciseTitle={exercise.title}
                  />
                </div>
              </div>
              
              <p className="text-gray-600 mb-4 line-clamp-3">
                {exercise.description}
              </p>
              
              <div className="flex gap-2">
                <Button asChild size="sm" className="flex-1">
                  <Link href={`/teacher/codelab/exercises/${exercise.id}`}>
                    View
                  </Link>
                </Button>
                <Button asChild variant="outline" size="sm" className="flex-1">
                  <Link href={`/teacher/codelab/exercises/${exercise.id}/edit`}>
                    Edit
                  </Link>
                </Button>
                <Button asChild variant="outline" size="sm" className="flex-1">
                  <Link href={`/teacher/codelab/exercises/${exercise.id}/tests`}>
                    Tests
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
