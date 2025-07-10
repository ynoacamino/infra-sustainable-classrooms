import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from '@/ui/form';
import { zodResolver } from '@hookform/resolvers/zod';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import { InferItem } from '@/ui/infer-field';
import type { Role } from '@/types/auth/role';
import RoleSelector from '@/components/auth/forms/role-selector';
import { Button } from '@/ui/button';
import {
  registerFormFields,
  registerFormSchema,
} from '@/lib/auth/forms/register-form';
import { Roles } from '@/lib/auth/enums/roles';

function RegisterForm() {
  const form = useForm<z.infer<typeof registerFormSchema>>({
    resolver: zodResolver(registerFormSchema),
    defaultValues: {
      names: '',
      email: '',
      password: '',
      confirmPassword: '',
      role: Roles.Student, // Default to the first role
    },
  });
  const rolField = registerFormFields.find((field) => field.name === 'role');
  const onSubmit = (values: z.infer<typeof registerFormSchema>) => {
    // Handle form submission logic here
    console.log('Form submitted with values:', values);
  };
  return (
    <Form {...form}>
      <form
        onSubmit={form.handleSubmit(onSubmit)}
        className="flex flex-col gap-y-4 w-full px-4 sm:px-6 md:px-8 max-w-3xl pb-20"
      >
        {registerFormFields
          .filter((field) => field.name !== 'role')
          .map((field) => (
            <FormField
              key={`form-register-${field.name}`}
              control={form.control}
              name={field.name}
              render={({ field: formField }) => (
                <InferItem {...field} {...formField} />
              )}
            />
          ))}
        <FormField
          control={form.control}
          name="role"
          render={({ field: formField }) => (
            <FormItem>
              <FormLabel>{rolField?.label}</FormLabel>
              <FormControl>
                <RoleSelector
                  value={formField.value as Role}
                  onChange={formField.onChange}
                />
              </FormControl>
              <FormDescription>{rolField?.description}</FormDescription>
              <FormMessage />
            </FormItem>
          )}
        />
        <Button type="submit" className="self-center">
          Register
        </Button>
      </form>
    </Form>
  );
}

export { RegisterForm };
