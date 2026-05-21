package main

import (
	"strings"
	"testing"
)

func TestAdvanceFocusedMovesCard(t *testing.T) {
	b := sampleBoard()
	next, action := advanceFocused(b, 0, 0)
	if !strings.Contains(action, "TUI-17") {
		t.Fatalf("unexpected action: %q", action)
	}
	if len(next.Columns[0].Cards) != 1 {
		t.Fatalf("expected backlog to lose a card, got %d", len(next.Columns[0].Cards))
	}
	if got := next.Columns[1].Cards[0].ID; got != "TUI-17" {
		t.Fatalf("expected moved card first in Doing, got %s", got)
	}
}

func TestDoneCardDoesNotMove(t *testing.T) {
	b := sampleBoard()
	next, action := advanceFocused(b, 3, 0)
	if action != "Done cards stay done" {
		t.Fatalf("unexpected action: %q", action)
	}
	if next.Columns[3].Cards[0].ID != "TUI-09" {
		t.Fatal("done card moved unexpectedly")
	}
}

func TestRenderContainsBoardRegions(t *testing.T) {
	out := render(newModel())
	for _, want := range []string{"Sprint Board", "Backlog", "Doing", "TUI-17", "enter advance"} {
		if !strings.Contains(out, want) {
			t.Fatalf("render missing %q in %q", want, out)
		}
	}
}
