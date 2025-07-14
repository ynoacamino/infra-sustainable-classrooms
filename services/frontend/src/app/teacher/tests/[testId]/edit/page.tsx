import { notFound } from 'next/navigation';
import { UpdateTestForm } from '@/components/knowledge/forms';
import { getTestAction } from '@/actions/knowledge/actions';
import Link from 'next/link';
import { Button } from '@/ui/button';
import { ArrowLeft } from 'lucide-react';

interface EditTestPageProps {
  params: Promise<{ testId: string }>;
}

export default async function EditTestPage({ params }: EditTestPageProps) {
  const resolvedParams = await params;
  const testId = parseInt(resolvedParams.testId);

  if (isNaN(testId)) {
    notFound();
  }

  const testResult = await getTestAction({ id: testId });

  if (!testResult.success) {
    if (testResult.error.status === 404) {
      notFound();
    }
    return (
      <div className="container mx-auto p-6">
        <div className="text-center">
          <h1 className="text-2xl font-bold mb-4 text-red-600">Error</h1>
          <p>Error loading test: {testResult.error.message}</p>
        </div>
      </div>
    );
  }

  const test = testResult.data.test;

  return (
    <div className="container mx-auto p-6">
      <div className="mb-6">
        <div className="flex items-center gap-4 mb-4">
          <Button variant="outline" size="sm" asChild>
            <Link href={`/teacher/tests/${testId}`}>
              <ArrowLeft className="h-4 w-4 mr-2" />
              Back to Test
            </Link>
          </Button>
          <div>
            <h1 className="text-3xl font-bold">Edit Test</h1>
            <p className="text-muted-foreground">Modify test details</p>
          </div>
        </div>
      </div>

      <UpdateTestForm test={test} />
    </div>
  );
}
