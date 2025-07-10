import { zodResolver } from '@hookform/resolvers/zod';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import { useRouter } from 'next/navigation';
import { Form, FormField } from '@/ui/form';
import { InferItem } from '@/ui/infer-field';
import { Button } from '@/ui/button';
import { loginFormFields, loginFormSchema } from '@/lib/auth/forms/login-form';
import { authService } from '@/services/auth/auth';

function LoginForm() {
  const router = useRouter();
  const form = useForm<z.infer<typeof loginFormSchema>>({
    resolver: zodResolver(loginFormSchema),
    defaultValues: {
      email: '',
      password: '',
    },
  });

  const onSubmit = async (values: z.infer<typeof loginFormSchema>) => {
    const { user, error } = await authService.login(values);
    if (error) {
      form.setError('root', { message: error });
      return;
    }
    if (user) {
      router.push('/dashboard');
    }
  };

  return (
    <Form {...form}>
      <form
        onSubmit={form.handleSubmit(onSubmit)}
        className="flex flex-col gap-y-4 w-full px-4 sm:px-6 md:px-8 max-w-3xl pb-20"
      >
        {loginFormFields.map((field) => (
          <FormField
            key={`form-login-${field.name}`}
            control={form.control}
            name={field.name}
            render={({ field: formField }) => (
              <InferItem {...field} {...formField} />
            )}
          />
        ))}
        <Button type="submit" className="self-center">
          Login
        </Button>
      </form>
    </Form>
  );
}

export { LoginForm };
