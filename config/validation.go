package config

import (
	"fmt"
)

func validateConfig(cfg *Config) error {
	if cfg == nil {
		panic("config during validation cant be <nil> pointer")
	}

	defaultCfg := defaultConfig()

	// ActivateTemperature
	if cfg.ActivateTemperature < 0 {
		cfg.ActivateTemperature = defaultCfg.ActivateTemperature
		return fmt.Errorf("`ActivateTemperature` can't be below 0: %w", ErrValidation)
	}

	if cfg.ActivateTemperature > 100 {
		cfg.ActivateTemperature = defaultCfg.ActivateTemperature
		return fmt.Errorf("`ActivateTemperature` cant't be higher than 100: %w", ErrValidation)
	}

	// Hysteresis
	if cfg.Hysteresis < 0 {
		cfg.Hysteresis = defaultCfg.Hysteresis
		return fmt.Errorf("`Hysteresis` can't be below 0: %w", ErrValidation)
	}

	if cfg.Hysteresis > 100 {
		cfg.Hysteresis = defaultCfg.Hysteresis
		return fmt.Errorf("`Hysteresis` cant't be higher than 100: %w", ErrValidation)
	}

	deactivateTemperature := cfg.ActivateTemperature - cfg.Hysteresis
	if deactivateTemperature < 0 {
		cfg.Hysteresis = defaultCfg.Hysteresis
		return fmt.Errorf("`ActivateTemperature` - `Hysteresis` cant't be below 0: %w", ErrValidation)
	}

	// MeasurementInterval
	if cfg.MeasurementInterval <= 0 {
		cfg.MeasurementInterval = defaultCfg.MeasurementInterval
		return fmt.Errorf("`MeasurementInterval` can't be less or equal 0ms: %w", ErrValidation)
	}

	return nil
}
