import { SubmitTestPayloadSchema } from '@/types/knowledge/schemas/payload';
import type { Field } from '@/types/shared/field';
import type z from 'zod';

export const submitTestFormSchema = SubmitTestPayloadSchema;

// Note: Este formulario es especial porque las preguntas son dinámicas
// Los campos se generarán dinámicamente basándose en las preguntas del test
// Cada pregunta tendrá un campo de radio button group con las opciones A, B, C, D

export const createSubmitTestFormFields = (
  questions: Array<{
    id: number;
    question_text: string;
    option_a: string;
    option_b: string;
    option_c: string;
    option_d: string;
  }>,
): Field<string>[] => {
  return questions.map((question, index) => ({
    name: `question_${question.id}`,
    label: `Question ${index + 1}`,
    type: 'select' as const,
    placeholder: 'Select your answer',
    description: question.question_text,
    options: [
      {
        key: `q${question.id}-a`,
        value: '0',
        textValue: `A) ${question.option_a}`,
      },
      {
        key: `q${question.id}-b`,
        value: '1',
        textValue: `B) ${question.option_b}`,
      },
      {
        key: `q${question.id}-c`,
        value: '2',
        textValue: `C) ${question.option_c}`,
      },
      {
        key: `q${question.id}-d`,
        value: '3',
        textValue: `D) ${question.option_d}`,
      },
    ],
  }));
};
