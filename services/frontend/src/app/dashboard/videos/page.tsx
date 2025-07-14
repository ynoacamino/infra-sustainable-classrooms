import { Suspense } from 'react';
import { VideoSearch } from '@/components/video_learning/student/video-search';
import { VideoRecommendations } from '@/components/video_learning/student/video-recommendations';
import { VideoCategories } from '@/components/video_learning/student/video-categories';
import H1 from '@/ui/h1';
import Section from '@/ui/section';
import { Skeleton } from '@/ui/skeleton';

export default function VideosPage() {
  return (
    <div className="space-y-6">
      <H1>Video Learning</H1>

      <Section title="Search Videos">
        <Suspense fallback={<Skeleton className="h-10 w-full" />}>
          <VideoSearch />
        </Suspense>
      </Section>

      <Section title="Categories">
        <Suspense
          fallback={
            <div className="grid grid-cols-2 md:grid-cols-4 gap-4">
              {Array.from({ length: 8 }).map((_, i) => (
                <Skeleton key={i} className="h-24 w-full" />
              ))}
            </div>
          }
        >
          <VideoCategories />
        </Suspense>
      </Section>

      <Section title="Recommended for You">
        <Suspense
          fallback={
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
              {Array.from({ length: 6 }).map((_, i) => (
                <Skeleton key={i} className="h-48 w-full" />
              ))}
            </div>
          }
        >
          <VideoRecommendations />
        </Suspense>
      </Section>
    </div>
  );
}
