package main

import (
	"regexp"
	"testing"
	"time"
)

func Test_generateActivationKey(t *testing.T) {
	pattern := "^[A-Z0-9]{4}-[A-Z0-9]{4}-[A-Z0-9]{4}-[A-Z0-9]{4}$"
	re := regexp.MustCompile(pattern)

	tests := []struct {
		name string
		want string
	}{
		{name: "#1", want: pattern},
		{name: "#2", want: pattern},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1 := generateActivationKey()
			if !re.MatchString(got1) {
				t.Errorf("generateActivationKey() = %v, want %v", got1, tt.want)
			}
			time.Sleep(time.Nanosecond)
			got2 := generateActivationKey()
			if !re.MatchString(got2) {
				t.Errorf("generateActivationKey() = %v, want %v", got2, tt.want)
			}
			if got1 == got2 {
				t.Errorf("resolts are equal, got1 = %s, got2 = %s, want different", got1, got2)
			}
		})
	}
}
