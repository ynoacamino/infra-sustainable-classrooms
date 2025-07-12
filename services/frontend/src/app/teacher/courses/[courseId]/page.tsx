import { textService } from "@/services/text/service";
import { cookies } from "next/headers";
import Link from "next/link";
import { Button } from "@/ui/button";
import { ArrowLeft, Edit, Plus } from "lucide-react";
import { notFound } from "next/navigation";

interface CoursePageProps {
  params: { courseId: string };
}

export default async function CoursePage({ params }: CoursePageProps) {
  const courseId = parseInt(params.courseId);
  
  if (isNaN(courseId)) {
    notFound();
  }

  const text = await textService(cookies());
  const courseResult = await text.getCourse({ course_id: courseId });

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
          <Link href="/teacher/courses">
            <ArrowLeft className="h-4 w-4 mr-2" />
            Back to Courses
          </Link>
        </Button>
        <h1 className="text-3xl font-bold">{course.title}</h1>
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
        {/* Course Details */}
        <div className="lg:col-span-2">
          <div className="bg-white rounded-lg border shadow-sm p-6">
            <div className="flex justify-between items-start mb-4">
              <div>
                <h2 className="text-2xl font-semibold">{course.title}</h2>
                <p className="text-gray-600 mt-2">{course.description}</p>
              </div>
              <Button variant="outline" size="sm" asChild>
                <Link href={`/teacher/courses/${course.id}/edit`}>
                  <Edit className="h-4 w-4 mr-2" />
                  Edit Course
                </Link>
              </Button>
            </div>
            {course.imageUrl && (
              <img 
                src={course.imageUrl} 
                alt={course.title}
                className="w-full h-64 object-cover rounded-md mt-4"
              />
            )}
          </div>
        </div>

        {/* Course Actions */}
        <div className="space-y-4">
          <div className="bg-white rounded-lg border shadow-sm p-6">
            <h3 className="text-lg font-semibold mb-4">Course Management</h3>
            <div className="space-y-3">
              <Button asChild className="w-full">
                <Link href={`/teacher/courses/${course.id}/sections`}>
                  View Sections
                </Link>
              </Button>
              <Button asChild variant="outline" className="w-full">
                <Link href={`/teacher/courses/${course.id}/sections/new`}>
                  <Plus className="h-4 w-4 mr-2" />
                  Add Section
                </Link>
              </Button>
              <Button asChild variant="outline" className="w-full">
                <Link href={`/teacher/courses/${course.id}/edit`}>
                  <Edit className="h-4 w-4 mr-2" />
                  Edit Course
                </Link>
              </Button>
            </div>
          </div>

          {/* Course Info */}
          <div className="bg-white rounded-lg border shadow-sm p-6">
            <h3 className="text-lg font-semibold mb-4">Course Information</h3>
            <div className="space-y-2 text-sm">
              <div>
                <strong>Created:</strong>{" "}
                {new Date(course.created_at).toLocaleDateString()}
              </div>
              <div>
                <strong>Last Updated:</strong>{" "}
                {new Date(course.updated_at).toLocaleDateString()}
              </div>
              <div>
                <strong>Course ID:</strong> {course.id}
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
