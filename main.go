package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/NikitaKissa/victus-thermal-guard/temperatures"
	"github.com/NikitaKissa/victus-thermal-guard/victus"
	victusdbus "github.com/NikitaKissa/victus-thermal-guard/victus/dbus"
)

func main() {
	log.SetFlags(0)

	if err := run(); err != nil {
		log.Print(err)
		os.Exit(1)
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
		log.Fatal(err)
	}

	controller := victus.NewController(backend)

	ticker := time.NewTicker(250 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
		}

		cpu, gpu, err := getTemperatures()
		if err != nil {
			log.Print(err)
			continue
		}

		maxT := math.Max(cpu, gpu)

		if maxT > 76 {
			if err := controller.SetFansMax(); err != nil {
				err = fmt.Errorf("unable to set fans `max`: %w", err)
				log.Print(err)
			}
		}

		if maxT < 70 {
			if err := controller.SetFansAuto(); err != nil {
				err = fmt.Errorf("unable to set fans `auto`: %w", err)
				log.Print(err)
			}
		}
	}
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
