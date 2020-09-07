package main

import "testing"

func Test_randUUID(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"TEST", "DSHIDFHASDUIHDUIASHDIUAS"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := randUUID(); got != tt.want {
				t.Errorf("randUUID() = %v, want %v", got, tt.want)
			}
		})
	}
}
