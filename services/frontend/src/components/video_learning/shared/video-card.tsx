'use client';

import { useState } from 'react';
import { Eye, Heart, User, Play } from 'lucide-react';
import { Button } from '@/ui/button';
import { Link } from '@/ui/link';
import type { Video } from '@/types/video_learning/models';
import { toggleVideoLikeAction } from '@/actions/video_learning/actions';
import { toast } from 'sonner';
import Image from 'next/image';
import { formatViews } from '@/lib/video_learning/utils';
import { mapToFile } from '@/lib/shared/files/utils';

interface VideoCardProps {
  video: Video;
  showActions?: boolean;
  onDelete?: (videoId: number) => void;
  isOwner?: boolean;
}

export function VideoCard({
  video,
  showActions = true,
  onDelete,
  isOwner = false,
}: VideoCardProps) {
  const [isLiked, setIsLiked] = useState(false);
  const [likeCount, setLikeCount] = useState(video.likes);
  const [isLoading, setIsLoading] = useState(false);

  const handleLike = async () => {
    setIsLoading(true);
    const result = await toggleVideoLikeAction({ id: video.id });
    if (!result.success) {
      setIsLoading(false);
      console.error('Failed to like video:', result.error);
      toast.error('Failed to like video');
      return;
    }
    setIsLiked((prev) => !prev);
    setLikeCount((prev) => (isLiked ? prev - 1 : prev + 1));
    setIsLoading(false);
  };

  const handleDelete = async () => {
    if (!onDelete || !isOwner) return;

    const confirmed = window.confirm(
      'Are you sure you want to delete this video?',
    );
    if (confirmed) {
      setIsLoading(true);
      try {
        onDelete(video.id);
      } catch (error) {
        console.error('Failed to delete video:', error);
      } finally {
        setIsLoading(false);
      }
    }
  };

  return (
    <div className="bg-card border rounded-lg overflow-hidden hover:shadow-md transition-shadow">
      {/* Thumbnail */}
      <div className="relative aspect-video bg-muted">
        <img
          src={mapToFile(video.thumbnail_url)}
          alt={video.title}
          className="w-full h-full object-cover"
        />
        <div className="absolute inset-0 bg-black/20 flex items-center justify-center opacity-0 hover:opacity-100 transition-opacity">
          <Link href={`/dashboard/videos/${video.id}`}>
            <Button size="icon" variant="secondary" className="rounded-full">
              <Play className="h-5 w-5" />
            </Button>
          </Link>
        </div>
      </div>

      {/* Content */}
      <div className="p-4">
        <Link href={`/dashboard/videos/${video.id}`}>
          <h3 className="font-semibold text-sm line-clamp-2 hover:text-primary transition-colors mb-2">
            {video.title}
          </h3>
        </Link>

        <div className="flex items-center gap-2 text-sm text-muted-foreground mb-3">
          <User className="h-4 w-4" />
          <span>{video.author}</span>
        </div>

        {/* Stats */}
        <div className="flex items-center justify-between text-sm text-muted-foreground mb-3">
          <div className="flex items-center gap-4">
            <div className="flex items-center gap-1">
              <Eye className="h-4 w-4" />
              <span>{formatViews(video.views)}</span>
            </div>
            <div className="flex items-center gap-1">
              <Heart
                className={`h-4 w-4 ${isLiked ? 'fill-red-500 text-red-500' : ''}`}
              />
              <span>{likeCount}</span>
            </div>
          </div>
        </div>

        {/* Actions */}
        {showActions && (
          <div className="flex gap-2">
            {!isOwner && (
              <Button
                variant="outline"
                size="sm"
                onClick={handleLike}
                disabled={isLoading}
                className="flex-1"
              >
                <Heart
                  className={`h-4 w-4 mr-1 ${isLiked ? 'fill-red-500 text-red-500' : ''}`}
                />
                {isLiked ? 'Liked' : 'Like'}
              </Button>
            )}

            {isOwner && (
              <>
                <Button variant="outline" size="sm" className="w-full">
                  Edit
                </Button>
                {/* gustavo god como vas  */}
                {/* maso ahi, no carga igual el bucket, sigue forbidden */}
                {/* mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm */}
                <Button
                  variant="destructive"
                  size="sm"
                  onClick={handleDelete}
                  disabled={isLoading}
                >
                  Delete
                </Button>
              </>
            )}
          </div>
        )}
      </div>
    </div>
  );
}
