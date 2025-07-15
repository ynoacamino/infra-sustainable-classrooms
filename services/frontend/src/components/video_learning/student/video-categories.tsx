'use client';

import { useState } from 'react';
import { VideoCard } from '@/components/video_learning/shared/video-card';
import { Button } from '@/ui/button';
import { Grid, List, ChevronRight } from 'lucide-react';
import { Link } from '@/ui/link';
import { useGetVideosGroupedByCategory } from '@/hooks/video_learning/useSWR';
import { Skeleton } from '@/ui/skeleton';

export function VideoCategories() {
  const [viewMode, setViewMode] = useState<'grid' | 'list'>('grid');
  const {
    isLoading,
    data: groupedVideos,
    error,
    // mutate,
  } = useGetVideosGroupedByCategory({ amount: 10 });

  if (isLoading) {
    return (
      <div className="space-y-6">
        <div className="flex items-center justify-between">
          <h3 className="text-lg font-semibold">Browse by Category</h3>
          <div className="flex gap-2">
            <Button variant="outline" size="sm" disabled>
              <Grid className="h-4 w-4" />
            </Button>
            <Button variant="outline" size="sm" disabled>
              <List className="h-4 w-4" />
            </Button>
          </div>
        </div>
        <div className="space-y-8">
          {[...Array(3)].map((_, index) => (
            <div key={index} className="space-y-4">
              <Skeleton className="h-6 w-1/3 mb-2" />
              <div className="grid gap-4 grid-cols-1 md:grid-cols-2 lg:grid-cols-4">
                {[...Array(4)].map((_, idx) => (
                  <Skeleton key={idx} className="h-40 w-full rounded-lg" />
                ))}
              </div>
            </div>
          ))}
        </div>
      </div>
    );
  }

  if (error || !groupedVideos) {
    return (
      <div className="text-center py-12">
        <div className="text-muted-foreground mb-4">
          <h3 className="text-lg font-semibold mb-2">
            Error loading categories
          </h3>
          <p className="text-sm">Please try again later</p>
        </div>
      </div>
    );
  }

  if (groupedVideos.length === 0) {
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
        {groupedVideos.map(({ category, videos }) => (
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
            <div
              className={`grid gap-4 ${
                viewMode === 'grid'
                  ? 'grid-cols-1 md:grid-cols-2 lg:grid-cols-4'
                  : 'grid-cols-1 md:grid-cols-2'
              }`}
            >
              {videos.map((video) => (
                <VideoCard key={video.id} video={video} showActions={true} />
              ))}
            </div>
            {videos.length === 0 && viewMode === 'grid' && (
              <div className="text-center py-8 bg-muted rounded-lg">
                <p className="text-muted-foreground text-sm">
                  No videos in this category yet
                </p>
              </div>
              // holaaa
            )}
          </div>
        ))}
      </div>
    </div>
  );
}
