package main

import (
	"testing"

	"../../src/utils/envloader"
)

func TestLoadConfig(t *testing.T) {
	got := envloader.LoadConfig()
	if got.ENV != "production" && got.ENV != "development" {
		t.Errorf("value of ENV is invalid.")
	}
}

func TestFailed(t *testing.T) {
	t.Errorf("FAILED")
}
