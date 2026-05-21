package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	"charm.land/huh/v2"
	"charm.land/lipgloss/v2"
)

type config struct {
	Name       string   `json:"name"`
	Mode       string   `json:"mode"`
	Features   []string `json:"features"`
	Accessible bool     `json:"accessible"`
}

var projectNameRE = regexp.MustCompile(`^[a-z][a-z0-9-]{2,31}$`)

func sampleConfig(accessible bool) config {
	return config{
		Name:       "release-radar",
		Mode:       "dashboard",
		Features:   []string{"keyboard-help", "markdown-preview", "structured-logs"},
		Accessible: accessible,
	}
}

func validateProjectName(value string) error {
	if !projectNameRE.MatchString(value) {
		return errors.New("project name must be 3-32 lowercase letters, digits, or hyphens and start with a letter")
	}
	return nil
}

func normalizeFeatures(features []string) []string {
	seen := map[string]bool{}
	out := make([]string, 0, len(features))
	for _, feature := range features {
		feature = strings.TrimSpace(feature)
		if feature == "" || seen[feature] {
			continue
		}
		seen[feature] = true
		out = append(out, feature)
	}
	return out
}

func renderConfig(cfg config) string {
	cfg.Features = normalizeFeatures(cfg.Features)
	data, _ := json.MarshalIndent(cfg, "", "  ")
	return titleStyle.Render("Generated Charm TUI config") + "\n\n" + boxStyle.Render(string(data))
}

func runForm(accessible bool) (config, error) {
	cfg := sampleConfig(accessible)
	confirmed := false

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Project name").
				Description("lowercase letters, digits, and hyphens").
				Value(&cfg.Name).
				Validate(validateProjectName),
			huh.NewSelect[string]().
				Title("Runtime mode").
				Options(
					huh.NewOption("Dashboard", "dashboard"),
					huh.NewOption("Setup wizard", "wizard"),
					huh.NewOption("SSH app", "ssh"),
				).
				Value(&cfg.Mode),
			huh.NewMultiSelect[string]().
				Title("Features").
				Options(
					huh.NewOption("Keyboard help", "keyboard-help"),
					huh.NewOption("Markdown preview", "markdown-preview"),
					huh.NewOption("Structured logs", "structured-logs"),
					huh.NewOption("VHS demo tape", "vhs-demo"),
				).
				Value(&cfg.Features),
			huh.NewConfirm().
				Title("Generate preview?").
				Affirmative("Generate").
				Negative("Cancel").
				Value(&confirmed),
		),
	).WithAccessible(accessible)

	if err := form.Run(); err != nil {
		return cfg, err
	}
	if !confirmed {
		return cfg, errors.New("preview cancelled")
	}
	cfg.Features = normalizeFeatures(cfg.Features)
	return cfg, nil
}

var (
	titleStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("12"))
	boxStyle   = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(1, 2)
)

func main() {
	printSample := flag.Bool("print-sample", false, "print deterministic sample output and exit")
	accessibleFlag := flag.Bool("accessible", false, "enable Huh accessible mode")
	flag.Parse()

	accessible := *accessibleFlag || os.Getenv("ACCESSIBLE") != ""
	if *printSample {
		fmt.Println(renderConfig(sampleConfig(accessible)))
		return
	}

	cfg, err := runForm(accessible)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(renderConfig(cfg))
}
