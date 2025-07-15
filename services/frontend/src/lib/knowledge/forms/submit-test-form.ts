import type { QuestionForm } from '@/types/knowledge/models';
import { QuestionFormSchema } from '@/types/knowledge/schemas/models';
import type { Field } from '@/types/shared/field';
import z from 'zod';

// Note: Este formulario es especial porque las preguntas son dinámicas
// Los campos se generarán dinámicamente basándose en las preguntas del test
// Cada pregunta tendrá un campo de radio button group con las opciones A, B, C, D

export const createSubmitTestFormSchema = (questions: QuestionForm[]) =>
  z.object({
    id: QuestionFormSchema.shape.id,
    answers: z.object(
      Object.fromEntries(
        questions.map((question) => [
          `question_${question.id}`,
          z.string().min(1, {
            message: `Please select an answer for question ${question.id}`,
          }),
        ]),
      ),
    ),
  });

export const createSubmitTestFormFields = (
  questions: QuestionForm[],
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
