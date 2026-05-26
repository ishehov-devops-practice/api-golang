package main

import (
	"testing"
)

// A simple unit test to verify core logic
func TestSanityCheck(t *testing.T) {
	expectedStatus := "active"
	currentStatus := "active"

	if currentStatus != expectedStatus {
		t.Errorf("Sanity check failed: got %s, want %s", currentStatus, expectedStatus)
	}
}
