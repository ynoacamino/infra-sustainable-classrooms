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

interface TwoDividedInputOptProps {
  maxLength: number;
}

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
