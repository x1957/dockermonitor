package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var (
	period = flag.Duration("period", 2*time.Second, "")
)

func main() {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	go func() {
		for _ = range time.Tick(*period) {
			containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
			if err != nil {
				log.Print(err)
				continue
			}

			for _, container := range containers {
				fmt.Printf("%s %s\n", container.ID[:10], container.Image)
			}
		}
	}()
	select {}
}
