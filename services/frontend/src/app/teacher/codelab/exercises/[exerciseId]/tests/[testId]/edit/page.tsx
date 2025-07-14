import { notFound } from 'next/navigation'
import { getExercise, getTestsByExercise } from '@/actions/codelab/actions'
import { UpdateTestForm } from '@/components/codelab/forms/update-test-form'

interface EditTestPageProps {
  params: Promise<{
    exerciseId: string
    testId: string
  }>
}

export default async function EditTestPage({ params }: EditTestPageProps) {
  const exerciseId = parseInt((await params).exerciseId)
  const testId = parseInt((await params).testId)
  
  if (isNaN(exerciseId) || isNaN(testId)) {
    notFound()
  }

  const [exerciseResult, testsResult] = await Promise.all([
    getExercise(exerciseId),
    getTestsByExercise(exerciseId),
  ])
  
  if (!exerciseResult?.success || !testsResult?.success) {
    notFound()
  }

  const exercise = exerciseResult.data
  const test = testsResult.data.find(t => t.id === testId)

  if (!test) {
    notFound()
  }

  return (
    <div className="container mx-auto py-8">
      <div className="max-w-4xl mx-auto">
        <div className="mb-8">
          <h1 className="text-3xl font-bold">Edit Test Case</h1>
          <p className="text-muted-foreground">
            Update test case for: {exercise.title}
          </p>
        </div>
        
        <UpdateTestForm test={test} />
      </div>
    </div>
  )
}
