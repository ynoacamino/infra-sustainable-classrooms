import { textService } from "@/services/text/service";
import { cookies } from "next/headers";
import Link from "next/link";
import { Button } from "@/ui/button";
import { UpdateSectionForm } from "@/components/text/forms";
import { ArrowLeft } from "lucide-react";
import { notFound } from "next/navigation";

interface EditSectionPageProps {
  params: { courseId: string; sectionId: string };
}

export default async function EditSectionPage({ params }: EditSectionPageProps) {
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
          <Link href={`/teacher/courses/${courseId}/sections/${sectionId}`}>
            <ArrowLeft className="h-4 w-4 mr-2" />
            Back to Section
          </Link>
        </Button>
        <h1 className="text-3xl font-bold">Edit Section: {section.title}</h1>
      </div>

      <div className="max-w-2xl mx-auto">
        <UpdateSectionForm section={section} />
      </div>
    </div>
  );
}
