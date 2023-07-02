package port

import (
	"context"
	"io"

	"shipsport/internal/core/domain"
)

type ShipsPortService interface {
	Execute(ctx context.Context, done chan struct{}, reader io.Reader)
}

//go:generate mockery --name=ShipsPortRepository --output=../automock
type ShipsPortRepository interface {
	UpdateOrCreate(shipPort domain.ShipsPort) error
}
