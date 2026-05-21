# Palette Studio

This demo shows how the Charm Skills guide an agent from a prompt to a terminal design preview tool.

It demonstrates Lip Gloss color systems, swatches, component states, responsive layout, and Bubble Tea key handling.

## Run

Interactive mode:

```bash
go run ./cmd/palette-studio
```

Deterministic sample output:

```bash
go run ./cmd/palette-studio --print-sample
```

Run tests:

```bash
go test ./...
```

## Interactions

- `n` / `right`: next theme.
- `p` / `left`: previous theme.
- `?`: toggle extra help.
- `q` / `ctrl+c`: quit.

## Prompt To Implementation

The prompt asks for visual theme exploration, so `charm-lipgloss-layout` owns most of the implementation: semantic colors, swatches, cards, alert states, and responsive preview layout.

`charm-tui-builder` supplies the small Bubble Tea state machine for cycling themes and reacting to resize events.

`charm-tui-qa` drives deterministic `--print-sample` output and tests that prove theme cycling and preview rendering continue to work.
