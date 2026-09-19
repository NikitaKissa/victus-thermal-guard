package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/NikitaKissa/victus-thermal-guard.git/temperatures"
	"golang.org/x/sync/errgroup"
)

func main() {
	ctx := context.Background()
	log.SetFlags(0)

	if err := temperatures.FindTelemetry(); err != nil {
		err = fmt.Errorf("error during searching for telemetry: %w", err)
		panic(err)
	}

	ticker := time.NewTicker(250 * time.Millisecond)

	for range ticker.C {
		cpu, gpu, err := asyncGetTemperatures(ctx)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("CPU: %v;  GPU: %v", cpu, gpu)
	}
}

func asyncGetTemperatures(ctx context.Context) (cpu int, gpu int, err error) {
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		cpu, err = temperatures.GetCPUTemperature()
		return err
	})

	g.Go(func() error {
		gpu, err = temperatures.GetGPUTemperature()
		return err
	})

	if err := g.Wait(); err != nil {
		return 0, 0, err
	}

	return cpu, gpu, nil
}
