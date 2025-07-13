'use client';

import { Form, FormField } from '@/ui/form';
import { zodResolver } from '@hookform/resolvers/zod';
import { useForm } from 'react-hook-form';
import { Button } from '@/ui/button';
import { toast } from 'sonner';
import type { QuestionForm, Test } from '@/types/knowledge/models';
import { submitTestAction } from '@/actions/knowledge/actions';
import { redirect } from 'next/navigation';
import { useState, useEffect } from 'react';
import { z } from 'zod';
import { 
  Dialog, 
  DialogContent, 
  DialogDescription, 
  DialogFooter, 
  DialogHeader, 
  DialogTitle, 
  DialogTrigger 
} from '@/ui/dialog';
import { AlertTriangle } from 'lucide-react';

// Simple schema for the test form
const submitTestFormSchema = z.object({
  test_id: z.number(),
}).passthrough(); // Allow additional fields for questions

interface SubmitTestFormProps {
  test: Test;
  questions: QuestionForm[];
}

function SubmitTestForm({ test, questions }: SubmitTestFormProps) {
  const [currentQuestionIndex, setCurrentQuestionIndex] = useState(0);
  const [showConfirmDialog, setShowConfirmDialog] = useState(false);
  const [isSubmitting, setIsSubmitting] = useState(false);
  
  // Create dynamic default values
  const defaultValues = questions.reduce(
    (acc, question) => ({
      ...acc,
      [`question_${question.id}`]: '',
    }),
    { test_id: test.id },
  );

  const form = useForm<any>({
    resolver: zodResolver(submitTestFormSchema),
    defaultValues,
  });

  const onSubmit = async (values: any) => {
    // Validate that all questions are answered
    const unansweredQuestions = questions.filter(
      (question) => !values[`question_${question.id}`] || values[`question_${question.id}`] === ''
    );

    if (unansweredQuestions.length > 0) {
      toast.error(`Please answer all questions. ${unansweredQuestions.length} questions remaining.`);
      return;
    }

    // Show confirmation dialog
    setShowConfirmDialog(true);
  };

  const handleConfirmSubmit = async () => {
    setIsSubmitting(true);
    setShowConfirmDialog(false);

    try {
      const values = form.getValues();
      
      // Transform form values to the expected format
      const answers = questions.map((question) => ({
        question_id: question.id,
        selected_answer: parseInt(values[`question_${question.id}`], 10),
      }));

      const payload = {
        test_id: test.id,
        answers,
      };

      const res = await submitTestAction(payload);
      if (!res.success) {
        form.setError('root', { message: res.error.message });
        console.error(res.error);
        toast.error(res.error.message);
        return;
      }
      
      toast.success('Test submitted successfully!');
      clearProgress(); // Clear saved progress
      redirect('/dashboard/tests');
    } finally {
      setIsSubmitting(false);
    }
  };

  console.log({questions})
  const currentQuestion = questions[currentQuestionIndex];
  console.log('Current Question:', currentQuestion);
  const totalQuestions = questions.length;

  const goToNextQuestion = () => {
    if (currentQuestionIndex < totalQuestions - 1) {
      setCurrentQuestionIndex(currentQuestionIndex + 1);
    }
  };

  const goToPreviousQuestion = () => {
    if (currentQuestionIndex > 0) {
      setCurrentQuestionIndex(currentQuestionIndex - 1);
    }
  };

  const goToQuestion = (index: number) => {
    setCurrentQuestionIndex(index);
  };

  // Calculate answered questions count
  const answeredCount = questions.filter(q => form.watch(`question_${q.id}`) !== '').length;

  // Auto-save progress to localStorage (could be extended to backend)
  useEffect(() => {
    const values = form.getValues();
    const progressKey = `test_progress_${test.id}`;
    
    // Save progress to localStorage
    try {
      localStorage.setItem(progressKey, JSON.stringify({
        values,
        currentQuestionIndex,
        timestamp: Date.now()
      }));
    } catch (error) {
      console.warn('Failed to save progress to localStorage:', error);
    }
  }, [form.watch(), currentQuestionIndex, test.id]);

  // Load saved progress on mount
  useEffect(() => {
    const progressKey = `test_progress_${test.id}`;
    
    try {
      const savedProgress = localStorage.getItem(progressKey);
      if (savedProgress) {
        const { values, currentQuestionIndex: savedIndex, timestamp } = JSON.parse(savedProgress);
        
        // Only restore if saved within last 24 hours
        if (Date.now() - timestamp < 24 * 60 * 60 * 1000) {
          // Restore form values
          Object.entries(values).forEach(([key, value]) => {
            if (key.startsWith('question_')) {
              form.setValue(key, value as string);
            }
          });
          
          // Restore current question index
          setCurrentQuestionIndex(savedIndex || 0);
          
          toast.info('Restored your previous progress');
        } else {
          // Clear old progress
          localStorage.removeItem(progressKey);
        }
      }
    } catch (error) {
      console.warn('Failed to load progress from localStorage:', error);
    }
  }, []);

  // Clear progress on successful submission
  const clearProgress = () => {
    const progressKey = `test_progress_${test.id}`;
    try {
      localStorage.removeItem(progressKey);
    } catch (error) {
      console.warn('Failed to clear progress from localStorage:', error);
    }
  };

  return (
    <Form {...form}>
      <form
        onSubmit={form.handleSubmit(onSubmit)}
        className="space-y-6"
      >
        {/* Progress indicator */}
        <div className="mb-6">
          <div className="flex justify-between items-center mb-2">
            <span className="text-sm font-medium text-gray-600">
              Question {currentQuestionIndex + 1} of {totalQuestions}
            </span>
            <span className="text-sm text-gray-500">
              {Math.round(((currentQuestionIndex + 1) / totalQuestions) * 100)}% Complete
            </span>
          </div>
          <div className="w-full bg-gray-200 rounded-full h-2">
            <div
              className="bg-blue-600 h-2 rounded-full transition-all duration-300"
              style={{
                width: `${((currentQuestionIndex + 1) / totalQuestions) * 100}%`,
              }}
            />
          </div>
        </div>

        {/* Current Question */}
        <div className="bg-white border border-gray-200 rounded-lg p-6">
          <h3 className="text-lg font-semibold mb-4">
            {currentQuestion.question_text}
          </h3>
          
          <FormField
            control={form.control}
            name={`question_${currentQuestion.id}`}
            render={({ field }) => (
              <div className="space-y-3">
                {[
                  { value: '0', label: 'A', text: currentQuestion.option_a },
                  { value: '1', label: 'B', text: currentQuestion.option_b },
                  { value: '2', label: 'C', text: currentQuestion.option_c },
                  { value: '3', label: 'D', text: currentQuestion.option_d },
                ].map((option) => (
                  <label
                    key={option.value}
                    className={`flex items-center p-3 border rounded-lg cursor-pointer transition-colors ${
                      field.value === option.value
                        ? 'border-blue-500 bg-blue-50'
                        : 'border-gray-200 hover:border-gray-300'
                    }`}
                  >
                    <input
                      type="radio"
                      value={option.value}
                      checked={field.value === option.value}
                      onChange={field.onChange}
                      className="sr-only"
                    />
                    <span className="flex items-center">
                      <span className="w-6 h-6 border-2 border-gray-300 rounded-full mr-3 flex items-center justify-center">
                        {field.value === option.value && (
                          <span className="w-3 h-3 bg-blue-500 rounded-full" />
                        )}
                      </span>
                      <span className="font-medium mr-2">{option.label}.</span>
                      <span>{option.text}</span>
                    </span>
                  </label>
                ))}
              </div>
            )}
          />
        </div>

        {/* Navigation and Submit */}
        <div className="flex justify-between items-center">
          <Button
            type="button"
            variant="outline"
            onClick={goToPreviousQuestion}
            disabled={currentQuestionIndex === 0}
          >
            Previous
          </Button>

          <div className="flex gap-2">
            {currentQuestionIndex < totalQuestions - 1 ? (
              <Button
                type="button"
                onClick={goToNextQuestion}
              >
                Next
              </Button>
            ) : (
              <Dialog open={showConfirmDialog} onOpenChange={setShowConfirmDialog}>
                <DialogTrigger asChild>
                  <Button
                    type="button"
                    disabled={isSubmitting}
                    className="bg-green-600 hover:bg-green-700"
                  >
                    {isSubmitting ? 'Submitting...' : 'Submit Test'}
                  </Button>
                </DialogTrigger>
                <DialogContent>
                  <DialogHeader>
                    <DialogTitle className="flex items-center gap-2">
                      <AlertTriangle className="h-5 w-5 text-amber-500" />
                      Confirm Test Submission
                    </DialogTitle>
                    <DialogDescription>
                      Are you sure you want to submit your test? Once submitted, you cannot change your answers.
                      <br />
                      <span className="font-medium mt-2 block">
                        Questions answered: {questions.filter(q => form.watch(`question_${q.id}`) !== '').length} of {questions.length}
                      </span>
                    </DialogDescription>
                  </DialogHeader>
                  <DialogFooter>
                    <Button
                      variant="outline"
                      onClick={() => setShowConfirmDialog(false)}
                      disabled={isSubmitting}
                    >
                      Cancel
                    </Button>
                    <Button
                      onClick={handleConfirmSubmit}
                      disabled={isSubmitting}
                      className="bg-green-600 hover:bg-green-700"
                    >
                      {isSubmitting ? 'Submitting...' : 'Yes, Submit Test'}
                    </Button>
                  </DialogFooter>
                </DialogContent>
              </Dialog>
            )}
          </div>
        </div>

        {/* Question overview */}
        <div className="mt-8 p-4 bg-gray-50 rounded-lg">
          <div className="flex items-center justify-between mb-3">
            <h4 className="font-medium">Question Overview</h4>
            <span className="text-sm text-gray-600">
              {answeredCount} of {totalQuestions} answered
            </span>
          </div>
          <div className="grid grid-cols-5 sm:grid-cols-10 gap-2">
            {questions.map((question, index) => {
              const fieldName = `question_${question.id}`;
              const isAnswered = form.watch(fieldName) !== '';
              const isCurrent = index === currentQuestionIndex;
              
              return (
                <button
                  key={question.id}
                  type="button"
                  onClick={() => goToQuestion(index)}
                  className={`w-8 h-8 rounded text-sm font-medium transition-colors focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 ${
                    isCurrent
                      ? 'bg-blue-600 text-white'
                      : isAnswered
                      ? 'bg-green-100 text-green-800 border border-green-300 hover:bg-green-200'
                      : 'bg-gray-100 text-gray-600 border border-gray-300 hover:bg-gray-200'
                  }`}
                  title={`Question ${index + 1}${isAnswered ? ' (answered)' : ' (not answered)'}`}
                >
                  {index + 1}
                </button>
              );
            })}
          </div>
          <div className="flex items-center justify-between mt-3 text-xs text-gray-500">
            <span>Click on a number to jump to that question</span>
            <div className="flex items-center gap-4">
              <div className="flex items-center gap-1">
                <div className="w-3 h-3 bg-green-100 border border-green-300 rounded"></div>
                <span>Answered</span>
              </div>
              <div className="flex items-center gap-1">
                <div className="w-3 h-3 bg-gray-100 border border-gray-300 rounded"></div>
                <span>Not answered</span>
              </div>
              <div className="flex items-center gap-1">
                <div className="w-3 h-3 bg-blue-600 rounded"></div>
                <span>Current</span>
              </div>
            </div>
          </div>
        </div>
      </form>
    </Form>
  );
}

export { SubmitTestForm };
