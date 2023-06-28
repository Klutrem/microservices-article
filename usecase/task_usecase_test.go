package usecase_test

import (
	"context"
	"errors"
	"testing"

	"main/domain"
	"main/domain/mocks"
	"main/lib"
	"main/usecase"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFetchByUserID(t *testing.T) {
	mockTaskRepository := new(mocks.TaskRepository)
	userObjectID := 2
	userID := string(rune(userObjectID))

	t.Run("success", func(t *testing.T) {

		mockTask := domain.Task{
			ID:     1,
			Title:  "Test Title",
			UserID: 2,
		}

		mockListTask := make([]domain.Task, 0)
		mockListTask = append(mockListTask, mockTask)

		mockTaskRepository.On("FetchByUserID", mock.Anything, userID).Return(mockListTask, nil).Once()

		u := usecase.NewTaskUsecase(mockTaskRepository, lib.Env{})

		list, err := u.FetchByUserID(context.Background(), userID)

		assert.NoError(t, err)
		assert.NotNil(t, list)
		assert.Len(t, list, len(mockListTask))

		mockTaskRepository.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockTaskRepository.On("FetchByUserID", mock.Anything, userID).Return(nil, errors.New("Unexpected")).Once()

		u := usecase.NewTaskUsecase(mockTaskRepository, lib.Env{})

		list, err := u.FetchByUserID(context.Background(), userID)

		assert.Error(t, err)
		assert.Nil(t, list)

		mockTaskRepository.AssertExpectations(t)
	})

}
