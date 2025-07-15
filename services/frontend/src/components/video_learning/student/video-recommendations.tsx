'use client';

import { useState } from 'react';
import { VideoCard } from '@/components/video_learning/shared/video-card';
import { Skeleton } from '@/ui/skeleton';
import { Button } from '@/ui/button';
import { RefreshCw } from 'lucide-react';
import { useSWRAll } from '@/lib/shared/swr/utils';
import { useGetRecommendations } from '@/hooks/video_learning/useSWR';

export function VideoRecommendations() {
  const {
    isLoading,
    data: [videos],
    errors,
    mutateAll,
  } = useSWRAll([useGetRecommendations({ amount: 6 })]);
  const [refreshing, setRefreshing] = useState(false);

  const handleRefresh = async () => {
    setRefreshing(true);
    mutateAll();
  };

  if (isLoading) {
    return (
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        {Array.from({ length: 6 }).map((_, i) => (
          <Skeleton key={i} className="h-64 w-full rounded-lg" />
        ))}
      </div>
    );
  }

  if (errors.length > 0 || !videos) {
    return (
      <div className="text-center py-12">
        <div className="text-muted-foreground mb-4">
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
              d="M12 9v3m0 0v3m0-3h3m-3 0H9m6.364 7.364a9 9 0 11-12.728-12.728l1.414 1.414a7 7 0 009.9 9.9l1.414 1.414z"
            />
          </svg>
          <h3 className="text-lg font-semibold mb-2">
            Error loading recommendations
          </h3>
          <p className="text-sm">Please try refreshing the page</p>
        </div>
        <Button onClick={handleRefresh} variant="outline">
          <RefreshCw className="h-4 w-4 mr-2" />
          Refresh
        </Button>
      </div>
    );
  }

  if (videos.length === 0) {
    return (
      <div className="text-center py-12">
        <div className="text-muted-foreground mb-4">
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
              d="M7 4V2a1 1 0 011-1h8a1 1 0 011 1v2h4a1 1 0 010 2h-1v12a2 2 0 01-2 2H6a2 2 0 01-2-2V6H3a1 1 0 010-2h4zM6 6v12h12V6H6z"
            />
          </svg>
          <h3 className="text-lg font-semibold mb-2">No recommendations yet</h3>
          <p className="text-sm">
            Watch some videos to get personalized recommendations
          </p>
        </div>
        <Button onClick={handleRefresh} variant="outline">
          <RefreshCw className="h-4 w-4 mr-2" />
          Check Again
        </Button>
      </div>
    );
  }

  return (
    <div className="space-y-4">
      <div className="flex items-center justify-between">
        <h3 className="text-lg font-semibold">
          Recommended Videos ({videos.length})
        </h3>
        <Button
          variant="outline"
          size="sm"
          onClick={handleRefresh}
          disabled={refreshing}
        >
          <RefreshCw
            className={`h-4 w-4 mr-2 ${refreshing ? 'animate-spin' : ''}`}
          />
          Refresh
        </Button>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        {videos.map((video) => (
          <VideoCard key={video.id} video={video} showActions={true} />
        ))}
      </div>
    </div>
  );
}
