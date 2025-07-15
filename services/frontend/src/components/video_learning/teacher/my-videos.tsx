'use client';

import { useState, useMemo } from 'react';

import { Skeleton } from '@/ui/skeleton';
import { Button } from '@/ui/button';
import { Input } from '@/ui/input';
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/ui/select';
import {
  Search,
  Filter,
  Plus,
  Grid,
  List,
  SortAsc,
  SortDesc,
} from 'lucide-react';
import { Link } from '@/ui/link';
import { toast } from 'sonner';
import { VideoCard } from '@/components/video_learning/shared/video-card';
import { useGetOwnVideos } from '@/hooks/video_learning/useSWR';
import { deleteVideoAction } from '@/actions/video_learning/actions';
import Image from 'next/image';

type SortOption = 'title' | 'views' | 'likes' | 'upload_date';
type SortDirection = 'asc' | 'desc';

export function MyVideos() {
  const [searchQuery, setSearchQuery] = useState('');
  const [sortBy, setSortBy] = useState<SortOption>('upload_date');
  const [sortDirection, setSortDirection] = useState<SortDirection>('desc');
  const [viewMode, setViewMode] = useState<'grid' | 'list'>('grid');
  const [page] = useState(1);

  const {
    isLoading,
    data: videos,
    error,
    mutate,
  } = useGetOwnVideos({
    page,
    page_size: 50, // Load more videos at once for better filtering
  });

  const filteredVideos = useMemo(() => {
    if (!videos) return [];

    let filtered = [...videos];

    // Apply search filter
    if (searchQuery) {
      filtered = filtered.filter((video) =>
        video.title.toLowerCase().includes(searchQuery.toLowerCase()),
      );
    }

    // Apply sorting
    filtered.sort((a, b) => {
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
          aVal = a.upload_date;
          bVal = b.upload_date;
          break;
        default:
          aVal = a.upload_date;
          bVal = b.upload_date;
      }

      if (sortDirection === 'asc') {
        return aVal > bVal ? 1 : -1;
      } else {
        return aVal < bVal ? 1 : -1;
      }
    });

    return filtered;
  }, [videos, searchQuery, sortBy, sortDirection]);

  const handleDelete = async (videoId: number) => {
    try {
      const result = await deleteVideoAction({ id: videoId });

      if (result.success) {
        mutate(); // Refresh the videos list
        toast.success('Video deleted successfully');
      } else {
        toast.error(result.error?.message || 'Failed to delete video');
      }
    } catch (error) {
      console.error('Failed to delete video:', error);
      toast.error('An error occurred while deleting the video');
    }
  };

  const formatDate = (timestamp: number) => {
    return new Date(timestamp * 1000).toLocaleDateString('en-US', {
      year: 'numeric',
      month: 'short',
      day: 'numeric',
    });
  };

  const toggleSortDirection = () => {
    setSortDirection((prev) => (prev === 'asc' ? 'desc' : 'asc'));
  };

  if (isLoading) {
    return (
      <div className="space-y-4">
        {/* Loading skeleton for controls */}
        <div className="flex flex-col sm:flex-row gap-4">
          <Skeleton className="h-10 flex-1" />
          <Skeleton className="h-10 w-32" />
          <Skeleton className="h-10 w-24" />
        </div>

        {/* Loading skeleton for videos */}
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {Array.from({ length: 6 }).map((_, i) => (
            <Skeleton key={i} className="h-64 w-full rounded-lg" />
          ))}
        </div>
      </div>
    );
  }

  if (error || !videos) {
    return (
      <div className="text-center py-12">
        <p className="text-muted-foreground">
          Failed to load videos. Please try again.
        </p>
      </div>
    );
  }

  return (
    <div className="space-y-6">
      {/* Controls */}
      <div className="flex flex-col sm:flex-row gap-4 items-center justify-between">
        <div className="flex flex-1 gap-2">
          <div className="relative flex-1">
            <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 text-muted-foreground h-4 w-4" />
            <Input
              placeholder="Search your videos..."
              value={searchQuery}
              onChange={(e) => setSearchQuery(e.target.value)}
              className="pl-10"
            />
          </div>

          <Select
            value={sortBy}
            onValueChange={(value: SortOption) => setSortBy(value)}
          >
            <SelectTrigger className="w-[140px]">
              <Filter className="h-4 w-4 mr-2" />
              <SelectValue />
            </SelectTrigger>
            <SelectContent>
              <SelectItem value="upload_date">Upload Date</SelectItem>
              <SelectItem value="title">Title</SelectItem>
              <SelectItem value="views">Views</SelectItem>
              <SelectItem value="likes">Likes</SelectItem>
            </SelectContent>
          </Select>

          <Button
            variant="outline"
            size="icon"
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

        <div className="flex gap-2">
          <Button
            variant={viewMode === 'grid' ? 'default' : 'outline'}
            size="icon"
            onClick={() => setViewMode('grid')}
          >
            <Grid className="h-4 w-4" />
          </Button>
          <Button
            variant={viewMode === 'list' ? 'default' : 'outline'}
            size="icon"
            onClick={() => setViewMode('list')}
          >
            <List className="h-4 w-4" />
          </Button>

          <Link href="/teacher/videos/upload">
            <Button>
              <Plus className="h-4 w-4 mr-2" />
              Upload Video
            </Button>
          </Link>
        </div>
      </div>

      {/* Stats */}
      <div className="grid grid-cols-1 md:grid-cols-4 gap-4">
        <div className="bg-card border rounded-lg p-4">
          <div className="text-2xl font-bold">{videos?.length || 0}</div>
          <div className="text-sm text-muted-foreground">Total Videos</div>
        </div>
        <div className="bg-card border rounded-lg p-4">
          <div className="text-2xl font-bold">
            {videos
              ? videos
                  .reduce((sum, video) => sum + video.views, 0)
                  .toLocaleString()
              : 0}
          </div>
          <div className="text-sm text-muted-foreground">Total Views</div>
        </div>
        <div className="bg-card border rounded-lg p-4">
          <div className="text-2xl font-bold">
            {videos
              ? videos
                  .reduce((sum, video) => sum + video.likes, 0)
                  .toLocaleString()
              : 0}
          </div>
          <div className="text-sm text-muted-foreground">Total Likes</div>
        </div>
        <div className="bg-card border rounded-lg p-4">
          <div className="text-2xl font-bold">
            {videos && videos.length > 0
              ? Math.round(
                  videos.reduce((sum, video) => sum + video.views, 0) /
                    videos.length,
                )
              : 0}
          </div>
          <div className="text-sm text-muted-foreground">Avg Views</div>
        </div>
      </div>

      {/* Videos */}
      {filteredVideos.length === 0 ? (
        <div className="text-center py-12">
          <div className="text-muted-foreground mb-4">
            {searchQuery ? (
              <>
                <Search className="h-12 w-12 mx-auto mb-4" />
                <h3 className="text-lg font-semibold mb-2">No videos found</h3>
                <p className="text-sm">Try adjusting your search terms</p>
              </>
            ) : (
              <>
                <Plus className="h-12 w-12 mx-auto mb-4" />
                <h3 className="text-lg font-semibold mb-2">
                  No videos uploaded yet
                </h3>
                <p className="text-sm mb-4">
                  Start by uploading your first video
                </p>
                <Link href="/teacher/videos/upload">
                  <Button>
                    <Plus className="h-4 w-4 mr-2" />
                    Upload Video
                  </Button>
                </Link>
              </>
            )}
          </div>
        </div>
      ) : (
        <div className="space-y-4">
          <div className="flex items-center justify-between">
            <h3 className="text-lg font-semibold">
              {searchQuery
                ? `Search Results (${filteredVideos.length})`
                : `My Videos (${filteredVideos.length})`}
            </h3>
          </div>

          {viewMode === 'grid' ? (
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
              {filteredVideos.map((video) => (
                <VideoCard
                  key={video.id}
                  video={{
                    ...video,
                    author: 'You', // Since it's the teacher's own video
                  }}
                  onDelete={handleDelete}
                  isOwner={true}
                  showActions={true}
                />
              ))}
            </div>
          ) : (
            <div className="space-y-4">
              {filteredVideos.map((video) => (
                <div key={video.id} className="bg-card border rounded-lg p-4">
                  <div className="flex gap-4">
                    <Image
                      src={video.thumbnail_url}
                      alt={video.title}
                      width={128}
                      height={80}
                      className="w-32 h-20 object-cover rounded flex-shrink-0"
                      onError={(e) => {
                        const target = e.target as HTMLImageElement;
                        target.src = '/placeholder-video.jpg';
                      }}
                    />
                    <div className="flex-1 space-y-2">
                      <div className="flex items-start justify-between">
                        <Link href={`/teacher/videos/${video.id}`}>
                          <h4 className="font-semibold hover:text-primary transition-colors">
                            {video.title}
                          </h4>
                        </Link>
                        <div className="flex gap-2">
                          <Link href={`/teacher/videos/${video.id}/edit`}>
                            <Button variant="outline" size="sm">
                              Edit
                            </Button>
                          </Link>
                          <Button
                            variant="destructive"
                            size="sm"
                            onClick={() => handleDelete(video.id)}
                          >
                            Delete
                          </Button>
                        </div>
                      </div>
                      <div className="flex items-center gap-4 text-sm text-muted-foreground">
                        <span>{video.views.toLocaleString()} views</span>
                        <span>{video.likes.toLocaleString()} likes</span>
                        <span>Uploaded {formatDate(video.upload_date)}</span>
                      </div>
                    </div>
                  </div>
                </div>
              ))}
            </div>
          )}

          {/* Load More Button - Removed since we're loading all videos at once now */}
        </div>
      )}
    </div>
  );
}
