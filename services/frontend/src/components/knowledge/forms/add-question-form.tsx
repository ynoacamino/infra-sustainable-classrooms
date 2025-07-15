'use client';

import { Form, FormField } from '@/ui/form';
import { zodResolver } from '@hookform/resolvers/zod';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import { InferItem } from '@/ui/infer-field';
import { Button } from '@/ui/button';
import { toast } from 'sonner';
import {
  addQuestionFormFields,
  addQuestionFormSchema,
} from '@/lib/knowledge/forms/add-question-form';
import { addQuestionAction } from '@/actions/knowledge/actions';
import { redirect } from 'next/navigation';

interface AddQuestionFormProps {
  testId: number;
}

function AddQuestionForm({ testId }: AddQuestionFormProps) {
  const form = useForm<z.infer<typeof addQuestionFormSchema>>({
    resolver: zodResolver(addQuestionFormSchema),
    defaultValues: {
      test_id: testId,
      question_text: '',
      option_a: '',
      option_b: '',
      option_c: '',
      option_d: '',
      correct_answer: 0,
    },
  });

  const onSubmit = async (values: z.infer<typeof addQuestionFormSchema>) => {
    const res = await addQuestionAction(values);
    if (!res.success) {
      form.setError('root', { message: res.error.message });
      console.error(res.error);
      toast.error(res.error.message);
      return;
    }
    form.reset({
      test_id: testId,
      question_text: '',
      option_a: '',
      option_b: '',
      option_c: '',
      option_d: '',
      correct_answer: 0,
    });
    toast.success('Question added successfully');
    redirect(`/teacher/tests/${testId}/questions`);
  };

  return (
    <Form {...form}>
      <form
        onSubmit={form.handleSubmit(onSubmit)}
        className="flex flex-col gap-y-4 w-full px-4 sm:px-6 md:px-8 max-w-3xl pb-20"
      >
        {addQuestionFormFields.map((field) => (
          <FormField
            key={`form-add-question-${field.name}`}
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
            {form.formState.isSubmitting ? 'Adding...' : 'Add Question'}
          </Button>
        </div>
      </form>
    </Form>
  );
}

export { AddQuestionForm };
