import { textService } from '@/services/text/service';
import { cookies } from 'next/headers';
import Link from 'next/link';
import { Button } from '@/ui/button';
import { ArrowLeft } from 'lucide-react';
import { notFound } from 'next/navigation';
import { CreateSectionForm } from '@/components/text/forms/create-section-form';

interface NewSectionPageProps {
  params: Promise<{ courseId: string }>;
}

export default async function NewSectionPage({ params }: NewSectionPageProps) {
  const asyncParams = await params;

  const courseId = parseInt(asyncParams.courseId);

  if (isNaN(courseId)) {
    notFound();
  }

  const text = await textService(cookies());
  const courseResult = await text.getCourse({ id: courseId });

  if (!courseResult.success) {
    if (courseResult.error.status === 404) {
      notFound();
    }
    return (
      <div className="container mx-auto px-4 py-8">
        <div className="text-center">
          <h1 className="text-2xl font-bold mb-4 text-red-600">Error</h1>
          <p>Error loading course: {courseResult.error.message}</p>
        </div>
      </div>
    );
  }

  const course = courseResult.data;

  return (
    <div className="container mx-auto px-4 py-8">
      <div className="flex items-center gap-4 mb-8">
        <Button variant="outline" size="sm" asChild>
          <Link href={`/teacher/courses/${courseId}/sections`}>
            <ArrowLeft className="h-4 w-4 mr-2" />
            Back to Sections
          </Link>
        </Button>
        <div>
          <h1 className="text-3xl font-bold">Add New Section</h1>
          <p className="text-gray-600">{course.title}</p>
        </div>
      </div>

      <div className="max-w-2xl mx-auto">
        <CreateSectionForm courseId={courseId} />
      </div>
    </div>
  );
}
