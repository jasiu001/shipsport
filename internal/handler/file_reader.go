package handler

import (
	"context"
	"fmt"
	"log"
	"os"

	"shipsport/internal/core/port"
)

type FileReader struct {
	service port.ShipsPortService
}

func NewFileReader(service port.ShipsPortService) *FileReader {
	return &FileReader{service: service}
}

// Handle validates incoming parameters, looks for path to file, checks file, opens and pass to service
func (fr *FileReader) Handle(ctx context.Context, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("wrong number of parameters: %d, should be 1", len(args))
	}
	if args[0] == "" {
		return fmt.Errorf("file path is empty")
	}
	_, err := os.Stat(args[0])
	if os.IsNotExist(err) {
		return fmt.Errorf("file from path %q does not exist", args[0])
	}

	file, err := os.Open(args[0])
	if err != nil {
		return fmt.Errorf("cannot open file: %w", err)
	}
	defer file.Close()

	done := make(chan struct{})
	go fr.service.Execute(ctx, done, file)
	for {
		select {
		case <-ctx.Done():
			log.Println("[stop] cancel file reader handler")
			return nil
		case <-done:
			log.Println("file reading completed")
			return nil
		}
	}
}
