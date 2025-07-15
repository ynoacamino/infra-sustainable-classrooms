'use client';

import { Form, FormField } from '@/ui/form';
import { zodResolver } from '@hookform/resolvers/zod';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import { InferItem } from '@/ui/infer-field';
import { Button } from '@/ui/button';
import { toast } from 'sonner';
import {
  createCategoryFormFields,
  createCategoryFormSchema,
} from '@/lib/video_learning/forms/create-category-form';
import { createCategoryAction } from '@/actions/video_learning/actions';
import type { VideoCategory } from '@/types/video_learning/models';

interface CreateCategoryFormProps {
  onSuccess?: (category: VideoCategory) => void;
  onCancel?: () => void;
}

function CreateCategoryForm({ onSuccess, onCancel }: CreateCategoryFormProps) {
  const form = useForm<z.infer<typeof createCategoryFormSchema>>({
    resolver: zodResolver(createCategoryFormSchema),
    defaultValues: {
      name: '',
    },
  });

  const onSubmit = async (values: z.infer<typeof createCategoryFormSchema>) => {
    const res = await createCategoryAction(values);
    if (!res.success) {
      form.setError('root', { message: res.error.message });
      console.error(res.error);
      toast.error(res.error.message);
      return;
    }
    form.reset();
    toast.success('Category created successfully');
    onSuccess?.(res.data);
  };

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4">
        {createCategoryFormFields.map((field) => (
          <FormField
            key={`form-create-category-${field.name}`}
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
            {form.formState.isSubmitting ? 'Creating...' : 'Create Category'}
          </Button>
        </div>
      </form>
    </Form>
  );
}

export { CreateCategoryForm };
