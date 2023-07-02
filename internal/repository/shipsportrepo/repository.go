package shipsportrepo

import (
	"log"
	"sync"

	"shipsport/internal/core/domain"
)

type Repository struct {
	storage sync.Map
}

func NewRepository() *Repository {
	return &Repository{}
}

// UpdateOrCreate add new ships port entity to memory storage or replace if ships port entity with id already exist
func (r *Repository) UpdateOrCreate(shipsPort domain.ShipsPort) error {
	log.Printf("entity %q with id %s added to storage", shipsPort.Name, shipsPort.Id)
	r.storage.Store(shipsPort.Id, shipsPort)

	return nil
}

// List returns slice with all ships port entities in the storage
func (r *Repository) List() []domain.ShipsPort {
	list := make([]domain.ShipsPort, 0)
	r.storage.Range(func(_, value any) bool {
		sp, ok := value.(domain.ShipsPort)
		if !ok {
			return true
		}
		list = append(list, sp)
		return true
	})

	return list
}
