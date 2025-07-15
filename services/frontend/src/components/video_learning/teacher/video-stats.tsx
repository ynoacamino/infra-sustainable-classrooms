'use client';

import { useState, useMemo } from 'react';
import { Eye, Heart, TrendingUp, BarChart3, PieChart } from 'lucide-react';
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '@/ui/card';
import { Progress } from '@/ui/progress';
import { Badge } from '@/ui/badge';
import type { OwnVideo } from '@/types/video_learning/models';
import { Skeleton } from '@/ui/skeleton';
import { useGetOwnVideos } from '@/hooks/video_learning/useSWR';
import Image from 'next/image';
import { formatViews } from '@/lib/video_learning/utils';

interface VideoStats {
  totalVideos: number;
  totalViews: number;
  totalLikes: number;
  averageViews: number;
  topPerformingVideo: OwnVideo | null;
}

export function VideoStats() {
  const [timeRange, setTimeRange] = useState<'7d' | '30d' | '90d' | 'all'>(
    '30d',
  );

  const {
    isLoading,
    data: videos,
    error,
  } = useGetOwnVideos({
    page: 1,
    page_size: 100, // Load more videos for better stats
  });

  const stats = useMemo(() => {
    if (!videos) return null;

    const totalVideos = videos.length;
    const totalViews = videos.reduce((sum, video) => sum + video.views, 0);
    const totalLikes = videos.reduce((sum, video) => sum + video.likes, 0);
    const averageViews =
      totalVideos > 0 ? Math.round(totalViews / totalVideos) : 0;

    // Find top performing video
    const topPerformingVideo =
      videos.length > 0
        ? videos.reduce((prev, current) =>
            prev.views > current.views ? prev : current,
          )
        : null;

    return {
      totalVideos,
      totalViews,
      totalLikes,
      averageViews,
      topPerformingVideo,
    };
  }, [videos]);

  if (isLoading) {
    return (
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
        {Array.from({ length: 4 }).map((_, i) => (
          <Skeleton key={i} className="h-32 w-full" />
        ))}
      </div>
    );
  }

  if (error || !videos || !stats) {
    return (
      <div className="text-center py-8">
        <BarChart3 className="h-12 w-12 mx-auto mb-4 text-muted-foreground" />
        <p className="text-muted-foreground">
          {error ? 'Failed to load statistics' : 'No statistics available'}
        </p>
      </div>
    );
  }

  return (
    <div className="space-y-6">
      {/* Time Range Filter */}
      <div className="flex gap-2">
        {(['7d', '30d', '90d', 'all'] as const).map((range) => (
          <Badge
            key={range}
            variant={timeRange === range ? 'default' : 'outline'}
            className="cursor-pointer"
            onClick={() => setTimeRange(range)}
          >
            {range === 'all' ? 'All Time' : range}
          </Badge>
        ))}
      </div>

      {/* Main Stats Cards */}
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">Total Videos</CardTitle>
            <Eye className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">{stats.totalVideos}</div>
            <p className="text-xs text-muted-foreground">
              {stats.totalVideos > 0 ? 'Published videos' : 'No videos yet'}
            </p>
          </CardContent>
        </Card>

        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">Total Views</CardTitle>
            <TrendingUp className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">
              {formatViews(stats.totalViews)}
            </div>
            <p className="text-xs text-muted-foreground">Across all videos</p>
          </CardContent>
        </Card>

        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">Total Likes</CardTitle>
            <Heart className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">
              {formatViews(stats.totalLikes)}
            </div>
            <p className="text-xs text-muted-foreground">Total engagement</p>
          </CardContent>
        </Card>

        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">Average Views</CardTitle>
            <BarChart3 className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">
              {formatViews(stats.averageViews)}
            </div>
            <p className="text-xs text-muted-foreground">Per video</p>
          </CardContent>
        </Card>
      </div>

      {/* Top Performing Video */}
      {stats.topPerformingVideo && (
        <Card>
          <CardHeader>
            <CardTitle className="flex items-center gap-2">
              <TrendingUp className="h-5 w-5" />
              Top Performing Video
            </CardTitle>
            <CardDescription>Your most viewed video</CardDescription>
          </CardHeader>
          <CardContent>
            <div className="flex gap-4">
              <Image
                src={stats.topPerformingVideo.thumbnail_url}
                alt={stats.topPerformingVideo.title}
                width={96}
                height={64}
                className="w-24 h-16 object-cover rounded"
                onError={(e) => {
                  const target = e.target as HTMLImageElement;
                  target.src = '/placeholder-video.jpg';
                }}
              />
              <div className="flex-1">
                <h4 className="font-semibold line-clamp-2 mb-2">
                  {stats.topPerformingVideo.title}
                </h4>
                <div className="flex items-center gap-4 text-sm text-muted-foreground">
                  <div className="flex items-center gap-1">
                    <Eye className="h-4 w-4" />
                    <span>
                      {formatViews(stats.topPerformingVideo.views)} views
                    </span>
                  </div>
                  <div className="flex items-center gap-1">
                    <Heart className="h-4 w-4" />
                    <span>
                      {formatViews(stats.topPerformingVideo.likes)} likes
                    </span>
                  </div>
                </div>
              </div>
            </div>
          </CardContent>
        </Card>
      )}

      {/* Video Performance Distribution */}
      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2">
            <PieChart className="h-5 w-5" />
            Performance Distribution
          </CardTitle>
          <CardDescription>How your videos are performing</CardDescription>
        </CardHeader>
        <CardContent>
          <div className="space-y-4">
            {videos.slice(0, 5).map((video) => {
              const percentage =
                stats.totalViews > 0
                  ? (video.views / stats.totalViews) * 100
                  : 0;
              return (
                <div key={video.id} className="space-y-2">
                  <div className="flex justify-between text-sm">
                    <span className="line-clamp-1">{video.title}</span>
                    <span>{formatViews(video.views)} views</span>
                  </div>
                  <Progress value={percentage} className="h-2" />
                </div>
              );
            })}
          </div>
        </CardContent>
      </Card>
    </div>
  );
}
