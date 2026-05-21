---
name: charm-bubbletea-components
description: Add, configure, test, and compose Charm Bubbles components in Bubble Tea applications. Use when Codex needs text input, text area, table, list, viewport, spinner, progress, paginator, timer, stopwatch, help, key bindings, file picker, focus management, or reusable Bubble Tea component patterns.
---

# Charm Bubble Tea Components

## Overview

Use Bubbles before writing custom terminal widgets. Components should own their local state while the parent Bubble Tea model owns routing, focus, orchestration, and domain state.

## Workflow

1. Read `references/component-patterns.md` before adding or replacing components.
2. Choose the closest existing Bubble: list, table, viewport, text input, text area, spinner, progress, paginator, timer, stopwatch, help, key, or file picker.
3. Add the component to the parent model with explicit focus and dimensions.
4. Route messages to the focused component, collect its returned command, and batch commands when multiple components update.
5. Render component views inside Lip Gloss layout constraints.
6. Verify component behavior with keyboard navigation, resize events, and any existing tests.

## Component Choice

- Use `textinput` for one-line editable values and search boxes.
- Use `textarea` for multiline input and chat/message composition.
- Use `list` for browse/filter/select flows with built-in pagination, help, spinner, and status messages.
- Use `table` for row/column data where navigation matters.
- Use `viewport` for scrollable markdown, logs, help, or long detail panes.
- Use `spinner` for unknown-duration work and `progress` for known progress.
- Use `key` plus `help` to make keybindings declarative and discoverable.
- Use `filepicker` for filesystem selection instead of shelling out to platform pickers.

## Implementation Rules

- Initialize each component with meaningful placeholder text, width, height, style, and keymap.
- Keep focus explicit. A parent model should know which component is active.
- Pass `tea.WindowSizeMsg` effects into component sizing; do not let components render beyond available space.
- Use Bubbles key bindings for matching and help text instead of string switches scattered through the model.
- Keep domain validation outside visual components unless the component API intentionally owns validation.
- Use `tea.Batch` for multiple component commands returned from one update.

## Avoid

- Avoid duplicating Bubbles behavior with custom structs unless the UX clearly exceeds the component.
- Avoid updating every component for every key when only one has focus.
- Avoid showing loading spinners without explaining what is loading.
- Avoid hidden keybindings. Include help or obvious labels for essential actions.
- Avoid component widths that ignore borders, sidebars, or terminal width.

## Verification

- Run tests around focus routing, component state transitions, and domain actions triggered by components.
- Manually verify arrow keys, vim-style keys if supported, tab/shift-tab focus, enter, escape, delete/backspace, paste, and resize.
- Verify lists and tables with zero, one, many, and filtered items.
- Verify text components with long Unicode input and pasted text.

## References

- `references/component-patterns.md`: component map, focus routing, update composition, and test checklist.
