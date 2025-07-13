import { textService } from "@/services/text/service";
import { cookies } from "next/headers";
import Link from "next/link";
import { Button } from "@/ui/button";

export default async function CoursesPage() {
  const text = await textService(cookies());
  const courses = await text.listCourses({});

  if (!courses.success) {
    return (
      <div className="flex flex-col items-center justify-center w-full h-full">
        <h1 className="text-2xl font-bold mb-4 text-red-600">Error</h1>
        <p>Error loading courses: {courses.error.message}</p>
      </div>
    );
  }

  return (
    <div className="container mx-auto px-4 py-8">
      <div className="flex justify-between items-center mb-8">
        <h1 className="text-3xl font-bold">My Courses</h1>
        <Button asChild>
          <Link href="/teacher/courses/new">Create New Course</Link>
        </Button>
      </div>

      {courses.data.length === 0 ? (
        <div className="text-center py-12">
          <h2 className="text-xl font-semibold mb-4">No courses yet</h2>
          <p className="text-gray-600 mb-6">Create your first course to get started</p>
          <Button asChild>
            <Link href="/teacher/courses/new">Create Course</Link>
          </Button>
        </div>
      ) : (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {courses.data.map(course => (
            <div key={course.id} className="bg-white rounded-lg border shadow-sm hover:shadow-lg transition-shadow overflow-hidden">
              <div className="p-6">
                <h3 className="text-xl font-semibold mb-2">{course.title}</h3>
                <p className="text-gray-600 mb-4">{course.description}</p>
              </div>
              {course.imageUrl && (
                <div className="px-6 pb-4">
                  <img 
                    src={course.imageUrl} 
                    alt={course.title}
                    className="w-full h-48 object-cover rounded-md"
                  />
                </div>
              )}
              <div className="px-6 pb-6 flex gap-2">
                <Button asChild variant="outline" size="sm">
                  <Link href={`/teacher/courses/${course.id}`}>View</Link>
                </Button>
                <Button asChild variant="outline" size="sm">
                  <Link href={`/teacher/courses/${course.id}/edit`}>Edit</Link>
                </Button>
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
}