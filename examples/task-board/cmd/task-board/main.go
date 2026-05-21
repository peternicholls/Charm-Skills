package main

import (
	"flag"
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type card struct {
	ID       string
	Title    string
	Owner    string
	Priority string
}

type column struct {
	Name  string
	Cards []card
}

type board struct {
	Columns []column
}

type model struct {
	board      board
	width      int
	height     int
	focusCol   int
	focusCard  int
	showHelp   bool
	lastAction string
}

func newModel() model {
	return model{
		board:      sampleBoard(),
		width:      100,
		height:     28,
		lastAction: "Ready",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyPressMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "?":
			m.showHelp = !m.showHelp
		case "h", "left":
			m.focusCol = clamp(m.focusCol-1, 0, len(m.board.Columns)-1)
			m.focusCard = m.clampedCardIndex()
		case "l", "right":
			m.focusCol = clamp(m.focusCol+1, 0, len(m.board.Columns)-1)
			m.focusCard = m.clampedCardIndex()
		case "j", "down":
			m.focusCard = clamp(m.focusCard+1, 0, max(0, len(m.board.Columns[m.focusCol].Cards)-1))
		case "k", "up":
			m.focusCard = clamp(m.focusCard-1, 0, max(0, len(m.board.Columns[m.focusCol].Cards)-1))
		case "enter":
			next, action := advanceFocused(m.board, m.focusCol, m.focusCard)
			m.board = next
			m.lastAction = action
			m.focusCard = m.clampedCardIndex()
		}
	}
	return m, nil
}

func (m model) clampedCardIndex() int {
	if len(m.board.Columns[m.focusCol].Cards) == 0 {
		return 0
	}
	return clamp(m.focusCard, 0, len(m.board.Columns[m.focusCol].Cards)-1)
}

func (m model) View() tea.View {
	v := tea.NewView(render(m))
	v.AltScreen = true
	return v
}

func advanceFocused(b board, colIndex, cardIndex int) (board, string) {
	if colIndex < 0 || colIndex >= len(b.Columns) {
		return b, "No column selected"
	}
	if colIndex == len(b.Columns)-1 {
		return b, "Done cards stay done"
	}
	cards := b.Columns[colIndex].Cards
	if cardIndex < 0 || cardIndex >= len(cards) {
		return b, "No card selected"
	}
	moving := cards[cardIndex]
	b.Columns[colIndex].Cards = append(cards[:cardIndex], cards[cardIndex+1:]...)
	b.Columns[colIndex+1].Cards = append([]card{moving}, b.Columns[colIndex+1].Cards...)
	return b, fmt.Sprintf("Moved %s to %s", moving.ID, b.Columns[colIndex+1].Name)
}

func render(m model) string {
	header := titleStyle.Render("Sprint Board") + " " + mutedStyle.Render(m.lastAction)
	boardView := renderBoard(m)
	help := "h/l columns • j/k cards • enter advance • ? help • q quit"
	if m.showHelp {
		help = "h/left previous column • l/right next column • j/down next card • k/up previous card • enter advance focused card • q quit"
	}
	return lipgloss.JoinVertical(lipgloss.Left, header, boardView, mutedStyle.Render(help))
}

func renderBoard(m model) string {
	if m.width < 80 {
		parts := make([]string, 0, len(m.board.Columns))
		for i := range m.board.Columns {
			parts = append(parts, renderColumn(m, i, max(28, m.width-4)))
		}
		return lipgloss.JoinVertical(lipgloss.Left, parts...)
	}
	colWidth := max(24, (m.width-8)/len(m.board.Columns))
	parts := make([]string, 0, len(m.board.Columns))
	for i := range m.board.Columns {
		parts = append(parts, renderColumn(m, i, colWidth))
	}
	return lipgloss.JoinHorizontal(lipgloss.Top, parts...)
}

func renderColumn(m model, index int, width int) string {
	col := m.board.Columns[index]
	rows := []string{columnTitleStyle.Width(width - 4).Render(fmt.Sprintf("%s (%d)", col.Name, len(col.Cards)))}
	if len(col.Cards) == 0 {
		rows = append(rows, mutedStyle.Render("No cards"))
	}
	for cardIndex, card := range col.Cards {
		style := cardStyle.Width(width - 4)
		if index == m.focusCol && cardIndex == m.focusCard {
			style = focusedCardStyle.Width(width - 4)
		}
		rows = append(rows, style.Render(fmt.Sprintf("%s  %s\n%s · %s", card.ID, card.Priority, card.Title, card.Owner)))
	}
	return columnStyle.Width(width).Render(lipgloss.JoinVertical(lipgloss.Left, rows...))
}

func sampleBoard() board {
	return board{Columns: []column{
		{Name: "Backlog", Cards: []card{
			{ID: "TUI-17", Title: "Add resize handling", Owner: "Ari", Priority: "P1"},
			{ID: "TUI-22", Title: "Write VHS tape", Owner: "Mina", Priority: "P2"},
		}},
		{Name: "Doing", Cards: []card{
			{ID: "TUI-31", Title: "Polish key help", Owner: "Sol", Priority: "P1"},
		}},
		{Name: "Review", Cards: []card{
			{ID: "TUI-35", Title: "Check accessible mode", Owner: "Bea", Priority: "P2"},
		}},
		{Name: "Done", Cards: []card{
			{ID: "TUI-09", Title: "Create initial model", Owner: "Kai", Priority: "P3"},
		}},
	}}
}

var (
	titleStyle       = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("12"))
	mutedStyle       = lipgloss.NewStyle().Faint(true)
	columnStyle      = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(0, 1).MarginRight(1)
	columnTitleStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("14"))
	cardStyle        = lipgloss.NewStyle().Border(lipgloss.NormalBorder()).Padding(0, 1).MarginTop(1)
	focusedCardStyle = cardStyle.BorderForeground(lipgloss.Color("10")).Foreground(lipgloss.Color("15"))
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	printSample := flag.Bool("print-sample", false, "print deterministic board output and exit")
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
