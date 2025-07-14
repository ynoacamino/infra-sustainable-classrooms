'use client';

import { useState, useEffect, useCallback } from 'react';
import { PlayCircle, Clock, Eye, TrendingUp } from 'lucide-react';
import { Button } from '@/ui/button';
import { Link } from '@/ui/link';
import { videoLearningService } from '@/services/video_learning/service';
import { cookies } from 'next/headers';
import type { Video } from '@/types/video_learning/models';
import { Skeleton } from '@/ui/skeleton';
import { toast } from 'sonner';
import Image from 'next/image';

interface VideoSidebarProps {
  videoId: number;
}

export function VideoSidebar({ videoId }: VideoSidebarProps) {
  const [similarVideos, setSimilarVideos] = useState<Video[]>([]);
  const [recommendedVideos, setRecommendedVideos] = useState<Video[]>([]);
  const [loading, setLoading] = useState(true);
  const [activeTab, setActiveTab] = useState<'similar' | 'recommended'>(
    'similar',
  );

  const loadSidebarVideos = useCallback(async () => {
    try {
      setLoading(true);
      const service = await videoLearningService(cookies());

      const [similarResult, recommendedResult] = await Promise.all([
        service.getSimilarVideos({ id: videoId, amount: 10 }),
        service.getRecommendations({ amount: 10 }),
      ]);

      if (similarResult.success) {
        setSimilarVideos(similarResult.data);
      }

      if (recommendedResult.success) {
        setRecommendedVideos(recommendedResult.data);
      }
    } catch (error) {
      console.error('Failed to load sidebar videos:', error);
      toast.error('Failed to load related videos');
    } finally {
      setLoading(false);
    }
  }, [videoId]);

  useEffect(() => {
    loadSidebarVideos();
  }, [loadSidebarVideos, videoId]);

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
              src={video.thumbnail_url}
              alt={video.title}
              fill
              className="w-full h-full object-cover group-hover:scale-105 transition-transform"
              onError={(e) => {
                const target = e.target as HTMLImageElement;
                target.src = '/placeholder-video.jpg';
              }}
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
        {loading ? (
          renderLoadingSkeleton()
        ) : (
          <div className="space-y-2">
            {activeTab === 'similar' && (
              <>
                {similarVideos.length > 0 ? (
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
                {recommendedVideos.length > 0 ? (
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
