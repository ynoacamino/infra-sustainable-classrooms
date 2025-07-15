import type {
  ControllerRenderProps,
  FieldPath,
  FieldValues,
} from 'react-hook-form';
import type { Field } from '@/types/shared/field';
import {
  FormControl,
  FormDescription,
  FormItem,
  FormLabel,
  FormMessage,
} from '@/ui/form';
import { SupportedFields } from '@/lib/shared/enums/field';
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/ui/select';
import { Textarea } from '@/ui/textarea';
import { Input } from '@/ui/input';
import { TwoDividedInputOpt } from '@/ui/two-divided-input-opt';
import dynamic from 'next/dynamic';
// import QuillEditor from '@/components/text/editor/quill-editor';

const JoditEditor = dynamic(
  () => import('@/components/text/editor/joddit-editor'),
  {
    ssr: false,
  },
);

/**
 * InferItem component
 *
 * A dynamic form field component that automatically renders the appropriate
 * input type based on the field configuration. Supports various input types
 * including text, select, textarea, OTP, file upload, and number inputs.
 *
 * @template FieldName - The field name type
 * @template TFieldValues - Form field values type
 * @template TName - Field path type
 * @param props - Combined field configuration and React Hook Form controller props
 * @param props.label - Label text for the field
 * @param props.description - Optional description text for the field
 * @param props.type - The type of input to render (from SupportedFields enum)
 * @returns The rendered form field component
 *
 * @example
 * ```tsx
 * // Text input
 * <InferItem
 *   label="Username"
 *   type={SupportedFields.TEXT}
 *   name="username"
 *   {...fieldProps}
 * />
 *
 * // Select dropdown
 * <InferItem
 *   label="Country"
 *   type={SupportedFields.SELECT}
 *   options={[
 *     { key: 'us', value: 'US', textValue: 'United States' },
 *     { key: 'ca', value: 'CA', textValue: 'Canada' }
 *   ]}
 *   name="country"
 *   {...fieldProps}
 * />
 *
 * // File upload
 * <InferItem
 *   label="Profile Picture"
 *   type={SupportedFields.FILE}
 *   name="avatar"
 *   {...fieldProps}
 * />
 * ```
 */
function InferItem<
  FieldName extends string,
  TFieldValues extends FieldValues = FieldValues,
  TName extends FieldPath<TFieldValues> = FieldPath<TFieldValues>,
>({
  label,
  description,
  ...props
}: Field<FieldName> & ControllerRenderProps<TFieldValues, TName>) {
  return (
    <FormItem>
      <FormLabel>{label}</FormLabel>
      <FormControl>
        {(() => {
          if (props.type === SupportedFields.SELECT) {
            return (
              <Select onValueChange={props.onChange} defaultValue={props.value}>
                <SelectTrigger className="w-full">
                  <SelectValue placeholder={props.placeholder} />
                </SelectTrigger>
                <SelectContent>
                  {props.options.map(({ key, value, textValue }) => (
                    <SelectItem key={`select-${key}-${value}`} value={value}>
                      {textValue || value}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>
            );
          } else if (props.type === SupportedFields.TEXTAREA) {
            return <Textarea {...props} />;
          } else if (props.type === SupportedFields.HTML_EDITOR) {
            return (
              <JoditEditor
                content={props.value as string}
                setContent={props.onChange as (content: string) => void}
                placeholder={props.placeholder}
              />
            );
          } else if (props.type === SupportedFields.OTP) {
            return <TwoDividedInputOpt {...props} />;
          } else if (props.type === SupportedFields.FILE) {
            return (
              <Input
                type={props.type}
                name={props.name}
                onChange={(e) => {
                  const file = e.target.files?.[0];
                  props.onChange(file);
                }}
              />
            );
          } else {
            return (
              <Input
                {...props}
                onChange={
                  props.type === SupportedFields.NUMBER
                    ? (e) => props.onChange(Number(e.target.value))
                    : props.onChange
                }
              />
            );
          }
        })()}
      </FormControl>
      {description && <FormDescription>{description}</FormDescription>}
      <FormMessage />
    </FormItem>
  );
}

export { InferItem };
