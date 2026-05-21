package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"

	"charm.land/bubbles/v2/table"
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type logEvent struct {
	Time     string
	Severity string
	Service  string
	Message  string
	TraceID  string
}

type model struct {
	width    int
	events   []logEvent
	filter   textinput.Model
	table    table.Model
	filtered []logEvent
}

func newModel() model {
	filter := textinput.New()
	filter.Placeholder = "filter logs..."
	filter.Prompt = "/ "
	filter.SetWidth(32)
	_ = filter.Focus()

	t := table.New(
		table.WithColumns([]table.Column{
			{Title: "Time", Width: 8},
			{Title: "Level", Width: 7},
			{Title: "Service", Width: 14},
			{Title: "Message", Width: 42},
		}),
		table.WithFocused(true),
		table.WithHeight(8),
		table.WithWidth(78),
	)

	m := model{
		width:  96,
		events: sampleEvents(),
		filter: filter,
		table:  t,
	}
	m.applyFilter()
	return m
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.table.SetWidth(max(60, msg.Width-4))
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "esc":
			m.filter.SetValue("")
			m.applyFilter()
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.filter, cmd = m.filter.Update(msg)
	if _, ok := msg.(tea.KeyPressMsg); ok {
		m.applyFilter()
	}
	m.table, _ = m.table.Update(msg)
	return m, cmd
}

func (m model) View() tea.View {
	v := tea.NewView(render(m))
	v.AltScreen = true
	return v
}

func (m *model) applyFilter() {
	m.filtered = filterEvents(m.events, m.filter.Value())
	m.table.SetRows(rowsFor(m.filtered))
}

func filterEvents(events []logEvent, query string) []logEvent {
	query = strings.ToLower(strings.TrimSpace(query))
	if query == "" {
		return append([]logEvent(nil), events...)
	}
	out := make([]logEvent, 0, len(events))
	for _, event := range events {
		haystack := strings.ToLower(strings.Join([]string{event.Time, event.Severity, event.Service, event.Message, event.TraceID}, " "))
		if strings.Contains(haystack, query) {
			out = append(out, event)
		}
	}
	return out
}

func rowsFor(events []logEvent) []table.Row {
	rows := make([]table.Row, 0, len(events))
	for _, event := range events {
		rows = append(rows, table.Row{event.Time, event.Severity, event.Service, truncate(event.Message, 42)})
	}
	return rows
}

func severityCounts(events []logEvent) map[string]int {
	counts := map[string]int{}
	for _, event := range events {
		counts[event.Severity]++
	}
	return counts
}

func render(m model) string {
	counts := severityCounts(m.filtered)
	summary := fmt.Sprintf("debug %d  info %d  warn %d  error %d", counts["debug"], counts["info"], counts["warn"], counts["error"])
	selected := selectedEvent(m)
	detail := detailStyle.Width(max(60, m.width-4)).Render(fmt.Sprintf("%s  %s  %s\ntrace %s\n%s", selected.Time, selected.Severity, selected.Service, selected.TraceID, selected.Message))
	return lipgloss.JoinVertical(
		lipgloss.Left,
		titleStyle.Render("Log Inspector"),
		m.filter.View(),
		mutedStyle.Render(summary),
		m.table.View(),
		detail,
		mutedStyle.Render("type to filter • up/down select • esc clear • ctrl+c quit"),
	)
}

func selectedEvent(m model) logEvent {
	if len(m.filtered) == 0 {
		return logEvent{Time: "--:--", Severity: "none", Service: "none", Message: "No matching events.", TraceID: "-"}
	}
	index := clamp(m.table.Cursor(), 0, len(m.filtered)-1)
	return m.filtered[index]
}

func sampleEvents() []logEvent {
	events := []logEvent{
		{"12:00:01", "info", "api-gateway", "deploy started for release 2026.05.22", "trc-1001"},
		{"12:00:07", "debug", "api-gateway", "loaded 18 route rules", "trc-1001"},
		{"12:01:13", "warn", "billing", "retrying invoice sync after timeout", "trc-2011"},
		{"12:01:44", "error", "worker", "queue depth exceeded policy threshold", "trc-3098"},
		{"12:02:05", "info", "docs", "published markdown bundle", "trc-4100"},
		{"12:02:22", "warn", "worker", "slow job detected in image pipeline", "trc-3099"},
	}
	sort.SliceStable(events, func(i, j int) bool { return events[i].Time < events[j].Time })
	return events
}

func truncate(value string, width int) string {
	if len(value) <= width {
		return value
	}
	if width <= 3 {
		return value[:width]
	}
	return value[:width-3] + "..."
}

func clamp(value, low, high int) int {
	if value < low {
		return low
	}
	if value > high {
		return high
	}
	return value
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var (
	titleStyle  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("12"))
	mutedStyle  = lipgloss.NewStyle().Faint(true)
	detailStyle = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(0, 1).MarginTop(1)
)

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
