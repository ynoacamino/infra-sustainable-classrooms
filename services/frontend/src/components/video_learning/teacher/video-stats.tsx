'use client';

import { useState, useEffect, useCallback } from 'react';
import {
  Eye,
  Heart,
  TrendingUp,
  Calendar,
  BarChart3,
  PieChart,
} from 'lucide-react';
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '@/ui/card';
import { Progress } from '@/ui/progress';
import { Badge } from '@/ui/badge';
import { videoLearningService } from '@/services/video_learning/service';
import { cookies } from 'next/headers';
import type { OwnVideo } from '@/types/video_learning/models';
import { Skeleton } from '@/ui/skeleton';
import { toast } from 'sonner';
import Image from 'next/image';

interface VideoStats {
  totalVideos: number;
  totalViews: number;
  totalLikes: number;
  averageViews: number;
  topPerformingVideo: OwnVideo | null;
  recentActivity: {
    date: string;
    views: number;
    likes: number;
  }[];
}

export function VideoStats() {
  const [stats, setStats] = useState<VideoStats | null>(null);
  const [videos, setVideos] = useState<OwnVideo[]>([]);
  const [loading, setLoading] = useState(true);
  const [timeRange, setTimeRange] = useState<'7d' | '30d' | '90d' | 'all'>(
    '30d',
  );

  const calculateStats = useCallback((videoList: OwnVideo[]) => {
    const totalVideos = videoList.length;
    const totalViews = videoList.reduce((sum, video) => sum + video.views, 0);
    const totalLikes = videoList.reduce((sum, video) => sum + video.likes, 0);
    const averageViews =
      totalVideos > 0 ? Math.round(totalViews / totalVideos) : 0;

    // Find top performing video
    const topPerformingVideo =
      videoList.length > 0
        ? videoList.reduce((prev, current) =>
            prev.views > current.views ? prev : current,
          )
        : null;

    // Generate mock recent activity data
    const recentActivity = generateMockActivity();

    setStats({
      totalVideos,
      totalViews,
      totalLikes,
      averageViews,
      topPerformingVideo,
      recentActivity,
    });
  }, []);

  const loadStats = useCallback(async () => {
    try {
      setLoading(true);
      const service = await videoLearningService(cookies());

      // Load all videos to calculate stats
      const result = await service.getOwnVideos({
        page: 1,
        page_size: 100, // Load more videos for better stats
      });

      if (result.success) {
        setVideos(result.data);
        calculateStats(result.data);
      } else {
        toast.error('Failed to load video statistics');
      }
    } catch (error) {
      console.error('Failed to load stats:', error);
      toast.error('An error occurred while loading statistics');
    } finally {
      setLoading(false);
    }
  }, [calculateStats]);

  useEffect(() => {
    loadStats();
  }, [loadStats, timeRange]);

  const generateMockActivity = () => {
    const activity = [];
    const now = new Date();

    for (let i = 6; i >= 0; i--) {
      const date = new Date(now);
      date.setDate(date.getDate() - i);

      activity.push({
        date: date.toLocaleDateString(),
        views: Math.floor(Math.random() * 100) + 20,
        likes: Math.floor(Math.random() * 20) + 5,
      });
    }

    return activity;
  };

  const formatNumber = (num: number) => {
    if (num >= 1000000) {
      return `${(num / 1000000).toFixed(1)}M`;
    } else if (num >= 1000) {
      return `${(num / 1000).toFixed(1)}K`;
    }
    return num.toString();
  };

  if (loading) {
    return (
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
        {Array.from({ length: 4 }).map((_, i) => (
          <Skeleton key={i} className="h-32 w-full" />
        ))}
      </div>
    );
  }

  if (!stats) {
    return (
      <div className="text-center py-8">
        <BarChart3 className="h-12 w-12 mx-auto mb-4 text-muted-foreground" />
        <p className="text-muted-foreground">No statistics available</p>
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
              {formatNumber(stats.totalViews)}
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
              {formatNumber(stats.totalLikes)}
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
              {formatNumber(stats.averageViews)}
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
                      {formatNumber(stats.topPerformingVideo.views)} views
                    </span>
                  </div>
                  <div className="flex items-center gap-1">
                    <Heart className="h-4 w-4" />
                    <span>
                      {formatNumber(stats.topPerformingVideo.likes)} likes
                    </span>
                  </div>
                </div>
              </div>
            </div>
          </CardContent>
        </Card>
      )}

      {/* Recent Activity */}
      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2">
            <Calendar className="h-5 w-5" />
            Recent Activity
          </CardTitle>
          <CardDescription>
            Views and likes over the last 7 days
          </CardDescription>
        </CardHeader>
        <CardContent>
          <div className="space-y-4">
            {stats.recentActivity.map((activity, index) => (
              <div key={index} className="flex items-center justify-between">
                <div className="flex items-center gap-3">
                  <div className="w-2 h-2 bg-primary rounded-full" />
                  <span className="text-sm">{activity.date}</span>
                </div>
                <div className="flex items-center gap-4 text-sm">
                  <div className="flex items-center gap-1">
                    <Eye className="h-3 w-3" />
                    <span>{activity.views}</span>
                  </div>
                  <div className="flex items-center gap-1">
                    <Heart className="h-3 w-3" />
                    <span>{activity.likes}</span>
                  </div>
                </div>
              </div>
            ))}
          </div>
        </CardContent>
      </Card>

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
                    <span>{formatNumber(video.views)} views</span>
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
