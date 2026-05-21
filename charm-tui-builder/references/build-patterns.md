# Charm TUI Build Patterns

## Contents

- Source Baseline
- Bubble Tea Shape
- Message Design
- Screen And Focus Pattern
- View State Checklist
- Terminal Fit

## Source Baseline

Use the current Charm library list from `https://charm.land/libs/`: Bubble Tea, Huh, Lip Gloss, Wish, Bubbles, Glamour, Log, and Harmonica.

Current module paths seen in official repositories and docs:

- Bubble Tea: `charm.land/bubbletea/v2`
- Lip Gloss: `charm.land/lipgloss/v2`
- Bubbles: `charm.land/bubbles/v2`
- Huh: `charm.land/huh/v2`
- Wish: `charm.land/wish/v2`
- Glamour: docs show `charm.land/glamour/v2`; verify with the target project's `go.mod` or `go list -m -versions` before editing imports.
- Log: current `go.mod` shows `charm.land/log/v2`; older examples may use `github.com/charmbracelet/log`.
- Harmonica: `github.com/charmbracelet/harmonica`

## Bubble Tea Shape

Use this structure unless the existing project has a stronger local pattern:

```go
type model struct {
    width  int
    height int
    err    error
    ready  bool
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.WindowSizeMsg:
        m.width = msg.Width
        m.height = msg.Height
        m.ready = true
    case tea.KeyPressMsg:
        switch msg.String() {
        case "ctrl+c":
            return m, tea.Quit
        }
    }
    return m, nil
}

func (m model) View() tea.View {
    if !m.ready {
        return tea.NewView("Loading...")
    }
    return tea.NewView(renderMain(m))
}
```

Prefer value models for simple apps. Use pointer receivers only when existing code or large component state makes copying harmful.

## Message Design

Create typed messages for every async result:

```go
type loadStartedMsg struct{}
type loadFinishedMsg struct{ items []Item }
type loadFailedMsg struct{ err error }
```

Return commands for work:

```go
func loadItemsCmd(client Client) tea.Cmd {
    return func() tea.Msg {
        items, err := client.Items()
        if err != nil {
            return loadFailedMsg{err: err}
        }
        return loadFinishedMsg{items: items}
    }
}
```

Never perform network, file, or subprocess work directly inside `Update`.

## Screen And Focus Pattern

Use a small enum for screens:

```go
type screen int

const (
    screenList screen = iota
    screenDetail
    screenConfirm
)
```

Route keys by screen first, then by focused component. Keep global keys tiny: quit, back/cancel, help toggle.

## View State Checklist

Every substantial screen should render:

- Initial/loading state
- Empty state
- Populated state
- Error state with recovery action
- Disabled or unavailable state when commands cannot run
- Narrow-width state without clipped control labels

## Terminal Fit

Use `lipgloss.Width`, `lipgloss.Height`, `lipgloss.JoinHorizontal`, `lipgloss.JoinVertical`, and style width/height constraints. Avoid `len(string)` for terminal cells because ANSI sequences and wide runes break it.

At narrow widths, collapse sidebars, shorten help text, and prefer vertical stacks. At low heights, preserve the active control and action feedback before decorative chrome.
