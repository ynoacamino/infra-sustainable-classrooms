'use client';

import { useState } from 'react';
import { Search, Filter, X } from 'lucide-react';
import { Button } from '@/ui/button';
import { Input } from '@/ui/input';
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/ui/select';
import { VideoCard } from '@/components/video_learning/shared/video-card';
import { useSWRAll } from '@/lib/shared/swr/utils';
import {
  useGetAllCategories,
  useSearchVideos,
} from '@/hooks/video_learning/useSWR';
import { PAGINATION_DEFAULT_SIZE } from '@/config/shared/const';

export function VideoSearch() {
  const [searchQuery, setSearchQuery] = useState('');
  const [selectedCategory, setSelectedCategory] = useState<number | undefined>(
    undefined,
  );
  const [page, setPage] = useState(1);

  const {
    isLoading,
    data: [categories, videos],
    errors,
    // mutateAll,
  } = useSWRAll([
    useGetAllCategories(),
    useSearchVideos({
      query: searchQuery,
      category_id: selectedCategory,
      page,
      page_size: PAGINATION_DEFAULT_SIZE,
    }),
  ]);

  const loadMore = () => {
    setPage((prev) => prev + 1);
  };

  const clearFilters = () => {
    setSearchQuery('');
    setSelectedCategory(undefined);
    setPage(1);
  };

  if (errors.length > 0 || !categories || !videos) {
    return (
      <div className="text-center py-12">
        <h3 className="text-lg font-semibold mb-2">Error loading data</h3>
        <p className="text-muted-foreground">
          Please try again later or contact support
        </p>
      </div>
    );
  }

  return (
    <div className="space-y-6">
      {/* Search Controls */}
      <div className="flex flex-col sm:flex-row gap-4">
        <div className="flex-1 relative">
          <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 text-muted-foreground h-4 w-4" />
          <Input
            placeholder="Search videos..."
            value={searchQuery}
            onChange={(e) => setSearchQuery(e.target.value)}
            className="pl-10"
          />
        </div>

        <div className="flex gap-2">
          <Select
            value={selectedCategory?.toString() ?? ''}
            onValueChange={(v) =>
              setSelectedCategory(v ? parseInt(v) : undefined)
            }
          >
            <SelectTrigger className="w-[180px]">
              <Filter className="h-4 w-4 mr-2" />
              <SelectValue placeholder="All Categories" />
            </SelectTrigger>
            <SelectContent>
              <SelectItem value="">All Categories</SelectItem>
              {categories.map((category) => (
                <SelectItem key={category.id} value={category.id.toString()}>
                  {category.name}
                </SelectItem>
              ))}
            </SelectContent>
          </Select>

          {(searchQuery || selectedCategory) && (
            <Button variant="outline" onClick={clearFilters}>
              <X className="h-4 w-4 mr-2" />
              Clear
            </Button>
          )}
        </div>
      </div>

      {/* Search Results */}
      {videos.length > 0 && (
        <div className="space-y-4">
          <div className="flex items-center justify-between">
            <h3 className="text-lg font-semibold">
              Search Results ({videos.length} videos)
            </h3>
          </div>

          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {videos.map((video) => (
              <VideoCard key={video.id} video={video} />
            ))}
          </div>

          <div className="flex justify-center pt-4">
            <Button variant="outline" onClick={loadMore} disabled={isLoading}>
              {isLoading ? 'Loading...' : 'Load More'}
            </Button>
          </div>
        </div>
      )}

      {/* Empty State */}
      {videos.length === 0 &&
        !isLoading &&
        (searchQuery || selectedCategory) && (
          <div className="text-center py-12">
            <Search className="h-12 w-12 mx-auto text-muted-foreground mb-4" />
            <h3 className="text-lg font-semibold mb-2">No videos found</h3>
            <p className="text-muted-foreground">
              Try adjusting your search terms or filters
            </p>
          </div>
        )}
    </div>
  );
}
