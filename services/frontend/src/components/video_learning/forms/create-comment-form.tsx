'use client';

import { Form, FormField } from '@/ui/form';
import { zodResolver } from '@hookform/resolvers/zod';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import { InferItem } from '@/ui/infer-field';
import { Button } from '@/ui/button';
import { toast } from 'sonner';
import {
  createCommentFormFields,
  createCommentFormSchema,
} from '@/lib/video_learning/forms/create-comment-form';
import { createCommentAction } from '@/actions/video_learning/actions';
import { Send } from 'lucide-react';
import type { VideoDetails } from '@/types/video_learning/models';

interface CreateCommentFormProps {
  video: VideoDetails;
  onSuccess?: () => void;
}

function CreateCommentForm({ video, onSuccess }: CreateCommentFormProps) {
  const form = useForm<z.infer<typeof createCommentFormSchema>>({
    resolver: zodResolver(createCommentFormSchema),
    defaultValues: {
      title: '',
      body: '',
      video_id: video.id,
    },
  });

  const onSubmit = async (values: z.infer<typeof createCommentFormSchema>) => {
    const result = await createCommentAction(values);

    if (!result.success) {
      form.setError('root', { message: result.error.message });
      toast.error(result.error.message);
      return;
    }

    form.reset({
      title: '',
      body: '',
      video_id: video.id,
    });
    toast.success('Comment posted successfully');
    onSuccess?.();
  };

  return (
    <div className="p-4 bg-muted/30 rounded-lg">
      <h4 className="font-medium mb-4">Add a comment</h4>

      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4">
          {createCommentFormFields.map((field) => (
            <FormField
              key={`form-create-comment-${field.name}`}
              control={form.control}
              name={field.name}
              render={({ field: formField }) => (
                <InferItem {...field} {...formField} />
              )}
            />
          ))}

          <div className="flex justify-end">
            <Button
              type="submit"
              disabled={form.formState.isSubmitting}
              className="flex items-center gap-2"
            >
              <Send className="h-4 w-4" />
              {form.formState.isSubmitting ? 'Posting...' : 'Post Comment'}
            </Button>
          </div>
        </form>
      </Form>
    </div>
  );
}

export { CreateCommentForm };
