'use client';

import { useTheme } from 'next-themes';
import { Toaster as Sonner, type ToasterProps } from 'sonner';

/**
 * Toaster component
 *
 * A toast notification component built on top of the Sonner library.
 * Automatically adapts to the current theme (light/dark) and provides
 * styled toast notifications with consistent design system integration.
 *
 * @param props - Toaster component props from Sonner
 * @returns The rendered toaster component
 *
 * @example
 * ```tsx
 * // Add to your app layout or root component
 * import { Toaster } from '@/ui/sonner';
 *
 * function App() {
 *   return (
 *     <div>
 *       <main>Your app content</main>
 *       <Toaster />
 *     </div>
 *   );
 * }
 *
 * // Use with toast functions
 * import { toast } from 'sonner';
 *
 * toast.success('Operation completed!');
 * toast.error('Something went wrong');
 * toast('Default message');
 * ```
 */
const Toaster = ({ ...props }: ToasterProps) => {
  const { theme = 'system' } = useTheme();

  return (
    <Sonner
      theme={theme as ToasterProps['theme']}
      className="toaster group"
      style={
        {
          '--normal-bg': 'var(--popover)',
          '--normal-text': 'var(--popover-foreground)',
          '--normal-border': 'var(--border)',
        } as React.CSSProperties
      }
      {...props}
    />
  );
};

export { Toaster };
