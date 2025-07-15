import { Suspense } from 'react';
import H1 from '@/ui/h1';
import Section from '@/ui/section';
import { Skeleton } from '@/ui/skeleton';
import { VideosByCategory } from '@/components/video_learning/student/videos-by-category';
import { notFound } from 'next/navigation';
import { videoLearningService } from '@/services/video_learning/service';
import { cookies } from 'next/headers';

interface CategoryPageProps {
  params: Promise<{
    categoryId: string;
  }>;
}

export default async function CategoryPage({ params }: CategoryPageProps) {
  const asyncParams = await params;
  const categoryId = parseInt(asyncParams.categoryId);

  if (isNaN(categoryId)) {
    return notFound();
  }

  const videoLearning = await videoLearningService(cookies());
  const categoryResult = await videoLearning.getCategory({ id: categoryId });

  if (!categoryResult.success) {
    if (categoryResult.error.status === 404) {
      return notFound();
    }
    return (
      <div className="container mx-auto px-4 py-8">
        <div className="text-center">
          <h1 className="text-2xl font-bold mb-4 text-red-600">Error</h1>
          <p>Error loading category: {categoryResult.error.message}</p>
        </div>
      </div>
    );
  }

  return (
    <div className="space-y-6">
      <H1>Category Videos</H1>

      <Section title="Videos in Category">
        <Suspense
          fallback={
            <div className="space-y-4">
              <div className="flex items-center justify-between">
                <Skeleton className="h-6 w-48" />
                <Skeleton className="h-10 w-32" />
              </div>
              <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
                {Array.from({ length: 9 }).map((_, i) => (
                  <Skeleton key={i} className="h-64 w-full" />
                ))}
              </div>
            </div>
          }
        >
          <VideosByCategory category={categoryResult.data} />
        </Suspense>
      </Section>
    </div>
  );
}
