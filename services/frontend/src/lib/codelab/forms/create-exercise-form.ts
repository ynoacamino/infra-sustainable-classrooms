import { CreateExercisePayloadSchema } from '@/types/codelab/schemas/payload';
import type { Field } from '@/types/shared/field';
import type z from 'zod';

export const createExerciseFormSchema = CreateExercisePayloadSchema;

export const createExerciseFormFields: Field<
  keyof z.infer<typeof createExerciseFormSchema>
>[] = [
  {
    name: 'title',
    label: 'Exercise Title',
    type: 'text',
    placeholder: 'e.g. Sum Two Numbers',
    description: 'Enter the exercise title (1-200 characters).',
  },
  {
    name: 'description',
    label: 'Exercise Description',
    type: 'textarea',
    placeholder: 'e.g. Write a function that returns the sum of two numbers...',
    description: 'Enter a detailed description of the exercise.',
  },
  {
    name: 'initial_code',
    label: 'Initial Code Template',
    type: 'textarea',
    placeholder:
      'def sum_two_numbers(a, b):\n    # Write your code here\n    pass',
    description:
      'Provide the initial code template that students will start with.',
  },
  {
    name: 'solution',
    label: 'Exercise Solution',
    type: 'textarea',
    placeholder: 'def sum_two_numbers(a, b):\n    return a + b',
    description: 'Provide the correct solution for the exercise.',
  },
  {
    name: 'difficulty',
    label: 'Difficulty Level',
    type: 'select',
    options: [
      { key: 'easy', value: 'easy', textValue: 'Easy' },
      { key: 'medium', value: 'medium', textValue: 'Medium' },
      { key: 'hard', value: 'hard', textValue: 'Hard' },
    ],
    description: 'Select the difficulty level of the exercise.',
  },
];
