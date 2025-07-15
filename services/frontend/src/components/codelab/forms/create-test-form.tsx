'use client';

import { Form, FormField } from '@/ui/form';
import { zodResolver } from '@hookform/resolvers/zod';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import { InferItem } from '@/ui/infer-field';
import { Button } from '@/ui/button';
import { toast } from 'sonner';
import {
  createTestFormFields,
  createTestFormSchema,
} from '@/lib/codelab/forms/create-test-form';
import { createTestAction } from '@/actions/codelab/actions';

interface CreateTestFormProps {
  exerciseId: number;
  onSuccess?: () => void;
}

function CreateTestForm({ exerciseId, onSuccess }: CreateTestFormProps) {
  const form = useForm<z.infer<typeof createTestFormSchema>>({
    resolver: zodResolver(createTestFormSchema),
    defaultValues: {
      input: '',
      output: '',
      public: true,
      exercise_id: exerciseId,
    },
  });

  const onSubmit = async (values: z.infer<typeof createTestFormSchema>) => {
    const res = await createTestAction(values);
    if (!res.success) {
      form.setError('root', { message: res.error.message });
      console.error(res.error);
      toast.error(res.error.message);
      return;
    }
    form.reset({
      input: '',
      output: '',
      public: true,
      exercise_id: exerciseId,
    });
    toast.success('Test case created successfully');
    onSuccess?.();
  };

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4">
        {createTestFormFields.map((field) => {
          // Skip hidden fields
          if (field.name === 'exercise_id') {
            return null;
          }

          return (
            <FormField
              key={field.name}
              control={form.control}
              name={field.name}
              render={({ field: formField }) => (
                <InferItem
                  {...field}
                  {...formField}
                  onChange={
                    field.name === 'public'
                      ? (value) => formField.onChange(value === 'true')
                      : formField.onChange
                  }
                />
              )}
            />
          );
        })}

        <Button type="submit" className="w-full">
          Create Test Case
        </Button>
      </form>
    </Form>
  );
}

export { CreateTestForm };
