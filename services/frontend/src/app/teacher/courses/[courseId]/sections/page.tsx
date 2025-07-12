import { textService } from "@/services/text/service";
import { cookies } from "next/headers";
import Link from "next/link";
import { Button } from "@/ui/button";
import { ArrowLeft, Edit, Plus, FileText } from "lucide-react";
import { notFound } from "next/navigation";

interface SectionsPageProps {
  params: { courseId: string };
}

export default async function SectionsPage({ params }: SectionsPageProps) {
  const courseId = parseInt(params.courseId);
  
  if (isNaN(courseId)) {
    notFound();
  }

  const text = await textService(cookies());
  const [courseResult, sectionsResult] = await Promise.all([
    text.getCourse({ course_id: courseId }),
    text.listSections({ course_id: courseId })
  ]);

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

  if (!sectionsResult.success) {
    return (
      <div className="container mx-auto px-4 py-8">
        <div className="text-center">
          <h1 className="text-2xl font-bold mb-4 text-red-600">Error</h1>
          <p>Error loading sections: {sectionsResult.error.message}</p>
        </div>
      </div>
    );
  }

  const course = courseResult.data;
  const sections = sectionsResult.data;

  return (
    <div className="container mx-auto px-4 py-8">
      <div className="flex items-center gap-4 mb-8">
        <Button variant="outline" size="sm" asChild>
          <Link href={`/teacher/courses/${course.id}`}>
            <ArrowLeft className="h-4 w-4 mr-2" />
            Back to Course
          </Link>
        </Button>
        <div>
          <h1 className="text-3xl font-bold">Sections</h1>
          <p className="text-gray-600">{course.title}</p>
        </div>
      </div>

      <div className="flex justify-between items-center mb-6">
        <h2 className="text-xl font-semibold">Course Sections</h2>
        <Button asChild>
          <Link href={`/teacher/courses/${courseId}/sections/new`}>
            <Plus className="h-4 w-4 mr-2" />
            Add Section
          </Link>
        </Button>
      </div>

      {sections.length === 0 ? (
        <div className="text-center py-12">
          <h3 className="text-lg font-semibold mb-4">No sections yet</h3>
          <p className="text-gray-600 mb-6">Add your first section to organize course content</p>
          <Button asChild>
            <Link href={`/teacher/courses/${courseId}/sections/new`}>
              <Plus className="h-4 w-4 mr-2" />
              Add Section
            </Link>
          </Button>
        </div>
      ) : (
        <div className="space-y-4">
          {sections
            .sort((a, b) => a.order - b.order)
            .map(section => (
            <div key={section.id} className="bg-white rounded-lg border shadow-sm p-6">
              <div className="flex justify-between items-start">
                <div className="flex-1">
                  <div className="flex items-center gap-3 mb-2">
                    <span className="bg-blue-100 text-blue-800 text-sm font-medium px-2.5 py-0.5 rounded">
                      Section {section.order}
                    </span>
                    <h3 className="text-xl font-semibold">{section.title}</h3>
                  </div>
                  <p className="text-gray-600 mb-4">{section.description}</p>
                  <div className="text-sm text-gray-500">
                    Created: {new Date(section.created_at).toLocaleDateString()}
                  </div>
                </div>
                <div className="flex gap-2 ml-4">
                  <Button asChild variant="outline" size="sm">
                    <Link href={`/teacher/courses/${courseId}/sections/${section.id}`}>
                      View
                    </Link>
                  </Button>
                  <Button asChild variant="outline" size="sm">
                    <Link href={`/teacher/courses/${courseId}/sections/${section.id}/edit`}>
                      <Edit className="h-4 w-4 mr-2" />
                      Edit
                    </Link>
                  </Button>
                  <Button asChild size="sm">
                    <Link href={`/teacher/courses/${courseId}/sections/${section.id}/articles`}>
                      <FileText className="h-4 w-4 mr-2" />
                      Articles
                    </Link>
                  </Button>
                </div>
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
}
