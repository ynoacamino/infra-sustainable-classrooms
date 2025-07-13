import { CreateCourseForm } from '@/components/text/forms';
import Link from 'next/link';
import { Button } from '@/ui/button';
import { ArrowLeft } from 'lucide-react';

export default function NewCoursePage() {
  return (
    <div className="container mx-auto px-4 py-8">
      <div className="flex items-center gap-4 mb-8">
        <Button variant="outline" size="sm" asChild>
          <Link href="/teacher/courses">
            <ArrowLeft className="h-4 w-4 mr-2" />
            Back to Courses
          </Link>
        </Button>
        <h1 className="text-3xl font-bold">Create New Course</h1>
      </div>

      <div className="max-w-2xl mx-auto">
        <CreateCourseForm />
      </div>
    </div>
  );
}
