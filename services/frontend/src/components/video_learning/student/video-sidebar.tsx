'use client';

import { useState } from 'react';
import { PlayCircle, Clock, Eye, TrendingUp } from 'lucide-react';
import { Button } from '@/ui/button';
import { Link } from '@/ui/link';
import type { Video, VideoDetails } from '@/types/video_learning/models';
import { Skeleton } from '@/ui/skeleton';
import {
  useGetSimilarVideos,
  useGetRecommendations,
} from '@/hooks/video_learning/useSWR';
import { useSWRAll } from '@/lib/shared/swr/utils';
import Image from 'next/image';
import { mapToFile } from '@/lib/shared/files/utils';

interface VideoSidebarProps {
  video: VideoDetails;
}

export function VideoSidebar({ video }: VideoSidebarProps) {
  const [activeTab, setActiveTab] = useState<'similar' | 'recommended'>(
    'similar',
  );

  const {
    isLoading,
    data: [similarVideos, recommendedVideosResult],
    errors,
  } = useSWRAll([
    useGetSimilarVideos({ id: video.id, amount: 10 }),
    useGetRecommendations({ amount: 10 }),
  ]);

  const formatViews = (views: number) => {
    if (views >= 1000000) {
      return `${(views / 1000000).toFixed(1)}M`;
    } else if (views >= 1000) {
      return `${(views / 1000).toFixed(1)}K`;
    }
    return views.toString();
  };

  const formatDuration = (seconds: number) => {
    const minutes = Math.floor(seconds / 60);
    const remainingSeconds = seconds % 60;
    return `${minutes}:${remainingSeconds.toString().padStart(2, '0')}`;
  };

  const renderVideoItem = (video: Video) => (
    <div key={video.id} className="group">
      <Link href={`/dashboard/videos/${video.id}`} className="block">
        <div className="flex gap-3 p-3 rounded-lg hover:bg-muted/50 transition-colors">
          {/* Thumbnail */}
          <div className="relative flex-shrink-0 w-24 h-16 bg-muted rounded overflow-hidden">
            <Image
              src={mapToFile(video.thumbnail_url)}
              alt={video.title}
              width={400}
              height={225}
              className="w-full h-full object-cover group-hover:scale-105 transition-transform"
            />
            <div className="absolute inset-0 bg-black/20 flex items-center justify-center opacity-0 group-hover:opacity-100 transition-opacity">
              <PlayCircle className="h-6 w-6 text-white" />
            </div>
          </div>

          {/* Video Info */}
          <div className="flex-1 min-w-0">
            <h4 className="font-medium text-sm line-clamp-2 group-hover:text-primary transition-colors mb-1">
              {video.title}
            </h4>
            <p className="text-xs text-muted-foreground mb-2">{video.author}</p>
            <div className="flex items-center gap-2 text-xs text-muted-foreground">
              <div className="flex items-center gap-1">
                <Eye className="h-3 w-3" />
                <span>{formatViews(video.views)}</span>
              </div>
              <div className="flex items-center gap-1">
                <Clock className="h-3 w-3" />
                <span>{formatDuration(120)}</span> {/* Placeholder duration */}
              </div>
            </div>
          </div>
        </div>
      </Link>
    </div>
  );

  const renderLoadingSkeleton = () => (
    <div className="space-y-3">
      {Array.from({ length: 5 }).map((_, i) => (
        <div key={i} className="flex gap-3 p-3">
          <Skeleton className="w-24 h-16 rounded" />
          <div className="flex-1 space-y-2">
            <Skeleton className="h-4 w-full" />
            <Skeleton className="h-3 w-16" />
            <Skeleton className="h-3 w-20" />
          </div>
        </div>
      ))}
    </div>
  );

  const recommendedVideos = recommendedVideosResult.videos || [];

  return (
    <div className="space-y-4">
      {/* Tab Navigation */}
      <div className="flex border-b">
        <Button
          variant={activeTab === 'similar' ? 'default' : 'ghost'}
          size="sm"
          onClick={() => setActiveTab('similar')}
          className="flex-1 rounded-b-none"
        >
          <TrendingUp className="h-4 w-4 mr-2" />
          Similar Videos
        </Button>
        <Button
          variant={activeTab === 'recommended' ? 'default' : 'ghost'}
          size="sm"
          onClick={() => setActiveTab('recommended')}
          className="flex-1 rounded-b-none"
        >
          <PlayCircle className="h-4 w-4 mr-2" />
          Recommended
        </Button>
      </div>

      {/* Video List */}
      <div className="max-h-[600px] overflow-y-auto">
        {isLoading || errors.length > 0 ? (
          renderLoadingSkeleton()
        ) : (
          <div className="space-y-2">
            {activeTab === 'similar' && (
              <>
                {similarVideos && similarVideos.length > 0 ? (
                  <div className="space-y-1">
                    {similarVideos.map(renderVideoItem)}
                  </div>
                ) : (
                  <div className="text-center py-8">
                    <TrendingUp className="h-12 w-12 mx-auto mb-4 text-muted-foreground" />
                    <p className="text-sm text-muted-foreground">
                      No similar videos found
                    </p>
                  </div>
                )}
              </>
            )}

            {activeTab === 'recommended' && (
              <>
                {recommendedVideos && recommendedVideos.length > 0 ? (
                  <div className="space-y-1">
                    {recommendedVideos.map(renderVideoItem)}
                  </div>
                ) : (
                  <div className="text-center py-8">
                    <PlayCircle className="h-12 w-12 mx-auto mb-4 text-muted-foreground" />
                    <p className="text-sm text-muted-foreground">
                      No recommendations available
                    </p>
                  </div>
                )}
              </>
            )}
          </div>
        )}
      </div>

      {/* Quick Actions */}
      <div className="pt-4 border-t">
        <div className="space-y-2">
          <Link href="/dashboard/videos" className="w-full">
            <Button variant="outline" size="sm" className="w-full">
              Browse All Videos
            </Button>
          </Link>
          <Link href="/dashboard/videos?category=popular" className="w-full">
            <Button variant="ghost" size="sm" className="w-full">
              Popular Videos
            </Button>
          </Link>
        </div>
      </div>
    </div>
  );
}
