package config

import (
	"testing"

	"github.com/NikitaKissa/victus-thermal-guard/config/parser"
)

func TestConfigMapToConfig(t *testing.T) {
	defaultCfg := defaultConfig()

	tests := []struct {
		name    string
		input   parser.ConfigMap
		want    Config
		wantErr bool
	}{
		{
			name: "valid data",
			input: parser.ConfigMap{
				"ActivateTemperature": "83",
				"Hysteresis":          "15",
				"MeasurementInterval": "200",
			},
			want: Config{
				ActivateTemperature: 83,
				Hysteresis:          15,
				MeasurementInterval: 200,
			},
		},
		{
			name: "all empty fields",
			input: parser.ConfigMap{
				"ActivateTemperature": "",
				"Hysteresis":          "",
				"MeasurementInterval": "",
			},
			want: defaultCfg,
		},
		{
			name: "empty `ActivateTemperature`",
			input: parser.ConfigMap{
				"ActivateTemperature": "",
				"Hysteresis":          "15",
				"MeasurementInterval": "200",
			},
			want: Config{
				ActivateTemperature: defaultCfg.ActivateTemperature,
				Hysteresis:          15,
				MeasurementInterval: 200,
			},
		},
		{
			name: "empty `Hysteresis`",
			input: parser.ConfigMap{
				"ActivateTemperature": "83",
				"Hysteresis":          "",
				"MeasurementInterval": "200",
			},
			want: Config{
				ActivateTemperature: 83,
				Hysteresis:          defaultCfg.Hysteresis,
				MeasurementInterval: 200,
			},
		},
		{
			name: "empty `MeasurementInterval`",
			input: parser.ConfigMap{
				"ActivateTemperature": "83",
				"Hysteresis":          "15",
				"MeasurementInterval": "",
			},
			want: Config{
				ActivateTemperature: 83,
				Hysteresis:          15,
				MeasurementInterval: defaultCfg.MeasurementInterval,
			},
		},
		{
			name: "invalid data",
			input: parser.ConfigMap{
				"ActivateTemperature": "80C",
				"Hysteresis":          "10C",
				"MeasurementInterval": "250ms",
			},
			want:    defaultCfg,
			wantErr: true,
		},
		{
			name:    "empty config map",
			input:   parser.ConfigMap{},
			want:    defaultCfg,
			wantErr: false,
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
