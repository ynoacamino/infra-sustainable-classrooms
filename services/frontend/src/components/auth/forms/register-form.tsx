'use client';

import { Form, FormField } from '@/ui/form';
import { zodResolver } from '@hookform/resolvers/zod';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import { InferItem } from '@/ui/infer-field';
import { Button } from '@/ui/button';
import { authFormFields, authFormSchema } from '@/lib/auth/forms/auth-form';
import { generateSecretAction } from '@/actions/auth/actions';
import { toast } from 'sonner';
import { redirect, RedirectType } from 'next/navigation';

function RegisterForm() {
  const form = useForm<z.infer<typeof authFormSchema>>({
    resolver: zodResolver(authFormSchema),
    defaultValues: {
      identifier: '',
    },
  });
  const onSubmit = async (values: z.infer<typeof authFormSchema>) => {
    const res = await generateSecretAction(values);
    if (res.success) {
      toast.success('Login successful!');
      redirect('/auth/register/save', RedirectType.push);
    } else {
      toast.error(`Login failed. Please try again. ${res.error.message}`);
    }
  };
  return (
    <Form {...form}>
      <form
        onSubmit={form.handleSubmit(onSubmit)}
        className="flex flex-col gap-y-4 w-full px-4 sm:px-6 md:px-8 max-w-3xl pb-20"
      >
        {authFormFields.map((field) => (
          <FormField
            key={`form-register-${field.name}`}
            control={form.control}
            name={field.name}
            render={({ field: formField }) => (
              <InferItem {...field} {...formField} />
            )}
          />
        ))}
        <Button type="submit" className="self-center">
          Register
        </Button>
      </form>
    </Form>
  );
}

export { RegisterForm };
