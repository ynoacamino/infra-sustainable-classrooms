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
