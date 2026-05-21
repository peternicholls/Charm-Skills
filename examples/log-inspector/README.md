# Log Inspector

This demo shows how the Charm Skills guide an agent from a prompt to a terminal log triage tool.

It demonstrates Bubbles `textinput` and `table`, Bubble Tea state routing, Lip Gloss layout, structured log-like sample data, and deterministic filtering logic.

## Run

Interactive mode:

```bash
go run ./cmd/log-inspector
```

Deterministic sample output:

```bash
go run ./cmd/log-inspector --print-sample
```

Run tests:

```bash
go test ./...
```

## Interactions

- Type to filter by service, severity, or message.
- `up` / `down`: move through matching events.
- `esc`: clear the filter.
- `ctrl+c`: quit.

## Prompt To Implementation

The prompt asks for filtering and tabular scanning, so `charm-bubbletea-components` maps the UI to Bubbles `textinput` and `table` rather than custom cursor math.

`charm-tui-builder` keeps raw events, filtered rows, and selected detail explicit in the model.

`charm-lipgloss-layout` creates the summary strip and detail panel without mixing layout into filtering logic.

`charm-tui-motion-observability` informs the event shape and severity counts while keeping logs in the UI rather than printing debug output over the TUI.

`charm-tui-qa` drives pure tests for filtering and a deterministic `--print-sample` mode.
