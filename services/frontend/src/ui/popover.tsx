import { cn } from '@/lib/shared/utils';
import * as PopoverPrimitive from '@radix-ui/react-popover';

/**
 * Popover component
 *
 * Root component for popovers. Provides context for all popover components.
 * Built on top of Radix UI Popover primitive.
 *
 * @param props - Root popover props
 * @returns The popover root component
 *
 * @example
 * ```tsx
 * <Popover>
 *   <PopoverTrigger>Open popover</PopoverTrigger>
 *   <PopoverContent>
 *     <p>Popover content goes here</p>
 *   </PopoverContent>
 * </Popover>
 * ```
 */
function Popover({
  ...props
}: React.ComponentProps<typeof PopoverPrimitive.Root>) {
  return <PopoverPrimitive.Root data-slot="popover" {...props} />;
}

/**
 * PopoverTrigger component
 *
 * Trigger element that opens the popover when activated.
 *
 * @param props - Trigger props
 * @returns The popover trigger component
 *
 * @example
 * ```tsx
 * <PopoverTrigger asChild>
 *   <Button>Click me</Button>
 * </PopoverTrigger>
 * ```
 */
function PopoverTrigger({
  ...props
}: React.ComponentProps<typeof PopoverPrimitive.Trigger>) {
  return <PopoverPrimitive.Trigger data-slot="popover-trigger" {...props} />;
}

/**
 * PopoverContent component
 *
 * Container for popover content. Appears when the trigger is activated.
 * Includes animations and positioning options.
 *
 * @param props - Content props
 * @param props.className - Additional CSS classes
 * @param props.align - Alignment relative to the trigger
 * @param props.sideOffset - Distance from the trigger element
 * @returns The popover content component
 *
 * @example
 * ```tsx
 * <PopoverContent align="start" sideOffset={8}>
 *   <div>
 *     <h3>Popover Title</h3>
 *     <p>This is some popover content.</p>
 *   </div>
 * </PopoverContent>
 * ```
 */
function PopoverContent({
  className,
  align = 'center',
  sideOffset = 4,
  ...props
}: React.ComponentProps<typeof PopoverPrimitive.Content>) {
  return (
    <PopoverPrimitive.Portal>
      <PopoverPrimitive.Content
        data-slot="popover-content"
        align={align}
        sideOffset={sideOffset}
        className={cn(
          'bg-popover text-popover-foreground data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95 data-[side=bottom]:slide-in-from-top-2 data-[side=left]:slide-in-from-right-2 data-[side=right]:slide-in-from-left-2 data-[side=top]:slide-in-from-bottom-2 z-50 w-72 origin-(--radix-popover-content-transform-origin) rounded-md border p-4 shadow-md outline-hidden',
          className,
        )}
        {...props}
      />
    </PopoverPrimitive.Portal>
  );
}

/**
 * PopoverAnchor component
 *
 * An optional anchor element for the popover positioning.
 * When used, the popover will position relative to this anchor instead of the trigger.
 *
 * @param props - Anchor props
 * @returns The popover anchor component
 *
 * @example
 * ```tsx
 * <PopoverAnchor>
 *   <div>Anchor element</div>
 * </PopoverAnchor>
 * ```
 */
function PopoverAnchor({
  ...props
}: React.ComponentProps<typeof PopoverPrimitive.Anchor>) {
  return <PopoverPrimitive.Anchor data-slot="popover-anchor" {...props} />;
}

export { Popover, PopoverTrigger, PopoverContent, PopoverAnchor };
