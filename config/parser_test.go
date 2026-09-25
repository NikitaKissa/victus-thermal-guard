package config

import (
	"testing"

	configparser "github.com/NikitaKissa/victus-thermal-guard/config/parser"
)

func TestConfigMapToConfig(t *testing.T) {
	tests := []struct {
		name    string
		input   configparser.ConfigMap
		want    Config
		wantErr bool
	}{
		{
			name: "valid data",
			input: configparser.ConfigMap{
				"ActivateTemperature": "80",
				"Hysteresis":          "10",
				"MeasurementInterval": "250",
			},
			want: Config{
				ActivateTemperature: 80,
				Hysteresis:          10,
				MeasurementInterval: 250,
			},
		},
		{
			name: "all empty fields",
			input: configparser.ConfigMap{
				"ActivateTemperature": "",
				"Hysteresis":          "",
				"MeasurementInterval": "",
			},
			want: Config{
				ActivateTemperature: defaultActivateTemperature,
				Hysteresis:          defaultHysteresis,
				MeasurementInterval: defaultMeasurementInterval,
			},
		},
		{
			name: "invalid data",
			input: configparser.ConfigMap{
				"ActivateTemperature": "80C",
				"Hysteresis":          "10C",
				"MeasurementInterval": "250ms",
			},
			want: Config{
				ActivateTemperature: defaultActivateTemperature,
				Hysteresis:          defaultHysteresis,
				MeasurementInterval: defaultMeasurementInterval,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		got, err := configMapToConfig(tt.input)

		if (err != nil) != tt.wantErr {
			t.Fatalf("configMapToConfig() error = %v, wantErr %v", err, tt.wantErr)
		}

		if got != tt.want {
			t.Errorf("configMapToConfig() = %v, want %v", got, tt.want)
		}
	}
}
