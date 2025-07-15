import { cn } from '@/lib/shared/utils';

/**
 * Input component
 *
 * A styled input component with consistent design system integration.
 * Provides focus states, validation styling, and file input support.
 *
 * @param props - Input component props
 * @param props.className - Additional CSS classes
 * @param props.type - Input type (text, email, password, file, etc.)
 * @returns The rendered input component
 *
 * @example
 * ```tsx
 * <Input type="text" placeholder="Enter your name" />
 *
 * <Input type="email" placeholder="Enter your email" />
 *
 * <Input type="file" accept="image/*" />
 *
 * <Input
 *   type="password"
 *   placeholder="Enter password"
 *   aria-invalid={hasError}
 * />
 * ```
 */
function Input({ className, type, ...props }: React.ComponentProps<'input'>) {
  return (
    <input
      type={type}
      data-slot="input"
      className={cn(
        'file:text-foreground placeholder:text-muted-foreground selection:bg-primary selection:text-primary-foreground dark:bg-input/30 border-input flex h-9 w-full min-w-0 rounded-md border bg-transparent px-3 py-1 text-base shadow-xs transition-[color,box-shadow] outline-none file:inline-flex file:h-7 file:border-0 file:bg-transparent file:text-sm file:font-medium disabled:pointer-events-none disabled:cursor-not-allowed disabled:opacity-50 md:text-sm',
        'focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px]',
        'aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 aria-invalid:border-destructive',
        className,
      )}
      {...props}
    />
  );
}

export { Input };
