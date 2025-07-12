import { textService } from "@/services/text/service";
import { cookies } from "next/headers";
import Link from "next/link";
import { Button } from "@/ui/button";
import { ArrowLeft, Edit, Plus } from "lucide-react";
import { notFound } from "next/navigation";

interface SectionPageProps {
  params: { courseId: string; sectionId: string };
}

export default async function SectionPage({ params }: SectionPageProps) {
  const courseId = parseInt(params.courseId);
  const sectionId = parseInt(params.sectionId);
  
  if (isNaN(courseId) || isNaN(sectionId)) {
    notFound();
  }

  const text = await textService(cookies());
  const sectionResult = await text.getSection({ section_id: sectionId });

  if (!sectionResult.success) {
    if (sectionResult.error.status === 404) {
      notFound();
    }
    return (
      <div className="container mx-auto px-4 py-8">
        <div className="text-center">
          <h1 className="text-2xl font-bold mb-4 text-red-600">Error</h1>
          <p>Error loading section: {sectionResult.error.message}</p>
        </div>
      </div>
    );
  }

  const section = sectionResult.data;

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
          <h1 className="text-3xl font-bold">{section.title}</h1>
          <p className="text-gray-600">Section {section.order}</p>
        </div>
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
        <div className="lg:col-span-2">
          <div className="bg-white rounded-lg border shadow-sm p-6">
            <h2 className="text-2xl font-semibold mb-4">{section.title}</h2>
            <p className="text-gray-600 mb-6">{section.description}</p>
          </div>
        </div>

        <div className="space-y-4">
          <div className="bg-white rounded-lg border shadow-sm p-6">
            <h3 className="text-lg font-semibold mb-4">Section Management</h3>
            <div className="space-y-3">
              <Button asChild className="w-full">
                <Link href={`/teacher/courses/${courseId}/sections/${sectionId}/articles`}>
                  View Articles
                </Link>
              </Button>
              <Button asChild variant="outline" className="w-full">
                <Link href={`/teacher/courses/${courseId}/sections/${sectionId}/articles/new`}>
                  <Plus className="h-4 w-4 mr-2" />
                  Add Article
                </Link>
              </Button>
              <Button asChild variant="outline" className="w-full">
                <Link href={`/teacher/courses/${courseId}/sections/${sectionId}/edit`}>
                  <Edit className="h-4 w-4 mr-2" />
                  Edit Section
                </Link>
              </Button>
            </div>
          </div>

          <div className="bg-white rounded-lg border shadow-sm p-6">
            <h3 className="text-lg font-semibold mb-4">Section Information</h3>
            <div className="space-y-2 text-sm">
              <div><strong>Order:</strong> {section.order}</div>
              <div><strong>Course ID:</strong> {section.course_id}</div>
              <div><strong>Created:</strong> {new Date(section.created_at).toLocaleDateString()}</div>
              <div><strong>Updated:</strong> {new Date(section.updated_at).toLocaleDateString()}</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
