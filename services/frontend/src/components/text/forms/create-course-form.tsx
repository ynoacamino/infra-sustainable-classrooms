'use client';

import { Form, FormField } from '@/ui/form';
import { zodResolver } from '@hookform/resolvers/zod';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import { InferItem } from '@/ui/infer-field';
import { Button } from '@/ui/button';
import { toast } from 'sonner';
import { createCourseFormFields, createCourseFormSchema } from '@/lib/text/forms/create-course-form';
import { createCourseAction } from '@/actions/text/actions';

function CreateCourseForm() {
  const form = useForm<z.infer<typeof createCourseFormSchema>>({
    resolver: zodResolver(createCourseFormSchema),
    defaultValues: {
      title: '',
      description: '',
      imageUrl: '',
    },
  });

  const onSubmit = async (values: z.infer<typeof createCourseFormSchema>) => {
    const res = await createCourseAction(values);
    if (!res.success) {
      form.setError('root', { message: res.error.message });
      console.error(res.error);
      toast.error(res.error.message);
      return;
    }
    form.reset();
    toast.success('Course created successfully');
  };

  return (
    <Form {...form}>
      <form
        onSubmit={form.handleSubmit(onSubmit)}
        className="flex flex-col gap-y-4 w-full px-4 sm:px-6 md:px-8 max-w-3xl pb-20"
      >
        {createCourseFormFields.map((field) => (
          <FormField
            key={`form-create-course-${field.name}`}
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
          Create Course
        </Button>
      </form>
    </Form>
  );
}

export { CreateCourseForm };
