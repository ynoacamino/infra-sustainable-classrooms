import { textService } from '@/services/text/service';
import { cookies } from 'next/headers';
import Link from 'next/link';
import { Button } from '@/ui/button';
import { UpdateArticleForm } from '@/components/text/forms';
import { ArrowLeft } from 'lucide-react';
import { notFound } from 'next/navigation';

interface EditArticlePageProps {
  params: { courseId: string; sectionId: string; articleId: string };
}

export default async function EditArticlePage({
  params,
}: EditArticlePageProps) {
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
      <div className="flex items-center gap-4 mb-8">
        <Button variant="outline" size="sm" asChild>
          <Link
            href={`/teacher/courses/${courseId}/sections/${sectionId}/articles/${articleId}`}
          >
            <ArrowLeft className="h-4 w-4 mr-2" />
            Back to Article
          </Link>
        </Button>
        <h1 className="text-3xl font-bold">Edit Article: {article.title}</h1>
      </div>

      <div className="max-w-2xl mx-auto">
        <UpdateArticleForm article={article} />
      </div>
    </div>
  );
}
