import { textService } from '@/services/text/service';
import { cookies } from 'next/headers';
import Link from 'next/link';
import { Button } from '@/ui/button';
import { BookOpen, Clock, ChevronRight } from 'lucide-react';
import Image from 'next/image';

export default async function CoursesPage() {
  const text = await textService(cookies());
  const coursesResult = await text.listCourses();

  if (!coursesResult.success) {
    return (
      <div className="container mx-auto px-4 py-8">
        <div className="text-center">
          <h1 className="text-2xl font-bold mb-4 text-red-600">Error</h1>
          <p>Error loading courses: {coursesResult.error.message}</p>
        </div>
      </div>
    );
  }

  const courses = coursesResult.data;

  return (
    <div className="container mx-auto px-4 py-8">
      <div className="mb-8">
        <h1 className="text-3xl font-bold mb-2">Available Courses</h1>
        <p className="text-gray-600">
          Explore and learn from our comprehensive course catalog
        </p>
      </div>

      {courses.length === 0 ? (
        <div className="text-center py-12">
          <BookOpen className="mx-auto h-12 w-12 text-gray-400 mb-4" />
          <h3 className="text-lg font-semibold mb-2">No courses available</h3>
          <p className="text-gray-600">Check back later for new courses!</p>
        </div>
      ) : (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {courses.map((course) => (
            <div
              key={course.id}
              className="bg-white rounded-lg border shadow-sm hover:shadow-md transition-shadow overflow-hidden"
            >
              {/* Course Image */}
              <div className="aspect-video bg-gradient-to-br from-blue-500 to-purple-600 relative">
                {course.imageUrl ? (
                  <Image
                    src={course.imageUrl}
                    alt={course.title}
                    width={400}
                    height={225}
                    className="w-full h-full object-cover"
                  />
                ) : (
                  <div className="w-full h-full flex items-center justify-center">
                    <BookOpen className="h-12 w-12 text-white" />
                  </div>
                )}
              </div>

              {/* Course Content */}
              <div className="p-6">
                <h3 className="text-xl font-semibold mb-2 line-clamp-2">
                  {course.title}
                </h3>
                <p className="text-gray-600 mb-4 line-clamp-3">
                  {course.description}
                </p>

                <div className="flex items-center justify-between text-sm text-gray-500 mb-4">
                  <div className="flex items-center gap-1">
                    <Clock className="h-4 w-4" />
                    <span>
                      Updated {new Date(course.updated_at).toLocaleDateString()}
                    </span>
                  </div>
                </div>

                <Button asChild className="w-full">
                  <Link href={`/dashboard/courses/${course.id}`}>
                    Start Learning
                    <ChevronRight className="h-4 w-4 ml-2" />
                  </Link>
                </Button>
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
}
