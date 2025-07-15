'use client';

import { Form, FormField } from '@/ui/form';
import { zodResolver } from '@hookform/resolvers/zod';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import { InferItem } from '@/ui/infer-field';
import { Button } from '@/ui/button';
import { toast } from 'sonner';
import {
  createTagFormFields,
  createTagFormSchema,
} from '@/lib/video_learning/forms/create-tag-form';
import type { VideoTag } from '@/types/video_learning/models';
import { createTagAction } from '@/actions/video_learning/actions';

interface CreateTagFormProps {
  onSuccess?: (tag: VideoTag) => void;
  onCancel?: () => void;
}

function CreateTagForm({ onSuccess, onCancel }: CreateTagFormProps) {
  const form = useForm<z.infer<typeof createTagFormSchema>>({
    resolver: zodResolver(createTagFormSchema),
    defaultValues: {
      name: '',
    },
  });

  const onSubmit = async (values: z.infer<typeof createTagFormSchema>) => {
    const res = await createTagAction(values);
    if (!res.success) {
      form.setError('root', { message: res.error.message });
      console.error(res.error);
      toast.error(res.error.message);
      return;
    }
    form.reset();
    toast.success('Tag created successfully');
    onSuccess?.(res.data);
  };

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4">
        {createTagFormFields.map((field) => (
          <FormField
            key={`form-create-tag-${field.name}`}
            control={form.control}
            name={field.name}
            render={({ field: formField }) => (
              <InferItem {...field} {...formField} />
            )}
          />
        ))}
        <div className="flex justify-end gap-2">
          <Button
            type="button"
            variant="outline"
            onClick={onCancel}
            disabled={form.formState.isSubmitting}
          >
            Cancel
          </Button>
          <Button type="submit" disabled={form.formState.isSubmitting}>
            {form.formState.isSubmitting ? 'Creating...' : 'Create Tag'}
          </Button>
        </div>
      </form>
    </Form>
  );
}

export { CreateTagForm };
