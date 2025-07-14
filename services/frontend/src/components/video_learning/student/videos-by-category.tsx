'use client';

import { useState, useEffect, useCallback } from 'react';
import { ArrowLeft, Filter, Grid, List, SortAsc, SortDesc } from 'lucide-react';
import { Button } from '@/ui/button';
import { Badge } from '@/ui/badge';
import { VideoCard } from '@/components/video_learning/shared/video-card';
import { videoLearningService } from '@/services/video_learning/service';
import { cookies } from 'next/headers';
import type { Video, VideoCategory } from '@/types/video_learning/models';
import { Skeleton } from '@/ui/skeleton';
import { Link } from '@/ui/link';
import { toast } from 'sonner';
import Image from 'next/image';

interface VideosByCategoryProps {
  categoryId: number;
}

type SortOption = 'title' | 'views' | 'likes' | 'upload_date';
type SortDirection = 'asc' | 'desc';

export function VideosByCategory({ categoryId }: VideosByCategoryProps) {
  const [videos, setVideos] = useState<Video[]>([]);
  const [category, setCategory] = useState<VideoCategory | null>(null);
  const [loading, setLoading] = useState(true);
  const [sortBy, setSortBy] = useState<SortOption>('views');
  const [sortDirection, setSortDirection] = useState<SortDirection>('desc');
  const [viewMode, setViewMode] = useState<'grid' | 'list'>('grid');
  const [hasMore, setHasMore] = useState(true);
  const [loadingMore, setLoadingMore] = useState(false);

  const loadCategoryAndVideos = useCallback(async () => {
    try {
      setLoading(true);
      const service = await videoLearningService(cookies());

      const [categoryResult, videosResult] = await Promise.all([
        service.getCategory({ id: categoryId }),
        service.getVideosByCategory({ id: categoryId, amount: 12 }),
      ]);

      if (categoryResult.success) {
        setCategory(categoryResult.data);
      }

      if (videosResult.success) {
        setVideos(videosResult.data);
        setHasMore(videosResult.data.length === 12);
      } else {
        toast.error('Failed to load videos');
      }
    } catch (error) {
      console.error('Failed to load category videos:', error);
      toast.error('An error occurred while loading videos');
    } finally {
      setLoading(false);
    }
  }, [categoryId]);

  useEffect(() => {
    loadCategoryAndVideos();
  }, [categoryId, loadCategoryAndVideos]);
  const sortVideos = useCallback(() => {
    setVideos((prev) => {
      const sorted = [...prev].sort((a, b) => {
        let aVal;
        let bVal;

        switch (sortBy) {
          case 'title':
            aVal = a.title.toLowerCase();
            bVal = b.title.toLowerCase();
            break;
          case 'views':
            aVal = a.views;
            bVal = b.views;
            break;
          case 'likes':
            aVal = a.likes;
            bVal = b.likes;
            break;
          case 'upload_date':
            // Note: Video model doesn't have upload_date, using views as fallback
            aVal = a.views;
            bVal = b.views;
            break;
          default:
            aVal = a.views;
            bVal = b.views;
        }

        if (sortDirection === 'asc') {
          return aVal > bVal ? 1 : -1;
        } else {
          return aVal < bVal ? 1 : -1;
        }
      });

      return sorted;
    });
  }, [sortBy, sortDirection]);

  useEffect(() => {
    sortVideos();
  }, [sortBy, sortDirection, sortVideos]);

  const loadMoreVideos = async () => {
    if (loadingMore || !hasMore) return;

    try {
      setLoadingMore(true);
      const service = await videoLearningService(cookies());
      const result = await service.getVideosByCategory({
        id: categoryId,
        amount: 12,
        // Note: This endpoint might need pagination support
      });

      if (result.success) {
        setVideos((prev) => [...prev, ...result.data]);
        setHasMore(result.data.length === 12);
      }
    } catch (error) {
      console.error('Failed to load more videos:', error);
      toast.error('Failed to load more videos');
    } finally {
      setLoadingMore(false);
    }
  };

  const handleLike = async (videoId: number) => {
    try {
      const service = await videoLearningService(cookies());
      await service.toogleVideoLike({ id: videoId });
    } catch (error) {
      console.error('Failed to like video:', error);
      throw error;
    }
  };

  const toggleSortDirection = () => {
    setSortDirection((prev) => (prev === 'asc' ? 'desc' : 'asc'));
  };

  if (loading) {
    return (
      <div className="space-y-6">
        <div className="flex items-center justify-between">
          <Skeleton className="h-8 w-48" />
          <Skeleton className="h-10 w-32" />
        </div>
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {Array.from({ length: 9 }).map((_, i) => (
            <Skeleton key={i} className="h-64 w-full" />
          ))}
        </div>
      </div>
    );
  }

  return (
    <div className="space-y-6">
      {/* Header */}
      <div className="flex items-center justify-between">
        <div className="flex items-center gap-4">
          <Link href="/dashboard/videos" variant="ghost">
            <ArrowLeft className="h-4 w-4 mr-2" />
            Back to Videos
          </Link>
          <div>
            <h2 className="text-xl font-semibold">
              {category?.name || 'Category Videos'}
            </h2>
            <p className="text-sm text-muted-foreground">
              {videos.length} videos in this category
            </p>
          </div>
        </div>

        <Badge variant="secondary" className="px-3 py-1">
          {category?.name}
        </Badge>
      </div>

      {/* Controls */}
      <div className="flex items-center justify-between gap-4">
        <div className="flex items-center gap-2">
          <Filter className="h-4 w-4 text-muted-foreground" />
          <span className="text-sm font-medium">Sort by:</span>
          <select
            value={sortBy}
            onChange={(e) => setSortBy(e.target.value as SortOption)}
            className="text-sm border rounded px-2 py-1"
          >
            <option value="views">Views</option>
            <option value="likes">Likes</option>
            <option value="title">Title</option>
            <option value="upload_date">Upload Date</option>
          </select>
          <Button
            variant="outline"
            size="sm"
            onClick={toggleSortDirection}
            title={`Sort ${sortDirection === 'asc' ? 'ascending' : 'descending'}`}
          >
            {sortDirection === 'asc' ? (
              <SortAsc className="h-4 w-4" />
            ) : (
              <SortDesc className="h-4 w-4" />
            )}
          </Button>
        </div>

        <div className="flex items-center gap-2">
          <Button
            variant={viewMode === 'grid' ? 'default' : 'outline'}
            size="sm"
            onClick={() => setViewMode('grid')}
          >
            <Grid className="h-4 w-4" />
          </Button>
          <Button
            variant={viewMode === 'list' ? 'default' : 'outline'}
            size="sm"
            onClick={() => setViewMode('list')}
          >
            <List className="h-4 w-4" />
          </Button>
        </div>
      </div>

      {/* Videos */}
      {videos.length === 0 ? (
        <div className="text-center py-12">
          <div className="text-muted-foreground mb-4">
            <Grid className="h-12 w-12 mx-auto mb-4" />
            <h3 className="text-lg font-semibold mb-2">No videos found</h3>
            <p className="text-sm">
              This category doesn&apos;t have any videos yet.
            </p>
          </div>
          <Link href="/dashboard/videos">
            <Button variant="outline">Browse All Videos</Button>
          </Link>
        </div>
      ) : (
        <div className="space-y-6">
          {viewMode === 'grid' ? (
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
              {videos.map((video) => (
                <VideoCard
                  key={video.id}
                  video={video}
                  onLike={handleLike}
                  showActions={true}
                />
              ))}
            </div>
          ) : (
            <div className="space-y-4">
              {videos.map((video) => (
                <div key={video.id} className="border rounded-lg p-4">
                  <div className="flex gap-4">
                    <Image
                      src={video.thumbnail_url}
                      alt={video.title}
                      width={128}
                      height={72}
                      className="w-32 h-20 object-cover rounded flex-shrink-0"
                      onError={(e) => {
                        const target = e.target as HTMLImageElement;
                        target.src = '/placeholder-video.jpg';
                      }}
                    />
                    <div className="flex-1">
                      <Link href={`/dashboard/videos/${video.id}`}>
                        <h3 className="font-semibold hover:text-primary transition-colors mb-2">
                          {video.title}
                        </h3>
                      </Link>
                      <p className="text-sm text-muted-foreground mb-2">
                        By {video.author}
                      </p>
                      <div className="flex items-center gap-4 text-sm text-muted-foreground">
                        <span>{video.views.toLocaleString()} views</span>
                        <span>{video.likes.toLocaleString()} likes</span>
                      </div>
                    </div>
                  </div>
                </div>
              ))}
            </div>
          )}

          {/* Load More Button */}
          {hasMore && viewMode === 'grid' && (
            <div className="flex justify-center">
              <Button
                variant="outline"
                onClick={loadMoreVideos}
                disabled={loadingMore}
              >
                {loadingMore ? 'Loading...' : 'Load More Videos'}
              </Button>
            </div>
          )}
        </div>
      )}
    </div>
  );
}
