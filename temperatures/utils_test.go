package temperatures

import "testing"

func TestByteToInt(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		want    int
		wantErr bool
	}{
		{
			name:  "positive temperature",
			input: []byte("42000"),
			want:  42000,
		},
		{
			name:  "negative temperature",
			input: []byte("-10000"),
			want:  -10000,
		},
		{
			name:  "with newline",
			input: []byte("42000\n"),
			want:  42000,
		},
		{
			name:    "invalid value",
			input:   []byte("hello"),
			wantErr: true,
		},
		{
			name:    "empty value",
			input:   []byte(""),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := bytesToInt(tt.input)

			if (err != nil) != tt.wantErr {
				t.Fatalf("bytesToInt() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err != nil {
				return
			}

			if got != tt.want {
				t.Errorf("bytesToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
