import type { HTMLAttributes, ReactNode } from 'react';

/**
 * Props for the H1 component
 */
interface H1Props extends HTMLAttributes<HTMLHeadingElement> {
  /** Content to be rendered inside the h1 element */
  children: ReactNode;
}

/**
 * H1 component
 *
 * A styled heading component that renders an h1 element with consistent
 * spacing and typography. Includes responsive padding and centered layout.
 *
 * @param props - H1 component props
 * @param props.children - Content to be rendered inside the heading
 * @param props.rest - Additional HTML attributes for the h1 element
 * @returns The rendered H1 component
 *
 * @example
 * ```tsx
 * <H1>Page Title</H1>
 *
 * <H1 className="custom-class">
 *   Custom Styled Title
 * </H1>
 * ```
 */
export default function H1({ children, ...rest }: H1Props) {
  return (
    <div className="w-full flex justify-center py-4 px-4 sm:px-6 md:px-8 max-w-5xl">
      <div className="flex w-full flex-col gap-4">
        <h1 className="text-[32px] font-bold" {...rest}>
          {children}
        </h1>
      </div>
    </div>
  );
}
