import Link from 'next/link';
import { Button } from '@/ui/button';
import { ArrowLeft } from 'lucide-react';
import { CreateExerciseForm } from '@/components/codelab/forms/create-exercise-form';

export default function NewExercisePage() {
  return (
    <div className="container mx-auto px-4 py-8">
      <div className="flex items-center gap-4 mb-8">
        <Button variant="outline" size="sm" asChild>
          <Link href="/teacher/codelab/exercises">
            <ArrowLeft className="h-4 w-4 mr-2" />
            Back to Exercises
          </Link>
        </Button>
        <h1 className="text-3xl font-bold">Create New Exercise</h1>
      </div>

      <div className="max-w-4xl mx-auto">
        <CreateExerciseForm />
      </div>
    </div>
  );
}
