'use client';

import { Form, FormField } from '@/ui/form';
import { zodResolver } from '@hookform/resolvers/zod';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import { InferItem } from '@/ui/infer-field';
import { Button } from '@/ui/button';
import { toast } from 'sonner';
import {
  updateQuestionFormFields,
  updateQuestionFormSchema,
} from '@/lib/knowledge/forms/update-question-form';
import type { Question } from '@/types/knowledge/models';
import { updateQuestionAction } from '@/actions/knowledge/actions';
import { redirect } from 'next/navigation';

interface UpdateQuestionFormProps {
  question: Question;
}

function UpdateQuestionForm({ question }: UpdateQuestionFormProps) {
  const form = useForm<z.infer<typeof updateQuestionFormSchema>>({
    resolver: zodResolver(updateQuestionFormSchema),
    defaultValues: {
      test_id: question.test_id,
      id: question.id,
      question_text: question.question_text,
      option_a: question.option_a,
      option_b: question.option_b,
      option_c: question.option_c,
      option_d: question.option_d,
      correct_answer: question.correct_answer,
    },
  });

  const onSubmit = async (values: z.infer<typeof updateQuestionFormSchema>) => {
    const res = await updateQuestionAction(values);
    if (!res.success) {
      form.setError('root', { message: res.error.message });
      console.error(res.error);
      toast.error(res.error.message);
      return;
    }
    toast.success('Question updated successfully');
    redirect(`/teacher/tests/${question.test_id}/questions`);
  };

  return (
    <Form {...form}>
      <form
        onSubmit={form.handleSubmit(onSubmit)}
        className="flex flex-col gap-y-4 w-full px-4 sm:px-6 md:px-8 max-w-3xl pb-20"
      >
        {updateQuestionFormFields.map((field) => (
          <FormField
            key={`form-update-question-${field.name}`}
            control={form.control}
            name={field.name}
            render={({ field: formField }) => (
              <InferItem
                {...field}
                {...formField}
                onChange={
                  field.name === 'correct_answer'
                    ? (value: string) => formField.onChange(parseInt(value, 10))
                    : formField.onChange
                }
              />
            )}
          />
        ))}

        <div className="flex gap-2 mt-6">
          <Button type="submit" disabled={form.formState.isSubmitting}>
            {form.formState.isSubmitting ? 'Updating...' : 'Update Question'}
          </Button>
        </div>
      </form>
    </Form>
  );
}

export { UpdateQuestionForm };
