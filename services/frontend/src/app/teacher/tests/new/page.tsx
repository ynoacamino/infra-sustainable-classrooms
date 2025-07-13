import { CreateTestForm } from '@/components/knowledge/forms';
import Link from 'next/link';
import { Button } from '@/ui/button';
import { ArrowLeft } from 'lucide-react';

export default function NewTestPage() {
  return (
    <div className="container mx-auto px-4 py-8">
      <div className="flex items-center gap-4 mb-8">
        <Button variant="outline" size="sm" asChild>
          <Link href="/teacher/tests">
            <ArrowLeft className="h-4 w-4 mr-2" />
            Back to Tests
          </Link>
        </Button>
        <div>
          <h1 className="text-3xl font-bold">Create New Test</h1>
          <p className="text-gray-600 mt-2">
            Create a new test for your students
          </p>
        </div>
      </div>

      <div className="max-w-3xl mx-auto">
        <div className="bg-white rounded-lg border border-gray-200 p-6">
          <CreateTestForm />
        </div>
      </div>
    </div>
  );
}
