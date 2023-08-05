package service

import (
	"context"
	"errors"
	"reflect"
	"route256/notifications/internal/domain/model"
	"route256/notifications/internal/domain/service/mocks"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestService_GetHistory(t *testing.T) {
	t.Parallel()

	t.Run("default", func(t *testing.T) {
		t.Parallel()

		userID := int64(1)

		expected := []*model.Notification{
			{
				UserID:  userID,
				OrderID: 1,
				Status:  "status",
			},
			{
				UserID:  userID,
				OrderID: 2,
				Status:  "status2",
			},
		}

		repoMock := mocks.NewRepository(t)
		repoMock.On("GetLst", mock.Anything, userID).
			Return(expected, nil)

		senderMock := mocks.NewSender(t)

		s := &Service{
			sender: senderMock,
			repo:   repoMock,
		}

		got, err := s.GetHistory(context.Background(), userID)

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("%v got != expected %v", got, expected)
		}

		if err != nil {
			t.Error("ProcessStatusChange() products service unexpected error")
			return
		}
	})

	t.Run("repo error", func(t *testing.T) {
		t.Parallel()

		userID := int64(1)

		repoMock := mocks.NewRepository(t)
		repoMock.On("GetLst", mock.Anything, userID).
			Return(nil, errors.New("some error"))

		senderMock := mocks.NewSender(t)

		s := &Service{
			sender: senderMock,
			repo:   repoMock,
		}

		_, err := s.GetHistory(context.Background(), userID)

		if err == nil {
			t.Error("ProcessStatusChange() products service unhandled error")
			return
		}
	})
}
