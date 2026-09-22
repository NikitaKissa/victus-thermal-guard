package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/NikitaKissa/victus-thermal-guard/temperatures"
	"github.com/NikitaKissa/victus-thermal-guard/victus"
	victusdbus "github.com/NikitaKissa/victus-thermal-guard/victus/dbus"
)

func main() {
	runtime.GOMAXPROCS(2)

	log.SetFlags(0)

	if err := run(); err != nil {
		log.Print(err)
		os.Exit(1)
	}
}

const shutdownTimeout = 3 * time.Second

func restoreFans(controller victus.Controller) {
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := controller.SetFansAuto(ctx); err != nil {
		log.Printf("restore fans to auto: %v", err)
	}
}

func run() error {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	if err := temperatures.FindTelemetry(); err != nil {
		return fmt.Errorf("find telemetry: %w", err)
	}

	backend, err := victusdbus.New()
	if err != nil {
		return err
	}
	defer backend.Close()

	controller := victus.NewController(backend)

	ticker := time.NewTicker(250 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			restoreFans(controller)
			return nil
		case <-ticker.C:
		}

		cpu, gpu, err := getTemperatures()
		if err != nil {
			log.Print(err)
			continue
		}

		maxT := math.Max(cpu, gpu)

		if err := setFans(ctx, controller, maxT); err != nil {
			log.Print(err)
		}
	}
}

func setFans(
	ctx context.Context,
	controller victus.Controller,
	temperature float64,
) error {
	if temperature >= 77 {
		return controller.SetFansMax(ctx)
	}

	if temperature < 68 {
		return controller.SetFansAuto(ctx)
	}

	return nil
}

func getTemperatures() (cpu float64, gpu float64, err error) {
	gpu, err = temperatures.GetGPUTemperature()
	if err != nil {
		return 0, 0, err
	}

	cpu, err = temperatures.GetCPUTemperature()
	if err != nil {
		return 0, 0, err
	}

	return cpu, gpu, nil
}
