import { notFound } from 'next/navigation'
import { getExercise } from '@/actions/codelab/actions'
import { CreateTestForm } from '@/components/codelab/forms/create-test-form'

interface NewTestPageProps {
  params: Promise<{
    exerciseId: string
  }>
}

export default async function NewTestPage({ params }: NewTestPageProps) {
  const exerciseId = parseInt((await params).exerciseId)
  
  if (isNaN(exerciseId)) {
    notFound()
  }

  const exerciseResult = await getExercise(exerciseId)
  
  if (!exerciseResult?.success) {
    notFound()
  }

  const exercise = exerciseResult.data

  return (
    <div className="container mx-auto py-8">
      <div className="max-w-4xl mx-auto">
        <div className="mb-8">
          <h1 className="text-3xl font-bold">Create Test Case</h1>
          <p className="text-muted-foreground">
            Add a test case for: {exercise.title}
          </p>
        </div>
        
        <CreateTestForm exerciseId={exerciseId} />
      </div>
    </div>
  )
}
