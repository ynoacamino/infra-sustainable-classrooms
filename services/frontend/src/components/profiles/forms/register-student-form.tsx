'use client';

import { Form, FormField } from '@/ui/form';
import { zodResolver } from '@hookform/resolvers/zod';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import { InferItem } from '@/ui/infer-field';
import { Button } from '@/ui/button';
import { verifyTotpAction } from '@/actions/auth/actions';
import { toast } from 'sonner';
import {
  registerStudentFormFields,
  registerStudentFormSchema,
} from '@/lib/profiles/forms/register-student-form';
import { createStudentProfileAction } from '@/actions/profiles/actions';
import { redirect, RedirectType } from 'next/navigation';

function RegisterStudentForm() {
  const form = useForm<z.infer<typeof registerStudentFormSchema>>({
    resolver: zodResolver(registerStudentFormSchema),
    defaultValues: {
      avatar_url:
        'https://ynoa-uploader.ynoacamino.site/uploads/1752353997_profile.webp',
      bio: '',
      email: '',
      first_name: '',
      last_name: '',
      phone: '',
      grade_level: '',
      major: '',
    },
  });
  const onSubmit = async (
    values: z.infer<typeof registerStudentFormSchema>,
  ) => {
    const res = await createStudentProfileAction(values);
    if (!res.success) {
      form.setError('root', { message: res.error.message });
      console.error(res.error);
      toast.error(res.error.message);
      return;
    }
    form.reset();
    toast.success('Student profile created successfully');
    redirect('/dashboard', RedirectType.push);
  };
  return (
    <Form {...form}>
      <form
        onSubmit={form.handleSubmit(onSubmit)}
        className="flex flex-col gap-y-4 w-full px-4 sm:px-6 md:px-8 max-w-3xl pb-20"
      >
        {registerStudentFormFields.map((field) => (
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
          Register Student
        </Button>
      </form>
    </Form>
  );
}

export { RegisterStudentForm };
