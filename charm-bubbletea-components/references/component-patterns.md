# Bubble Tea Component Patterns

## Contents

- Import
- Parent-Owned Focus
- Key Bindings
- Size Management
- List Guidance
- Viewport Guidance
- Progress And Spinner Guidance
- Component QA

## Import

Use Bubbles v2 paths:

```go
import "charm.land/bubbles/v2/list"
import "charm.land/bubbles/v2/textinput"
import "charm.land/bubbles/v2/key"
import "charm.land/bubbles/v2/help"
```

Verify exact package paths in the target project before editing imports.

## Parent-Owned Focus

Track focus in the parent:

```go
type focus int

const (
    focusSearch focus = iota
    focusResults
    focusDetails
)

type model struct {
    focus focus
    search textinput.Model
    results list.Model
}
```

Route updates:

```go
var cmd tea.Cmd

switch m.focus {
case focusSearch:
    m.search, cmd = m.search.Update(msg)
case focusResults:
    m.results, cmd = m.results.Update(msg)
}

return m, cmd
```

If multiple components must respond to the same message, collect commands and return `tea.Batch(cmds...)`.

## Key Bindings

Use the Bubbles `key` package for matching and help:

```go
type keyMap struct {
    Up   key.Binding
    Down key.Binding
}

var keys = keyMap{
    Up: key.NewBinding(key.WithKeys("k", "up"), key.WithHelp("up", "move up")),
    Down: key.NewBinding(key.WithKeys("j", "down"), key.WithHelp("down", "move down")),
}
```

Use `help.Model` to render short/full help when the app has more than a few actions.

## Size Management

On `tea.WindowSizeMsg`, compute available space after headers, footers, margins, and borders, then set component sizes:

```go
case tea.WindowSizeMsg:
    m.width = msg.Width
    m.height = msg.Height
    m.results.SetSize(msg.Width-4, msg.Height-6)
```

Do not render a component at its default size in a full-screen app.

## List Guidance

Use `list.Model` for searchable/browsable collections. Provide:

- Item title and description that scan independently.
- Empty state text.
- Filtering behavior that matches the task.
- Status text for selection, filtering, and item counts.

Keep domain actions in the parent model: the list chooses the item, the parent decides what it means.

## Viewport Guidance

Use `viewport.Model` for long markdown, logs, help, or output. Recompute content when source text or width changes. Preserve scroll position where possible, but clamp it after content shrink.

## Progress And Spinner Guidance

Use spinner for unknown duration. Use progress for known percentage or step count. Stop timers/spinners when work completes or the screen changes.

## Component QA

- Check zero, one, and many items.
- Check long labels and wide Unicode.
- Check keyboard-only operation.
- Check paste into text components.
- Check focus after deleting selected items.
- Check resize while a component has focus.
