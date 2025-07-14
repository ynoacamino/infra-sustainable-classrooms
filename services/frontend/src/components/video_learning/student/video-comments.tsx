'use client';

import { useState, useEffect, useCallback } from 'react';
import {
  Send,
  Heart,
  User,
  Calendar,
  MessageCircle,
  Trash2,
  Reply,
} from 'lucide-react';
import { Button } from '@/ui/button';
import { Textarea } from '@/ui/textarea';
import { Input } from '@/ui/input';
import { Separator } from '@/ui/separator';
import { videoLearningService } from '@/services/video_learning/service';
import { cookies } from 'next/headers';
import type { Comment } from '@/types/video_learning/models';
import { Skeleton } from '@/ui/skeleton';
import { toast } from 'sonner';

interface VideoCommentsProps {
  videoId: number;
}

export function VideoComments({ videoId }: VideoCommentsProps) {
  const [comments, setComments] = useState<Comment[]>([]);
  const [loading, setLoading] = useState(true);
  const [newComment, setNewComment] = useState({ title: '', body: '' });
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [page, setPage] = useState(1);
  const [hasMore, setHasMore] = useState(true);
  const [loadingMore, setLoadingMore] = useState(false);

  const loadComments = useCallback(
    async (resetPage = true) => {
      try {
        if (resetPage) {
          setLoading(true);
          setPage(1);
        } else {
          setLoadingMore(true);
        }

        const service = await videoLearningService(cookies());
        const result = await service.getComments({
          video_id: videoId,
          page: resetPage ? 1 : page,
          page_size: 10,
        });

        if (result.success) {
          if (resetPage) {
            setComments(result.data as unknown as Comment[]);
          } else {
            setComments((prev) => [...prev, ...result.data] as Comment[]);
          }
          setHasMore(result.data.length === 10);
        } else {
          toast.error('Failed to load comments');
        }
      } catch (error) {
        console.error('Failed to load comments:', error);
        toast.error('An error occurred while loading comments');
      } finally {
        setLoading(false);
        setLoadingMore(false);
      }
    },
    [page, videoId],
  );

  useEffect(() => {
    loadComments();
  }, [loadComments, videoId]);

  const handleSubmitComment = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!newComment.title.trim() || !newComment.body.trim()) {
      toast.error('Please fill in both title and comment');
      return;
    }

    setIsSubmitting(true);
    try {
      const service = await videoLearningService(cookies());
      const result = await service.createComment({
        video_id: videoId,
        title: newComment.title,
        body: newComment.body,
      });

      if (result.success) {
        setNewComment({ title: '', body: '' });
        toast.success('Comment posted successfully');
        // Reload comments to show the new one
        loadComments();
      } else {
        toast.error('Failed to post comment');
      }
    } catch (error) {
      console.error('Failed to post comment:', error);
      toast.error('An error occurred while posting comment');
    } finally {
      setIsSubmitting(false);
    }
  };

  const handleDeleteComment = async (commentId: number) => {
    const confirmed = window.confirm(
      'Are you sure you want to delete this comment?',
    );
    if (!confirmed) return;

    try {
      const service = await videoLearningService(cookies());
      const result = await service.deleteComment({ id: commentId });

      if (result.success) {
        setComments((prev) =>
          prev.filter((comment) => comment.id !== commentId),
        );
        toast.success('Comment deleted successfully');
      } else {
        toast.error('Failed to delete comment');
      }
    } catch (error) {
      console.error('Failed to delete comment:', error);
      toast.error('An error occurred while deleting comment');
    }
  };

  const handleLoadMore = () => {
    setPage((prev) => prev + 1);
    loadComments(false);
  };

  const formatDate = (timestamp: number) => {
    return new Date(timestamp * 1000).toLocaleDateString('en-US', {
      year: 'numeric',
      month: 'short',
      day: 'numeric',
      hour: '2-digit',
      minute: '2-digit',
    });
  };

  if (loading) {
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

  return (
    <div className="space-y-6">
      {/* Comments Header */}
      <div className="flex items-center gap-2 mb-4">
        <MessageCircle className="h-5 w-5" />
        <h3 className="text-lg font-semibold">Comments ({comments.length})</h3>
      </div>

      {/* New Comment Form */}
      <form
        onSubmit={handleSubmitComment}
        className="space-y-4 p-4 bg-muted/30 rounded-lg"
      >
        <h4 className="font-medium">Add a comment</h4>

        <div className="space-y-3">
          <Input
            placeholder="Comment title"
            value={newComment.title}
            onChange={(e) =>
              setNewComment((prev) => ({ ...prev, title: e.target.value }))
            }
            disabled={isSubmitting}
          />

          <Textarea
            placeholder="Write your comment..."
            value={newComment.body}
            onChange={(e) =>
              setNewComment((prev) => ({ ...prev, body: e.target.value }))
            }
            disabled={isSubmitting}
            rows={3}
          />
        </div>

        <div className="flex justify-end">
          <Button
            type="submit"
            disabled={isSubmitting}
            className="flex items-center gap-2"
          >
            <Send className="h-4 w-4" />
            {isSubmitting ? 'Posting...' : 'Post Comment'}
          </Button>
        </div>
      </form>

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
                      <span>{formatDate(comment.date)}</span>
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
            {hasMore && (
              <div className="flex justify-center pt-4">
                <Button
                  variant="outline"
                  onClick={handleLoadMore}
                  disabled={loadingMore}
                >
                  {loadingMore ? 'Loading...' : 'Load More Comments'}
                </Button>
              </div>
            )}
          </>
        )}
      </div>
    </div>
  );
}
