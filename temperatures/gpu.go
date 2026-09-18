package temperatures

func GetGPUTemperature() int {
	filepath := telemetryPaths.GPU
	return getTemperature(filepath)
}
