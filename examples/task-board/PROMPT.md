# Prompt

Build a compact terminal kanban board for a small engineering team. It should show cards across columns, support keyboard navigation, move cards between columns, and render cleanly at narrow and wide terminal sizes. Use the Charm Skills so the app is idiomatic, visually clear, and easy to verify with deterministic sample data.

## Expected Skills

- `charm-tui-builder`
- `charm-lipgloss-layout`
- `charm-bubbletea-components`
- `charm-tui-qa`

## Acceptance Criteria

- Runs locally with embedded task data and no network calls.
- Uses Bubble Tea for state and key handling.
- Uses Lip Gloss for responsive board layout and focused-card styling.
- Supports moving focus across columns and cards.
- Supports advancing a card to the next workflow column.
- Provides keyboard help and quit behavior.
- Includes deterministic sample output for CI and quick inspection.
- Includes tests for card movement, selection clamping, and rendering.
