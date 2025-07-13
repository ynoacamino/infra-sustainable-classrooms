import { AddQuestionPayloadSchema } from '@/types/knowledge/schemas/payload';
import type { Field } from '@/types/shared/field';
import type z from 'zod';

export const addQuestionFormSchema = AddQuestionPayloadSchema;

export const addQuestionFormFields: Field<
  keyof z.infer<typeof addQuestionFormSchema>
>[] = [
  {
    name: 'question_text',
    label: 'Question Text',
    type: 'textarea',
    placeholder:
      'e.g. What is the correct syntax for creating a variable in JavaScript?',
    description: 'Enter the question text (5-500 characters).',
  },
  {
    name: 'option_a',
    label: 'Option A',
    type: 'text',
    placeholder: 'e.g. var x = 5;',
    description: 'Enter option A (1-200 characters).',
  },
  {
    name: 'option_b',
    label: 'Option B',
    type: 'text',
    placeholder: 'e.g. let x = 5;',
    description: 'Enter option B (1-200 characters).',
  },
  {
    name: 'option_c',
    label: 'Option C',
    type: 'text',
    placeholder: 'e.g. const x = 5;',
    description: 'Enter option C (1-200 characters).',
  },
  {
    name: 'option_d',
    label: 'Option D',
    type: 'text',
    placeholder: 'e.g. All of the above',
    description: 'Enter option D (1-200 characters).',
  },
  {
    name: 'correct_answer',
    label: 'Correct Answer',
    type: 'select',
    placeholder: 'Select the correct answer',
    description: 'Select the correct answer option.',
    options: [
      { key: 'option-a', value: '0', textValue: 'Option A' },
      { key: 'option-b', value: '1', textValue: 'Option B' },
      { key: 'option-c', value: '2', textValue: 'Option C' },
      { key: 'option-d', value: '3', textValue: 'Option D' },
    ],
  },
];
