'use client';

import { useState } from 'react';
import {
  Heart,
  Share2,
  Download,
  Eye,
  User,
  Calendar,
  Tag,
  Folder,
} from 'lucide-react';
import { Button } from '@/ui/button';
import { Badge } from '@/ui/badge';
import { Separator } from '@/ui/separator';
import type { VideoDetails } from '@/types/video_learning/models';
import { Skeleton } from '@/ui/skeleton';
import { toast } from 'sonner';
import { useSWRAll } from '@/lib/shared/swr/utils';
import {
  useGetCategory,
  useGetTagsByVideo,
} from '@/hooks/video_learning/useSWR';
import { toggleVideoLikeAction } from '@/actions/video_learning/actions';
import { formatViews } from '@/lib/video_learning/utils';
import { formatDate } from '@/lib/shared/utils';

interface VideoInfoProps {
  video: VideoDetails;
}

export function VideoInfo({ video }: VideoInfoProps) {
  const {
    isLoading,
    data: [category, tags],
    errors,
    // mutateAll,
  } = useSWRAll([
    useGetCategory({ id: video.category_id }),
    useGetTagsByVideo({ id: video.id }),
  ]);
  const [isLiked, setIsLiked] = useState(false);
  const [likeCount, setLikeCount] = useState(0);
  const [isLiking, setIsLiking] = useState(false);

  const handleLike = async () => {
    if (!video || isLiking) return;

    setIsLiking(true);
    try {
      const result = await toggleVideoLikeAction({ id: video.id });

      if (result.success) {
        setIsLiked(!isLiked);
        setLikeCount((prev) => (isLiked ? prev - 1 : prev + 1));
        toast.success(
          isLiked ? 'Removed from favorites' : 'Added to favorites',
        );
      } else {
        toast.error('Failed to update like status');
      }
    } catch (error) {
      console.error('Failed to like video:', error);
      toast.error('An error occurred while updating like status');
    } finally {
      setIsLiking(false);
    }
  };

  const handleShare = async () => {
    if (!video) return;

    try {
      const url = window.location.href;
      if (navigator.share) {
        await navigator.share({
          title: video.title,
          text: video.description,
          url: url,
        });
      } else {
        await navigator.clipboard.writeText(url);
        toast.success('Link copied to clipboard');
      }
    } catch (error) {
      console.error('Failed to share:', error);
      toast.error('Failed to share video');
    }
  };

  if (isLoading) {
    return (
      <div className="space-y-4">
        <Skeleton className="h-8 w-3/4" />
        <div className="flex items-center gap-4">
          <Skeleton className="h-6 w-20" />
          <Skeleton className="h-6 w-20" />
          <Skeleton className="h-6 w-32" />
        </div>
        <Skeleton className="h-20 w-full" />
        <div className="flex gap-2">
          <Skeleton className="h-8 w-20" />
          <Skeleton className="h-8 w-20" />
        </div>
      </div>
    );
  }

  if (!video || errors.length > 0) {
    return (
      <div className="text-center py-8">
        <p className="text-muted-foreground">
          Failed to load video information
        </p>
      </div>
    );
  }

  return (
    <div className="space-y-6">
      {/* Video Title */}
      <div>
        <h1 className="text-2xl font-bold mb-2">{video.title}</h1>
        <div className="flex items-center gap-4 text-sm text-muted-foreground">
          <div className="flex items-center gap-1">
            <Eye className="h-4 w-4" />
            <span>{formatViews(video.views)} views</span>
          </div>
          <div className="flex items-center gap-1">
            <Calendar className="h-4 w-4" />
            <span>{formatDate(video.upload_date)}</span>
          </div>
          <div className="flex items-center gap-1">
            <User className="h-4 w-4" />
            <span>{video.author}</span>
          </div>
        </div>
      </div>

      {/* Action Buttons */}
      <div className="flex items-center gap-3">
        <Button
          variant={isLiked ? 'default' : 'outline'}
          onClick={handleLike}
          disabled={isLiking}
          className="flex items-center gap-2"
        >
          <Heart className={`h-4 w-4 ${isLiked ? 'fill-current' : ''}`} />
          <span>{likeCount}</span>
        </Button>

        <Button
          variant="outline"
          onClick={handleShare}
          className="flex items-center gap-2"
        >
          <Share2 className="h-4 w-4" />
          <span>Share</span>
        </Button>

        <Button variant="outline" className="flex items-center gap-2">
          <Download className="h-4 w-4" />
          <span>Download</span>
        </Button>
      </div>

      <Separator />

      {/* Video Description */}
      <div className="space-y-4">
        <h3 className="text-lg font-semibold">Description</h3>
        <div className="bg-muted/50 rounded-lg p-4">
          <p className="text-sm leading-relaxed whitespace-pre-wrap">
            {video.description}
          </p>
        </div>
      </div>

      {/* Category and Tags */}
      <div className="space-y-4">
        {category && (
          <div className="flex items-center gap-2">
            <Folder className="h-4 w-4 text-muted-foreground" />
            <span className="text-sm font-medium">Category:</span>
            <Badge variant="secondary">{category.name}</Badge>
          </div>
        )}

        {tags.length > 0 && (
          <div className="flex items-start gap-2">
            <Tag className="h-4 w-4 text-muted-foreground mt-0.5" />
            <div className="flex-1">
              <span className="text-sm font-medium">Tags:</span>
              <div className="flex flex-wrap gap-2 mt-1">
                {tags.map((tag) => (
                  <Badge key={tag.id} variant="outline" className="text-xs">
                    {tag.name}
                  </Badge>
                ))}
              </div>
            </div>
          </div>
        )}
      </div>

      {/* Video Stats */}
      <div className="grid grid-cols-3 gap-4 p-4 bg-muted/30 rounded-lg">
        <div className="text-center">
          <div className="text-2xl font-bold">{formatViews(video.views)}</div>
          <div className="text-sm text-muted-foreground">Views</div>
        </div>
        <div className="text-center">
          <div className="text-2xl font-bold">{likeCount}</div>
          <div className="text-sm text-muted-foreground">Likes</div>
        </div>
        <div className="text-center">
          <div className="text-2xl font-bold">
            {formatDate(video.upload_date)}
          </div>
          <div className="text-sm text-muted-foreground">Uploaded</div>
        </div>
      </div>
    </div>
  );
}
