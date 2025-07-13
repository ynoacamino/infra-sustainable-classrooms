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
  updateProfileFormFields,
  updateProfilePayloadSchema,
} from '@/lib/profiles/forms/update-profile-form';
import {
  createTeacherProfileAction,
  updateProfileAction,
} from '@/actions/profiles/actions';
import type { CompleteProfile } from '@/types/profiles/models';
import { redirect } from 'next/navigation';

function UpdateProfileForm({ profile }: { profile: CompleteProfile }) {
  const form = useForm<z.infer<typeof updateProfilePayloadSchema>>({
    resolver: zodResolver(updateProfilePayloadSchema),
    defaultValues: profile,
  });
  const onSubmit = async (
    values: z.infer<typeof updateProfilePayloadSchema>,
  ) => {
    const res = await updateProfileAction(values);
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
        {updateProfileFormFields.map((field) => (
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
          Update Profile
        </Button>
      </form>
    </Form>
  );
}

export { UpdateProfileForm };
