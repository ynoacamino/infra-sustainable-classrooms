'use client';

import { Form, FormField } from '@/ui/form';
import { zodResolver } from '@hookform/resolvers/zod';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import { InferItem } from '@/ui/infer-field';
import { Button } from '@/ui/button';
import { toast } from 'sonner';
import { updateCourseFormFields, updateCourseFormSchema } from '@/lib/text/forms/update-course-form';
import { updateCourseAction } from '@/actions/text/actions';
import type { Course } from '@/types/text/models';

function UpdateCourseForm({ course }: { course: Course }) {
  const form = useForm<z.infer<typeof updateCourseFormSchema>>({
    resolver: zodResolver(updateCourseFormSchema),
    defaultValues: {
      title: course.title,
      description: course.description,
      imageUrl: course.imageUrl,
      course_id: course.id,
    },
  });

  const onSubmit = async (values: z.infer<typeof updateCourseFormSchema>) => {
    const res = await updateCourseAction(values);
    if (!res.success) {
      form.setError('root', { message: res.error.message });
      console.error(res.error);
      toast.error(res.error.message);
      return;
    }
    toast.success('Course updated successfully');
  };

  return (
    <Form {...form}>
      <form
        onSubmit={form.handleSubmit(onSubmit)}
        className="flex flex-col gap-y-4 w-full px-4 sm:px-6 md:px-8 max-w-3xl pb-20"
      >
        {updateCourseFormFields.map((field) => (
          <FormField
            key={`form-update-course-${field.name}`}
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
          Update Course
        </Button>
      </form>
    </Form>
  );
}

export { UpdateCourseForm };
