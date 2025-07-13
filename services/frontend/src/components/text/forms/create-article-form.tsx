'use client';

import { Form, FormField } from '@/ui/form';
import { zodResolver } from '@hookform/resolvers/zod';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import { InferItem } from '@/ui/infer-field';
import { Button } from '@/ui/button';
import { toast } from 'sonner';
import { createArticleFormFields, createArticleFormSchema } from '@/lib/text/forms/create-article-form';
import { createArticleAction } from '@/actions/text/actions';
import { redirect } from 'next/navigation';

interface CreateArticleFormProps {
  sectionId: number;
}

function CreateArticleForm({ sectionId }: CreateArticleFormProps) {
  const form = useForm<z.infer<typeof createArticleFormSchema>>({
    resolver: zodResolver(createArticleFormSchema),
    defaultValues: {
      title: '',
      content: '',
      section_id: sectionId,
    },
  });

  const onSubmit = async (values: z.infer<typeof createArticleFormSchema>) => {
    const res = await createArticleAction(values);
    if (!res.success) {
      form.setError('root', { message: res.error.message });
      console.error(res.error);
      toast.error(res.error.message);
      return;
    }
    form.reset();
    toast.success('Article created successfully');
  };

  return (
    <Form {...form}>
      <form
        onSubmit={form.handleSubmit(onSubmit)}
        className="flex flex-col gap-y-4 w-full px-4 sm:px-6 md:px-8 max-w-3xl pb-20"
      >
        {createArticleFormFields.map((field) => (
          <FormField
            key={`form-create-article-${field.name}`}
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
          Create Article
        </Button>
      </form>
    </Form>
  );
}

export { CreateArticleForm };
