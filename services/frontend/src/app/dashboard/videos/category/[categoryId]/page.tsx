import { Suspense } from 'react';
import H1 from '@/ui/h1';
import Section from '@/ui/section';
import { Skeleton } from '@/ui/skeleton';
import { VideosByCategory } from '@/components/video_learning/student/videos-by-category';

interface CategoryPageProps {
  params: {
    categoryId: string;
  };
}

export default function CategoryPage({ params }: CategoryPageProps) {
  const categoryId = parseInt(params.categoryId);

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
          <VideosByCategory categoryId={categoryId} />
        </Suspense>
      </Section>
    </div>
  );
}
