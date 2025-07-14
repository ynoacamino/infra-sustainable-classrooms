'use client';

import { Form, FormField } from '@/ui/form';
import { zodResolver } from '@hookform/resolvers/zod';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import { InferItem } from '@/ui/infer-field';
import { Button } from '@/ui/button';
import { toast } from 'sonner';
import {
  registerTeacherFormFields,
  registerTeacherFormSchema,
} from '@/lib/profiles/forms/register-teacher-form';
import { createTeacherProfileAction } from '@/actions/profiles/actions';
import { redirect } from 'next/navigation';

function RegisterTeacherForm() {
  const form = useForm<z.infer<typeof registerTeacherFormSchema>>({
    resolver: zodResolver(registerTeacherFormSchema),
    defaultValues: {
      avatar_url:
        'https://ynoa-uploader.ynoacamino.site/uploads/1752353997_profile.webp',
      bio: '',
      email: '',
      first_name: '',
      last_name: '',
      phone: '',
      position: '',
    },
  });
  const onSubmit = async (
    values: z.infer<typeof registerTeacherFormSchema>,
  ) => {
    const res = await createTeacherProfileAction(values);
    if (!res.success) {
      form.setError('root', { message: res.error.message });
      console.error(res.error);
      toast.error(res.error.message);
      return;
    }
    form.reset();
    toast.success('Teacher profile created successfully');
    redirect('/dashboard');
  };
  return (
    <Form {...form}>
      <form
        onSubmit={form.handleSubmit(onSubmit)}
        className="flex flex-col gap-y-4 w-full px-4 sm:px-6 md:px-8 max-w-3xl pb-20"
      >
        {registerTeacherFormFields.map((field) => (
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
          Register Teacher
        </Button>
      </form>
    </Form>
  );
}

export { RegisterTeacherForm };
