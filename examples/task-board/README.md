# Task Board

This demo shows how the Charm Skills guide an agent from a prompt to a small kanban-style terminal app.

It demonstrates Bubble Tea state transitions, Lip Gloss layout, keyboard navigation, and deterministic board rendering.

## Run

Interactive mode:

```bash
go run ./cmd/task-board
```

Deterministic sample output:

```bash
go run ./cmd/task-board --print-sample
```

Run tests:

```bash
go test ./...
```

## Interactions

- `h` / `left`: previous column.
- `l` / `right`: next column.
- `j` / `down`: next card.
- `k` / `up`: previous card.
- `enter`: advance the focused card to the next column.
- `?`: toggle full help.
- `q` / `ctrl+c`: quit.

## Prompt To Implementation

The prompt asks for a kanban workflow, so `charm-tui-builder` shapes the app around explicit board state and typed key transitions.

`charm-lipgloss-layout` drives the columns, cards, focused style, and narrow-width vertical fallback.

`charm-bubbletea-components` informs the keyboard/help pattern even though this demo keeps the board itself custom because cards are small domain objects.

`charm-tui-qa` drives the deterministic `--print-sample` mode plus tests for card movement and render output.
