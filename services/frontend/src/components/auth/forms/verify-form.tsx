'use client';

import { Form, FormField } from '@/ui/form';
import { zodResolver } from '@hookform/resolvers/zod';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import { InferItem } from '@/ui/infer-field';
import { Button } from '@/ui/button';
import { verifyTotpAction } from '@/actions/auth/actions';
import {
  verifyFormFields,
  verifyFormSchema,
} from '@/lib/auth/forms/verify-form';
import { toast } from 'sonner';

function VerifyForm() {
  const form = useForm<z.infer<typeof verifyFormSchema>>({
    resolver: zodResolver(verifyFormSchema),
    defaultValues: {
      identifier: '',
    },
  });
  const onSubmit = async (values: z.infer<typeof verifyFormSchema>) => {
    const res = await verifyTotpAction(values);
    if (!res.success) {
      form.setError('root', { message: res.error.message });
      toast.error(res.error.message);
      return;
    }
    form.reset();
    toast.success('Login successful');
  };
  return (
    <Form {...form}>
      <form
        onSubmit={form.handleSubmit(onSubmit)}
        className="flex flex-col gap-y-4 w-full px-4 sm:px-6 md:px-8 max-w-3xl pb-20"
      >
        {verifyFormFields.map((field) => (
          <FormField
            key={`form-register-${field.name}`}
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
          Login
        </Button>
      </form>
    </Form>
  );
}

export { VerifyForm };
