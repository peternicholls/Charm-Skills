package main

import (
	"strings"
	"testing"
)

func TestFilterEventsMatchesSeverityServiceAndMessage(t *testing.T) {
	events := sampleEvents()
	cases := map[string]string{
		"error":   "worker",
		"billing": "invoice sync",
		"bundle":  "docs",
	}
	for query, want := range cases {
		got := filterEvents(events, query)
		if len(got) == 0 {
			t.Fatalf("expected results for %q", query)
		}
		if !strings.Contains(strings.ToLower(got[0].Service+" "+got[0].Message), strings.ToLower(want)) {
			t.Fatalf("query %q returned unexpected event %#v", query, got[0])
		}
	}
}

func TestSeverityCounts(t *testing.T) {
	counts := severityCounts(sampleEvents())
	if counts["warn"] != 2 || counts["error"] != 1 || counts["info"] != 2 {
		t.Fatalf("unexpected counts: %#v", counts)
	}
}

func TestRowsForTruncatesLongMessages(t *testing.T) {
	rows := rowsFor([]logEvent{{Time: "12:00:00", Severity: "info", Service: "api", Message: strings.Repeat("x", 80)}})
	if len(rows) != 1 {
		t.Fatalf("expected one row, got %d", len(rows))
	}
	if len(rows[0][3]) > 42 {
		t.Fatalf("message was not truncated: %q", rows[0][3])
	}
}

func TestRenderIncludesCoreRegions(t *testing.T) {
	out := render(newModel())
	for _, want := range []string{"Log Inspector", "debug", "api-gateway", "type to filter"} {
		if !strings.Contains(out, want) {
			t.Fatalf("render missing %q in %q", want, out)
		}
	}
}
