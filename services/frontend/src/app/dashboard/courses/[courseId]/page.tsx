import { textService } from '@/services/text/service';
import { cookies } from 'next/headers';
import Link from 'next/link';
import { Button } from '@/ui/button';
import { ArrowLeft, BookOpen, FileText, ChevronRight } from 'lucide-react';
import { notFound } from 'next/navigation';
import Image from 'next/image';
import { capitalize } from '@/lib/shared/utils';

interface StudentCoursePageProps {
  params: Promise<{ courseId: string }>;
}

export default async function StudentCoursePage({
  params,
}: StudentCoursePageProps) {
  const asyncParams = await params;
  const courseId = parseInt(asyncParams.courseId);

  if (isNaN(courseId)) {
    notFound();
  }

  const text = await textService(cookies());
  const [courseResult, sectionsResult] = await Promise.all([
    text.getCourse({ id: courseId }),
    text.listSections({ course_id: courseId }),
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

  // Obtener artículos para cada sección
  const sectionsWithArticles = await Promise.all(
    sections.map(async (section) => {
      const articlesResult = await text.listArticles({
        section_id: section.id,
      });
      return {
        ...section,
        articles: articlesResult.success ? articlesResult.data : [],
      };
    }),
  );

  return (
    <div className="container mx-auto px-4 py-8">
      {/* Header Navigation */}
      <div className="flex items-center gap-4 mb-8">
        <Button variant="outline" size="sm" asChild>
          <Link href="/dashboard/courses">
            <ArrowLeft className="h-4 w-4 mr-2" />
            Back to Courses
          </Link>
        </Button>
      </div>

      {/* Course Header */}
      <div className="bg-white rounded-lg border shadow-sm p-8 mb-8">
        <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
          {/* Course Image */}
          <div className="lg:col-span-1">
            <div className="aspect-video bg-gradient-to-br from-blue-500 to-purple-600 rounded-lg overflow-hidden">
              {course.imageUrl ? (
                <Image
                  src={course.imageUrl}
                  alt={course.title}
                  width={600}
                  height={400}
                  className="w-full h-full object-cover"
                />
              ) : (
                <div className="w-full h-full flex items-center justify-center">
                  <BookOpen className="h-16 w-16 text-white" />
                </div>
              )}
            </div>
          </div>

          {/* Course Info */}
          <div className="lg:col-span-2">
            <h1 className="text-3xl font-bold mb-4">
              {capitalize(course.title)}
            </h1>
            <p className="text-gray-600 text-lg mb-6">{course.description}</p>

            <div className="flex flex-wrap gap-4 text-sm text-gray-500">
              <div>
                <strong>Sections:</strong> {sections.length}
              </div>
              <div>
                <strong>Articles:</strong>{' '}
                {sectionsWithArticles.reduce(
                  (total, section) => total + section.articles.length,
                  0,
                )}
              </div>
              <div>
                <strong>Created:</strong>{' '}
                {new Date(course.created_at).toLocaleDateString()}
              </div>
            </div>
          </div>
        </div>
      </div>

      {/* Course Content */}
      <div className="space-y-8">
        <h2 className="text-2xl font-bold">Course Content</h2>

        {sectionsWithArticles.length === 0 ? (
          <div className="text-center py-12">
            <BookOpen className="mx-auto h-12 w-12 text-gray-400 mb-4" />
            <h3 className="text-lg font-semibold mb-2">No content available</h3>
            <p className="text-gray-600">
              This course doesn&apos;t have any sections yet.
            </p>
          </div>
        ) : (
          <div className="space-y-6">
            {sectionsWithArticles
              .sort((a, b) => a.order - b.order)
              .map((section) => (
                <div
                  key={section.id}
                  className="bg-white rounded-lg border shadow-sm overflow-hidden"
                >
                  {/* Section Header */}
                  <div className="bg-gray-50 px-6 py-4 border-b">
                    <div className="flex items-center gap-3">
                      <span className="bg-blue-100 text-blue-800 text-sm font-medium px-3 py-1 rounded-full">
                        Section {section.order}
                      </span>
                      <h3 className="text-xl font-semibold">
                        {capitalize(section.title)}
                      </h3>
                    </div>
                    <p className="text-gray-600 mt-2">{section.description}</p>
                  </div>

                  {/* Articles List */}
                  <div className="p-6">
                    {section.articles.length === 0 ? (
                      <p className="text-gray-500 italic">
                        No articles in this section yet.
                      </p>
                    ) : (
                      <div className="space-y-3">
                        {section.articles
                          .sort((a, b) => a.id - b.id)
                          .map((article, articleIndex) => (
                            <Link
                              key={article.id}
                              href={`/dashboard/courses/${courseId}/articles/${article.id}`}
                              className="flex items-center gap-4 p-4 rounded-lg border hover:bg-gray-50 transition-colors group"
                            >
                              <div className="flex-shrink-0">
                                <div className="w-8 h-8 bg-green-100 text-green-800 rounded-full flex items-center justify-center text-sm font-medium">
                                  {articleIndex + 1}
                                </div>
                              </div>

                              <div className="flex-1 min-w-0">
                                <h4 className="text-lg font-medium group-hover:text-blue-600 transition-colors">
                                  {capitalize(article.title)}
                                </h4>
                              </div>

                              <div className="flex-shrink-0">
                                <div className="flex items-center gap-2 text-sm text-gray-500">
                                  <FileText className="h-4 w-4" />
                                  <span>Read</span>
                                  <ChevronRight className="h-4 w-4 group-hover:translate-x-1 transition-transform" />
                                </div>
                              </div>
                            </Link>
                          ))}
                      </div>
                    )}
                  </div>
                </div>
              ))}
          </div>
        )}
      </div>
    </div>
  );
}
