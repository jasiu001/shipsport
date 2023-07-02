package shipsport

import (
	"context"
	"strings"
	"testing"
	"time"

	mocks "shipsport/internal/core/automock"
	"shipsport/internal/core/domain"
)

func TestService_Execute(t *testing.T) {
	// given
	repositoryMock := mocks.NewShipsPortRepository(t)
	repositoryMock.On("UpdateOrCreate", domain.ShipsPort{
		Id:      "sid",
		Name:    "a",
		City:    "b",
		Country: "c",
	}).Return(nil)

	service := NewService(repositoryMock)
	done := make(chan struct{})

	ctx, cancel := context.WithTimeout(context.TODO(), 1*time.Second)
	defer cancel()

	// when
	go service.Execute(context.TODO(), done, strings.NewReader(`{"sid":{"name":"a", "city":"b", "country":"c"}}`))

	// then
	for {
		select {
		case <-done:
			return
		case <-ctx.Done():
			t.Log("timeout for test")
			t.Fail()
			return
		}
	}
}

// TODO add more tests for no happy path
