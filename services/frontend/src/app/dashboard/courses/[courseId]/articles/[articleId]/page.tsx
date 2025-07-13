import { textService } from '@/services/text/service';
import { cookies } from 'next/headers';
import Link from 'next/link';
import { Button } from '@/ui/button';
import { ArrowLeft, BookOpen, Clock } from 'lucide-react';
import { notFound } from 'next/navigation';

interface StudentArticlePageProps {
  params: { courseId: string; articleId: string };
}

export default async function StudentArticlePage({
  params,
}: StudentArticlePageProps) {
  const courseId = parseInt(params.courseId);
  const articleId = parseInt(params.articleId);

  if (isNaN(courseId) || isNaN(articleId)) {
    notFound();
  }

  const text = await textService(cookies());
  const [courseResult, articleResult] = await Promise.all([
    text.getCourse({ id: courseId }),
    text.getArticle({ id: articleId }),
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

  if (!articleResult.success) {
    if (articleResult.error.status === 404) {
      notFound();
    }
    return (
      <div className="container mx-auto px-4 py-8">
        <div className="text-center">
          <h1 className="text-2xl font-bold mb-4 text-red-600">Error</h1>
          <p>Error loading article: {articleResult.error.message}</p>
        </div>
      </div>
    );
  }

  const course = courseResult.data;
  const article = articleResult.data;

  // Get section info
  const sectionResult = await text.getSection({
    id: article.section_id,
  });
  const section = sectionResult.success ? sectionResult.data : null;

  return (
    <div className="container mx-auto px-4 py-8">
      {/* Header Navigation */}
      <div className="flex items-center gap-4 mb-8">
        <Button variant="outline" size="sm" asChild>
          <Link href={`/dashboard/courses/${courseId}`}>
            <ArrowLeft className="h-4 w-4 mr-2" />
            Back to Course
          </Link>
        </Button>
      </div>

      {/* Article Header */}
      <div className="bg-white rounded-lg border shadow-sm p-8 mb-8">
        <div className="mb-6">
          {/* Breadcrumb */}
          <div className="flex items-center gap-2 text-sm text-gray-500 mb-4">
            <Link
              href={`/dashboard/courses/${courseId}`}
              className="hover:text-blue-600"
            >
              {course.title}
            </Link>
            {section && (
              <>
                <span>/</span>
                <span>
                  Section {section.order}: {section.title}
                </span>
              </>
            )}
            <span>/</span>
            <span className="text-gray-900">{article.title}</span>
          </div>

          <h1 className="text-3xl font-bold mb-4">{article.title}</h1>

          <div className="flex items-center gap-6 text-sm text-gray-500">
            <div className="flex items-center gap-2">
              <BookOpen className="h-4 w-4" />
              <span>{course.title}</span>
            </div>
            {section && (
              <div className="flex items-center gap-2">
                <span className="bg-blue-100 text-blue-800 px-2 py-1 rounded text-xs font-medium">
                  Section {section.order}
                </span>
                <span>{section.title}</span>
              </div>
            )}
            <div className="flex items-center gap-2">
              <Clock className="h-4 w-4" />
              <span>
                Updated {new Date(article.updated_at).toLocaleDateString()}
              </span>
            </div>
          </div>
        </div>
      </div>

      {/* Article Content */}
      <div className="bg-white rounded-lg border shadow-sm p-8">
        <div className="prose prose-lg max-w-none">
          {/* Format article content with proper paragraph breaks */}
          {article.content.split('\n').map((paragraph, index) =>
            paragraph.trim() ? (
              <p key={index} className="mb-4 text-gray-700 leading-relaxed">
                {paragraph.trim()}
              </p>
            ) : (
              <div key={index} className="mb-4"></div>
            ),
          )}
        </div>

        {/* Article Footer */}
        <div className="mt-8 pt-6 border-t">
          <div className="flex items-center justify-between">
            <div className="text-sm text-gray-500">
              <div>
                Created: {new Date(article.created_at).toLocaleDateString()}
              </div>
              <div>
                Last updated:{' '}
                {new Date(article.updated_at).toLocaleDateString()}
              </div>
            </div>

            <Button asChild>
              <Link href={`/dashboard/courses/${courseId}`}>
                Continue Learning
              </Link>
            </Button>
          </div>
        </div>
      </div>
    </div>
  );
}
