import { textService } from '@/services/text/service';
import { cookies } from 'next/headers';
import Link from 'next/link';
import { Button } from '@/ui/button';
import { ArrowLeft, Edit } from 'lucide-react';
import { notFound } from 'next/navigation';

interface ArticlePageProps {
  params: { courseId: string; sectionId: string; articleId: string };
}

export default async function ArticlePage({ params }: ArticlePageProps) {
  const courseId = parseInt(params.courseId);
  const sectionId = parseInt(params.sectionId);
  const articleId = parseInt(params.articleId);

  if (isNaN(courseId) || isNaN(sectionId) || isNaN(articleId)) {
    notFound();
  }

  const text = await textService(cookies());
  const articleResult = await text.getArticle({ article_id: articleId });

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

  const article = articleResult.data;

  return (
    <div className="container mx-auto px-4 py-8">
      <div className="flex flex-col items-start justify-center gap-4 mb-8">
        <Button variant="outline" size="sm" asChild>
          <Link href={`/teacher/courses/${courseId}/sections/${sectionId}`}>
            <ArrowLeft className="h-4 w-4 mr-2" />
            Back to Articles
          </Link>
        </Button>
        <h1 className="text-3xl font-bold">{article.title}</h1>
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-4 gap-8">
        <div className="lg:col-span-3">
          <div className="bg-white rounded-lg border shadow-sm p-8">
            <h1 className="text-3xl font-bold mb-6">{article.title}</h1>
            <div className="prose max-w-none">
              <div className="whitespace-pre-wrap">{article.content}</div>
            </div>
          </div>
        </div>

        <div className="space-y-4">
          <div className="bg-white rounded-lg border shadow-sm p-6">
            <h3 className="text-lg font-semibold mb-4">Article Management</h3>
            <div className="space-y-3">
              <Button asChild variant="outline" className="w-full">
                <Link
                  href={`/teacher/courses/${courseId}/sections/${sectionId}/articles/${articleId}/edit`}
                >
                  <Edit className="h-4 w-4 mr-2" />
                  Edit Article
                </Link>
              </Button>
            </div>
          </div>

          <div className="bg-white rounded-lg border shadow-sm p-6">
            <h3 className="text-lg font-semibold mb-4">Article Information</h3>
            <div className="space-y-2 text-sm">
              <div>
                <strong>Article ID:</strong> {article.id}
              </div>
              <div>
                <strong>Section ID:</strong> {article.section_id}
              </div>
              <div>
                <strong>Created:</strong>{' '}
                {new Date(article.created_at).toLocaleDateString()}
              </div>
              <div>
                <strong>Updated:</strong>{' '}
                {new Date(article.updated_at).toLocaleDateString()}
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
