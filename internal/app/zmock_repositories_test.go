// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package app_test

import (
	"context"
	"github.com/bruli/rasberryTelegram/internal/app"
	"github.com/bruli/rasberryTelegram/internal/domain/status"
	"sync"
)

// Ensure, that StatusRepositoryMock does implement app.StatusRepository.
// If this is not the case, regenerate this file with moq.
var _ app.StatusRepository = &StatusRepositoryMock{}

// StatusRepositoryMock is a mock implementation of app.StatusRepository.
//
//	func TestSomethingThatUsesStatusRepository(t *testing.T) {
//
//		// make and configure a mocked app.StatusRepository
//		mockedStatusRepository := &StatusRepositoryMock{
//			FindStatusFunc: func(ctx context.Context) (status.Status, error) {
//				panic("mock out the FindStatus method")
//			},
//		}
//
//		// use mockedStatusRepository in code that requires app.StatusRepository
//		// and then make assertions.
//
//	}
type StatusRepositoryMock struct {
	// FindStatusFunc mocks the FindStatus method.
	FindStatusFunc func(ctx context.Context) (status.Status, error)

	// calls tracks calls to the methods.
	calls struct {
		// FindStatus holds details about calls to the FindStatus method.
		FindStatus []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
	}
	lockFindStatus sync.RWMutex
}

// FindStatus calls FindStatusFunc.
func (mock *StatusRepositoryMock) FindStatus(ctx context.Context) (status.Status, error) {
	if mock.FindStatusFunc == nil {
		panic("StatusRepositoryMock.FindStatusFunc: method is nil but StatusRepository.FindStatus was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockFindStatus.Lock()
	mock.calls.FindStatus = append(mock.calls.FindStatus, callInfo)
	mock.lockFindStatus.Unlock()
	return mock.FindStatusFunc(ctx)
}

// FindStatusCalls gets all the calls that were made to FindStatus.
// Check the length with:
//
//	len(mockedStatusRepository.FindStatusCalls())
func (mock *StatusRepositoryMock) FindStatusCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockFindStatus.RLock()
	calls = mock.calls.FindStatus
	mock.lockFindStatus.RUnlock()
	return calls
}