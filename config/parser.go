package config

import (
	"errors"
	"fmt"
	"strconv"

	configparser "github.com/NikitaKissa/victus-thermal-guard/config/parser"
)

var (
	ErrSyntax = errors.New("syntax error during parsing config")
)

func configMapToConfig(cfgMap configparser.ConfigMap) (Config, error) {
	var cfg Config
	errs := make([]error, 0, 3)

	// Activate Temperature
	activateTemperature := cfgMap["ActivateTemperature"]
	if activateTemperature == "" {
		cfg.ActivateTemperature = defaultActivateTemperature
	} else {
		activateTemperatureVal, err := stringToFloat(activateTemperature)
		if err != nil {
			cfg.ActivateTemperature = defaultActivateTemperature
			errs = append(
				errs,
				fmt.Errorf("parsing to float `ActivateTemperature`: %w", err),
			)
		} else {
			cfg.ActivateTemperature = activateTemperatureVal
		}
	}

	// Hysteresis
	hysteresis := cfgMap["Hysteresis"]
	if hysteresis == "" {
		cfg.Hysteresis = defaultHysteresis
	} else {
		hysteresisVal, err := stringToFloat(hysteresis)
		if err != nil {
			cfg.Hysteresis = defaultHysteresis
			errs = append(
				errs,
				fmt.Errorf("parsing to float `Hysteresis`: %w", err),
			)
		} else {
			cfg.Hysteresis = hysteresisVal
		}
	}

	// Measurement Interval
	measurementInterval := cfgMap["MeasurementInterval"]
	if activateTemperature == "" {
		cfg.MeasurementInterval = defaultMeasurementInterval
	} else {
		measurementIntervalVal, err := stringToInt(measurementInterval)
		if err != nil {
			cfg.MeasurementInterval = defaultMeasurementInterval
			errs = append(
				errs,
				fmt.Errorf("parsing to int `MeasurementInterval`: %w", err),
			)
		} else {
			cfg.MeasurementInterval = measurementIntervalVal
		}
	}

	// errors processing

	if len(errs) == 0 {
		return cfg, nil
	}

	err := fmt.Errorf("%w\n", ErrSyntax)
	for _, val := range errs {
		err = fmt.Errorf(
			"%w\n%v",
			err,
			val,
		)
	}

	return cfg, err
}

func stringToFloat(s string) (float64, error) {
	f, err := strconv.ParseFloat(s, 64)
	return f, err
}

func stringToInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	return i, err
}
