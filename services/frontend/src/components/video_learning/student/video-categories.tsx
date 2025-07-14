'use client';

import { useState, useEffect, useCallback } from 'react';
import { VideoCard } from '@/components/video_learning/shared/video-card';
import { videoLearningService } from '@/services/video_learning/service';
import { cookies } from 'next/headers';
import type { Video, VideoCategory } from '@/types/video_learning/models';
import { Skeleton } from '@/ui/skeleton';
import { Button } from '@/ui/button';
import { Grid, List, ChevronRight } from 'lucide-react';
import { Link } from '@/ui/link';

export function VideoCategories() {
  const [categories, setCategories] = useState<VideoCategory[]>([]);
  const [categoryVideos, setCategoryVideos] = useState<Record<number, Video[]>>(
    {},
  );
  const [loading, setLoading] = useState(true);
  const [loadingVideos, setLoadingVideos] = useState<Record<number, boolean>>(
    {},
  );
  const [error, setError] = useState<string | null>(null);
  const [viewMode, setViewMode] = useState<'grid' | 'list'>('grid');

  const loadCategories = useCallback(async () => {
    try {
      setError(null);
      const service = await videoLearningService(cookies());
      const result = await service.getAllCategories();

      if (result.success) {
        setCategories(result.data);
        // Load a few videos for each category
        result.data.forEach((category) => {
          loadCategoryVideos(category.id);
        });
      } else {
        setError('Failed to load categories');
      }
    } catch (error) {
      console.error('Failed to load categories:', error);
      setError('An error occurred while loading categories');
    } finally {
      setLoading(false);
    }
  }, [setLoading]);

  useEffect(() => {
    loadCategories();
  }, [loadCategories]);

  const loadCategoryVideos = async (categoryId: number) => {
    try {
      setLoadingVideos((prev) => ({ ...prev, [categoryId]: true }));
      const service = await videoLearningService(cookies());
      const result = await service.getVideosByCategory({
        id: categoryId,
        amount: 4,
      });

      if (result.success) {
        setCategoryVideos((prev) => ({ ...prev, [categoryId]: result.data }));
      }
    } catch (error) {
      console.error(`Failed to load videos for category ${categoryId}:`, error);
    } finally {
      setLoadingVideos((prev) => ({ ...prev, [categoryId]: false }));
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

  if (loading) {
    return (
      <div className="grid grid-cols-2 md:grid-cols-4 gap-4">
        {Array.from({ length: 8 }).map((_, i) => (
          <Skeleton key={i} className="h-24 w-full rounded-lg" />
        ))}
      </div>
    );
  }

  if (error) {
    return (
      <div className="text-center py-12">
        <div className="text-destructive mb-4">
          <svg
            className="h-12 w-12 mx-auto mb-4"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth={2}
              d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"
            />
          </svg>
          <p className="text-sm">{error}</p>
        </div>
        <Button onClick={loadCategories} variant="outline">
          Try Again
        </Button>
      </div>
    );
  }

  if (categories.length === 0) {
    return (
      <div className="text-center py-12">
        <div className="text-muted-foreground mb-4">
          <Grid className="h-12 w-12 mx-auto mb-4" />
          <h3 className="text-lg font-semibold mb-2">
            No categories available
          </h3>
          <p className="text-sm">
            Categories will appear here when they&apos;re created
          </p>
        </div>
      </div>
    );
  }

  return (
    <div className="space-y-6">
      {/* View Mode Toggle */}
      <div className="flex items-center justify-between">
        <h3 className="text-lg font-semibold">Browse by Category</h3>
        <div className="flex gap-2">
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

      {/* Categories */}
      <div className="space-y-8">
        {categories.map((category) => (
          <div key={category.id} className="space-y-4">
            <div className="flex items-center justify-between">
              <h4 className="text-md font-medium">{category.name}</h4>
              <Link
                href={`/dashboard/videos/category/${category.id}`}
                className="text-sm text-primary hover:underline flex items-center gap-1"
              >
                View All
                <ChevronRight className="h-4 w-4" />
              </Link>
            </div>

            {/* Category Videos */}
            {loadingVideos[category.id] ? (
              <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
                {Array.from({ length: 4 }).map((_, i) => (
                  <Skeleton key={i} className="h-48 w-full rounded-lg" />
                ))}
              </div>
            ) : categoryVideos[category.id]?.length > 0 ? (
              <div
                className={`grid gap-4 ${
                  viewMode === 'grid'
                    ? 'grid-cols-1 md:grid-cols-2 lg:grid-cols-4'
                    : 'grid-cols-1 md:grid-cols-2'
                }`}
              >
                {categoryVideos[category.id].map((video) => (
                  <VideoCard
                    key={video.id}
                    video={video}
                    onLike={handleLike}
                    showActions={true}
                  />
                ))}
              </div>
            ) : (
              <div className="text-center py-8 bg-muted rounded-lg">
                <p className="text-muted-foreground text-sm">
                  No videos in this category yet
                </p>
              </div>
            )}
          </div>
        ))}
      </div>
    </div>
  );
}
