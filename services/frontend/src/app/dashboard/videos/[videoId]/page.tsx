import { Suspense } from 'react';
import { VideoPlayer } from '@/components/video_learning/student/video-player';
import { VideoComments } from '@/components/video_learning/student/video-comments';
import { VideoSidebar } from '@/components/video_learning/student/video-sidebar';
import { VideoInfo } from '@/components/video_learning/student/video-info';
import { Skeleton } from '@/ui/skeleton';
import { notFound } from 'next/navigation';
import { videoLearningService } from '@/services/video_learning/service';
import { cookies } from 'next/headers';

interface VideoDetailPageProps {
  params: {
    videoId: string;
  };
}

export default async function VideoDetailPage({
  params,
}: VideoDetailPageProps) {
  const videoId = parseInt(params.videoId);
  if (isNaN(videoId)) {
    notFound();
  }
  const videoLearning = await videoLearningService(cookies());

  const videoResult = await videoLearning.getVideo({ id: videoId });

  if (!videoResult.success) {
    if (videoResult.error.status === 404) {
      notFound();
    }
    return (
      <div className="container mx-auto px-4 py-8">
        <div className="text-center">
          <h1 className="text-2xl font-bold mb-4 text-red-600">Error</h1>
          <p>Error loading video: {videoResult.error.message}</p>
        </div>
      </div>
    );
  }

  const video = videoResult.data;

  return (
    <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
      {/* Main video content */}
      <div className="lg:col-span-2 space-y-6">
        <Suspense fallback={<Skeleton className="aspect-video w-full" />}>
          <VideoPlayer video={video} />
        </Suspense>

        <Suspense fallback={<Skeleton className="h-32 w-full" />}>
          <VideoInfo video={video} />
        </Suspense>

        <Suspense
          fallback={
            <div className="space-y-4">
              {Array.from({ length: 3 }).map((_, i) => (
                <Skeleton key={i} className="h-24 w-full" />
              ))}
            </div>
          }
        >
          <VideoComments video={video} />
        </Suspense>
      </div>

      {/* Sidebar with related videos */}
      <div className="lg:col-span-1">
        <Suspense
          fallback={
            <div className="space-y-4">
              {Array.from({ length: 5 }).map((_, i) => (
                <Skeleton key={i} className="h-24 w-full" />
              ))}
            </div>
          }
        >
          <VideoSidebar video={video} />
        </Suspense>
      </div>
    </div>
  );
}
