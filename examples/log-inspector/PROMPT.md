# Prompt

Build a terminal log inspector for support engineers. It should show a filter box, severity counts, a table of matching log events, and a detail preview for the selected event. Use Charm Skills so the app uses reusable components, keyboard-friendly navigation, deterministic sample data, and clear verification.

## Expected Skills

- `charm-tui-builder`
- `charm-bubbletea-components`
- `charm-lipgloss-layout`
- `charm-tui-motion-observability`
- `charm-tui-qa`

## Acceptance Criteria

- Runs locally with embedded log data and no network calls.
- Uses Bubble Tea for the interactive loop.
- Uses Bubbles text input and table components.
- Uses Lip Gloss for status, preview, and layout.
- Supports filtering by service, severity, or message text.
- Includes deterministic sample output for CI and quick inspection.
- Includes tests for filtering, table rows, and severity counts.
