package temperatures

func GetCPUTemperature() int {
	filepath := telemetryPaths.CPU
	return getTemperature(filepath)
}
