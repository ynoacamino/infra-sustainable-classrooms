import type {
  TestSchema,
  QuestionSchema,
  QuestionFormSchema,
  AnswerSchema,
  SubmissionSchema,
} from '@/types/knowledge/schemas/models';
import type z from 'zod';

export type Test = z.infer<typeof TestSchema>;

export type Question = z.infer<typeof QuestionSchema>;

export type QuestionForm = z.infer<typeof QuestionFormSchema>;

export type Answer = z.infer<typeof AnswerSchema>;

export type Submission = z.infer<typeof SubmissionSchema>;
