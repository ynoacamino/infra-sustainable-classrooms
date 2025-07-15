import { cn } from '@/lib/shared/utils';

/**
 * Skeleton component
 *
 * A loading placeholder component that displays an animated skeleton
 * while content is loading. Provides a visual indicator that content
 * is being fetched or processed.
 *
 * @param props - Skeleton component props
 * @param props.className - Additional CSS classes
 * @returns The rendered skeleton component
 *
 * @example
 * ```tsx
 * // Basic skeleton
 * <Skeleton className="h-4 w-32" />
 *
 * // Card skeleton
 * <div className="space-y-2">
 *   <Skeleton className="h-4 w-full" />
 *   <Skeleton className="h-4 w-3/4" />
 *   <Skeleton className="h-4 w-1/2" />
 * </div>
 *
 * // Avatar skeleton
 * <Skeleton className="h-12 w-12 rounded-full" />
 *
 * // Button skeleton
 * <Skeleton className="h-10 w-24 rounded-md" />
 * ```
 */
function Skeleton({ className, ...props }: React.ComponentProps<'div'>) {
  return (
    <div
      data-slot="skeleton"
      className={cn('bg-gray-300 animate-pulse rounded-md', className)}
      {...props}
    />
  );
}

export { Skeleton };
