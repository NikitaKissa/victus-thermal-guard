package temperatures

import "fmt"

func GetCPUTemperature() (float64, error) {
	filepath := telemetryPaths.CPU
	telemetry, err := getTemperature(filepath)
	if err != nil {
		return 0, fmt.Errorf(
			"get cpu temperature: %w",
			err,
		)
	}

	return telemetry, nil
}
