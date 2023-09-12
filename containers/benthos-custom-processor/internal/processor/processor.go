package processor

import (
	"context"
	"time"

	"github.com/benthosdev/benthos/v4/public/service"
)

type Data struct {
	lastTime time.Time
}

type CustomProcessor struct {
	seconds int
	data    map[string]*Data
}

func NewProcessor(seconds int) *CustomProcessor {
	return &CustomProcessor{
		seconds: seconds,
		data:    map[string]*Data{},
	}
}

func (p *CustomProcessor) ProcessBatch(_ context.Context, batch service.MessageBatch) ([]service.MessageBatch, error) {
	// data := &Data{
	// 	lastTime: time.Now(),
	// }
	return []service.MessageBatch{batch}, nil
}

func (p *CustomProcessor) Close(ctx context.Context) error {
	return nil
}
