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
} from '@/lib/knowledge/forms/create-test-form';
import { createTestAction } from '@/actions/knowledge/actions';
import { redirect } from 'next/navigation';

function CreateTestForm() {
  const form = useForm<z.infer<typeof createTestFormSchema>>({
    resolver: zodResolver(createTestFormSchema),
    defaultValues: {
      title: '',
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
    form.reset();
    toast.success('Test created successfully');
    redirect('/teacher/tests');
  };

  return (
    <Form {...form}>
      <form
        onSubmit={form.handleSubmit(onSubmit)}
        className="flex flex-col gap-y-4 w-full px-4 sm:px-6 md:px-8 max-w-3xl pb-20"
      >
        {createTestFormFields.map((field) => (
          <FormField
            key={`form-create-test-${field.name}`}
            control={form.control}
            name={field.name}
            render={({ field: formField }) => (
              <InferItem {...field} {...formField} />
            )}
          />
        ))}

        <div className="flex gap-2 mt-6">
          <Button type="submit" disabled={form.formState.isSubmitting}>
            {form.formState.isSubmitting ? 'Creating...' : 'Create Test'}
          </Button>
        </div>
      </form>
    </Form>
  );
}

export { CreateTestForm };
