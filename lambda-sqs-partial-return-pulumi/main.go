package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		function, queue, err := NewSqsLambda(ctx)
		if err != nil {
			return err
		}

		ctx.Export("function", function.ID())
		ctx.Export("queue", queue.Url)
		return nil
	})
}
