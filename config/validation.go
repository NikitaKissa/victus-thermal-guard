package config

import (
	"fmt"

	"github.com/NikitaKissa/victus-thermal-guard/types"
)

const (
	MeasurementIntervalMin = 50
)

func validateAndRepairConfig(cfg *Config) error {
	if cfg == nil {
		panic("config during validation cant be <nil> pointer")
	}

	defaultCfg := defaultConfig()

	errs := make(types.Errors, 0, 4)

	// ActivateTemperature
	if cfg.ActivateTemperature < 0 {
		cfg.ActivateTemperature = defaultCfg.ActivateTemperature
		errs.Add(fmt.Errorf("`ActivateTemperature` can't be below 0: %w", ErrValidation))
	}

	if cfg.ActivateTemperature > 100 {
		cfg.ActivateTemperature = defaultCfg.ActivateTemperature
		errs.Add(fmt.Errorf("`ActivateTemperature` cant't be higher than 100: %w", ErrValidation))
	}

	// Hysteresis
	if cfg.Hysteresis < 0 {
		cfg.Hysteresis = defaultCfg.Hysteresis
		errs.Add(fmt.Errorf("`Hysteresis` can't be below 0: %w", ErrValidation))
	}

	if cfg.Hysteresis > 100 {
		cfg.Hysteresis = defaultCfg.Hysteresis
		errs.Add(fmt.Errorf("`Hysteresis` cant't be higher than 100: %w", ErrValidation))
	}

	deactivateTemperature := cfg.ActivateTemperature - cfg.Hysteresis
	if deactivateTemperature < 0 {
		cfg.Hysteresis = defaultCfg.Hysteresis
		errs.Add(fmt.Errorf("`ActivateTemperature` - `Hysteresis` cant't be below 0: %w", ErrValidation))
	}

	// MeasurementInterval
	if cfg.MeasurementInterval < MeasurementIntervalMin {
		cfg.MeasurementInterval = defaultCfg.MeasurementInterval
		errs.Add(fmt.Errorf("`MeasurementInterval` can't be less than %dms: %w", MeasurementIntervalMin, ErrValidation))
	}

	return errs.ToError(ErrValidation)
}
