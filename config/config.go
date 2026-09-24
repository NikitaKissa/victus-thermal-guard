package config

import (
	"fmt"
	"os"

	configparser "github.com/NikitaKissa/victus-thermal-guard/config/parser"
)

type Config struct {
	ActivateTemperature float64
	Hysteresis          float64
	MeasurementInterval int
}

const (
	defaultActivateTemperature float64 = 80
	defaultHysteresis          float64 = 10
	defaultMeasurementInterval int     = 250
)

func NewConfig(path string) (Config, error) {
	buf, err := os.ReadFile(path)
	if err != nil {
		return Config{
			ActivateTemperature: defaultActivateTemperature,
			Hysteresis:          defaultHysteresis,
			MeasurementInterval: defaultMeasurementInterval,
		}, fmt.Errorf("unable to open config file: %w", err)
	}

	configMap, err := configparser.Parse(buf)
	if err != nil {
		return Config{
			ActivateTemperature: defaultActivateTemperature,
			Hysteresis:          defaultHysteresis,
			MeasurementInterval: defaultMeasurementInterval,
		}, fmt.Errorf("unable to parse config file: %w", err)
	}

	config := configMapToConfig(configMap)
	return config, nil
}
