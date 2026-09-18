package main

import (
	"fmt"

	"github.com/NikitaKissa/victus-thermal-guard.git/temperatures"
)

func main() {
	if err := temperatures.FindTelemetry(); err != nil {
		err = fmt.Errorf("error during searching for telemetry: %w", err)
		panic(err)
	}

	fmt.Println(temperatures.GetCPUTemperature())
	fmt.Println(temperatures.GetGPUTemperature())
}
