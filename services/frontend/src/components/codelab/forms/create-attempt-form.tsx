'use client';

import { Form, FormField } from '@/ui/form';
import { zodResolver } from '@hookform/resolvers/zod';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import { Button } from '@/ui/button';
import { toast } from 'sonner';
import {
  createAttemptFormFields,
  createAttemptFormSchema,
} from '@/lib/codelab/forms/create-attempt-form';
import { createAttemptAction } from '@/actions/codelab/actions';
import Editor from '@monaco-editor/react';
import { FormControl, FormItem, FormLabel, FormMessage } from '@/ui/form';

interface CreateAttemptFormProps {
  exerciseId: number;
  onSuccess?: () => void;
}

function CreateAttemptForm({ exerciseId, onSuccess }: CreateAttemptFormProps) {
  const form = useForm<z.infer<typeof createAttemptFormSchema>>({
    resolver: zodResolver(createAttemptFormSchema),
    defaultValues: {
      exercise_id: exerciseId,
      code: '',
      success: false, // Will be determined by backend
    },
  });

  const onSubmit = async (values: z.infer<typeof createAttemptFormSchema>) => {
    const res = await createAttemptAction(values);
    if (!res.success) {
      form.setError('root', { message: res.error.message });
      console.error(res.error);
      toast.error(res.error.message);
      return;
    }
    toast.success('Code submitted successfully');
    onSuccess?.();
  };

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4">
        {createAttemptFormFields.map((field) => {
          if (field.name === 'code') {
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
                          height="300px"
                          defaultLanguage="python"
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

          return null; // Skip other fields as they're hidden
        })}

        <div className="flex gap-2">
          <Button type="submit" className="flex-1">
            Submit Code
          </Button>
          <Button 
            type="button" 
            variant="outline" 
            onClick={() => form.reset()}
            className="flex-1"
          >
            Reset
          </Button>
        </div>
      </form>
    </Form>
  );
}

export { CreateAttemptForm };
