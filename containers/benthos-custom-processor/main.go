package main

import (
	"context"
	"errors"

	_ "github.com/benthosdev/benthos/v4/public/components/aws"
	service "github.com/benthosdev/benthos/v4/public/service"
	"github.com/leometzger/kinesis-benthos-processing/internal/processor"
)

func main() {
	spec := service.
		NewConfigSpec().
		Field(service.NewIntField("seconds").Description("The number of seconds to wait between consider valid."))

	service.RegisterBatchProcessor(
		"customplug",
		spec,
		func(conf *service.ParsedConfig, mgr *service.Resources) (service.BatchProcessor, error) {
			seconds, err := conf.FieldInt("seconds")
			if err != nil {
				return nil, err
			}

			if seconds < 1 {
				return nil, errors.New("less than 1 second is an invalid configuration")
			}

			return processor.NewProcessor(seconds), nil
		},
	)

	service.RunCLI(context.Background())
}
