'use client';

import { Form, FormField } from '@/ui/form';
import { zodResolver } from '@hookform/resolvers/zod';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import { InferItem } from '@/ui/infer-field';
import { Button } from '@/ui/button';
import { toast } from 'sonner';
import {
  updateArticleFormFields,
  updateArticleFormSchema,
} from '@/lib/text/forms/update-article-form';
import { updateArticleAction } from '@/actions/text/actions';
import type { Article } from '@/types/text/models';
import { redirect } from 'next/navigation';

interface UpdateArticleFormProps {
  article: Article;
}

function UpdateArticleForm({ article }: UpdateArticleFormProps) {
  const form = useForm<z.infer<typeof updateArticleFormSchema>>({
    resolver: zodResolver(updateArticleFormSchema),
    defaultValues: {
      title: article.title,
      content: article.content,
      id: article.id,
    },
  });

  const onSubmit = async (values: z.infer<typeof updateArticleFormSchema>) => {
    const res = await updateArticleAction(values);
    if (!res.success) {
      form.setError('root', { message: res.error.message });
      console.error(res.error);
      toast.error(res.error.message);
      return;
    }
    toast.success('Article updated successfully');
    redirect(`/teacher/courses/`);
  };

  return (
    <Form {...form}>
      <form
        onSubmit={form.handleSubmit(onSubmit)}
        className="flex flex-col gap-y-4 w-full px-4 sm:px-6 md:px-8 max-w-3xl pb-20"
      >
        {updateArticleFormFields.map((field) => (
          <FormField
            key={`form-update-article-${field.name}`}
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
          Update Article
        </Button>
      </form>
    </Form>
  );
}

export { UpdateArticleForm };
