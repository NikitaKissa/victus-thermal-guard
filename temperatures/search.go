package temperatures

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type TelemetryPaths struct {
	CPU string
	GPU string
}

const (
	cpuChip  = "k10temp"
	cpuLabel = "Tctl"
	gpuChip  = "amdgpu"
	gpuLabel = "edge"
)

func readTrimmed(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}

func findInput(chip, label string) (string, error) {
	dirs, err := filepath.Glob("/sys/class/hwmon/hwmon*")
	if err != nil {
		return "", err
	}

	var found []string
	for _, dir := range dirs {
		name, err := readTrimmed(filepath.Join(dir, "name"))
		if err != nil || name != chip {
			continue
		}

		labelFiles, err := filepath.Glob(filepath.Join(dir, "temp*_label"))
		if err != nil {
			return "", err
		}

		for _, lf := range labelFiles {
			l, err := readTrimmed(lf)
			if err != nil || l != label {
				continue
			}
			found = append(found, strings.TrimSuffix(lf, "_label")+"_input")
		}
	}

	switch len(found) {
	case 0:
		return "", fmt.Errorf("sensor %s/%s not found", chip, label)
	case 1:
		return found[0], nil
	default:
		return "", fmt.Errorf("sensor %s/%s is ambiguous: %v", chip, label, found)
	}
}

var telemetryPaths TelemetryPaths

func FindTelemetry() error {
	cpu, err := findInput(cpuChip, cpuLabel)
	if err != nil {
		return fmt.Errorf("cpu: %w", err)
	}

	gpu, err := findInput(gpuChip, gpuLabel)
	if err != nil {
		return fmt.Errorf("gpu: %w", err)
	}

	telemetryPaths = TelemetryPaths{CPU: cpu, GPU: gpu}
	return nil
}
