package main

import (
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {
	tests := []struct {
		name     string
		tm       time.Time
		expected string
	}{
		{
			name:     "it returns a string in the specified format",
			tm:       time.Date(2020, 12, 17, 10, 0, 0, 0, time.UTC),
			expected: "17 Dec 2020 at 10:00",
		},
		{
			name:     "if input is the zero time, it returns the empty string",
			tm:       time.Time{},
			expected: "",
		},
		{

			name:     "it uses UTC time zone",
			tm:       time.Date(2020, 12, 17, 10, 0, 0, 0, time.FixedZone("CET", 1*60*60)),
			expected: "17 Dec 2020 at 09:00",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := humanDate(test.tm)

			if result != test.expected {
				t.Errorf("expected %q, got %q", test.expected, result)
			}
		})
	}
}
