'use client';

import { Button } from '@/ui/button';
import { Trash2 } from 'lucide-react';
import { useState } from 'react';
import { toast } from 'sonner';
import { deleteQuestionAction } from '@/actions/knowledge/actions';
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/ui/dialog';

interface DeleteQuestionButtonProps {
  testId: number;
  questionId: number;
  questionText: string;
  onSuccess?: () => void;
}

export function DeleteQuestionButton({
  testId,
  questionId,
  questionText,
  onSuccess,
}: DeleteQuestionButtonProps) {
  const [isLoading, setIsLoading] = useState(false);
  const [isOpen, setIsOpen] = useState(false);

  const handleDelete = async () => {
    try {
      setIsLoading(true);
      const result = await deleteQuestionAction(testId, questionId);

      if (result.success) {
        toast.success('Question deleted successfully');
        onSuccess?.();
        setIsOpen(false);
        // Refresh the page to update the questions list
        window.location.reload();
      } else {
        toast.error(result.error.message || 'Failed to delete question');
      }
    } catch (error) {
      console.error('Error deleting question:', error);
      toast.error('An error occurred while deleting the question');
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <Dialog open={isOpen} onOpenChange={setIsOpen}>
      <DialogTrigger asChild>
        <Button
          variant="outline"
          size="sm"
          className="text-red-600 hover:text-red-700 hover:bg-red-50"
          disabled={isLoading}
        >
          <Trash2 className="h-4 w-4 mr-1" />
          Delete
        </Button>
      </DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Delete Question</DialogTitle>
          <DialogDescription>
            Are you sure you want to delete this question? This action cannot be
            undone.
            <br />
            <br />
            <strong>Question:</strong>{' '}
            {questionText.length > 100
              ? `${questionText.substring(0, 100)}...`
              : questionText}
          </DialogDescription>
        </DialogHeader>
        <DialogFooter>
          <Button
            variant="outline"
            onClick={() => setIsOpen(false)}
            disabled={isLoading}
          >
            Cancel
          </Button>
          <Button
            onClick={handleDelete}
            disabled={isLoading}
            className="bg-red-600 hover:bg-red-700"
          >
            {isLoading ? 'Deleting...' : 'Delete Question'}
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}
