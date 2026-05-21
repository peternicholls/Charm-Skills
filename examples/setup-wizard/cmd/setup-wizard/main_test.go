package main

import (
	"strings"
	"testing"
)

func TestValidateProjectName(t *testing.T) {
	valid := []string{"app", "release-radar", "a12"}
	for _, value := range valid {
		if err := validateProjectName(value); err != nil {
			t.Fatalf("expected %q to be valid: %v", value, err)
		}
	}

	invalid := []string{"Aaa", "ab", "-bad", "bad_name"}
	for _, value := range invalid {
		if err := validateProjectName(value); err == nil {
			t.Fatalf("expected %q to be invalid", value)
		}
	}
}

func TestNormalizeFeatures(t *testing.T) {
	got := normalizeFeatures([]string{"logs", "logs", " ", "help"})
	want := []string{"logs", "help"}
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got %v, want %v", got, want)
		}
	}
}

func TestRenderConfigIncludesSample(t *testing.T) {
	out := renderConfig(sampleConfig(true))
	for _, want := range []string{"Generated Charm TUI config", "release-radar", "keyboard-help", `"accessible": true`} {
		if !strings.Contains(out, want) {
			t.Fatalf("rendered config missing %q in %q", want, out)
		}
	}
}
