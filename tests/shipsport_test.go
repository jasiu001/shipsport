package tests

import (
	"context"
	"testing"

	"shipsport/internal/core/service/shipsport"
	"shipsport/internal/handler"
	"shipsport/internal/repository/shipsportrepo"

	"github.com/stretchr/testify/require"
)

func TestShipsPortFlow(t *testing.T) {
	// given
	repository := shipsportrepo.NewRepository()
	service := shipsport.NewService(repository)
	hdl := handler.NewFileReader(service)

	// when
	err := hdl.Handle(context.TODO(), []string{"./payload/ports.json"})

	// then
	require.NoError(t, err)
	require.Len(t, repository.List(), 3)
}

// TODO add more tests for no happy path

// TODO add functional tests which check what data are persisted in storage for both create/update operation
