import type { Field } from '@/types/shared/field';
import {
  InputOTP,
  InputOTPGroup,
  InputOTPSeparator,
  InputOTPSlot,
} from '@/ui/input-otp';
import { useMemo } from 'react';
import type {
  ControllerRenderProps,
  FieldPath,
  FieldValues,
} from 'react-hook-form';

/**
 * Props for the TwoDividedInputOpt component
 */
interface TwoDividedInputOptProps {
  /** Maximum length of the OTP input */
  maxLength: number;
}

/**
 * TwoDividedInputOpt component
 *
 * A specialized OTP input component that divides the input slots into two groups
 * with a visual separator in between. Automatically calculates the distribution
 * of slots based on the maximum length.
 *
 * @template FieldName - The field name type
 * @template TFieldValues - Form field values type
 * @template TName - Field path type
 * @param props - Combined props from TwoDividedInputOptProps, Field, and ControllerRenderProps
 * @param props.maxLength - Maximum length of the OTP input
 * @returns The rendered two-divided OTP input component
 *
 * @example
 * ```tsx
 * // 6-digit OTP divided into two groups of 3
 * <TwoDividedInputOpt
 *   maxLength={6}
 *   name="verificationCode"
 *   {...fieldProps}
 * />
 *
 * // 8-digit OTP divided into groups of 4
 * <TwoDividedInputOpt
 *   maxLength={8}
 *   name="securityCode"
 *   {...fieldProps}
 * />
 * ```
 */
function TwoDividedInputOpt<
  FieldName extends string,
  TFieldValues extends FieldValues = FieldValues,
  TName extends FieldPath<TFieldValues> = FieldPath<TFieldValues>,
>({
  maxLength,
  ...props
}: TwoDividedInputOptProps &
  Omit<Field<FieldName>, 'label' | 'description'> &
  ControllerRenderProps<TFieldValues, TName>) {
  const leftSlots = useMemo(() => {
    return Array.from({ length: Math.floor(maxLength / 2) }, (_, i) => ({
      id: crypto.randomUUID(),
      index: i,
    }));
  }, [maxLength]);

  const rightSlots = useMemo(() => {
    return Array.from({ length: Math.ceil(maxLength / 2) }, (_, i) => ({
      id: crypto.randomUUID(),
      index: i + Math.floor(maxLength / 2),
    }));
  }, [maxLength]);
  return (
    <InputOTP maxLength={maxLength} {...props}>
      <InputOTPGroup>
        {leftSlots.map(({ id }, i) => (
          <InputOTPSlot key={`input-otp-${id}-${i}`} index={i} />
        ))}
      </InputOTPGroup>
      <InputOTPSeparator />
      <InputOTPGroup>
        {rightSlots.map(({ id }, i) => (
          <InputOTPSlot
            key={`input-otp-${id}-${i + Math.floor(maxLength / 2)}`}
            index={i + Math.floor(maxLength / 2)}
          />
        ))}
      </InputOTPGroup>
    </InputOTP>
  );
}

export { TwoDividedInputOpt };
