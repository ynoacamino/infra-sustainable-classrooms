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
} from '@/lib/codelab/forms/update-test-form';
import { updateTestAction } from '@/actions/codelab/actions';
import type { Test } from '@/types/codelab/models';

interface UpdateTestFormProps {
  test: Test;
  onSuccess?: () => void;
}

function UpdateTestForm({ test, onSuccess }: UpdateTestFormProps) {
  const form = useForm<z.infer<typeof updateTestFormSchema>>({
    resolver: zodResolver(updateTestFormSchema),
    defaultValues: {
      input: test.input,
      output: test.output,
      public: test.public,
    },
  });

  const onSubmit = async (values: z.infer<typeof updateTestFormSchema>) => {
    const res = await updateTestAction(test.id, values);
    if (!res.success) {
      form.setError('root', { message: res.error.message });
      console.error(res.error);
      toast.error(res.error.message);
      return;
    }
    toast.success('Test case updated successfully');
    onSuccess?.();
  };

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4">
        {updateTestFormFields.map((field) => (
          <FormField
            key={field.name}
            control={form.control}
            name={field.name}
            render={({ field: formField }) => (
              <InferItem {...field} {...formField} />
            )}
          />
        ))}

        <Button type="submit" className="w-full">
          Update Test Case
        </Button>
      </form>
    </Form>
  );
}

export { UpdateTestForm };
