import type { SimpleResponse } from '@/services/shared/response';
import type {
  Question,
  QuestionForm,
  Submission,
  Test,
} from '@/types/knowledge/models';

export type GetTestFormResponse = {
  test: Test;
  questions: QuestionForm[];
};

export type SubmitTestResponse = SimpleResponse & {
  submission_id: number;
  score: number;
};

export type GetSubmissionResultResponse = {
  submission: Submission;
  questions: {
    question: Question;
    selected_answer: number;
    is_correct: boolean;
  }[];
};
