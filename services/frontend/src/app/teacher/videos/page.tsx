'use client';

import { Suspense } from 'react';
import { VideoUpload } from '@/components/video_learning/teacher/video-upload';
import { MyVideos } from '@/components/video_learning/teacher/my-videos';
import { VideoManagement } from '@/components/video_learning/teacher/video-management';
import { VideoStats } from '@/components/video_learning/teacher/video-stats';
import H1 from '@/ui/h1';
import Section from '@/ui/section';
import { Skeleton } from '@/ui/skeleton';
import { useSWRAll } from '@/lib/shared/swr/utils';
import {
  useGetAllCategories,
  useGetAllTags,
} from '@/hooks/video_learning/useSWR';

export default function TeacherVideosPage() {
  const {
    // isLoading,
    data: [categories, tags],
    // errors,
    // mutateAll,
  } = useSWRAll([useGetAllCategories(), useGetAllTags()]);
  // Fetch categories and tags on server side
  if (!categories || !tags) {
    return (
      <div className="space-y-6">
        <H1>Video Management</H1>
        <Section title="Loading...">
          <Skeleton className="h-96 w-full" />
        </Section>
      </div>
    );
  }

  return (
    <div className="space-y-6">
      <H1>Video Management</H1>

      <Section title="Upload New Video">
        <Suspense fallback={<Skeleton className="h-96 w-full" />}>
          <VideoUpload />
        </Suspense>
      </Section>

      <Section title="Video Statistics">
        <Suspense
          fallback={
            <div className="grid grid-cols-1 md:grid-cols-4 gap-4">
              {Array.from({ length: 4 }).map((_, i) => (
                <Skeleton key={i} className="h-24 w-full" />
              ))}
            </div>
          }
        >
          <VideoStats />
        </Suspense>
      </Section>

      <Section title="My Videos">
        <Suspense
          fallback={
            <div className="space-y-4">
              {Array.from({ length: 5 }).map((_, i) => (
                <Skeleton key={i} className="h-32 w-full" />
              ))}
            </div>
          }
        >
          <MyVideos />
        </Suspense>
      </Section>

      <Section title="Content Management">
        <Suspense
          fallback={
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              {Array.from({ length: 4 }).map((_, i) => (
                <Skeleton key={i} className="h-48 w-full" />
              ))}
            </div>
          }
        >
          <VideoManagement categories={categories} tags={tags} />
        </Suspense>
      </Section>
    </div>
  );
}
