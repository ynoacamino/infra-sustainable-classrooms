'use client';

import { Form, FormField } from '@/ui/form';
import { zodResolver } from '@hookform/resolvers/zod';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import { InferItem } from '@/ui/infer-field';
import { Button } from '@/ui/button';
import { toast } from 'sonner';
import {
  createExerciseFormFields,
  createExerciseFormSchema,
} from '@/lib/codelab/forms/create-exercise-form';
import { createExerciseAction } from '@/actions/codelab/actions';
import { redirect } from 'next/navigation';
import Editor from '@monaco-editor/react';
import { FormControl, FormItem, FormLabel, FormMessage } from '@/ui/form';

const template = `function solution(input) {
    // Your solution here
    // The input is a string, your output should be a string as well
}
`;

function CreateExerciseForm() {
  const form = useForm<z.infer<typeof createExerciseFormSchema>>({
    resolver: zodResolver(createExerciseFormSchema),
    defaultValues: {
      title: '',
      description: '',
      initial_code: template,
      solution: template,
      difficulty: 'easy',
      created_by: 1, // TODO: Get from auth context
    },
  });

  const onSubmit = async (values: z.infer<typeof createExerciseFormSchema>) => {
    const res = await createExerciseAction(values);
    if (!res.success) {
      form.setError('root', { message: res.error.message });
      console.error(res.error);
      toast.error(res.error.message);
      return;
    }
    form.reset();
    toast.success('Exercise created successfully');
    redirect(`/teacher/codelab/exercises`);
  };

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-6">
        {createExerciseFormFields.map((field) => {
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
                          height="400px"
                          language="javascript"
                          theme="vs-dark"
                          value={formField.value}
                          onChange={(value) => formField.onChange(value || '')}
                          options={{
                            minimap: { enabled: true },
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

          // Skip hidden fields
          if (field.name === 'created_by') {
            return null;
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
          Create Exercise
        </Button>
      </form>
    </Form>
  );
}

export { CreateExerciseForm };
