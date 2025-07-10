import {
  courseFormFields,
  coursesFormSchema,
} from '@/lib/courses/admin/forms/courses-form';
import { Button } from '@/ui/button';
import { Form, FormField } from '@/ui/form';
import { InferItem } from '@/ui/infer-field';
import { zodResolver } from '@hookform/resolvers/zod';
import { useForm } from 'react-hook-form';
import z from 'zod';

function CoursesForm() {
  const form = useForm<z.infer<typeof coursesFormSchema>>({
    resolver: zodResolver(coursesFormSchema),
    defaultValues: {
      title: '',
      description: '',
      image: '',
      category: '',
      teacher: '',
      modules: NaN,
    },
  });

  const onSubmit = (values: z.infer<typeof coursesFormSchema>) => {
    console.log(values);
  };
  return (
    <Form {...form}>
      <form
        onSubmit={form.handleSubmit(onSubmit)}
        className="flex flex-col gap-y-4"
      >
        {courseFormFields.map((field) => (
          <FormField
            key={`form-courses-${field.name}`}
            control={form.control}
            name={field.name}
            render={({ field: formField }) => (
              <InferItem {...field} {...formField} />
            )}
          />
        ))}
        <Button type="submit" className="self-center">
          Save course
        </Button>
      </form>
    </Form>
  );
}

export { CoursesForm };
