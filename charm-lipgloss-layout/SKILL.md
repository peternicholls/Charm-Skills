---
name: charm-lipgloss-layout
description: Design, refactor, and polish terminal UI layout and styling with Charm Lip Gloss. Use when Codex needs to improve visual hierarchy, color, spacing, responsive terminal layouts, tables/lists/trees, adaptive light/dark color behavior, ANSI-aware measurement, or the View rendering layer of Bubble Tea and Charm applications.
---

# Charm Lip Gloss Layout

## Overview

Use this when the work is about how a terminal UI looks, scans, wraps, and adapts. Lip Gloss should carry layout and style decisions instead of hand-built ANSI strings or fragile padding.

## Workflow

1. Read `references/layout-patterns.md` before changing a non-trivial view.
2. Identify the user's primary scanning task, then assign hierarchy: title, active region, secondary context, help, status.
3. Define reusable styles near the view or style package used by the project. Keep style names semantic: `titleStyle`, `activeItemStyle`, `mutedStyle`.
4. Measure and fit with Lip Gloss functions, not byte length.
5. Check light/dark background behavior and reduced-color terminals where the project supports them.
6. Verify manually at narrow, normal, and wide widths.

## Use Lip Gloss For

- Borders, padding, margins, alignment, width, height, foreground/background color, bold/faint/italic text, and hyperlinks.
- Horizontal and vertical composition.
- Tables, lists, and trees through Lip Gloss subpackages when the output is primarily presentational.
- Runtime color decisions with terminal-aware helpers when the app can inspect stdin/stdout.
- Rendering text in a Bubble Tea `View` without side effects.

## Design Rules

- Make the active region obvious without relying on color alone.
- Keep borders purposeful: one frame for the main focus or none. Avoid nested decorative boxes.
- Use color as state and hierarchy, not wallpaper. Pair color with text, icon, position, or shape.
- Reserve high-contrast accents for current selection, errors, destructive actions, or completion.
- Preserve useful content before decorative spacing when width or height shrinks.
- Style data density by task: dashboards and admin tools should be compact; setup wizards can breathe more.

## Avoid

- Avoid negative margins, brittle manual spacing, and repeated string padding.
- Avoid assuming one terminal theme. Test dark and light backgrounds when possible.
- Avoid using Unicode-only signals for critical state without ASCII fallback or labels.
- Avoid wide fixed columns unless the app explicitly requires a minimum terminal width and explains it.
- Avoid adding a new palette when the project already has semantic styles.

## Verification

- Run existing view/golden tests when present.
- Add or update view tests for deterministic render helpers.
- Manually inspect at `60x20`, `80x24`, and `120x32` or comparable project sizes.
- Check that status/error text remains visible after resizing.
- Use `charm-tui-qa` when screenshots, videos, or visual regression artifacts are needed.

## References

- `references/layout-patterns.md`: layout recipes, adaptive colors, measurement, and view QA checklist.
