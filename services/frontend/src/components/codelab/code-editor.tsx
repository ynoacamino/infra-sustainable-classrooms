'use client';

import { useState } from 'react';
import { Editor } from '@monaco-editor/react';
import { Button } from '@/ui/button';
import { Card, CardContent, CardHeader, CardTitle } from '@/ui/card';
import { Badge } from '@/ui/badge';
import { Play, Save, RotateCcw, CheckCircle, XCircle, Clock } from 'lucide-react';
import { toast } from 'sonner';
import { createAttemptAction } from '@/actions/codelab/actions';
import type { Test } from '@/types/codelab/models';

interface CodeEditorProps {
  exerciseId: number;
  initialCode: string;
  testCases?: Test[];
}

interface TestResult {
  passed: boolean;
  input?: string;
  expected?: string;
  actual?: string;
  error?: string;
}

export function CodeEditor({ exerciseId, initialCode, testCases = [] }: CodeEditorProps) {
  const [code, setCode] = useState(initialCode);
  const [isRunning, setIsRunning] = useState(false);
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [results, setResults] = useState<TestResult[]>([]);
  const [lastSubmissionTime, setLastSubmissionTime] = useState<Date | null>(null);

  const handleEditorChange = (value: string | undefined) => {
    setCode(value || '');
  };

  const handleSubmit = async () => {
    if (!code.trim()) {
      toast.error('Please write some code first');
      return;
    }

    setIsSubmitting(true);

    try {
      const passedTestsCount = results.filter(r => r.passed).length;
      const isSuccessful = passedTestsCount === results.length;
      
      const result = await createAttemptAction({
        exercise_id: exerciseId,
        code: code,
        success: isSuccessful,
      });

      if (result?.success) {
        setLastSubmissionTime(new Date());
        if (isSuccessful) {
          toast.success('🎉 Perfect! All tests passed and solution submitted!');
        } else {
          toast.success('Solution submitted, but some tests failed. Keep trying!');
        }
      } else {
        toast.error('Failed to submit solution');
      }
    } catch (error) {
      console.error('Error submitting code:', error);
      toast.error('Failed to submit solution');
    } finally {
      setIsSubmitting(false);
    }
  };

  const handleReset = () => {
    if (confirm('Are you sure you want to reset your code? This will restore the initial template.')) {
      setCode(initialCode);
      setResults([]);
      toast.info('Code reset to initial template');
    }
  };

  const passedTests = results.filter(r => r.passed).length;
  const totalTests = results.length;

  return (
    <div className="space-y-4">
      {/* Code Editor */}
      <div className="border rounded-lg overflow-hidden">
        <div className="bg-muted px-4 py-2 border-b">
          <div className="flex items-center justify-between">
            <span className="text-sm font-medium">Solution Code</span>
            <div className="flex gap-2">
              <Button
                variant="outline"
                size="sm"
                onClick={handleReset}
                disabled={isRunning || isSubmitting}
              >
                <RotateCcw className="w-4 h-4 mr-2" />
                Reset
              </Button>
              <Button
                size="sm"
                onClick={handleSubmit}
                disabled={isRunning || isSubmitting}
              >
                {isSubmitting ? (
                  <Clock className="w-4 h-4 mr-2 animate-spin" />
                ) : (
                  <Save className="w-4 h-4 mr-2" />
                )}
                {isSubmitting ? 'Submitting...' : 'Submit'}
              </Button>
            </div>
          </div>
        </div>
        
        <Editor
          height="400px"
          defaultLanguage="javascript"
          value={code}
          onChange={handleEditorChange}
          theme="vs-dark"
          options={{
            minimap: { enabled: false },
            scrollBeyondLastLine: false,
            fontSize: 14,
            lineNumbers: 'on',
            roundedSelection: false,
            automaticLayout: true,
            wordWrap: 'on',
          }}
        />
      </div>

      {/* Test Results */}
      {results.length > 0 && (
        <Card>
          <CardHeader>
            <CardTitle className="flex items-center justify-between">
              <span className="flex items-center gap-2">
                Test Results
                <Badge variant={passedTests === totalTests ? 'default' : 'secondary'}>
                  {passedTests}/{totalTests} passed
                </Badge>
              </span>
            </CardTitle>
          </CardHeader>
          <CardContent>
            <div className="space-y-3">
              {results.map((result, index) => (
                <div
                  key={index}
                  className={`border rounded-lg p-3 ${
                    result.passed ? 'border-green-200 bg-green-50' : 'border-red-200 bg-red-50'
                  }`}
                >
                  <div className="flex items-center gap-2 mb-2">
                    {result.passed ? (
                      <CheckCircle className="w-4 h-4 text-green-600" />
                    ) : (
                      <XCircle className="w-4 h-4 text-red-600" />
                    )}
                    <span className="font-medium text-sm">
                      Test Case {index + 1} {result.passed ? 'Passed' : 'Failed'}
                    </span>
                  </div>
                  
                  {!result.passed && (
                    <div className="grid grid-cols-1 md:grid-cols-3 gap-3 text-xs">
                      {result.input && (
                        <div>
                          <div className="font-medium text-muted-foreground mb-1">Input:</div>
                          <pre className="bg-white p-2 rounded border">{result.input}</pre>
                        </div>
                      )}
                      {result.expected && (
                        <div>
                          <div className="font-medium text-muted-foreground mb-1">Expected:</div>
                          <pre className="bg-white p-2 rounded border">{result.expected}</pre>
                        </div>
                      )}
                      {result.actual && (
                        <div>
                          <div className="font-medium text-muted-foreground mb-1">Your Output:</div>
                          <pre className="bg-white p-2 rounded border text-red-600">{result.actual}</pre>
                        </div>
                      )}
                    </div>
                  )}
                  
                  {result.error && (
                    <div className="mt-2">
                      <div className="font-medium text-muted-foreground mb-1 text-xs">Error:</div>
                      <pre className="bg-white p-2 rounded border text-red-600 text-xs">{result.error}</pre>
                    </div>
                  )}
                </div>
              ))}
            </div>
          </CardContent>
        </Card>
      )}

      {/* Submission Status */}
      {lastSubmissionTime && (
        <Card>
          <CardContent className="py-3">
            <div className="flex items-center gap-2 text-sm text-muted-foreground">
              <CheckCircle className="w-4 h-4 text-green-600" />
              Last submitted: {lastSubmissionTime.toLocaleString()}
            </div>
          </CardContent>
        </Card>
      )}
    </div>
  );
}
