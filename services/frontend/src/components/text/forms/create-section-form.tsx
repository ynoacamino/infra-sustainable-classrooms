'use client';

import { Form, FormField } from '@/ui/form';
import { zodResolver } from '@hookform/resolvers/zod';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import { InferItem } from '@/ui/infer-field';
import { Button } from '@/ui/button';
import { toast } from 'sonner';
import {
  createSectionFormFields,
  createSectionFormSchema,
} from '@/lib/text/forms/create-section-form';
import { createSectionAction } from '@/actions/text/actions';
import { redirect } from 'next/navigation';

interface CreateSectionFormProps {
  courseId: number;
}

function CreateSectionForm({ courseId }: CreateSectionFormProps) {
  const form = useForm<z.infer<typeof createSectionFormSchema>>({
    resolver: zodResolver(createSectionFormSchema),
    defaultValues: {
      title: '',
      description: '',
      order: undefined,
      course_id: courseId,
    },
  });

  const onSubmit = async (values: z.infer<typeof createSectionFormSchema>) => {
    const res = await createSectionAction(values);
    if (!res.success) {
      form.setError('root', { message: res.error.message });
      console.error(res.error);
      toast.error(res.error.message);
      return;
    }
    form.reset();
    toast.success('Section created successfully');
    redirect(`/teacher/courses/${courseId}`);
  };

  return (
    <Form {...form}>
      <form
        onSubmit={form.handleSubmit(onSubmit)}
        className="flex flex-col gap-y-4 w-full px-4 sm:px-6 md:px-8 max-w-3xl pb-20"
      >
        {createSectionFormFields.map((field) => (
          <FormField
            key={`form-create-section-${field.name}`}
            control={form.control}
            name={field.name}
            render={({ field: formField }) => (
              <InferItem {...field} {...formField} />
            )}
          />
        ))}
        <Button
          type="submit"
          className="self-center"
          disabled={form.formState.isSubmitting}
        >
          Create Section
        </Button>
      </form>
    </Form>
  );
}

export { CreateSectionForm };
