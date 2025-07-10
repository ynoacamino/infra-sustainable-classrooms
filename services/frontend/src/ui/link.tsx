import { cn } from '@/lib/shared/utils';
import { Button, buttonVariants } from '@/ui/button';
import { type VariantProps } from 'class-variance-authority';
import Link from 'next/link';
import React from 'react';

function LinkComp({
  className,
  variant,
  size,
  children,
  href,
  ...props
}: React.ComponentProps<'a'> & VariantProps<typeof buttonVariants>) {
  return (
    <Button
      asChild
      variant={variant}
      size={size}
      className={cn('text-sm font-medium py-1', className)}
    >
      <Link href={href ?? ''} {...props}>
        {children}
      </Link>
    </Button>
  );
}

export { LinkComp as Link };
