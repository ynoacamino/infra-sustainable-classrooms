'use client';

import { Form, FormField } from '@/ui/form';
import { zodResolver } from '@hookform/resolvers/zod';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import { InferItem } from '@/ui/infer-field';
import { Button } from '@/ui/button';
import { toast } from 'sonner';
import {
  updateTestFormFields,
  updateTestFormSchema,
} from '@/lib/knowledge/forms/update-test-form';
import type { Test } from '@/types/knowledge/models';
import { updateTestAction } from '@/actions/knowledge/actions';
import { redirect } from 'next/navigation';

interface UpdateTestFormProps {
  test: Test;
}

function UpdateTestForm({ test }: UpdateTestFormProps) {
  const form = useForm<z.infer<typeof updateTestFormSchema>>({
    resolver: zodResolver(updateTestFormSchema),
    defaultValues: {
      id: test.id,
      title: test.title,
    },
  });

  const onSubmit = async (values: z.infer<typeof updateTestFormSchema>) => {
    const res = await updateTestAction(values);
    if (!res.success) {
      form.setError('root', { message: res.error.message });
      console.error(res.error);
      toast.error(res.error.message);
      return;
    }
    toast.success('Test updated successfully');
    redirect('/teacher/tests');
  };

  return (
    <Form {...form}>
      <form
        onSubmit={form.handleSubmit(onSubmit)}
        className="flex flex-col gap-y-4 w-full px-4 sm:px-6 md:px-8 max-w-3xl pb-20"
      >
        {updateTestFormFields.map((field) => (
          <FormField
            key={`form-update-test-${field.name}`}
            control={form.control}
            name={field.name}
            render={({ field: formField }) => (
              <InferItem {...field} {...formField} />
            )}
          />
        ))}

        <div className="flex gap-2 mt-6">
          <Button type="submit" disabled={form.formState.isSubmitting}>
            {form.formState.isSubmitting ? 'Updating...' : 'Update Test'}
          </Button>
        </div>
      </form>
    </Form>
  );
}

export { UpdateTestForm };
