'use client';

import { Button } from '@/ui/button';
import { Trash2 } from 'lucide-react';
import { useState } from 'react';
import { toast } from 'sonner';
import { useRouter } from 'next/navigation';
import { deleteTestAction } from '@/actions/codelab/actions';

interface DeleteTestButtonProps {
  testId: number;
  index: number;
}

export function DeleteTestButton({ testId, index }: DeleteTestButtonProps) {
  const [isDeleting, setIsDeleting] = useState(false);
  const router = useRouter();

  const handleDelete = async () => {
    if (
      !confirm(
        `Are you sure you want to delete Test Case #${index + 1}? This action cannot be undone.`,
      )
    ) {
      return;
    }

    setIsDeleting(true);

    try {
      const result = await deleteTestAction({ id: testId });

      if (result?.success) {
        toast.success('Test case deleted successfully');
        router.refresh();
      } else {
        toast.error('Failed to delete test case');
      }
    } catch (error) {
      console.error('Error deleting test:', error);
      toast.error('Failed to delete test case');
    } finally {
      setIsDeleting(false);
    }
  };

  return (
    <Button
      variant="outline"
      size="sm"
      className="text-destructive hover:text-destructive"
      onClick={handleDelete}
      disabled={isDeleting}
    >
      <Trash2 className="w-4 h-4 mr-2" />
      {isDeleting ? 'Deleting...' : 'Delete'}
    </Button>
  );
}
