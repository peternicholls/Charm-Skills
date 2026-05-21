# Prompt

Build a compact Bubble Tea deployment dashboard for release operators. It should show a list of deployments, a detail/log pane, progress, status, keyboard help, and clean quit/back behavior. Use the Charm Skills to make it visually polished, responsive to terminal size, component-driven, and easy to verify with deterministic sample data.

## Expected Skills

- `charm-tui-builder`
- `charm-lipgloss-layout`
- `charm-bubbletea-components`
- `charm-tui-motion-observability`
- `charm-tui-qa`

## Acceptance Criteria

- Runs locally with simulated deployment data and no network calls.
- Uses Bubble Tea for the model/update/view loop.
- Uses Lip Gloss for layout and style.
- Uses Bubbles list, viewport, help, and key bindings.
- Shows deployment status, progress, and logs.
- Provides keyboard help and quit behavior.
- Includes tests for deterministic state and rendering helpers.
- Includes a deterministic sample output path for quick inspection.
