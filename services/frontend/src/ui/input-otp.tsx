import * as React from 'react';
import { OTPInput, OTPInputContext } from 'input-otp';
import { MinusIcon } from 'lucide-react';
import { cn } from '@/lib/shared/utils';

/**
 * InputOTP component
 *
 * A customizable OTP (One-Time Password) input component built on top of the input-otp library.
 * Provides a user-friendly interface for entering verification codes.
 *
 * @param props - InputOTP component props
 * @param props.className - Additional CSS classes for the input
 * @param props.containerClassName - Additional CSS classes for the container
 * @returns The rendered OTP input component
 *
 * @example
 * ```tsx
 * <InputOTP maxLength={6} onComplete={(value) => console.log(value)}>
 *   <InputOTPGroup>
 *     <InputOTPSlot index={0} />
 *     <InputOTPSlot index={1} />
 *     <InputOTPSlot index={2} />
 *   </InputOTPGroup>
 *   <InputOTPSeparator />
 *   <InputOTPGroup>
 *     <InputOTPSlot index={3} />
 *     <InputOTPSlot index={4} />
 *     <InputOTPSlot index={5} />
 *   </InputOTPGroup>
 * </InputOTP>
 * ```
 */

function InputOTP({
  className,
  containerClassName,
  ...props
}: React.ComponentProps<typeof OTPInput> & {
  containerClassName?: string;
}) {
  return (
    <OTPInput
      data-slot="input-otp"
      containerClassName={cn(
        'flex items-center gap-2 has-disabled:opacity-50',
        containerClassName,
      )}
      className={cn('disabled:cursor-not-allowed', className)}
      {...props}
    />
  );
}

/**
 * InputOTPGroup component
 *
 * Groups related OTP input slots together. Used to visually separate
 * different sections of an OTP code.
 *
 * @param props - Standard div props
 * @param props.className - Additional CSS classes
 * @returns The rendered OTP group component
 *
 * @example
 * ```tsx
 * <InputOTPGroup>
 *   <InputOTPSlot index={0} />
 *   <InputOTPSlot index={1} />
 *   <InputOTPSlot index={2} />
 * </InputOTPGroup>
 * ```
 */
function InputOTPGroup({ className, ...props }: React.ComponentProps<'div'>) {
  return (
    <div
      data-slot="input-otp-group"
      className={cn('flex items-center', className)}
      {...props}
    />
  );
}

/**
 * InputOTPSlot component
 *
 * Individual slot for a single character in the OTP input.
 * Displays the character and cursor/focus states.
 *
 * @param props - Slot component props
 * @param props.index - The index of this slot in the OTP sequence
 * @param props.className - Additional CSS classes
 * @returns The rendered OTP slot component
 *
 * @example
 * ```tsx
 * <InputOTPSlot index={0} />
 * <InputOTPSlot index={1} />
 * ```
 */
function InputOTPSlot({
  index,
  className,
  ...props
}: React.ComponentProps<'div'> & {
  /** The index of this slot in the OTP sequence */
  index: number;
}) {
  const inputOTPContext = React.useContext(OTPInputContext);
  const { char, hasFakeCaret, isActive } = inputOTPContext?.slots[index] ?? {};

  return (
    <div
      data-slot="input-otp-slot"
      data-active={isActive}
      className={cn(
        'data-[active=true]:border-ring data-[active=true]:ring-ring/50 data-[active=true]:aria-invalid:ring-destructive/20 dark:data-[active=true]:aria-invalid:ring-destructive/40 aria-invalid:border-destructive data-[active=true]:aria-invalid:border-destructive dark:bg-input/30 border-input relative flex h-9 w-9 items-center justify-center border-y border-r text-sm shadow-xs transition-all outline-none first:rounded-l-md first:border-l last:rounded-r-md data-[active=true]:z-10 data-[active=true]:ring-[3px]',
        className,
      )}
      {...props}
    >
      {char}
      {hasFakeCaret && (
        <div className="pointer-events-none absolute inset-0 flex items-center justify-center">
          <div className="animate-caret-blink bg-foreground h-4 w-px duration-1000" />
        </div>
      )}
    </div>
  );
}

/**
 * InputOTPSeparator component
 *
 * Visual separator between groups of OTP input slots.
 * Displays a minus icon to separate different sections.
 *
 * @param props - Standard div props
 * @returns The rendered OTP separator component
 *
 * @example
 * ```tsx
 * <InputOTPGroup>
 *   <InputOTPSlot index={0} />
 *   <InputOTPSlot index={1} />
 * </InputOTPGroup>
 * <InputOTPSeparator />
 * <InputOTPGroup>
 *   <InputOTPSlot index={2} />
 *   <InputOTPSlot index={3} />
 * </InputOTPGroup>
 * ```
 */
function InputOTPSeparator({ ...props }: React.ComponentProps<'div'>) {
  return (
    <div data-slot="input-otp-separator" role="separator" {...props}>
      <MinusIcon />
    </div>
  );
}

export { InputOTP, InputOTPGroup, InputOTPSlot, InputOTPSeparator };
