'use client';

import * as LabelPrimitive from '@radix-ui/react-label';
import { Slot } from '@radix-ui/react-slot';
import {
  Controller,
  FormProvider,
  useFormContext,
  useFormState,
  type ControllerProps,
  type FieldPath,
  type FieldValues,
} from 'react-hook-form';

import { createContext, useContext, useId } from 'react';
import { cn } from '@/lib/shared/utils';
import { Label } from '@/ui/label';

/**
 * Form component
 *
 * Root form component that provides form context using React Hook Form.
 * This is an alias for React Hook Form's FormProvider.
 *
 * @example
 * ```tsx
 * const form = useForm();
 *
 * <Form {...form}>
 *   <form onSubmit={form.handleSubmit(onSubmit)}>
 *     <FormField
 *       control={form.control}
 *       name="username"
 *       render={({ field }) => (
 *         <FormItem>
 *           <FormLabel>Username</FormLabel>
 *           <FormControl>
 *             <Input {...field} />
 *           </FormControl>
 *         </FormItem>
 *       )}
 *     />
 *   </form>
 * </Form>
 * ```
 */
const Form = FormProvider;

/**
 * Context value for FormField
 * @template TFieldValues - Form field values type
 * @template TName - Field name type
 */
type FormFieldContextValue<
  TFieldValues extends FieldValues = FieldValues,
  TName extends FieldPath<TFieldValues> = FieldPath<TFieldValues>,
> = {
  /** Name of the form field */
  name: TName;
};

const FormFieldContext = createContext<FormFieldContextValue>(
  {} as FormFieldContextValue,
);

/**
 * FormField component
 *
 * Wrapper around React Hook Form's Controller that provides field context.
 *
 * @template TFieldValues - Form field values type
 * @template TName - Field name type
 * @param props - Controller props from React Hook Form
 * @returns The form field component
 *
 * @example
 * ```tsx
 * <FormField
 *   control={form.control}
 *   name="email"
 *   render={({ field }) => (
 *     <FormItem>
 *       <FormLabel>Email</FormLabel>
 *       <FormControl>
 *         <Input type="email" {...field} />
 *       </FormControl>
 *       <FormMessage />
 *     </FormItem>
 *   )}
 * />
 * ```
 */
const FormField = <
  TFieldValues extends FieldValues = FieldValues,
  TName extends FieldPath<TFieldValues> = FieldPath<TFieldValues>,
>({
  ...props
}: ControllerProps<TFieldValues, TName>) => {
  return (
    <FormFieldContext.Provider value={{ name: props.name }}>
      <Controller {...props} />
    </FormFieldContext.Provider>
  );
};

/**
 * useFormField hook
 *
 * Hook that provides access to form field state and IDs.
 * Must be used within a FormField component.
 *
 * @returns Object containing field state and generated IDs
 * @throws Error if used outside of FormField context
 *
 * @example
 * ```tsx
 * function CustomFormInput() {
 *   const { error, formItemId } = useFormField();
 *
 *   return (
 *     <input
 *       id={formItemId}
 *       aria-invalid={!!error}
 *       className={error ? 'error' : ''}
 *     />
 *   );
 * }
 * ```
 */
const useFormField = () => {
  const fieldContext = useContext(FormFieldContext);
  const itemContext = useContext(FormItemContext);
  const { getFieldState } = useFormContext();
  const formState = useFormState({ name: fieldContext.name });
  const fieldState = getFieldState(fieldContext.name, formState);

  if (!fieldContext) {
    throw new Error('useFormField should be used within <FormField>');
  }

  const { id } = itemContext;

  return {
    id,
    name: fieldContext.name,
    formItemId: `${id}-form-item`,
    formDescriptionId: `${id}-form-item-description`,
    formMessageId: `${id}-form-item-message`,
    ...fieldState,
  };
};

/**
 * Context value for FormItem
 */
type FormItemContextValue = {
  /** Unique identifier for the form item */
  id: string;
};

const FormItemContext = createContext<FormItemContextValue>(
  {} as FormItemContextValue,
);

/**
 * FormItem component
 *
 * Container for form field components. Provides context with unique ID.
 *
 * @param props - Standard div props
 * @param props.className - Additional CSS classes
 * @returns The form item component
 *
 * @example
 * ```tsx
 * <FormItem>
 *   <FormLabel>Username</FormLabel>
 *   <FormControl>
 *     <Input />
 *   </FormControl>
 *   <FormDescription>Enter your username</FormDescription>
 *   <FormMessage />
 * </FormItem>
 * ```
 */
function FormItem({ className, ...props }: React.ComponentProps<'div'>) {
  const id = useId();

  return (
    <FormItemContext.Provider value={{ id }}>
      <div
        data-slot="form-item"
        className={cn('grid gap-2', className)}
        {...props}
      />
    </FormItemContext.Provider>
  );
}

/**
 * FormLabel component
 *
 * Label component for form fields. Automatically associates with form controls
 * and shows error styling when the field has validation errors.
 *
 * @param props - Label props
 * @param props.className - Additional CSS classes
 * @returns The form label component
 *
 * @example
 * ```tsx
 * <FormLabel>Email Address</FormLabel>
 * ```
 */
function FormLabel({
  className,
  ...props
}: React.ComponentProps<typeof LabelPrimitive.Root>) {
  const { error, formItemId } = useFormField();

  return (
    <Label
      data-slot="form-label"
      data-error={!!error}
      className={cn('data-[error=true]:text-destructive', className)}
      htmlFor={formItemId}
      {...props}
    />
  );
}

/**
 * FormControl component
 *
 * Wrapper for form control elements (inputs, selects, etc.).
 * Automatically sets accessibility attributes and IDs.
 *
 * @param props - Slot props
 * @returns The form control component
 *
 * @example
 * ```tsx
 * <FormControl>
 *   <Input type="email" />
 * </FormControl>
 * ```
 */
function FormControl({ ...props }: React.ComponentProps<typeof Slot>) {
  const { error, formItemId, formDescriptionId, formMessageId } =
    useFormField();

  return (
    <Slot
      data-slot="form-control"
      id={formItemId}
      aria-describedby={
        !error
          ? `${formDescriptionId}`
          : `${formDescriptionId} ${formMessageId}`
      }
      aria-invalid={!!error}
      {...props}
    />
  );
}

/**
 * FormDescription component
 *
 * Provides helpful description text for form fields.
 * Automatically linked to form controls for accessibility.
 *
 * @param props - Paragraph props
 * @param props.className - Additional CSS classes
 * @returns The form description component
 *
 * @example
 * ```tsx
 * <FormDescription>
 *   We'll never share your email with anyone else.
 * </FormDescription>
 * ```
 */
function FormDescription({ className, ...props }: React.ComponentProps<'p'>) {
  const { formDescriptionId } = useFormField();

  return (
    <p
      data-slot="form-description"
      id={formDescriptionId}
      className={cn('text-muted-foreground text-sm', className)}
      {...props}
    />
  );
}

/**
 * FormMessage component
 *
 * Displays validation error messages for form fields.
 * Automatically shows error messages from form validation.
 *
 * @param props - Paragraph props
 * @param props.className - Additional CSS classes
 * @returns The form message component or null if no error
 *
 * @example
 * ```tsx
 * <FormMessage />
 * ```
 */
function FormMessage({ className, ...props }: React.ComponentProps<'p'>) {
  const { error, formMessageId } = useFormField();
  const body = error ? String(error?.message ?? '') : props.children;

  if (!body) {
    return null;
  }

  return (
    <p
      data-slot="form-message"
      id={formMessageId}
      className={cn('text-destructive text-sm', className)}
      {...props}
    >
      {body}
    </p>
  );
}

export {
  useFormField,
  Form,
  FormItem,
  FormLabel,
  FormControl,
  FormDescription,
  FormMessage,
  FormField,
};
