import * as React from 'react';
import { Slot } from '@radix-ui/react-slot';
import { cva, type VariantProps } from 'class-variance-authority';
import { cn } from '@/lib/shared/utils';

/**
 * Button variants configuration using class-variance-authority
 * Defines the visual styles for different button variants and sizes
 */
const buttonVariants = cva(
  "inline-flex items-center justify-center gap-2 whitespace-nowrap rounded-md text-sm font-medium transition-all disabled:pointer-events-none disabled:opacity-50 [&_svg]:pointer-events-none [&_svg:not([class*='size-'])]:size-4 shrink-0 [&_svg]:shrink-0 outline-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 aria-invalid:border-destructive",
  {
    variants: {
      /** Visual variant of the button */
      variant: {
        /** Default primary button style */
        default:
          'bg-primary text-primary-foreground shadow-xs hover:bg-primary/90',
        /** Destructive/dangerous action button style */
        destructive:
          'bg-destructive text-white shadow-xs hover:bg-destructive/90 focus-visible:ring-destructive/20 dark:focus-visible:ring-destructive/40 dark:bg-destructive/60',
        /** Outlined button style */
        outline:
          'border bg-background shadow-xs hover:bg-accent hover:text-accent-foreground dark:bg-input/30 dark:border-input dark:hover:bg-input/50',
        /** Secondary button style */
        secondary:
          'bg-secondary text-secondary-foreground shadow-xs hover:bg-secondary/80',
        /** Selected state button style */
        selected:
          'bg-background text-accent-foreground shadow-xs border-2 border-primary',
        /** Ghost button style (minimal styling) */
        ghost:
          'hover:bg-accent hover:text-accent-foreground dark:hover:bg-accent/50',
        /** Link-styled button */
        link: 'text-primary underline-offset-4 hover:underline',
      },
      /** Size variant of the button */
      size: {
        /** Default button size */
        default: 'h-9 px-4 py-2 has-[>svg]:px-3',
        /** Small button size */
        sm: 'h-8 rounded-md gap-1.5 px-3 has-[>svg]:px-2.5',
        /** Large button size */
        lg: 'h-10 rounded-md px-6 has-[>svg]:px-4',
        /** Icon-only button size */
        icon: 'size-9',
      },
    },
    defaultVariants: {
      variant: 'default',
      size: 'default',
    },
  },
);

/**
 * Button component props interface
 */
interface ButtonProps
  extends React.ComponentProps<'button'>,
    VariantProps<typeof buttonVariants> {
  /** Whether to render as a child component (using Radix Slot) */
  asChild?: boolean;
}

/**
 * Button component
 *
 * A versatile button component with multiple variants and sizes.
 * Supports rendering as different elements using the `asChild` prop.
 *
 * @param props - Button component props
 * @param props.className - Additional CSS classes
 * @param props.variant - Visual variant of the button
 * @param props.size - Size variant of the button
 * @param props.asChild - Whether to render as a child component
 * @returns The rendered button component
 *
 * @example
 * ```tsx
 * <Button variant="default" size="lg">
 *   Click me
 * </Button>
 *
 * <Button variant="destructive" size="sm">
 *   Delete
 * </Button>
 *
 * <Button asChild>
 *   <a href="/link">Link Button</a>
 * </Button>
 * ```
 */
function Button({
  className,
  variant,
  size,
  asChild = false,
  ...props
}: ButtonProps) {
  const Comp = asChild ? Slot : 'button';

  return (
    <Comp
      data-slot="button"
      className={cn(buttonVariants({ variant, size, className }))}
      {...props}
    />
  );
}

export { Button, buttonVariants };
