package main

import (
	"strings"
	"testing"
)

func TestTopicsAreSorted(t *testing.T) {
	got := topics()
	if len(got) < 2 {
		t.Fatalf("expected at least two topics, got %v", got)
	}
	for i := 1; i < len(got); i++ {
		if got[i-1] > got[i] {
			t.Fatalf("topics not sorted: %v", got)
		}
	}
}

func TestMarkdownForDefaultsToWelcome(t *testing.T) {
	doc, ok := markdownFor("")
	if !ok {
		t.Fatal("expected default topic to exist")
	}
	if !strings.Contains(doc, "Charm Skills Docs") {
		t.Fatalf("unexpected default doc: %q", doc)
	}
}

func TestRenderTopicFallsBackForMissingTopic(t *testing.T) {
	out := renderTopic("missing", 80)
	if !strings.Contains(out, "Missing") || !strings.Contains(out, "topic") {
		t.Fatalf("missing fallback heading in %q", out)
	}
	if !strings.Contains(out, "welcome") {
		t.Fatalf("missing available topics in %q", out)
	}
}

func TestDefaultServerConfigIsSafe(t *testing.T) {
	cfg := defaultServerConfig("127.0.0.1:23234")
	if cfg.AllowShell {
		t.Fatal("default server config must not allow shell access")
	}
	if cfg.Addr == "" || cfg.DefaultDoc == "" {
		t.Fatalf("incomplete config: %#v", cfg)
	}
}
