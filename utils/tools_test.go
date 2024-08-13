package utils

import "testing"

func TestTemperatureCharacterization(t *testing.T) {
	tests := []struct {
		name     string
		temp     int32
		expected string
	}{
		{"Hot temperature", 90, "hot"},
		{"Cold temperature", 40, "cold"},
		{"Moderate temperature", 70, "moderate"},
		{"Edge case - hot", 87, "hot"},
		{"Edge case - cold", 49, "cold"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TemperatureCharacterization(tt.temp)
			if result != tt.expected {
				t.Errorf("TemperatureCharacterization(%d) = %s; want %s", tt.temp, result, tt.expected)
			}
		})
	}
}
