import type { SupportedFields } from '@/lib/shared/enums/field';
import type { SelectOption } from '@/types/shared/select-option';

export type SupportFieldType =
  (typeof SupportedFields)[keyof typeof SupportedFields];

type BaseField<T extends string> = {
  name: T;
  label: string;
  description?: string;
  placeholder?: string;
};

export type Field<T extends string> = BaseField<T> &
  (
    | {
        type:
          | typeof SupportedFields.EMAIL
          | typeof SupportedFields.PASSWORD
          | typeof SupportedFields.TEXTAREA
          | typeof SupportedFields.HTML_EDITOR
          | typeof SupportedFields.TEXT
          | typeof SupportedFields.NUMBER
          | typeof SupportedFields.FILE;
      }
    | {
        type: typeof SupportedFields.SELECT;
        options: SelectOption[];
      }
    | {
        type: typeof SupportedFields.OTP;
        maxLength: number;
        pattern: string;
      }
  );
