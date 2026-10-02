package config

import (
	"fmt"
	"strconv"

	"github.com/NikitaKissa/victus-thermal-guard/config/parser"
	"github.com/NikitaKissa/victus-thermal-guard/types"
)

func configMapToConfig(cfgMap parser.ConfigMap) (Config, error) {
	cfg := defaultConfig()
	errs := make(types.Errors, 0, 3)

	// Activate Temperature
	activateTemperature, err := getFloatFromMap(cfgMap, "ActivateTemperature")
	if err != nil {
		err = fmt.Errorf("parsing to float `ActivateTemperature`: %w", err)
		errs.Add(err)
	}

	if activateTemperature != nil {
		cfg.ActivateTemperature = *activateTemperature
	}

	// Hysteresis
	hysteresis, err := getFloatFromMap(cfgMap, "Hysteresis")
	if err != nil {
		err = fmt.Errorf("parsing to float `Hysteresis`: %w", err)
		errs.Add(err)
	}

	if hysteresis != nil {
		cfg.Hysteresis = *hysteresis
	}

	// Measurement Interval
	measurementInterval, err := getIntFromMap(cfgMap, "MeasurementInterval")
	if err != nil {
		err = fmt.Errorf("parsing to int `MeasurementInterval`: %w", err)
		errs.Add(err)
	}

	if measurementInterval != nil {
		cfg.MeasurementInterval = *measurementInterval
	}

	return cfg, errs.ToError(ErrSyntax)
}

func stringToFloat(s string) (float64, error) {
	f, err := strconv.ParseFloat(s, 64)
	return f, err
}

func getFloatFromMap(cfgMap parser.ConfigMap, key string) (*float64, error) {
	strValue := cfgMap[key]
	if strValue == "" {
		return nil, nil
	}

	value, err := stringToFloat(strValue)
	if err != nil {
		return nil, err
	}

	return &value, nil
}

func stringToInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	return i, err
}

func getIntFromMap(cfgMap parser.ConfigMap, key string) (*int, error) {
	strValue := cfgMap[key]
	if strValue == "" {
		return nil, nil
	}

	value, err := stringToInt(strValue)
	if err != nil {
		return nil, err
	}

	return &value, nil
}
