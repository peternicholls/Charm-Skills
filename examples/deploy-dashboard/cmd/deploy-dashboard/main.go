package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/list"
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type deployment struct {
	Name     string
	Env      string
	Status   string
	Progress int
	Logs     []string
}

func (d deployment) Title() string {
	return d.Name
}

func (d deployment) Description() string {
	return fmt.Sprintf("%s · %s · %s", d.Env, d.Status, progressBar(d.Progress, 10))
}

func (d deployment) FilterValue() string {
	return d.Name + " " + d.Env + " " + d.Status
}

type focusArea int

const (
	focusList focusArea = iota
	focusLogs
)

type keyMap struct {
	Up   key.Binding
	Down key.Binding
	Tab  key.Binding
	Help key.Binding
	Quit key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Up, k.Down, k.Tab, k.Help, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{{k.Up, k.Down, k.Tab}, {k.Help, k.Quit}}
}

var keys = keyMap{
	Up:   key.NewBinding(key.WithKeys("k", "up"), key.WithHelp("k/up", "up")),
	Down: key.NewBinding(key.WithKeys("j", "down"), key.WithHelp("j/down", "down")),
	Tab:  key.NewBinding(key.WithKeys("tab"), key.WithHelp("tab", "focus")),
	Help: key.NewBinding(key.WithKeys("?"), key.WithHelp("?", "help")),
	Quit: key.NewBinding(key.WithKeys("q", "ctrl+c"), key.WithHelp("q", "quit")),
}

type model struct {
	width  int
	height int
	focus  focusArea
	keys   keyMap
	help   help.Model
	list   list.Model
	logs   viewport.Model
}

func newModel() model {
	items := make([]list.Item, 0, len(sampleDeployments()))
	for _, item := range sampleDeployments() {
		items = append(items, item)
	}
	deployList := list.New(items, list.NewDefaultDelegate(), 34, 14)
	deployList.Title = "Deployments"
	deployList.SetShowStatusBar(false)
	deployList.SetFilteringEnabled(false)
	deployList.DisableQuitKeybindings()

	logs := viewport.New(viewport.WithWidth(48), viewport.WithHeight(14))
	m := model{
		width:  96,
		height: 24,
		keys:   keys,
		help:   help.New(),
		list:   deployList,
		logs:   logs,
	}
	m.syncLogs()
	return m
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.resize()
	case tea.KeyPressMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.keys.Tab):
			if m.focus == focusList {
				m.focus = focusLogs
			} else {
				m.focus = focusList
			}
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		default:
			var cmd tea.Cmd
			if m.focus == focusLogs {
				m.logs, cmd = m.logs.Update(msg)
			} else {
				before := m.list.Index()
				m.list, cmd = m.list.Update(msg)
				if m.list.Index() != before {
					m.syncLogs()
				}
			}
			return m, cmd
		}
	}
	return m, nil
}

func (m *model) resize() {
	footerHeight := 2
	bodyHeight := max(8, m.height-footerHeight-3)
	if m.width < 74 {
		m.list.SetSize(max(20, m.width-4), min(10, bodyHeight/2))
		m.logs.SetWidth(max(20, m.width-4))
		m.logs.SetHeight(max(6, bodyHeight-m.list.Height()-1))
		return
	}
	listWidth := min(38, max(26, m.width/3))
	m.list.SetSize(listWidth, bodyHeight)
	m.logs.SetWidth(max(24, m.width-listWidth-8))
	m.logs.SetHeight(bodyHeight)
}

func (m *model) syncLogs() {
	item, ok := m.list.SelectedItem().(deployment)
	if !ok {
		m.logs.SetContent("No deployment selected.")
		return
	}
	m.logs.SetContent(strings.Join(item.Logs, "\n"))
}

func (m model) View() tea.View {
	v := tea.NewView(render(m))
	v.AltScreen = true
	return v
}

func render(m model) string {
	header := titleStyle.Render("Deploy Control")
	status := mutedStyle.Render("Release operator dashboard · simulated data")
	selected, _ := m.list.SelectedItem().(deployment)
	detail := renderDetail(selected, m.width)
	logPane := panelStyle.Width(m.logs.Width()).Height(m.logs.Height()).Render(m.logs.View())
	listPane := panelStyle.Width(m.list.Width()).Height(m.list.Height()).Render(m.list.View())

	var body string
	if m.width < 74 {
		body = lipgloss.JoinVertical(lipgloss.Left, detail, listPane, logPane)
	} else {
		right := lipgloss.JoinVertical(lipgloss.Left, detail, logPane)
		body = lipgloss.JoinHorizontal(lipgloss.Top, listPane, right)
	}

	footer := mutedStyle.Render(m.help.View(m.keys))
	return lipgloss.JoinVertical(lipgloss.Left, header, status, body, footer)
}

func renderDetail(d deployment, width int) string {
	if d.Name == "" {
		return panelStyle.Render("No deployment selected.")
	}
	line := fmt.Sprintf("%s  %s  %s", d.Name, d.Env, progressBar(d.Progress, 20))
	if width < 74 {
		line = fmt.Sprintf("%s\n%s\n%s", d.Name, d.Env, progressBar(d.Progress, 20))
	}
	return detailStyle.Render(line)
}

func progressBar(progress int, width int) string {
	progress = clamp(progress, 0, 100)
	filled := progress * width / 100
	return "[" + strings.Repeat("#", filled) + strings.Repeat("-", width-filled) + fmt.Sprintf("] %3d%%", progress)
}

func sampleDeployments() []deployment {
	return []deployment{
		{
			Name:     "api-gateway",
			Env:      "prod",
			Status:   "rolling",
			Progress: 72,
			Logs: []string{
				"12:01 pulled image ghcr.io/acme/api:2026.05.21",
				"12:02 drained 2 old pods",
				"12:03 started canary pod api-gateway-7df",
				"12:04 health checks passing in 3 regions",
			},
		},
		{
			Name:     "billing-worker",
			Env:      "prod",
			Status:   "waiting",
			Progress: 35,
			Logs: []string{
				"12:00 queued after api-gateway",
				"12:01 migration lock still active",
				"12:04 waiting for approval window",
			},
		},
		{
			Name:     "docs-site",
			Env:      "staging",
			Status:   "complete",
			Progress: 100,
			Logs: []string{
				"11:55 built static assets",
				"11:56 uploaded release bundle",
				"11:57 smoke checks complete",
			},
		},
	}
}

var (
	titleStyle  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("12"))
	mutedStyle  = lipgloss.NewStyle().Faint(true)
	panelStyle  = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(0, 1).MarginRight(1)
	detailStyle = lipgloss.NewStyle().Border(lipgloss.NormalBorder()).Padding(0, 1).MarginBottom(1)
)

func clamp(value, low, high int) int {
	if value < low {
		return low
	}
	if value > high {
		return high
	}
	return value
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
