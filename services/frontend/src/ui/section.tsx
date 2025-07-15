import { cn } from '@/lib/shared/utils';
import type { ReactNode, HTMLAttributes } from 'react';

/**
 * Props for the Section component
 */
interface SectionProps extends HTMLAttributes<HTMLElement> {
  /** Optional title to display at the top of the section */
  title?: string;
  /** Content to be rendered inside the section */
  children: ReactNode;
}

/**
 * Section component
 *
 * A layout component that provides consistent spacing and styling for content sections.
 * Includes optional title display and responsive padding.
 *
 * @param props - Section component props
 * @param props.title - Optional title to display at the top of the section
 * @param props.children - Content to be rendered inside the section
 * @param props.className - Additional CSS classes
 * @param props.rest - Additional HTML attributes for the section element
 * @returns The rendered section component
 *
 * @example
 * ```tsx
 * <Section title="User Profile">
 *   <p>User profile content goes here</p>
 * </Section>
 *
 * <Section className="custom-section">
 *   <div>Content without title</div>
 * </Section>
 *
 * <Section title="Settings" id="settings-section">
 *   <form>Settings form content</form>
 * </Section>
 * ```
 */
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
