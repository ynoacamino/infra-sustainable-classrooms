import { cn } from '@/lib/shared/utils';
import * as React from 'react';

/**
 * Textarea component
 *
 * A styled textarea component with consistent design system integration.
 * Provides focus states, validation styling, and automatic field sizing.
 * Supports all standard textarea attributes and behaviors.
 *
 * @param props - Textarea component props
 * @param props.className - Additional CSS classes
 * @returns The rendered textarea component
 *
 * @example
 * ```tsx
 * <Textarea placeholder="Enter your message..." />
 *
 * <Textarea
 *   placeholder="Description"
 *   rows={4}
 *   aria-invalid={hasError}
 * />
 *
 * <Textarea
 *   className="custom-textarea"
 *   defaultValue="Default text content"
 *   disabled
 * />
 * ```
 */
function Textarea({ className, ...props }: React.ComponentProps<'textarea'>) {
  return (
    <textarea
      data-slot="textarea"
      className={cn(
        'border-input placeholder:text-muted-foreground focus-visible:border-ring focus-visible:ring-ring/50 aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 aria-invalid:border-destructive dark:bg-input/30 flex field-sizing-content min-h-16 w-full rounded-md border bg-transparent px-3 py-2 text-base shadow-xs transition-[color,box-shadow] outline-none focus-visible:ring-[3px] disabled:cursor-not-allowed disabled:opacity-50 md:text-sm',
        className,
      )}
      {...props}
    />
  );
}

export { Textarea };
