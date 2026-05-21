package main

import (
	"strings"
	"testing"
)

func TestThemeIndexWraps(t *testing.T) {
	if got := nextIndex(2, 3); got != 0 {
		t.Fatalf("nextIndex wrap got %d", got)
	}
	if got := previousIndex(0, 3); got != 2 {
		t.Fatalf("previousIndex wrap got %d", got)
	}
}

func TestThemesHaveNames(t *testing.T) {
	for _, theme := range themes() {
		if theme.Name == "" {
			t.Fatalf("theme missing name: %#v", theme)
		}
	}
}

func TestRenderIncludesPreviewElements(t *testing.T) {
	out := render(newModel())
	for _, want := range []string{"Palette Studio", "Signal", "Release summary", "Promote", "Rollback"} {
		if !strings.Contains(out, want) {
			t.Fatalf("render missing %q in %q", want, out)
		}
	}
}
