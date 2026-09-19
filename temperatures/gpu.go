package temperatures

import "fmt"

func GetGPUTemperature() (int, error) {
	filepath := telemetryPaths.GPU
	telemetry, err := getTemperature(filepath)
	if err != nil {
		return 0, fmt.Errorf(
			"get gpu temperature: %w",
			err,
		)
	}

	return telemetry, nil
}
