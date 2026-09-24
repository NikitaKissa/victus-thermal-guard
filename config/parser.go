package config

import (
	"strconv"

	configparser "github.com/NikitaKissa/victus-thermal-guard/config/parser"
)

func configMapToConfig(cfgMap configparser.ConfigMap) Config {
	var cfg Config

	// Activate Temperature
	activateTemperature := cfgMap["ActivateTemperature"]
	if activateTemperature == "" {
		cfg.ActivateTemperature = defaultActivateTemperature
	}
	activateTemperatureVal, err := stringToFloat(activateTemperature)
	if err != nil {
		cfg.ActivateTemperature = defaultActivateTemperature
	}
	cfg.ActivateTemperature = activateTemperatureVal

	// Hysteresis
	hysteresis := cfgMap["Hysteresis"]
	if hysteresis == "" {
		cfg.Hysteresis = defaultHysteresis
	}
	hysteresisVal, err := stringToFloat(hysteresis)
	if err != nil {
		cfg.Hysteresis = defaultHysteresis
	}
	cfg.Hysteresis = hysteresisVal

	// Measurement Interval
	measurementInterval := cfgMap["MeasurementInterval"]
	if activateTemperature == "" {
		cfg.MeasurementInterval = defaultMeasurementInterval
	}
	mesurementIntervalVal, err := stringToInt(measurementInterval)
	if err != nil {
		cfg.MeasurementInterval = defaultMeasurementInterval
	}
	cfg.MeasurementInterval = mesurementIntervalVal

	return cfg
}

func stringToFloat(s string) (float64, error) {
	f, err := strconv.ParseFloat(s, 64)
	return f, err
}

func stringToInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	return i, err
}
