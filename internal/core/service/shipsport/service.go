package shipsport

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"

	"shipsport/internal/core/domain"
	"shipsport/internal/core/port"
)

var ErrorEOF = errors.New("end of reader")

type Service struct {
	stream     chan Entry
	repository port.ShipsPortRepository
}

func NewService(repository port.ShipsPortRepository) *Service {
	return &Service{
		stream:     make(chan Entry),
		repository: repository,
	}
}

type Entry struct {
	Error    error
	ShipPort domain.ShipsPort
}

// Execute reads data from reader, the read data is stored in the repository
func (s *Service) Execute(ctx context.Context, done chan struct{}, reader io.Reader) {
	defer func() {
		done <- struct{}{}
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Println("[stop] cancel ships port service executor")
				return
			case data := <-s.stream:
				switch {
				case errors.Is(data.Error, ErrorEOF):
					log.Println("end of reader")
					return
				case data.Error != nil:
					log.Printf("error from stream: %s", data.Error)
					return
				}

				if err := s.repository.UpdateOrCreate(data.ShipPort); err != nil {
					log.Printf("error from repository: %s", err)
					return
				}
			}
		}
	}()
	s.read(reader)
}

func (s *Service) read(reader io.Reader) {
	defer close(s.stream)

	decoder := json.NewDecoder(reader)

	if _, err := decoder.Token(); err != nil {
		s.stream <- Entry{Error: fmt.Errorf("wrong opening delimiter: %w", err)}
		return
	}

	for decoder.More() {
		var key json.Token
		key, err := decoder.Token()
		if err != nil {
			s.stream <- Entry{Error: fmt.Errorf("wrong opening delimiter of entity: %w", err)}
			return
		}

		var harbor domain.ShipsPort
		if err = decoder.Decode(&harbor); err != nil {
			s.stream <- Entry{Error: fmt.Errorf("cannot decode entity: %w", err)}
			return
		}

		id, ok := key.(string)
		if !ok {
			s.stream <- Entry{Error: fmt.Errorf("unsupported key type: %w", err)}
			return
		}
		harbor.Id = id
		s.stream <- Entry{ShipPort: harbor}
	}

	s.stream <- Entry{Error: fmt.Errorf("%w", ErrorEOF)}
}
