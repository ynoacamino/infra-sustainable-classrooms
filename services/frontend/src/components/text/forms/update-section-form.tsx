'use client';

import { Form, FormField } from '@/ui/form';
import { zodResolver } from '@hookform/resolvers/zod';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import { InferItem } from '@/ui/infer-field';
import { Button } from '@/ui/button';
import { toast } from 'sonner';
import {
  updateSectionFormFields,
  updateSectionFormSchema,
} from '@/lib/text/forms/update-section-form';
import { updateSectionAction } from '@/actions/text/actions';
import type { Section } from '@/types/text/models';
import { redirect } from 'next/navigation';

interface UpdateSectionFormProps {
  section: Section;
}

function UpdateSectionForm({ section }: UpdateSectionFormProps) {
  const form = useForm<z.infer<typeof updateSectionFormSchema>>({
    resolver: zodResolver(updateSectionFormSchema),
    defaultValues: {
      title: section.title,
      description: section.description,
      order: section.order,
      section_id: section.id,
    },
  });

  const onSubmit = async (values: z.infer<typeof updateSectionFormSchema>) => {
    const res = await updateSectionAction(values);
    if (!res.success) {
      form.setError('root', { message: res.error.message });
      console.error(res.error);
      toast.error(res.error.message);
      return;
    }
    toast.success('Section updated successfully');
    redirect(`/teacher/courses/${section.course_id}`);
  };

  return (
    <Form {...form}>
      <form
        onSubmit={form.handleSubmit(onSubmit)}
        className="flex flex-col gap-y-4 w-full px-4 sm:px-6 md:px-8 max-w-3xl pb-20"
      >
        {updateSectionFormFields.map((field) => (
          <FormField
            key={`form-update-section-${field.name}`}
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
          Update Section
        </Button>
      </form>
    </Form>
  );
}

export { UpdateSectionForm };
