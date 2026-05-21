---
name: charm-tui-builder
description: Build idiomatic terminal user interfaces with the Charm ecosystem. Use when Codex needs to design, implement, refactor, or extend Go CLI/TUI applications using Bubble Tea, Lip Gloss, Bubbles, Huh, Glamour, Wish, Log, Harmonica, or related Charm libraries; choose the right Charm tool; structure Bubble Tea models, messages, commands, and views; or turn a plain CLI workflow into a feature-rich terminal UI.
---

# Charm TUI Builder

## Overview

Use this as the entry skill for Charm terminal UI work. Start by selecting the right Charm libraries, then shape the app around Bubble Tea's model/update/view loop, terminal constraints, and explicit verification.

## Workflow

1. Identify whether the user needs an inline prompt, full-screen TUI, SSH app, markdown renderer, or non-interactive CLI polish.
2. Read `references/build-patterns.md` before implementing architecture, state flow, routing, async commands, or multi-screen behavior.
3. Prefer Bubble Tea for interactive application state, Lip Gloss for layout/style, Bubbles for common components, Huh for form-like input, Glamour for markdown, Wish for SSH delivery, Log for structured logs, and Harmonica for spring motion.
4. Inspect the existing project before adding dependencies. Reuse current model, message, style, and component patterns when present.
5. Implement the smallest coherent user flow first: initial model, key map, viewport sizing, render states, error states, and quit/cancel behavior.
6. Verify with `go test ./...`, `go vet ./...` when available, and a manual terminal run at narrow, normal, and wide widths.

## Library Choice

- Use Bubble Tea when UI behavior depends on evolving state, keyboard/mouse input, async work, resize handling, or multiple views.
- Use Huh when the workflow is a bounded form or wizard and the user mostly answers fields.
- Use Bubbles when a standard component already exists: list, table, viewport, text input, text area, spinner, progress, paginator, timer, stopwatch, help, key bindings, or file picker.
- Use Lip Gloss for every non-trivial view layout. Avoid hand-padding large strings when a style, join, table, list, tree, or adaptive color can express the intent.
- Use Glamour when displaying markdown content; do not hand-roll markdown styling.
- Use Wish when users should connect over SSH, each session needs its own TUI, or SSH keys/pty dimensions are part of the product.
- Use Log for human-readable structured logs; never print debug traces into an active TUI surface.
- Use Harmonica only for motion that clarifies state change. Keep motion optional and bounded.

## Suite Routing

- Use `charm-lipgloss-layout` for visual hierarchy, style, measurement, responsive layout, tables/lists/trees, and polish.
- Use `charm-bubbletea-components` for Bubbles component selection, focus, key maps, and component update/view composition.
- Use `charm-huh-forms` for prompts, forms, dynamic fields, validation, and accessible input.
- Use `charm-glamour-markdown` for Markdown rendering, wrapping, styling, and viewport integration.
- Use `charm-wish-ssh-apps` for SSH-accessible TUIs, Wish middleware, auth, PTY, and session behavior.
- Use `charm-tui-motion-observability` for Harmonica motion and Charm Log instrumentation.
- Use `charm-tui-qa` for tests, manual terminal UX checks, screenshots, videos, VHS, and completion evidence.

## Implementation Rules

- Keep state explicit in the model. Avoid package-level mutable UI state unless the existing code already uses it deliberately.
- Treat `Update` as the only place that mutates UI state. Keep `View` pure and fast.
- Use typed messages for async results, ticks, window size, loading outcomes, and domain events.
- Return commands for I/O, timers, process execution, and async work; do not block inside `Update`.
- Track terminal width and height in the model and render every view for constrained widths.
- Provide discoverable key help and consistent escape paths: `esc`/`ctrl+c` for cancel or back, `q` only when safe and expected.
- Preserve screen reader or plain-output paths for forms and critical workflows.
- Prefer clear empty, loading, error, success, and disabled states over decorative output.

## Avoid

- Avoid writing a custom event loop when Bubble Tea should own input, rendering, and commands.
- Avoid concatenating ANSI escape sequences by hand when Lip Gloss or a Bubble component can render the view.
- Avoid layouts that only work at the developer's terminal size.
- Avoid making every key global. Scope key handling by screen and focus.
- Avoid logging to stdout/stderr while a full-screen TUI is active unless the program has a debug/log file path.
- Avoid adding Charm libraries speculatively. Each dependency should own a visible part of the UX.

## Verification

- Run unit tests for update logic and domain reducers.
- Add golden tests or snapshot-style tests for stable view output where the project already uses them.
- Run the TUI manually in at least `80x24`, a narrow width around `60` columns, and a wide width above `120` columns.
- Exercise keyboard-only navigation, cancel/back paths, loading and error states, and resize behavior.
- Use `charm-tui-qa` when the task requires screenshots, VHS recordings, visual review, or richer terminal UX checks.

## References

- `references/build-patterns.md`: architecture patterns, module paths, message design, and view-state checklist.
