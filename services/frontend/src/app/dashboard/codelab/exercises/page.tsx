import { codelabService } from '@/services/codelab/service';
import { cookies } from 'next/headers';
import Link from 'next/link';
import { Button } from '@/ui/button';
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '@/ui/card';
import { Badge } from '@/ui/badge';
import { Code, Filter } from 'lucide-react';

interface ExercisesPageProps {
  searchParams: Promise<{
    difficulty?: string;
  }>;
}

export default async function ExercisesPage({
  searchParams,
}: ExercisesPageProps) {
  const { difficulty } = await searchParams;
  const codelab = await codelabService(cookies());
  const exercises = await codelab.listExercisesForStudents();

  if (!exercises.success) {
    return (
      <div className="flex flex-col items-center justify-center w-full h-full">
        <h1 className="text-2xl font-bold mb-4 text-red-600">Error</h1>
        <p>Error loading exercises: {exercises.error.message}</p>
      </div>
    );
  }

  // Filter exercises by difficulty if specified
  const filteredExercises = difficulty
    ? exercises.data.filter((ex) => ex.difficulty === difficulty)
    : exercises.data;

  const difficulties = ['easy', 'medium', 'hard'] as const;

  return (
    <div className="container mx-auto px-4 py-8">
      <div className="flex items-center gap-4 mb-8">
        <div>
          <h1 className="text-3xl font-bold">Coding Exercises</h1>
          <p className="text-muted-foreground">
            Challenge yourself with programming problems
          </p>
        </div>
      </div>

      {/* Filters */}
      <div className="flex gap-2 mb-6">
        <div className="flex items-center gap-2">
          <Filter className="h-4 w-4 text-muted-foreground" />
          <span className="text-sm font-medium">Filter by difficulty:</span>
        </div>
        <Button variant={!difficulty ? 'default' : 'outline'} size="sm" asChild>
          <Link href="/dashboard/codelab/exercises">
            All ({exercises.data.length})
          </Link>
        </Button>
        {difficulties.map((difficulty) => {
          const count = exercises.data.filter(
            (ex) => ex.difficulty === difficulty,
          ).length;
          return (
            <Button
              key={difficulty}
              variant={difficulty === difficulty ? 'default' : 'outline'}
              size="sm"
              asChild
            >
              <Link
                href={`/dashboard/codelab/exercises?difficulty=${difficulty}`}
              >
                {difficulty.charAt(0).toUpperCase() + difficulty.slice(1)} (
                {count})
              </Link>
            </Button>
          );
        })}
      </div>

      {/* Exercises Grid */}
      {filteredExercises.length === 0 ? (
        <Card>
          <CardContent className="py-12">
            <div className="text-center">
              <Code className="h-12 w-12 text-muted-foreground mx-auto mb-4" />
              <h3 className="text-lg font-medium mb-2">
                {difficulty
                  ? `No ${difficulty} exercises found`
                  : 'No exercises available'}
              </h3>
              <p className="text-muted-foreground mb-4">
                {difficulty
                  ? 'Try selecting a different difficulty level'
                  : 'Check back later for new coding challenges'}
              </p>
              {difficulty && (
                <Button asChild variant="outline">
                  <Link href="/dashboard/codelab/exercises">
                    View All Exercises
                  </Link>
                </Button>
              )}
            </div>
          </CardContent>
        </Card>
      ) : (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {filteredExercises.map((exercise) => (
            <Card
              key={exercise.id}
              className="hover:shadow-md transition-shadow"
            >
              <CardHeader>
                <div className="flex justify-between items-start mb-2">
                  <CardTitle className="text-lg">{exercise.title}</CardTitle>
                  <Badge
                    variant={
                      exercise.difficulty === 'easy'
                        ? 'default'
                        : exercise.difficulty === 'medium'
                          ? 'secondary'
                          : 'destructive'
                    }
                  >
                    {exercise.difficulty}
                  </Badge>
                </div>
                <CardDescription className="line-clamp-3">
                  {exercise.description}
                </CardDescription>
              </CardHeader>
              <CardContent>
                <div className="flex gap-2">
                  <Button asChild className="flex-1">
                    <Link href={`/dashboard/codelab/exercises/${exercise.id}`}>
                      Start Challenge
                    </Link>
                  </Button>
                </div>
                <div className="mt-3 text-xs text-muted-foreground">
                  Created {new Date(exercise.created_at).toLocaleDateString()}
                </div>
              </CardContent>
            </Card>
          ))}
        </div>
      )}

      {/* Results Summary */}
      {filteredExercises.length > 0 && (
        <div className="mt-8 text-center text-sm text-muted-foreground">
          Showing {filteredExercises.length} of {exercises.data.length}{' '}
          exercises
          {difficulty && <span> (filtered by {difficulty})</span>}
        </div>
      )}
    </div>
  );
}
