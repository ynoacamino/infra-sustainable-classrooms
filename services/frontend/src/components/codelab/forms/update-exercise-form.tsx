'use client';

import { Form, FormField } from '@/ui/form';
import { zodResolver } from '@hookform/resolvers/zod';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import { InferItem } from '@/ui/infer-field';
import { Button } from '@/ui/button';
import { toast } from 'sonner';
import {
  updateExerciseFormFields,
  updateExerciseFormSchema,
} from '@/lib/codelab/forms/update-exercise-form';
import { updateExerciseAction } from '@/actions/codelab/actions';
import { redirect } from 'next/navigation';
import Editor from '@monaco-editor/react';
import { FormControl, FormItem, FormLabel, FormMessage } from '@/ui/form';
import type { Exercise } from '@/types/codelab/models';

interface UpdateExerciseFormProps {
  exercise: Exercise;
}

function UpdateExerciseForm({ exercise }: UpdateExerciseFormProps) {
  const form = useForm<z.infer<typeof updateExerciseFormSchema>>({
    resolver: zodResolver(updateExerciseFormSchema),
    defaultValues: {
      title: exercise.title,
      description: exercise.description,
      initial_code: exercise.initial_code,
      solution: exercise.solution,
      difficulty: exercise.difficulty,
    },
  });

  const onSubmit = async (values: z.infer<typeof updateExerciseFormSchema>) => {
    const res = await updateExerciseAction({
      id: exercise.id,
      exercise: values,
    });
    if (!res.success) {
      form.setError('root', { message: res.error.message });
      console.error(res.error);
      toast.error(res.error.message);
      return;
    }
    toast.success('Exercise updated successfully');
    redirect(`/teacher/codelab/exercises/${exercise.id}`);
  };

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-6">
        {updateExerciseFormFields.map((field) => {
          // Handle code editor fields specially
          if (field.name === 'initial_code' || field.name === 'solution') {
            return (
              <FormField
                key={field.name}
                control={form.control}
                name={field.name}
                render={({ field: formField }) => (
                  <FormItem>
                    <FormLabel>{field.label}</FormLabel>
                    <FormControl>
                      <div className="border rounded-md overflow-hidden">
                        <Editor
                          height="200px"
                          language="javascript"
                          theme="vs-dark"
                          value={formField.value}
                          onChange={(value) => formField.onChange(value || '')}
                          options={{
                            minimap: { enabled: false },
                            lineNumbers: 'on',
                            roundedSelection: false,
                            scrollBeyondLastLine: false,
                            automaticLayout: true,
                          }}
                        />
                      </div>
                    </FormControl>
                    {field.description && (
                      <p className="text-sm text-muted-foreground">
                        {field.description}
                      </p>
                    )}
                    <FormMessage />
                  </FormItem>
                )}
              />
            );
          }

          // Regular fields
          return (
            <FormField
              key={field.name}
              control={form.control}
              name={field.name}
              render={({ field: formField }) => (
                <InferItem {...field} {...formField} />
              )}
            />
          );
        })}

        <Button type="submit" className="w-full">
          Update Exercise
        </Button>
      </form>
    </Form>
  );
}

export { UpdateExerciseForm };
