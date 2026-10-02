package config

import (
	"fmt"
	"os"
	"time"

	"github.com/NikitaKissa/victus-thermal-guard/config/parser"
)

type Config struct {
	ActivateTemperature float64
	Hysteresis          float64
	MeasurementInterval time.Duration
}

func defaultConfig() Config {
	return Config{
		ActivateTemperature: 80,
		Hysteresis:          10,
		MeasurementInterval: 250 * time.Millisecond,
	}
}

func Load(path string) (Config, error) {
	buf, err := os.ReadFile(path)
	if err != nil {
		return defaultConfig(), fmt.Errorf("unable to open config file: %w", err)
	}

	configMap, err := parser.Parse(buf)
	if err != nil {
		return defaultConfig(), fmt.Errorf("unable to parse config file: %w", err)
	}

	config, err := configMapToConfig(configMap)
	if err != nil {
		return config, err
	}

	if err := validateAndRepairConfig(&config); err != nil {
		return config, err
	}

	return config, nil
}

func StringifyConfig(cfg Config) string {
	return fmt.Sprintf(
		"ActivateTemperature=%v; Hysteresis=%v; MeasurementInterval=%v",
		cfg.ActivateTemperature,
		cfg.Hysteresis,
		cfg.MeasurementInterval,
	)
}
