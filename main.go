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
	period  = flag.Duration("period", 2*time.Second, "")
	timeout = flag.Duration("timeout", 1*time.Second, "")
)

func main() {
	flag.Parse()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	go func() {
		for _ = range time.Tick(*period) {
			if err := ping(ctx, cli); err != nil {
				log.Printf("Fail: %v", err)
			} else {
				log.Print("OK")
			}
		}
	}()
	select {}
}

func ping(ctx context.Context, cli *client.Client) error {
	ctx, cancel := context.WithTimeout(ctx, *timeout)
	defer cancel()
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		return err
	}
	for _, container := range containers {
		fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	}
	return nil
}
