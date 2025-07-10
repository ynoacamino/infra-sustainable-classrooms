import { ReactNode, HTMLAttributes } from 'react';
import { cn } from '@/modules/core/lib/utils';

interface SectionProps extends HTMLAttributes<HTMLElement> {
  title?: string;
  children: ReactNode;
}

export default function Section({
  title,
  children,
  className,
  ...rest
}: SectionProps) {
  return (
    <section
      className={cn(
        'flex flex-col my-2 sm:my-4 md:my-6 w-full gap-4 px-4 sm:px-6 md:px-8 max-w-5xl',
        className,
      )}
      {...rest}
    >
      {title && <h2 className="font-bold text-[22px] py-3">{title}</h2>}
      {children}
    </section>
  );
}
