'use client';

import { useState } from 'react';
import {
  Heart,
  User,
  Calendar,
  MessageCircle,
  Trash2,
  Reply,
} from 'lucide-react';
import { Button } from '@/ui/button';
import { Separator } from '@/ui/separator';
import { Skeleton } from '@/ui/skeleton';
import { toast } from 'sonner';
import { useGetComments } from '@/hooks/video_learning/useSWR';
import { CreateCommentForm } from '@/components/video_learning/forms/create-comment-form';
import { deleteCommentAction } from '@/actions/video_learning/actions';
import { useSWRAll } from '@/lib/shared/swr/utils';
import type { VideoDetails } from '@/types/video_learning/models';

interface VideoCommentsProps {
  video: VideoDetails;
}

export function VideoComments({ video }: VideoCommentsProps) {
  const [page, setPage] = useState(1);
  const {
    isLoading,
    errors,
    data: [commentsResult],
    mutateAll: mutateComments,
  } = useSWRAll([
    useGetComments({
      video_id: video.id,
      page,
    }),
  ]);

  const handleDeleteComment = async (commentId: number) => {
    const confirmed = window.confirm(
      'Are you sure you want to delete this comment?',
    );
    if (!confirmed) return;

    const result = await deleteCommentAction({ id: commentId });

    if (!result.success) {
      toast.error(result.error.message);
      return;
    }

    mutateComments();
    toast.success('Comment deleted successfully');
  };

  const handleLoadMore = () => {
    setPage((prev) => prev + 1);
  };

  if (isLoading) {
    return (
      <div className="space-y-6">
        <div className="flex items-center gap-2 mb-4">
          <MessageCircle className="h-5 w-5" />
          <h3 className="text-lg font-semibold">Comments</h3>
        </div>

        <div className="space-y-4">
          {Array.from({ length: 3 }).map((_, i) => (
            <div key={i} className="space-y-3">
              <div className="flex items-center gap-2">
                <Skeleton className="h-4 w-20" />
                <Skeleton className="h-4 w-32" />
              </div>
              <Skeleton className="h-6 w-3/4" />
              <Skeleton className="h-16 w-full" />
            </div>
          ))}
        </div>
      </div>
    );
  }

  if (errors.length > 0 || !commentsResult) {
    return (
      <div className="text-center py-12">
        <MessageCircle className="h-12 w-12 mx-auto mb-4 text-muted-foreground" />
        <h3 className="text-lg font-semibold mb-2">Error loading comments</h3>
        <p className="text-sm text-muted-foreground">
          Please try again later or contact support.
        </p>
      </div>
    );
  }

  const comments = commentsResult.comments || [];
  return (
    <div className="space-y-6">
      {/* Comments Header */}
      <div className="flex items-center gap-2 mb-4">
        <MessageCircle className="h-5 w-5" />
        <h3 className="text-lg font-semibold">Comments ({comments.length})</h3>
      </div>

      {/* New Comment Form */}
      <CreateCommentForm video={video} onSuccess={() => mutateComments()} />

      <Separator />

      {/* Comments List */}
      <div className="space-y-4">
        {comments.length === 0 ? (
          <div className="text-center py-8">
            <MessageCircle className="h-12 w-12 mx-auto mb-4 text-muted-foreground" />
            <h4 className="text-lg font-semibold mb-2">No comments yet</h4>
            <p className="text-muted-foreground">
              Be the first to comment on this video!
            </p>
          </div>
        ) : (
          <>
            {comments.map((comment) => (
              <div key={comment.id} className="space-y-3 p-4 border rounded-lg">
                {/* Comment Header */}
                <div className="flex items-center justify-between">
                  <div className="flex items-center gap-3">
                    <div className="flex items-center gap-1 text-sm text-muted-foreground">
                      <User className="h-4 w-4" />
                      <span className="font-medium">{comment.author}</span>
                    </div>
                    <div className="flex items-center gap-1 text-sm text-muted-foreground">
                      <Calendar className="h-4 w-4" />
                      <span>
                        {new Date(comment.date * 1000).toLocaleDateString()}
                      </span>
                    </div>
                  </div>

                  {/* Comment Actions */}
                  <div className="flex items-center gap-1">
                    <Button
                      variant="ghost"
                      size="sm"
                      onClick={() => handleDeleteComment(comment.id)}
                      className="text-destructive hover:text-destructive"
                    >
                      <Trash2 className="h-4 w-4" />
                    </Button>
                  </div>
                </div>

                {/* Comment Title */}
                <h5 className="font-semibold">{comment.title}</h5>

                {/* Comment Body */}
                <p className="text-sm leading-relaxed whitespace-pre-wrap">
                  {comment.body}
                </p>

                {/* Comment Actions */}
                <div className="flex items-center gap-2 pt-2">
                  <Button
                    variant="ghost"
                    size="sm"
                    className="text-muted-foreground hover:text-primary"
                  >
                    <Heart className="h-4 w-4 mr-1" />
                    Like
                  </Button>
                  <Button
                    variant="ghost"
                    size="sm"
                    className="text-muted-foreground hover:text-primary"
                  >
                    <Reply className="h-4 w-4 mr-1" />
                    Reply
                  </Button>
                </div>
              </div>
            ))}

            {/* Load More Button */}
            {comments.length > 0 && comments.length % 10 === 0 && (
              <div className="flex justify-center pt-4">
                <Button
                  variant="outline"
                  onClick={handleLoadMore}
                  disabled={isLoading}
                >
                  {isLoading ? 'Loading...' : 'Load More Comments'}
                </Button>
              </div>
            )}
          </>
        )}
      </div>
    </div>
  );
}
