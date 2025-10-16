package xmath

import "testing"

func TestClamp(t *testing.T) {
	testsInt := []struct {
		name     string
		value    int
		lower    int
		upper    int
		expected int
	}{
		{"Within range int", 5, 0, 10, 5},
		{"Below range int", -5, 0, 10, 0},
		{"Above range int", 15, 0, 10, 10},
		{"Equal to lower int", 0, 0, 10, 0},
		{"Equal to upper int", 10, 0, 10, 10},
	}

	for _, tt := range testsInt {
		t.Run(tt.name, func(t *testing.T) {
			if got := Clamp(tt.value, tt.lower, tt.upper); got != tt.expected {
				t.Errorf("Clamp(%v, %v, %v) = %v, want %v", tt.value, tt.lower, tt.upper, got, tt.expected)
			}
		})
	}

	testsFloat := []struct {
		name     string
		value    float64
		lower    float64
		upper    float64
		expected float64
	}{
		{"Within range float64", 5.5, 0.0, 10.0, 5.5},
		{"Below range float64", -5.5, 0.0, 10.0, 0.0},
		{"Above range float64", 15.5, 0.0, 10.0, 10.0},
		{"Equal to lower float64", 0.0, 0.0, 10.0, 0.0},
		{"Equal to upper float64", 10.0, 0.0, 10.0, 10.0},
	}

	for _, tt := range testsFloat {
		t.Run(tt.name, func(t *testing.T) {
			if got := Clamp(tt.value, tt.lower, tt.upper); got != tt.expected {
				t.Errorf("Clamp(%v, %v, %v) = %v, want %v", tt.value, tt.lower, tt.upper, got, tt.expected)
			}
		})
	}
}

func TestSign(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected interface{}
	}{
		{"Positive int", 5, 1},
		{"Negative int", -3, -1},
		{"Zero int", 0, 0},
		{"Positive float64", 5.5, 1.0},
		{"Negative float64", -3.3, -1.0},
		{"Zero float64", 0.0, 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch v := tt.input.(type) {
			case int:
				if got := Sign(v); got != tt.expected.(int) {
					t.Errorf("Sign(%v) = %v, want %v", tt.input, got, tt.expected)
				}
			case float64:
				if got := Sign(v); got != tt.expected.(float64) {
					t.Errorf("Sign(%v) = %v, want %v", tt.input, got, tt.expected)
				}
			}
		})
	}
}

func TestAbs(t *testing.T) {
	testsInt := []struct {
		name     string
		input    int
		expected int
	}{
		{"Positive int", 5, 5},
		{"Negative int", -5, 5},
		{"Zero int", 0, 0},
		{"Max negative int", -9223372036854775807, 9223372036854775807}, // Testing with max int64
	}

	for _, tt := range testsInt {
		t.Run(tt.name, func(t *testing.T) {
			if got := Abs(tt.input); got != tt.expected {
				t.Errorf("Abs(%v) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}

	testsFloat := []struct {
		name     string
		input    float64
		expected float64
	}{
		{"Positive float64", 5.5, 5.5},
		{"Negative float64", -5.5, 5.5},
		{"Zero float64", 0.0, 0.0},
		{"Small negative float64", -0.000001, 0.000001},
	}

	for _, tt := range testsFloat {
		t.Run(tt.name, func(t *testing.T) {
			if got := Abs(tt.input); got != tt.expected {
				t.Errorf("Abs(%v) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}
