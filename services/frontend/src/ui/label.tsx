import { cn } from '@/lib/shared/utils';
import * as LabelPrimitive from '@radix-ui/react-label';

/**
 * Label component
 *
 * A styled label component built on top of Radix UI Label primitive.
 * Provides consistent typography and accessibility features for form labels.
 * Automatically handles disabled states and peer interactions.
 *
 * @param props - Label component props
 * @param props.className - Additional CSS classes
 * @returns The rendered label component
 *
 * @example
 * ```tsx
 * <Label htmlFor="email">Email Address</Label>
 * <Input id="email" type="email" />
 *
 * <Label className="custom-label">
 *   Custom Styled Label
 * </Label>
 *
 * // With form field association
 * <div className="form-field">
 *   <Label htmlFor="username">Username</Label>
 *   <Input id="username" type="text" />
 * </div>
 * ```
 */
function Label({
  className,
  ...props
}: React.ComponentProps<typeof LabelPrimitive.Root>) {
  return (
    <LabelPrimitive.Root
      data-slot="label"
      className={cn(
        'flex items-center gap-2 text-sm leading-none font-medium select-none group-data-[disabled=true]:pointer-events-none group-data-[disabled=true]:opacity-50 peer-disabled:cursor-not-allowed peer-disabled:opacity-50',
        className,
      )}
      {...props}
    />
  );
}

export { Label };
