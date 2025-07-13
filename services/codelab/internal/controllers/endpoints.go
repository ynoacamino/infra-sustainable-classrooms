package controllers

import (
	"context"

	"github.com/ynoacamino/infra-sustainable-classrooms/services/codelab/gen/codelab"
	codelabdb "github.com/ynoacamino/infra-sustainable-classrooms/services/codelab/gen/database"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/gen/profiles"

	"github.com/dop251/goja"
)

func (s *codelabsvrc) CreateExercise(ctx context.Context, payload *codelab.CreateExercisePayload) (res *codelab.SimpleResponse, err error) {
	profile, err := s.profilesServiceRepo.GetCompleteProfile(context.Background(), &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil || profile == nil {
		return nil, codelab.Unauthorized("Unauthorized access")
	}

	if profile.Role != "teacher" {
		return nil, codelab.PermissionDenied("Only teachers can create exercises")
	}

	err = s.exercisesRepo.CreateExercise(ctx, codelabdb.CreateExerciseParams{
		Title:       payload.Title,
		Description: payload.Description,
		InitialCode: payload.InitialCode,
		Solution:    payload.Solution,
		Difficulty:  payload.Difficulty,
		CreatedBy:   profile.UserID,
	})
	if err != nil {
		return nil, codelab.InternalError("Failed to create exercise: " + err.Error())
	}

	return &codelab.SimpleResponse{
		Message: "Exercise created successfully",
		Success: true,
	}, nil
}

func (s *codelabsvrc) GetExercise(ctx context.Context, payload *codelab.GetExercisePayload) (res *codelab.Exercise, err error) {
	profile, err := s.profilesServiceRepo.GetCompleteProfile(context.Background(), &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil || profile == nil {
		return nil, codelab.Unauthorized("Unauthorized access")
	}

	if profile.Role != "teacher" {
		return nil, codelab.PermissionDenied("Only teachers can view exercise with solution")
	}

	exercise, err := s.exercisesRepo.GetExerciseById(ctx, payload.ID)
	if err != nil {
		return nil, codelab.NotFound("Exercise not found")
	}

	return &codelab.Exercise{
		ID:          exercise.ID,
		Title:       exercise.Title,
		Description: exercise.Description,
		InitialCode: exercise.InitialCode,
		Solution:    exercise.Solution,
		Difficulty:  exercise.Difficulty,
		CreatedBy:   exercise.CreatedBy,
		CreatedAt:   exercise.CreatedAt.Time.UnixMilli(),
		UpdatedAt:   exercise.UpdatedAt.Time.UnixMilli(),
	}, nil
}

func (s *codelabsvrc) ListExercises(ctx context.Context, payload *codelab.ListExercisesPayload) (res []*codelab.Exercise, err error) {
	profile, err := s.profilesServiceRepo.GetCompleteProfile(context.Background(), &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil || profile == nil {
		return nil, codelab.Unauthorized("Unauthorized access")
	}

	if profile.Role != "teacher" {
		return nil, codelab.PermissionDenied("Only teachers can view exercises with solutions")
	}

	exercises, err := s.exercisesRepo.ListExercises(ctx)
	if err != nil {
		return nil, codelab.InternalError("Failed to list exercises: " + err.Error())
	}

	result := make([]*codelab.Exercise, len(exercises))
	for i, exercise := range exercises {
		result[i] = &codelab.Exercise{
			ID:          exercise.ID,
			Title:       exercise.Title,
			Description: exercise.Description,
			InitialCode: exercise.InitialCode,
			Solution:    exercise.Solution,
			Difficulty:  exercise.Difficulty,
			CreatedBy:   exercise.CreatedBy,
			CreatedAt:   exercise.CreatedAt.Time.UnixMilli(),
			UpdatedAt:   exercise.UpdatedAt.Time.UnixMilli(),
		}
	}

	return result, nil
}

func (s *codelabsvrc) UpdateExercise(ctx context.Context, payload *codelab.UpdateExercisePayload2) (res *codelab.SimpleResponse, err error) {
	profile, err := s.profilesServiceRepo.GetCompleteProfile(context.Background(), &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil || profile == nil {
		return nil, codelab.Unauthorized("Unauthorized access")
	}

	if profile.Role != "teacher" {
		return nil, codelab.PermissionDenied("Only teachers can update exercises")
	}

	// Verify the exercise exists
	_, err = s.exercisesRepo.GetExerciseById(ctx, payload.ID)
	if err != nil {
		return nil, codelab.NotFound("Exercise not found")
	}

	err = s.exercisesRepo.UpdateExercise(ctx, codelabdb.UpdateExerciseParams{
		ID:          payload.ID,
		Title:       payload.Exercise.Title,
		Description: payload.Exercise.Description,
		InitialCode: payload.Exercise.InitialCode,
		Solution:    payload.Exercise.Solution,
		Difficulty:  payload.Exercise.Difficulty,
	})
	if err != nil {
		return nil, codelab.InternalError("Failed to update exercise: " + err.Error())
	}

	return &codelab.SimpleResponse{
		Message: "Exercise updated successfully",
		Success: true,
	}, nil
}

func (s *codelabsvrc) DeleteExercise(ctx context.Context, payload *codelab.DeleteExercisePayload) (res *codelab.SimpleResponse, err error) {
	profile, err := s.profilesServiceRepo.GetCompleteProfile(context.Background(), &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil || profile == nil {
		return nil, codelab.Unauthorized("Unauthorized access")
	}

	if profile.Role != "teacher" {
		return nil, codelab.PermissionDenied("Only teachers can delete exercises")
	}

	// Verify the exercise exists
	_, err = s.exercisesRepo.GetExerciseById(ctx, payload.ID)
	if err != nil {
		return nil, codelab.NotFound("Exercise not found")
	}

	err = s.exercisesRepo.DeleteExercise(ctx, payload.ID)
	if err != nil {
		return nil, codelab.InternalError("Failed to delete exercise: " + err.Error())
	}

	return &codelab.SimpleResponse{
		Message: "Exercise deleted successfully",
		Success: true,
	}, nil
}

func (s *codelabsvrc) CreateTest(ctx context.Context, payload *codelab.CreateTestPayload) (res *codelab.SimpleResponse, err error) {
	profile, err := s.profilesServiceRepo.GetCompleteProfile(context.Background(), &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil || profile == nil {
		return nil, codelab.Unauthorized("Unauthorized access")
	}

	if profile.Role != "teacher" {
		return nil, codelab.PermissionDenied("Only teachers can create tests")
	}

	// Verify the exercise exists
	_, err = s.exercisesRepo.GetExerciseById(ctx, payload.ExerciseID)
	if err != nil {
		return nil, codelab.NotFound("Exercise not found")
	}

	err = s.testsRepo.CreateTest(ctx, codelabdb.CreateTestParams{
		Input:      payload.Input,
		Output:     payload.Output,
		Public:     payload.Public,
		ExerciseID: payload.ExerciseID,
	})
	if err != nil {
		return nil, codelab.InternalError("Failed to create test: " + err.Error())
	}

	return &codelab.SimpleResponse{
		Message: "Test created successfully",
		Success: true,
	}, nil
}

func (s *codelabsvrc) GetTestsByExercise(ctx context.Context, payload *codelab.GetTestsByExercisePayload) (res []*codelab.Test, err error) {
	profile, err := s.profilesServiceRepo.GetCompleteProfile(context.Background(), &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil || profile == nil {
		return nil, codelab.Unauthorized("Unauthorized access")
	}

	if profile.Role != "teacher" {
		return nil, codelab.PermissionDenied("Only teachers can view all tests")
	}

	// Verify the exercise exists
	_, err = s.exercisesRepo.GetExerciseById(ctx, payload.ExerciseID)
	if err != nil {
		return nil, codelab.NotFound("Exercise not found")
	}

	tests, err := s.testsRepo.GetTestsByExercise(ctx, payload.ExerciseID)
	if err != nil {
		return nil, codelab.InternalError("Failed to get tests: " + err.Error())
	}

	result := make([]*codelab.Test, len(tests))
	for i, test := range tests {
		result[i] = &codelab.Test{
			ID:         test.ID,
			Input:      test.Input,
			Output:     test.Output,
			Public:     test.Public,
			ExerciseID: test.ExerciseID,
			CreatedAt:  test.CreatedAt.Time.UnixMilli(),
			UpdatedAt:  test.UpdatedAt.Time.UnixMilli(),
		}
	}

	return result, nil
}

func (s *codelabsvrc) UpdateTest(ctx context.Context, payload *codelab.UpdateTestPayload2) (res *codelab.SimpleResponse, err error) {
	profile, err := s.profilesServiceRepo.GetCompleteProfile(context.Background(), &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil || profile == nil {
		return nil, codelab.Unauthorized("Unauthorized access")
	}

	if profile.Role != "teacher" {
		return nil, codelab.PermissionDenied("Only teachers can update tests")
	}

	err = s.testsRepo.UpdateTest(ctx, codelabdb.UpdateTestParams{
		ID:     payload.ID,
		Input:  payload.Test.Input,
		Output: payload.Test.Output,
		Public: payload.Test.Public,
	})
	if err != nil {
		return nil, codelab.InternalError("Failed to update test: " + err.Error())
	}

	return &codelab.SimpleResponse{
		Message: "Test updated successfully",
		Success: true,
	}, nil
}

func (s *codelabsvrc) DeleteTest(ctx context.Context, payload *codelab.DeleteTestPayload) (res *codelab.SimpleResponse, err error) {
	profile, err := s.profilesServiceRepo.GetCompleteProfile(context.Background(), &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil || profile == nil {
		return nil, codelab.Unauthorized("Unauthorized access")
	}

	if profile.Role != "teacher" {
		return nil, codelab.PermissionDenied("Only teachers can delete tests")
	}

	err = s.testsRepo.DeleteTest(ctx, payload.ID)
	if err != nil {
		return nil, codelab.InternalError("Failed to delete test: " + err.Error())
	}

	return &codelab.SimpleResponse{
		Message: "Test deleted successfully",
		Success: true,
	}, nil
}

func (s *codelabsvrc) GetExerciseForStudent(ctx context.Context, payload *codelab.GetExerciseForStudentPayload) (res *codelab.ExerciseForStudents, err error) {
	profile, err := s.profilesServiceRepo.GetCompleteProfile(context.Background(), &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil || profile == nil {
		return nil, codelab.Unauthorized("Unauthorized access")
	}

	exercise, err := s.exercisesRepo.GetExerciseToResolveById(ctx, payload.ID)
	if err != nil {
		return nil, codelab.NotFound("Exercise not found")
	}

	return &codelab.ExerciseForStudents{
		ID:          exercise.ID,
		Title:       exercise.Title,
		Description: exercise.Description,
		InitialCode: exercise.InitialCode,
		Difficulty:  exercise.Difficulty,
		CreatedBy:   exercise.CreatedBy,
		CreatedAt:   exercise.CreatedAt.Time.UnixMilli(),
		UpdatedAt:   exercise.UpdatedAt.Time.UnixMilli(),
	}, nil
}

func (s *codelabsvrc) ListExercisesForStudents(ctx context.Context, payload *codelab.ListExercisesForStudentsPayload) (res []*codelab.ExerciseForStudentsListView, err error) {
	profile, err := s.profilesServiceRepo.GetCompleteProfile(context.Background(), &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil || profile == nil {
		return nil, codelab.Unauthorized("Unauthorized access")
	}

	exercises, err := s.exercisesRepo.ListExercisesToResolve(ctx)
	if err != nil {
		return nil, codelab.InternalError("Failed to list exercises: " + err.Error())
	}

	result := make([]*codelab.ExerciseForStudentsListView, len(exercises))
	for i, exercise := range exercises {
		result[i] = &codelab.ExerciseForStudentsListView{
			ID:          exercise.ID,
			Title:       exercise.Title,
			Description: exercise.Description,
			Difficulty:  exercise.Difficulty,
			CreatedBy:   exercise.CreatedBy,
			CreatedAt:   exercise.CreatedAt.Time.UnixMilli(),
			UpdatedAt:   exercise.UpdatedAt.Time.UnixMilli(),
		}
	}

	return result, nil
}

func (s *codelabsvrc) CreateAttempt(ctx context.Context, payload *codelab.CreateAttemptPayload) (res *codelab.SimpleResponse, err error) {
	profile, err := s.profilesServiceRepo.GetCompleteProfile(context.Background(), &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil || profile == nil {
		return nil, codelab.Unauthorized("Unauthorized access")
	}

	if profile.Role != "student" {
		return nil, codelab.PermissionDenied("Only students can create attempts")
	}

	exercise, err := s.exercisesRepo.GetExerciseToResolveById(ctx, payload.ExerciseID)
	if err != nil {
		return nil, codelab.NotFound("Exercise not found")
	}

	// Ensure answer exists for this user/exercise combination
	_, err = s.answersRepo.CheckIfAnswerExists(ctx, codelabdb.CheckIfAnswerExistsParams{
		ExerciseID: exercise.ID,
		UserID:     profile.UserID,
	})
	if err != nil {
		// Create answer if it doesn't exist
		err = s.answersRepo.CreateAnswer(ctx, codelabdb.CreateAnswerParams{
			ExerciseID: exercise.ID,
			UserID:     profile.UserID,
			Completed:  false,
		})
		if err != nil {
			return nil, codelab.InternalError("Failed to create answer: " + err.Error())
		}
	}

	// Get the answer
	answer, err := s.answersRepo.GetAnswerByUserAndExercise(ctx, codelabdb.GetAnswerByUserAndExerciseParams{
		ExerciseID: exercise.ID,
		UserID:     profile.UserID,
	})
	if err != nil {
		return nil, codelab.InternalError("Failed to get answer: " + err.Error())
	}

	// Get hidden tests for validation
	privateTests, err := s.testsRepo.GetHiddenTestsByExercise(ctx, exercise.ID)
	if err != nil {
		return nil, codelab.InternalError("Failed to get private tests: " + err.Error())
	}

	// Execute and validate code
	vm := goja.New()
	_, err = vm.RunString(payload.Code)
	if err != nil {
		return nil, codelab.InvalidInput("Code execution failed: " + err.Error())
	}

	solution, ok := goja.AssertFunction(vm.Get("solution"))
	if !ok {
		return nil, codelab.InvalidInput("Code must define a 'solution' function")
	}

	// Test the code against all hidden tests
	for _, test := range privateTests {
		result, err := solution(goja.Undefined(), vm.ToValue(test.Input))
		if err != nil {
			errorLog := err.Error()

			err = s.attemptsRepo.CreateAttempt(ctx, codelabdb.CreateAttemptParams{
				AnswerID: answer.ID,
				Code:     payload.Code,
				Success:  false,
			})
			if err != nil {
				return nil, codelab.InternalError("Failed to create attempt: " + err.Error())
			}

			return &codelab.SimpleResponse{
				Message: "Test failed: " + errorLog,
				Success: false,
			}, nil
		}

		success := result.Equals(vm.ToValue(test.Output))
		if !success {
			err = s.attemptsRepo.CreateAttempt(ctx, codelabdb.CreateAttemptParams{
				AnswerID: answer.ID,
				Code:     payload.Code,
				Success:  false,
			})
			if err != nil {
				return nil, codelab.InternalError("Failed to create attempt: " + err.Error())
			}

			return &codelab.SimpleResponse{
				Message: "Test failed: expected " + test.Output + ", got " + result.String(),
				Success: false,
			}, nil
		}
	}

	err = s.attemptsRepo.CreateAttempt(ctx, codelabdb.CreateAttemptParams{
		AnswerID: answer.ID,
		Code:     payload.Code,
		Success:  true,
	})
	if err != nil {
		return nil, codelab.InternalError("Failed to create attempt: " + err.Error())
	}

	err = s.answersRepo.UpdateAnswerCompleted(ctx, codelabdb.UpdateAnswerCompletedParams{
		ExerciseID: answer.ExerciseID,
		UserID:     profile.UserID,
		Completed:  true,
	})
	if err != nil {
		return nil, codelab.InternalError("Failed to update answer completion: " + err.Error())
	}

	return &codelab.SimpleResponse{
		Message: "All tests passed successfully",
		Success: true,
	}, nil
}

func (s *codelabsvrc) GetAttemptsByUserAndExercise(ctx context.Context, payload *codelab.GetAttemptsByUserAndExercisePayload) (res []*codelab.Attempt, err error) {
	profile, err := s.profilesServiceRepo.GetCompleteProfile(context.Background(), &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil || profile == nil {
		return nil, codelab.Unauthorized("Unauthorized access")
	}

	if profile.Role != "student" {
		return nil, codelab.PermissionDenied("Only students can access this endpoint")
	}

	// Students can only see their own attempts
	if profile.UserID != payload.UserID {
		return nil, codelab.PermissionDenied("Students can only view their own attempts")
	}

	attempts, err := s.attemptsRepo.GetAttemptsByUserAndExercise(ctx, codelabdb.GetAttemptsByUserAndExerciseParams{
		UserID:     payload.UserID,
		ExerciseID: payload.ExerciseID,
	})
	if err != nil {
		return nil, codelab.InternalError("Failed to get attempts: " + err.Error())
	}

	result := make([]*codelab.Attempt, len(attempts))
	for i, attempt := range attempts {
		result[i] = &codelab.Attempt{
			ID:        attempt.ID,
			AnswerID:  attempt.AnswerID,
			Code:      attempt.Code,
			Success:   attempt.Success,
			CreatedAt: attempt.CreatedAt.Time.UnixMilli(),
		}
	}

	return result, nil
}

func (s *codelabsvrc) GetAnswerByUserAndExercise(ctx context.Context, payload *codelab.GetAnswerByUserAndExercisePayload) (res *codelab.Answer, err error) {
	profile, err := s.profilesServiceRepo.GetCompleteProfile(context.Background(), &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil || profile == nil {
		return nil, codelab.Unauthorized("Unauthorized access")
	}

	if profile.Role != "student" {
		return nil, codelab.PermissionDenied("Only students can access this endpoint")
	}

	// Students can only see their own answers
	if profile.UserID != payload.UserID {
		return nil, codelab.PermissionDenied("Students can only view their own answers")
	}

	answer, err := s.answersRepo.GetAnswerByUserAndExercise(ctx, codelabdb.GetAnswerByUserAndExerciseParams{
		ExerciseID: payload.ExerciseID,
		UserID:     payload.UserID,
	})
	if err != nil {
		return nil, codelab.NotFound("Answer not found")
	}

	return &codelab.Answer{
		ID:         answer.ID,
		ExerciseID: answer.ExerciseID,
		UserID:     answer.UserID,
		Completed:  answer.Completed,
		CreatedAt:  answer.CreatedAt.Time.UnixMilli(),
		UpdatedAt:  answer.UpdatedAt.Time.UnixMilli(),
	}, nil
}
