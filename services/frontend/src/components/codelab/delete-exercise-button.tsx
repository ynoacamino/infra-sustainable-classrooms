'use client';

import { Button } from '@/ui/button';
import { Trash2 } from 'lucide-react';
import { useState } from 'react';
import { toast } from 'sonner';
import { useRouter } from 'next/navigation';
import { deleteExerciseAction } from '@/actions/codelab/actions';

interface DeleteExerciseButtonProps {
  exerciseId: number;
  exerciseTitle: string;
  redirectTo?: string;
}

export function DeleteExerciseButton({ 
  exerciseId, 
  exerciseTitle, 
  redirectTo = '/teacher/codelab/exercises' 
}: DeleteExerciseButtonProps) {
  const [isDeleting, setIsDeleting] = useState(false);
  const router = useRouter();

  const handleDelete = async () => {
    if (!confirm(`Are you sure you want to delete "${exerciseTitle}"? This will also delete all associated test cases and student attempts. This action cannot be undone.`)) {
      return;
    }

    setIsDeleting(true);
    
    try {
      const result = await deleteExerciseAction(exerciseId);
      
      if (result?.success) {
        toast.success('Exercise deleted successfully');
        router.push(redirectTo);
      } else {
        toast.error('Failed to delete exercise');
      }
    } catch (error) {
      console.error('Error deleting exercise:', error);
      toast.error('Failed to delete exercise');
    } finally {
      setIsDeleting(false);
    }
  };

  return (
    <Button
      variant="outline"
      className="text-destructive hover:text-destructive"
      onClick={handleDelete}
      disabled={isDeleting}
    >
      <Trash2 className="w-4 h-4 mr-2" />
      {isDeleting ? 'Deleting...' : 'Delete Exercise'}
    </Button>
  );
}
