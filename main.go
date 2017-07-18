package main

import (
	"context"
	"fmt"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

const (
	Nanosecond  time.Duration = 1
	Microsecond               = 1000 * Nanosecond
	Millisecond               = 1000 * Microsecond
	Second                    = 1000 * Millisecond
	Minute                    = 60 * Second
	Hour                      = 60 * Minute
)

func main() {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	go func() {
		for _ = range time.Tick(2 * Second) {
			containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
			if err != nil {
				panic(err)
			}

			for _, container := range containers {
				fmt.Printf("%s %s\n", container.ID[:10], container.Image)
			}
		}
	}()
	select {}
}
