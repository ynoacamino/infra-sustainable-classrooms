import { notFound } from 'next/navigation';
import { getExercise } from '@/actions/codelab/actions';
import { UpdateExerciseForm } from '@/components/codelab/forms/update-exercise-form';

interface EditExercisePageProps {
  params: Promise<{
    exerciseId: string;
  }>;
}

export default async function EditExercisePage({
  params,
}: EditExercisePageProps) {
  const exerciseId = parseInt((await params).exerciseId);

  if (isNaN(exerciseId)) {
    notFound();
  }

  const exerciseResult = await getExercise({ id: exerciseId });

  if (!exerciseResult?.success) {
    notFound();
  }

  const exercise = exerciseResult.data;

  return (
    <div className="container mx-auto py-8">
      <div className="max-w-4xl mx-auto">
        <div className="mb-8">
          <h1 className="text-3xl font-bold">Edit Exercise</h1>
          <p className="text-muted-foreground">
            Update the exercise details and configuration
          </p>
        </div>

        <UpdateExerciseForm exercise={exercise} />
      </div>
    </div>
  );
}
