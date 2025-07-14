import {
  AnswerSchema,
  QuestionSchema,
  SubmissionSchema,
  TestSchema,
} from '@/types/knowledge/schemas/models';
import z from 'zod';

// === TEACHER PAYLOAD SCHEMAS ===

// Test payload schemas
export const CreateTestPayloadSchema = TestSchema.omit({
  id: true,
  created_at: true,
  created_by: true,
  question_count: true,
});

export const GetTestPayloadSchema = TestSchema.pick({
  id: true,
});

export const UpdateTestPayloadSchema = TestSchema.pick({
  id: true,
  title: true,
}).partial({
  title: true,
});

export const DeleteTestPayloadSchema = TestSchema.pick({
  id: true,
});

export const GetTestQuestionsPayloadSchema = TestSchema.pick({
  id: true,
});

// Question payload schemas
export const AddQuestionPayloadSchema = QuestionSchema.pick({
  test_id: true,
  question_text: true,
  option_a: true,
  option_b: true,
  option_c: true,
  option_d: true,
  correct_answer: true,
});

export const GetQuestionPayloadSchema = QuestionSchema.pick({
  id: true,
  test_id: true,
});

export const UpdateQuestionPayloadSchema = QuestionSchema.pick({
  id: true,
  test_id: true,
  option_a: true,
  option_b: true,
  option_c: true,
  option_d: true,
  question_text: true,
  correct_answer: true,
}).partial({
  option_a: true,
  option_b: true,
  option_c: true,
  option_d: true,
  question_text: true,
  correct_answer: true,
});

export const DeleteQuestionPayloadSchema = QuestionSchema.pick({
  id: true,
  test_id: true,
});

export const GetTestFormPayloadSchema = TestSchema.pick({
  id: true,
});

export const SubmitTestPayloadSchema = TestSchema.pick({
  id: true,
}).extend({
  answers: z.array(AnswerSchema).describe('List of answers for the test'),
});

export const GetSubmissionPayloadSchema = SubmissionSchema.pick({
  id: true,
});

export const GetSubmissionResultPayloadSchema = SubmissionSchema.pick({
  id: true,
});
