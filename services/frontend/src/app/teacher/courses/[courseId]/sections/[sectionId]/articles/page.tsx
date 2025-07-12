import { textService } from "@/services/text/service";
import { cookies } from "next/headers";
import Link from "next/link";
import { Button } from "@/ui/button";
import { ArrowLeft, Edit, Plus, Eye } from "lucide-react";
import { notFound } from "next/navigation";

interface ArticlesPageProps {
  params: { courseId: string; sectionId: string };
}

export default async function ArticlesPage({ params }: ArticlesPageProps) {
  const courseId = parseInt(params.courseId);
  const sectionId = parseInt(params.sectionId);
  
  if (isNaN(courseId) || isNaN(sectionId)) {
    notFound();
  }

  const text = await textService(cookies());
  const [sectionResult, articlesResult] = await Promise.all([
    text.getSection({ section_id: sectionId }),
    text.listArticles({ section_id: sectionId })
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
      <div className="flex items-center gap-4 mb-8">
        <Button variant="outline" size="sm" asChild>
          <Link href={`/teacher/courses/${courseId}/sections/${sectionId}`}>
            <ArrowLeft className="h-4 w-4 mr-2" />
            Back to Section
          </Link>
        </Button>
        <div>
          <h1 className="text-3xl font-bold">Articles</h1>
          <p className="text-gray-600">{section.title}</p>
        </div>
      </div>

      <div className="flex justify-between items-center mb-6">
        <h2 className="text-xl font-semibold">Section Articles</h2>
        <Button asChild>
          <Link href={`/teacher/courses/${courseId}/sections/${sectionId}/articles/new`}>
            <Plus className="h-4 w-4 mr-2" />
            Add Article
          </Link>
        </Button>
      </div>

      {articles.length === 0 ? (
        <div className="text-center py-12">
          <h3 className="text-lg font-semibold mb-4">No articles yet</h3>
          <p className="text-gray-600 mb-6">Add your first article to this section</p>
          <Button asChild>
            <Link href={`/teacher/courses/${courseId}/sections/${sectionId}/articles/new`}>
              <Plus className="h-4 w-4 mr-2" />
              Add Article
            </Link>
          </Button>
        </div>
      ) : (
        <div className="space-y-4">
          {articles.map(article => (
            <div key={article.id} className="bg-white rounded-lg border shadow-sm p-6">
              <div className="flex justify-between items-start">
                <div className="flex-1">
                  <h3 className="text-xl font-semibold mb-2">{article.title}</h3>
                  <p className="text-gray-600 mb-4 line-clamp-3">{article.content}</p>
                  <div className="text-sm text-gray-500">
                    Created: {new Date(article.created_at).toLocaleDateString()}
                  </div>
                </div>
                <div className="flex gap-2 ml-4">
                  <Button asChild variant="outline" size="sm">
                    <Link href={`/teacher/courses/${courseId}/sections/${sectionId}/articles/${article.id}`}>
                      <Eye className="h-4 w-4 mr-2" />
                      View
                    </Link>
                  </Button>
                  <Button asChild variant="outline" size="sm">
                    <Link href={`/teacher/courses/${courseId}/sections/${sectionId}/articles/${article.id}/edit`}>
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
  );
}
