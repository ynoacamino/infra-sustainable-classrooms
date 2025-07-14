import { textService } from '@/services/text/service';
import { cookies } from 'next/headers';
import Link from 'next/link';
import { Button } from '@/ui/button';
import { ArrowLeft, Edit, Plus, FileText } from 'lucide-react';
import { notFound } from 'next/navigation';

interface SectionPageProps {
  params: Promise<{ courseId: string; sectionId: string }>;
}

export default async function SectionPage({ params }: SectionPageProps) {
  const asyncParams = await params;

  const courseId = parseInt(asyncParams.courseId);
  const sectionId = parseInt(asyncParams.sectionId);

  if (isNaN(courseId) || isNaN(sectionId)) {
    notFound();
  }

  const text = await textService(cookies());
  const [sectionResult, articlesResult] = await Promise.all([
    text.getSection({ id: sectionId }),
    text.listArticles({ section_id: sectionId }),
  ]);

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

  if (!articlesResult.success) {
    return (
      <div className="container mx-auto px-4 py-8">
        <div className="text-center">
          <h1 className="text-2xl font-bold mb-4 text-red-600">Error</h1>
          <p>Error loading articles: {articlesResult.error.message}</p>
        </div>
      </div>
    );
  }

  const section = sectionResult.data;
  const articles = articlesResult.data;

  return (
    <div className="container mx-auto px-4 py-8">
      <div className="flex flex-col items-start justify-center gap-4 mb-8">
        <Button variant="outline" size="sm" asChild>
          <Link href={`/teacher/courses/${courseId}`}>
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
        {/* Section Details */}
        <div className="lg:col-span-2 flex flex-col gap-6">
          <div className="bg-white rounded-lg border shadow-sm p-6">
            <div className="flex justify-between items-start mb-4">
              <div>
                <h2 className="text-2xl font-semibold">{section.title}</h2>
                <p className="text-gray-600 mt-2">{section.description}</p>
                <div className="text-sm text-gray-500 mt-2">
                  Section {section.order}
                </div>
              </div>
              <Button variant="outline" size="sm" asChild>
                <Link
                  href={`/teacher/courses/${courseId}/sections/${sectionId}/edit`}
                >
                  <Edit className="h-4 w-4 mr-2" />
                  Edit Section
                </Link>
              </Button>
            </div>
          </div>

          <div className="flex flex-col gap-2">
            <span className="text-2xl font-semibold block">Articles</span>
            <div>
              {articles.length === 0 ? (
                <div className="text-center py-12">
                  <h3 className="text-lg font-semibold mb-4">
                    No articles yet
                  </h3>
                  <p className="text-gray-600 mb-6">
                    Add your first article to this section
                  </p>
                  <Button asChild>
                    <Link
                      href={`/teacher/courses/${courseId}/sections/${sectionId}/articles/new`}
                    >
                      <Plus className="h-4 w-4 mr-2" />
                      Add Article
                    </Link>
                  </Button>
                </div>
              ) : (
                <div className="space-y-4">
                  {articles
                    .sort((a, b) => a.id - b.id)
                    .map((article) => (
                      <div
                        key={article.id}
                        className="bg-white rounded-lg border shadow-sm p-6"
                      >
                        <div className="flex justify-between items-start">
                          <div className="flex-1">
                            <div className="flex items-center gap-3 mb-2">
                              <span className="bg-green-100 text-green-800 text-sm font-medium px-2.5 py-0.5 rounded">
                                Article #{article.id}
                              </span>
                              <h3 className="text-xl font-semibold">
                                {article.title}
                              </h3>
                            </div>
                            <div
                              className="text-gray-600 mb-4 prose mt-10"
                              dangerouslySetInnerHTML={{
                                __html: article.content.substring(0, 60),
                              }}
                            ></div>
                            <div className="text-sm text-gray-500">
                              Created:{' '}
                              {new Date(
                                article.created_at,
                              ).toLocaleDateString()}
                            </div>
                          </div>
                          <div className="flex gap-2 ml-4">
                            <Button asChild variant="outline" size="sm">
                              <Link
                                href={`/teacher/courses/${courseId}/sections/${sectionId}/articles/${article.id}`}
                              >
                                View
                              </Link>
                            </Button>
                            <Button asChild variant="outline" size="sm">
                              <Link
                                href={`/teacher/courses/${courseId}/sections/${sectionId}/articles/${article.id}/edit`}
                              >
                                <Edit className="h-4 w-4 mr-2" />
                                Edit
                              </Link>
                            </Button>
                          </div>
                        </div>
                      </div>
                    ))}
                </div>
              )}
            </div>
          </div>
        </div>

        {/* Section Actions */}
        <div className="space-y-4">
          <div className="bg-white rounded-lg border shadow-sm p-6">
            <h3 className="text-lg font-semibold mb-4">Section Management</h3>
            <div className="space-y-3">
              <Button asChild className="w-full">
                <Link
                  href={`/teacher/courses/${courseId}/sections/${sectionId}/articles`}
                >
                  <FileText className="h-4 w-4 mr-2" />
                  View Articles
                </Link>
              </Button>
              <Button asChild variant="outline" className="w-full">
                <Link
                  href={`/teacher/courses/${courseId}/sections/${sectionId}/articles/new`}
                >
                  <Plus className="h-4 w-4 mr-2" />
                  Add Article
                </Link>
              </Button>
              <Button asChild variant="outline" className="w-full">
                <Link
                  href={`/teacher/courses/${courseId}/sections/${sectionId}/edit`}
                >
                  <Edit className="h-4 w-4 mr-2" />
                  Edit Section
                </Link>
              </Button>
            </div>
          </div>

          {/* Section Info */}
          <div className="bg-white rounded-lg border shadow-sm p-6">
            <h3 className="text-lg font-semibold mb-4">Section Information</h3>
            <div className="space-y-2 text-sm">
              <div>
                <strong>Order:</strong> {section.order}
              </div>
              <div>
                <strong>Course ID:</strong> {section.course_id}
              </div>
              <div>
                <strong>Created:</strong>{' '}
                {new Date(section.created_at).toLocaleDateString()}
              </div>
              <div>
                <strong>Updated:</strong>{' '}
                {new Date(section.updated_at).toLocaleDateString()}
              </div>
              <div>
                <strong>Section ID:</strong> {section.id}
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
