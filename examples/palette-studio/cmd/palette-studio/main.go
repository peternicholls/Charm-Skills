package main

import (
	"flag"
	"fmt"
	"image/color"
	"os"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type theme struct {
	Name    string
	Accent  color.Color
	Success color.Color
	Warning color.Color
	Danger  color.Color
	Muted   color.Color
}

type model struct {
	themes   []theme
	index    int
	width    int
	showHelp bool
}

func newModel() model {
	return model{themes: themes(), width: 96}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
	case tea.KeyPressMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "n", "right":
			m.index = nextIndex(m.index, len(m.themes))
		case "p", "left":
			m.index = previousIndex(m.index, len(m.themes))
		case "?":
			m.showHelp = !m.showHelp
		}
	}
	return m, nil
}

func (m model) View() tea.View {
	v := tea.NewView(render(m))
	v.AltScreen = true
	return v
}

func nextIndex(index, length int) int {
	if length == 0 {
		return 0
	}
	return (index + 1) % length
}

func previousIndex(index, length int) int {
	if length == 0 {
		return 0
	}
	return (index - 1 + length) % length
}

func render(m model) string {
	active := m.themes[m.index]
	header := lipgloss.NewStyle().Bold(true).Foreground(active.Accent).Render("Palette Studio") + " " + muted(active).Render(active.Name)
	preview := renderPreview(active, m.width)
	help := "n next • p previous • ? help • q quit"
	if m.showHelp {
		help = "n/right next theme • p/left previous theme • ? toggle help • q/ctrl+c quit"
	}
	return lipgloss.JoinVertical(lipgloss.Left, header, preview, muted(active).Render(help))
}

func renderPreview(t theme, width int) string {
	swatches := []string{
		swatch("accent", t.Accent),
		swatch("success", t.Success),
		swatch("warning", t.Warning),
		swatch("danger", t.Danger),
	}
	card := cardStyle(t).Render(strings.Join([]string{
		lipgloss.NewStyle().Bold(true).Foreground(t.Accent).Render("Release summary"),
		"api-gateway rolled out to 3 regions",
		buttonStyle(t).Render("Promote") + " " + dangerStyle(t).Render("Rollback"),
	}, "\n"))
	alerts := lipgloss.JoinVertical(lipgloss.Left,
		alertStyle(t, t.Success).Render("OK  checks passed"),
		alertStyle(t, t.Warning).Render("WARN  slow migration"),
		alertStyle(t, t.Danger).Render("ERR  queue threshold"),
	)
	if width < 78 {
		return lipgloss.JoinVertical(lipgloss.Left, lipgloss.JoinHorizontal(lipgloss.Top, swatches...), card, alerts)
	}
	return lipgloss.JoinVertical(lipgloss.Left, lipgloss.JoinHorizontal(lipgloss.Top, swatches...), lipgloss.JoinHorizontal(lipgloss.Top, card, alerts))
}

func swatch(label string, color color.Color) string {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("15")).
		Background(color).
		Padding(0, 1).
		MarginRight(1).
		Render(label)
}

func muted(t theme) lipgloss.Style {
	return lipgloss.NewStyle().Foreground(t.Muted)
}

func cardStyle(t theme) lipgloss.Style {
	return lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(t.Accent).Padding(1, 2).MarginRight(2).Width(42)
}

func buttonStyle(t theme) lipgloss.Style {
	return lipgloss.NewStyle().Foreground(lipgloss.Color("0")).Background(t.Success).Bold(true).Padding(0, 1)
}

func dangerStyle(t theme) lipgloss.Style {
	return lipgloss.NewStyle().Foreground(lipgloss.Color("15")).Background(t.Danger).Bold(true).Padding(0, 1)
}

func alertStyle(t theme, color color.Color) lipgloss.Style {
	return lipgloss.NewStyle().Border(lipgloss.NormalBorder(), false, false, false, true).BorderForeground(color).PaddingLeft(1).MarginBottom(1).Width(28)
}

func themes() []theme {
	return []theme{
		{Name: "Signal", Accent: lipgloss.Color("12"), Success: lipgloss.Color("10"), Warning: lipgloss.Color("11"), Danger: lipgloss.Color("9"), Muted: lipgloss.Color("8")},
		{Name: "Orchid", Accent: lipgloss.Color("13"), Success: lipgloss.Color("14"), Warning: lipgloss.Color("11"), Danger: lipgloss.Color("9"), Muted: lipgloss.Color("8")},
		{Name: "Forest", Accent: lipgloss.Color("2"), Success: lipgloss.Color("10"), Warning: lipgloss.Color("3"), Danger: lipgloss.Color("1"), Muted: lipgloss.Color("8")},
	}
}

func main() {
	printSample := flag.Bool("print-sample", false, "print deterministic sample output and exit")
	flag.Parse()

	m := newModel()
	if *printSample {
		fmt.Println(render(m))
		return
	}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
