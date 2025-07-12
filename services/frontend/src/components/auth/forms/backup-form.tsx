'use client';

import { Form, FormField } from '@/ui/form';
import { zodResolver } from '@hookform/resolvers/zod';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import { InferItem } from '@/ui/infer-field';
import { Button } from '@/ui/button';
import { verifyBackupCodeAction } from '@/actions/auth/actions';
import { toast } from 'sonner';
import {
  backupFormFields,
  backupFormSchema,
} from '@/lib/auth/forms/backup-form';

function BackupForm() {
  const form = useForm<z.infer<typeof backupFormSchema>>({
    resolver: zodResolver(backupFormSchema),
    defaultValues: {
      identifier: '',
      backup_code: '',
    },
  });

  const onSubmit = async (values: z.infer<typeof backupFormSchema>) => {
    const res = await verifyBackupCodeAction(values);
    if (!res.success) {
      form.setError('root', { message: res.error.message });
      toast.error(res.error.message);
      return;
    }
    form.reset();
    toast.success('Login successful with backup code');
  };

  return (
    <Form {...form}>
      <form
        onSubmit={form.handleSubmit(onSubmit)}
        className="flex flex-col gap-y-4 w-full px-4 sm:px-6 md:px-8 max-w-3xl pb-20"
      >
        {backupFormFields.map((fieldConfig) => (
          <FormField
            key={`form-backup-${fieldConfig.name}`}
            control={form.control}
            name={fieldConfig.name}
            render={({ field }) => <InferItem {...fieldConfig} {...field} />}
          />
        ))}
        <Button
          type="submit"
          className="self-center"
          disabled={form.formState.isSubmitting}
        >
          Verify Backup Code
        </Button>
      </form>
    </Form>
  );
}

export { BackupForm };
