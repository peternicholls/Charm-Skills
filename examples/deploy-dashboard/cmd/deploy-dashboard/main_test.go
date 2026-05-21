package main

import (
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
)

func TestProgressBarClamps(t *testing.T) {
	if got := progressBar(-10, 5); !strings.Contains(got, "0%") {
		t.Fatalf("expected low clamp, got %q", got)
	}
	if got := progressBar(150, 5); !strings.Contains(got, "100%") {
		t.Fatalf("expected high clamp, got %q", got)
	}
}

func TestSelectionSyncsLogs(t *testing.T) {
	m := newModel()
	initial := m.logs.View()

	updated, _ := m.Update(tea.KeyPressMsg{Code: tea.KeyDown})
	next := updated.(model)

	if next.list.Index() != 1 {
		t.Fatalf("expected selection index 1, got %d", next.list.Index())
	}
	if next.logs.View() == initial {
		t.Fatal("expected logs to update after selection changed")
	}
	if !strings.Contains(next.logs.View(), "migration lock") {
		t.Fatalf("expected billing-worker logs, got %q", next.logs.View())
	}
}

func TestRenderIncludesCoreRegions(t *testing.T) {
	m := newModel()
	out := render(m)
	for _, want := range []string{"Deploy Control", "Deployments", "api-gateway", "health checks"} {
		if !strings.Contains(out, want) {
			t.Fatalf("render missing %q in %q", want, out)
		}
	}
}
