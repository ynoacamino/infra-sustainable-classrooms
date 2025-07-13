package controllers

import (
	"context"
	"fmt"

	knowledgedb "github.com/ynoacamino/infra-sustainable-classrooms/services/knowledge/gen/database"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/knowledge/gen/knowledge"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/gen/profiles"
)

// Test management endpoints

func (s *knowledgesvrc) CreateTest(ctx context.Context, payload *knowledge.CreateTestPayload) (res *knowledge.SimpleResponse, err error) {
	profile, err := s.profilesServiceRepo.GetCompleteProfile(ctx, &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil {
		return nil, knowledge.Unauthorized("Failed to retrieve user profile: " + err.Error())
	}

	if profile.Role != "teacher" {
		return nil, knowledge.Unauthorized("Only teachers can create tests")
	}

	err = s.testRepo.CreateTest(ctx, knowledgedb.CreateTestParams{
		Title:     payload.Title,
		CreatedBy: profile.UserID,
	})
	if err != nil {
		return nil, knowledge.InvalidInput("Failed to create test: " + err.Error())
	}

	return &knowledge.SimpleResponse{
		Success: true,
		Message: "Test created successfully",
	}, nil
}

func (s *knowledgesvrc) GetMyTests(ctx context.Context, payload *knowledge.GetMyTestsPayload) (res *knowledge.TestsResponse, err error) {
	profile, err := s.profilesServiceRepo.GetCompleteProfile(ctx, &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil {
		return nil, knowledge.Unauthorized("Failed to retrieve user profile: " + err.Error())
	}

	if profile.Role != "teacher" {
		return nil, knowledge.Unauthorized("Only teachers can view their tests")
	}

	tests, err := s.testRepo.GetMyTests(ctx, profile.UserID)
	if err != nil {
		return nil, knowledge.InvalidInput("Failed to get tests: " + err.Error())
	}

	var response []*knowledge.Test
	for _, test := range tests {
		response = append(response, &knowledge.Test{
			ID:        test.ID,
			Title:     test.Title,
			CreatedBy: test.CreatedBy,
			CreatedAt: test.CreatedAt.Time.Unix(),
		})
	}

	return &knowledge.TestsResponse{
		Tests: response,
	}, nil
}

func (s *knowledgesvrc) UpdateTest(ctx context.Context, payload *knowledge.UpdateTestPayload) (res *knowledge.SimpleResponse, err error) {
	profile, err := s.profilesServiceRepo.GetCompleteProfile(ctx, &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil {
		return nil, knowledge.Unauthorized("Failed to retrieve user profile: " + err.Error())
	}

	if profile.Role != "teacher" {
		return nil, knowledge.Unauthorized("Only teachers can update tests")
	}

	// Check if test exists and belongs to user
	test, err := s.testRepo.GetTestById(ctx, payload.TestID)
	if err != nil {
		return nil, knowledge.InvalidInput("Test not found: " + err.Error())
	}

	if test.CreatedBy != profile.UserID {
		return nil, knowledge.Unauthorized("You can only update your own tests")
	}

	err = s.testRepo.UpdateTest(ctx, knowledgedb.UpdateTestParams{
		ID:    payload.TestID,
		Title: payload.Title,
	})
	if err != nil {
		return nil, knowledge.InvalidInput("Failed to update test: " + err.Error())
	}

	return &knowledge.SimpleResponse{
		Success: true,
		Message: "Test updated successfully",
	}, nil
}

func (s *knowledgesvrc) DeleteTest(ctx context.Context, payload *knowledge.DeleteTestPayload) (res *knowledge.SimpleResponse, err error) {
	profile, err := s.profilesServiceRepo.GetCompleteProfile(ctx, &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil {
		return nil, knowledge.Unauthorized("Failed to retrieve user profile: " + err.Error())
	}

	if profile.Role != "teacher" {
		return nil, knowledge.Unauthorized("Only teachers can delete tests")
	}

	// Check if test exists and belongs to user
	test, err := s.testRepo.GetTestById(ctx, payload.TestID)
	if err != nil {
		return nil, knowledge.InvalidInput("Test not found: " + err.Error())
	}

	if test.CreatedBy != profile.UserID {
		return nil, knowledge.Unauthorized("You can only delete your own tests")
	}

	err = s.testRepo.DeleteTest(ctx, payload.TestID)
	if err != nil {
		return nil, knowledge.InvalidInput("Failed to delete test: " + err.Error())
	}

	return &knowledge.SimpleResponse{
		Success: true,
		Message: "Test deleted successfully",
	}, nil
}

// Question management endpoints

func (s *knowledgesvrc) GetTestQuestions(ctx context.Context, payload *knowledge.GetTestQuestionsPayload) (res *knowledge.QuestionsResponse, err error) {
	profile, err := s.profilesServiceRepo.GetCompleteProfile(ctx, &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil {
		return nil, knowledge.Unauthorized("Failed to retrieve user profile: " + err.Error())
	}

	if profile.Role != "teacher" {
		return nil, knowledge.Unauthorized("Only teachers can view test questions")
	}

	// Check if test exists and belongs to user
	test, err := s.testRepo.GetTestById(ctx, payload.TestID)
	if err != nil {
		return nil, knowledge.InvalidInput("Test not found: " + err.Error())
	}

	if test.CreatedBy != profile.UserID {
		return nil, knowledge.Unauthorized("You can only view questions for your own tests")
	}

	questions, err := s.questionRepo.GetQuestionsByTestId(ctx, payload.TestID)
	if err != nil {
		return nil, knowledge.InvalidInput("Failed to get questions: " + err.Error())
	}

	var response []*knowledge.Question
	for _, question := range questions {
		response = append(response, &knowledge.Question{
			ID:            question.ID,
			TestID:        question.TestID,
			QuestionText:  question.QuestionText,
			OptionA:       question.OptionA,
			OptionB:       question.OptionB,
			OptionC:       question.OptionC,
			OptionD:       question.OptionD,
			CorrectAnswer: int(question.CorrectAnswer),
			QuestionOrder: int(question.QuestionOrder),
		})
	}

	return &knowledge.QuestionsResponse{
		Questions: response,
	}, nil
}

func (s *knowledgesvrc) AddQuestion(ctx context.Context, payload *knowledge.AddQuestionPayload) (res *knowledge.SimpleResponse, err error) {
	profile, err := s.profilesServiceRepo.GetCompleteProfile(ctx, &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil {
		return nil, knowledge.Unauthorized("Failed to retrieve user profile: " + err.Error())
	}

	if profile.Role != "teacher" {
		return nil, knowledge.Unauthorized("Only teachers can add questions")
	}

	// Check if test exists and belongs to user
	test, err := s.testRepo.GetTestById(ctx, payload.TestID)
	if err != nil {
		return nil, knowledge.InvalidInput("Test not found: " + err.Error())
	}

	if test.CreatedBy != profile.UserID {
		return nil, knowledge.Unauthorized("You can only add questions to your own tests")
	}

	// Get current questions count to set order
	questions, err := s.questionRepo.GetQuestionsByTestId(ctx, payload.TestID)
	if err != nil {
		return nil, knowledge.InvalidInput("Failed to get questions count: " + err.Error())
	}

	err = s.questionRepo.CreateQuestion(ctx, knowledgedb.CreateQuestionParams{
		TestID:        payload.TestID,
		QuestionText:  payload.QuestionText,
		OptionA:       payload.OptionA,
		OptionB:       payload.OptionB,
		OptionC:       payload.OptionC,
		OptionD:       payload.OptionD,
		CorrectAnswer: int32(payload.CorrectAnswer),
		QuestionOrder: int32(len(questions) + 1),
	})
	if err != nil {
		return nil, knowledge.InvalidInput("Failed to create question: " + err.Error())
	}

	return &knowledge.SimpleResponse{
		Success: true,
		Message: "Question added successfully",
	}, nil
}

func (s *knowledgesvrc) UpdateQuestion(ctx context.Context, payload *knowledge.UpdateQuestionPayload) (res *knowledge.SimpleResponse, err error) {
	profile, err := s.profilesServiceRepo.GetCompleteProfile(ctx, &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil {
		return nil, knowledge.Unauthorized("Failed to retrieve user profile: " + err.Error())
	}

	if profile.Role != "teacher" {
		return nil, knowledge.Unauthorized("Only teachers can update questions")
	}

	// Check if test exists and belongs to user
	test, err := s.testRepo.GetTestById(ctx, payload.TestID)
	if err != nil {
		return nil, knowledge.InvalidInput("Test not found: " + err.Error())
	}

	if test.CreatedBy != profile.UserID {
		return nil, knowledge.Unauthorized("You can only update questions for your own tests")
	}

	// Check if question exists
	question, err := s.questionRepo.GetQuestionById(ctx, payload.QuestionID)
	if err != nil {
		return nil, knowledge.InvalidInput("Question not found: " + err.Error())
	}

	if question.TestID != payload.TestID {
		return nil, knowledge.InvalidInput("Question does not belong to this test")
	}

	err = s.questionRepo.UpdateQuestion(ctx, knowledgedb.UpdateQuestionParams{
		ID:            payload.QuestionID,
		QuestionText:  payload.QuestionText,
		OptionA:       payload.OptionA,
		OptionB:       payload.OptionB,
		OptionC:       payload.OptionC,
		OptionD:       payload.OptionD,
		CorrectAnswer: int32(payload.CorrectAnswer),
	})
	if err != nil {
		return nil, knowledge.InvalidInput("Failed to update question: " + err.Error())
	}

	return &knowledge.SimpleResponse{
		Success: true,
		Message: "Question updated successfully",
	}, nil
}

func (s *knowledgesvrc) DeleteQuestion(ctx context.Context, payload *knowledge.DeleteQuestionPayload) (res *knowledge.SimpleResponse, err error) {
	profile, err := s.profilesServiceRepo.GetCompleteProfile(ctx, &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil {
		return nil, knowledge.Unauthorized("Failed to retrieve user profile: " + err.Error())
	}

	if profile.Role != "teacher" {
		return nil, knowledge.Unauthorized("Only teachers can delete questions")
	}

	// Check if test exists and belongs to user
	test, err := s.testRepo.GetTestById(ctx, payload.TestID)
	if err != nil {
		return nil, knowledge.InvalidInput("Test not found: " + err.Error())
	}

	if test.CreatedBy != profile.UserID {
		return nil, knowledge.Unauthorized("You can only delete questions from your own tests")
	}

	// Check if question exists
	question, err := s.questionRepo.GetQuestionById(ctx, payload.QuestionID)
	if err != nil {
		return nil, knowledge.InvalidInput("Question not found: " + err.Error())
	}

	if question.TestID != payload.TestID {
		return nil, knowledge.InvalidInput("Question does not belong to this test")
	}

	err = s.questionRepo.DeleteQuestion(ctx, payload.QuestionID)
	if err != nil {
		return nil, knowledge.InvalidInput("Failed to delete question: " + err.Error())
	}

	return &knowledge.SimpleResponse{
		Success: true,
		Message: "Question deleted successfully",
	}, nil
}

// Student endpoints

func (s *knowledgesvrc) GetAvailableTests(ctx context.Context, payload *knowledge.GetAvailableTestsPayload) (res *knowledge.TestsResponse, err error) {
	profile, err := s.profilesServiceRepo.GetCompleteProfile(ctx, &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil {
		return nil, knowledge.Unauthorized("Failed to retrieve user profile: " + err.Error())
	}

	if profile.Role != "student" {
		return nil, knowledge.Unauthorized("Only students can view available tests")
	}

	tests, err := s.testRepo.GetAvailableTests(ctx, profile.UserID)
	if err != nil {
		return nil, knowledge.InvalidInput("Failed to get available tests: " + err.Error())
	}

	var response []*knowledge.Test
	for _, test := range tests {
		questionCount := int(test.QuestionCount)
		response = append(response, &knowledge.Test{
			ID:            test.ID,
			Title:         test.Title,
			CreatedBy:     test.CreatedBy,
			CreatedAt:     test.CreatedAt.Time.Unix(),
			QuestionCount: &questionCount,
		})
	}

	return &knowledge.TestsResponse{
		Tests: response,
	}, nil
}

func (s *knowledgesvrc) GetTestForm(ctx context.Context, payload *knowledge.GetTestFormPayload) (res *knowledge.FormResponse, err error) {
	profile, err := s.profilesServiceRepo.GetCompleteProfile(ctx, &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil {
		return nil, knowledge.Unauthorized("Failed to retrieve user profile: " + err.Error())
	}

	if profile.Role != "student" {
		return nil, knowledge.Unauthorized("Only students can take tests")
	}

	// Check if user has already completed this test
	completed, err := s.submissionRepo.CheckUserCompletedTest(ctx, knowledgedb.CheckUserCompletedTestParams{
		UserID: profile.UserID,
		TestID: payload.TestID,
	})
	if err != nil {
		return nil, knowledge.InvalidInput("Failed to check completion status: " + err.Error())
	}

	if completed {
		return nil, knowledge.InvalidInput("You have already completed this test")
	}

	// Get test
	test, err := s.testRepo.GetTestById(ctx, payload.TestID)
	if err != nil {
		return nil, knowledge.InvalidInput("Test not found: " + err.Error())
	}

	// Get questions (without correct answers)
	questions, err := s.questionRepo.GetQuestionsByTestId(ctx, payload.TestID)
	if err != nil {
		return nil, knowledge.InvalidInput("Failed to get questions: " + err.Error())
	}

	var questionForms []*knowledge.QuestionForm
	for _, question := range questions {
		questionForms = append(questionForms, &knowledge.QuestionForm{
			ID:            question.ID,
			QuestionText:  question.QuestionText,
			OptionA:       question.OptionA,
			OptionB:       question.OptionB,
			OptionC:       question.OptionC,
			OptionD:       question.OptionD,
			QuestionOrder: int(question.QuestionOrder),
		})
	}

	return &knowledge.FormResponse{
		Test: &knowledge.Test{
			ID:        test.ID,
			Title:     test.Title,
			CreatedBy: test.CreatedBy,
			CreatedAt: test.CreatedAt.Time.Unix(),
		},
		Questions: questionForms,
	}, nil
}

func (s *knowledgesvrc) SubmitTest(ctx context.Context, payload *knowledge.SubmitTestPayload) (res *knowledge.SubmitResponse, err error) {
	profile, err := s.profilesServiceRepo.GetCompleteProfile(ctx, &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil {
		return nil, knowledge.Unauthorized("Failed to retrieve user profile: " + err.Error())
	}

	if profile.Role != "student" {
		return nil, knowledge.Unauthorized("Only students can submit tests")
	}

	// Check if user has already completed this test
	completed, err := s.submissionRepo.CheckUserCompletedTest(ctx, knowledgedb.CheckUserCompletedTestParams{
		UserID: profile.UserID,
		TestID: payload.TestID,
	})
	if err != nil {
		return nil, knowledge.InvalidInput("Failed to check completion status: " + err.Error())
	}

	if completed {
		return nil, knowledge.InvalidInput("You have already completed this test")
	}

	// Get all questions for this test
	questions, err := s.questionRepo.GetQuestionsByTestId(ctx, payload.TestID)
	if err != nil {
		return nil, knowledge.InvalidInput("Failed to get questions: " + err.Error())
	}

	// Create a map of question ID to question for easy lookup
	questionMap := make(map[int64]knowledgedb.Question)
	for _, question := range questions {
		questionMap[question.ID] = question
	}

	// Calculate score
	correctCount := 0
	totalQuestions := len(questions)

	// Validate all answers and calculate score
	for _, answer := range payload.Answers {
		question, exists := questionMap[answer.QuestionID]
		if !exists {
			return nil, knowledge.InvalidInput(fmt.Sprintf("Question ID %d not found in this test", answer.QuestionID))
		}

		if answer.SelectedAnswer == int(question.CorrectAnswer) {
			correctCount++
		}
	}

	score := float64(correctCount) / float64(totalQuestions) * 100

	submission, err := s.submissionRepo.CreateSubmission(ctx, knowledgedb.CreateSubmissionParams{
		TestID: payload.TestID,
		UserID: profile.UserID,
		Score:  s.Float64ToPgNumeric(score),
	})
	if err != nil {
		return nil, knowledge.InvalidInput("Failed to create submission: " + err.Error())
	}

	// Create answer submissions
	for _, answer := range payload.Answers {
		question := questionMap[answer.QuestionID]
		isCorrect := answer.SelectedAnswer == int(question.CorrectAnswer)

		err = s.answerRepo.CreateAnswerSubmission(ctx, knowledgedb.CreateAnswerSubmissionParams{
			SubmissionID:   submission.ID,
			QuestionID:     answer.QuestionID,
			SelectedAnswer: int32(answer.SelectedAnswer),
			IsCorrect:      isCorrect,
		})
		if err != nil {
			return nil, knowledge.InvalidInput("Failed to create answer submission: " + err.Error())
		}
	}

	return &knowledge.SubmitResponse{
		Success:      true,
		Message:      "Test submitted successfully",
		SubmissionID: submission.ID,
		Score:        score,
	}, nil
}

func (s *knowledgesvrc) GetMySubmissions(ctx context.Context, payload *knowledge.GetMySubmissionsPayload) (res *knowledge.SubmissionsResponse, err error) {
	profile, err := s.profilesServiceRepo.GetCompleteProfile(ctx, &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil {
		return nil, knowledge.Unauthorized("Failed to retrieve user profile: " + err.Error())
	}

	if profile.Role != "student" {
		return nil, knowledge.Unauthorized("Only students can view their submissions")
	}

	submissions, err := s.submissionRepo.GetUserSubmissions(ctx, profile.UserID)
	if err != nil {
		return nil, knowledge.InvalidInput("Failed to get submissions: " + err.Error())
	}

	var response []*knowledge.Submission
	for _, submission := range submissions {
		// Convert pgtype.Numeric to float64 for score
		score := 0.0
		if submission.Score.Valid {
			scoreFloat, err := submission.Score.Float64Value()
			if err == nil {
				score = scoreFloat.Float64
			}
		}

		response = append(response, &knowledge.Submission{
			ID:          submission.ID,
			TestID:      submission.TestID,
			TestTitle:   submission.TestTitle,
			Score:       score,
			SubmittedAt: submission.SubmittedAt.Time.Unix(),
		})
	}

	return &knowledge.SubmissionsResponse{
		Submissions: response,
	}, nil
}

func (s *knowledgesvrc) GetSubmissionResult(ctx context.Context, payload *knowledge.GetSubmissionResultPayload) (res *knowledge.SubmissionResult, err error) {
	profile, err := s.profilesServiceRepo.GetCompleteProfile(ctx, &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil {
		return nil, knowledge.Unauthorized("Failed to retrieve user profile: " + err.Error())
	}

	if profile.Role != "student" {
		return nil, knowledge.Unauthorized("Only students can view their submission results")
	}

	// Get submission
	submission, err := s.submissionRepo.GetSubmissionById(ctx, payload.SubmissionID)
	if err != nil {
		return nil, knowledge.InvalidInput("Submission not found: " + err.Error())
	}

	// Check if submission belongs to user
	if submission.UserID != profile.UserID {
		return nil, knowledge.Unauthorized("You can only view your own submissions")
	}

	// Get test info
	test, err := s.testRepo.GetTestById(ctx, submission.TestID)
	if err != nil {
		return nil, knowledge.InvalidInput("Test not found: " + err.Error())
	}

	// Get answers with questions
	answers, err := s.answerRepo.GetAnswersBySubmission(ctx, payload.SubmissionID)
	if err != nil {
		return nil, knowledge.InvalidInput("Failed to get answers: " + err.Error())
	}

	// Convert answers to question results
	var questionResults []*knowledge.QuestionResult
	for _, answer := range answers {
		questionResults = append(questionResults, &knowledge.QuestionResult{
			Question: &knowledge.Question{
				ID:            answer.QuestionID,
				TestID:        submission.TestID,
				QuestionText:  answer.QuestionText,
				OptionA:       answer.OptionA,
				OptionB:       answer.OptionB,
				OptionC:       answer.OptionC,
				OptionD:       answer.OptionD,
				CorrectAnswer: int(answer.CorrectAnswer),
			},
			SelectedAnswer: int(answer.SelectedAnswer),
			IsCorrect:      answer.IsCorrect,
		})
	}

	// Convert score
	score := 0.0
	if submission.Score.Valid {
		scoreFloat, err := submission.Score.Float64Value()
		if err == nil {
			score = scoreFloat.Float64
		}
	}

	return &knowledge.SubmissionResult{
		Submission: &knowledge.Submission{
			ID:          submission.ID,
			TestID:      submission.TestID,
			TestTitle:   test.Title,
			Score:       score,
			SubmittedAt: submission.SubmittedAt.Time.Unix(),
		},
		Questions: questionResults,
	}, nil
}
