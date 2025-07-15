import { UpdateTestPayloadSchema } from '@/types/codelab/schemas/payload';
import type { Field } from '@/types/shared/field';
import type z from 'zod';

// Extract just the test fields for the form
export const updateTestFormSchema = UpdateTestPayloadSchema.shape.test;

export const updateTestFormFields: Field<
  keyof z.infer<typeof updateTestFormSchema>
>[] = [
  {
    name: 'input',
    label: 'Test Input',
    type: 'text',
    placeholder: 'e.g. 5, 3',
    description: 'Enter the input for this test case.',
  },
  {
    name: 'output',
    label: 'Expected Output',
    type: 'text',
    placeholder: 'e.g. 8',
    description: 'Enter the expected output for this test case.',
  },
  {
    name: 'public',
    label: 'Public Test',
    type: 'select',
    options: [
      { key: 'true', value: 'true', textValue: 'Yes (visible to students)' },
      { key: 'false', value: 'false', textValue: 'No (hidden from students)' },
    ],
    description: 'Choose whether this test case is visible to students.',
  },
];
