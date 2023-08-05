package service

import (
	"errors"
	"fmt"
	"route256/notifications/internal/domain/model"
	"route256/notifications/internal/domain/service/mocks"
	"testing"

	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
)

func TestService_ProcessStatusChange(t *testing.T) {
	t.Parallel()

	t.Run("default", func(t *testing.T) {
		t.Parallel()

		order := &model.Notification{
			UserID:  1,
			OrderID: 1,
			Status:  "status",
		}

		notification := fmt.Sprintf("order id=%d, status changed to %s", order.OrderID, order.Status)

		repoMock := mocks.NewRepository(t)
		repoMock.On("Add", mock.Anything, order.UserID, order.OrderID, order.Status).
			Return(nil)

		senderMock := mocks.NewSender(t)
		senderMock.On("Send", mock.Anything, notification).Return(nil)

		s := &Service{
			sender: senderMock,
			repo:   repoMock,
		}

		err := s.ProcessStatusChange(context.Background(), order)

		if err != nil {
			t.Error("ProcessStatusChange() products service unexpected error")
			return
		}
	})

	t.Run("error", func(t *testing.T) {
		t.Parallel()

		order := &model.Notification{
			OrderID: 1,
			Status:  "status",
		}

		notification := fmt.Sprintf("order id=%d, status changed to %s", order.OrderID, order.Status)

		repoMock := mocks.NewRepository(t)
		repoMock.On("Add", mock.Anything, order.UserID, order.OrderID, order.Status).
			Return(nil)

		senderMock := mocks.NewSender(t)
		senderMock.On("Send", mock.Anything, notification).Return(errors.New("some error"))

		s := &Service{
			sender: senderMock,
			repo:   repoMock,
		}

		err := s.ProcessStatusChange(context.Background(), order)

		if err == nil {
			t.Error("ProcessStatusChange() products service unhandled error")
			return
		}
	})

	t.Run("repo failed", func(t *testing.T) {
		t.Parallel()

		order := &model.Notification{
			UserID:  1,
			OrderID: 1,
			Status:  "status",
		}

		repoMock := mocks.NewRepository(t)
		repoMock.On("Add", mock.Anything, order.UserID, order.OrderID, order.Status).
			Return(errors.New("some error"))

		senderMock := mocks.NewSender(t)

		s := &Service{
			sender: senderMock,
			repo:   repoMock,
		}

		err := s.ProcessStatusChange(context.Background(), order)

		if err == nil {
			t.Error("ProcessStatusChange() products service unhandled error")
			return
		}
	})
}
