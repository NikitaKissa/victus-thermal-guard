package configparser

import "testing"

func TestParse(t *testing.T) {
	input := `# /etc/thermal-guard.cfg

ActivateTemperature=79
Hysteresis=10  # so DeactivateTemperature=69
MeasurementInterval=250 # in milliseconds
EmptyField=`

	want := ConfigMap{
		"ActivateTemperature": "79",
		"Hysteresis":          "10",
		"MeasurementInterval": "250",
		"EmptyField":          "",
	}

	got, err := Parse([]byte(input))
	if err != nil {
		t.Fatalf("Parse() err = %v, wantErr false", err)
	}

	for key, wantValue := range want {
		gotValue, ok := got[key]
		if !ok {
			t.Errorf("Parse() missing key %q", key)
			continue
		}

		if gotValue != wantValue {
			t.Errorf("Parse()[%q] = %q, want %q", key, gotValue, wantValue)
		}
	}
}
