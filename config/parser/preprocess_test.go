package configparser

import (
	"testing"
)

func TestCleanRow(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "upper case",
			input: "TEMPERATURE=79",
			want:  "TEMPERATURE=79",
		},
		{
			name:  "lower case",
			input: "temperature=79",
			want:  "temperature=79",
		},
		{
			name:  "with tab",
			input: "	temperature=79",
			want:  "temperature=79",
		},
		{
			name:  "perfectly valid data",
			input: "ActivateTemperature=79",
			want:  "ActivateTemperature=79",
		},
		{
			name:  "with spaces",
			input: "ActivateTemperature = 79",
			want:  "ActivateTemperature=79",
		},
		{
			name:  "with newline",
			input: "ActivateTemperature=79\n",
			want:  "ActivateTemperature=79",
		},
		{
			name:  "with 3 newlines",
			input: "ActivateTemperature=79\n\n\n",
			want:  "ActivateTemperature=79",
		},
		{
			name:  "only comment",
			input: "#ActivateTemperature=79",
			want:  "",
		},
		{
			name:  "with comment",
			input: "Hysteresis=10  # so DeactivateTemperature=69",
			want:  "Hysteresis=10",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cleanRow(tt.input)

			if got != tt.want {
				t.Errorf("cleanRow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitAndCleanRows(t *testing.T) {
	input := `# /etc/thermal-guard.cfg

ActivateTemperature=79
Hysteresis=10  # so DeactivateTemperature=69
MeasurementInterval=250 # in milliseconds`

	want := []string{
		"ActivateTemperature=79",
		"Hysteresis=10",
		"MeasurementInterval=250",
	}

	got := splitAndCleanRows(input)

	for i, wantRow := range want {
		if got[i] != wantRow {
			t.Errorf("splitAndCleanRows()[%d] = %q, want %q", i, got[i], wantRow)
		}
	}
}
