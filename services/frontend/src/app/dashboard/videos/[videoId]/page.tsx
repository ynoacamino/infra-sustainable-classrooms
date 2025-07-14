import { Suspense } from 'react';
import { VideoPlayer } from '@/components/video_learning/student/video-player';
import { VideoComments } from '@/components/video_learning/student/video-comments';
import { VideoSidebar } from '@/components/video_learning/student/video-sidebar';
import { VideoInfo } from '@/components/video_learning/student/video-info';
import { Skeleton } from '@/ui/skeleton';

interface VideoDetailPageProps {
  params: {
    videoId: string;
  };
}

export default function VideoDetailPage({ params }: VideoDetailPageProps) {
  const videoId = parseInt(params.videoId);

  return (
    <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
      {/* Main video content */}
      <div className="lg:col-span-2 space-y-6">
        <Suspense fallback={<Skeleton className="aspect-video w-full" />}>
          <VideoPlayer videoId={videoId} />
        </Suspense>

        <Suspense fallback={<Skeleton className="h-32 w-full" />}>
          <VideoInfo videoId={videoId} />
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
          <VideoComments videoId={videoId} />
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
          <VideoSidebar videoId={videoId} />
        </Suspense>
      </div>
    </div>
  );
}
