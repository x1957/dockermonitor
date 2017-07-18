package main

import (
	"context"
	"log"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/x1957/dockermonitor/prome"
	"github.com/x1957/dockermonitor/agent"
)

const (
	Nanosecond  time.Duration = 1
	Microsecond               = 1000 * Nanosecond
	Millisecond               = 1000 * Microsecond
	Second                    = 1000 * Millisecond
	Minute                    = 60 * Second
	Hour                      = 60 * Minute
)

func record(name string, value int64, agents []agent.Agent) {
	for _, ag := range agents {
		ag.RecordGauge(name, value)
	}
}

func recordDuration(name string, value time.Duration, agents []agent.Agent) {
	for _, ag := range agents {
		ag.RecordDuration(name, value)
	}
}

func main() {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), Minute)
	var start time.Time
	var ps = time.Duration(0)
	info := time.Duration(0)
	containerSize := 0
	errs := 0
	var agents []agent.Agent
	promeAgent := prome.NewPrometheus()
	promeAgent.Run()
	agents = append(agents, promeAgent)
	// add agents
	go func() {
		for _ = range time.Tick(2 * Second) {
			start = time.Now()
			containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
			if err != nil {
				errs ++
				// record error
				record("docker_error_cnt", int64(errs), agents)
				continue
			}
			ps = time.Since(start)
			// record ps latency
			recordDuration("docker_ps_latency", ps, agents)
			start = time.Now()
			// record container size
			containerSize = len(containers)
			record("docker_containter_size", int64(containerSize), agents)
			_, err = cli.Info(ctx)
			if err != nil {
				errs ++
				// record error
				record("docker_error_cnt", int64(errs), agents)
				continue
			}
			// record info latency
			info = time.Since(start)
			recordDuration("docker_info_lantency", info, agents)
			// reset erros
			errs = 0
			record("docker_error_cnt", int64(errs), agents)
			log.Printf("OK")
		}
	}()
	select {}
}
