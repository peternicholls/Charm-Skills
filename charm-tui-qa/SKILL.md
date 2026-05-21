---
name: charm-tui-qa
description: Verify, record, and polish Charm terminal user interfaces with automated and manual terminal UX checks. Use when Codex needs to test Bubble Tea update/view behavior, run Go tests, capture screenshots or videos with VHS, inspect terminal rendering at multiple sizes, validate keyboard-only workflows, accessibility/plain-output paths, error/loading/empty states, or produce completion evidence for Charm CLI/TUI work.
---

# Charm TUI QA

## Overview

Use this before claiming Charm terminal UI work is complete. The goal is evidence: tests for behavior, terminal runs for interaction, and screenshots or VHS recordings when visual quality matters.

## Workflow

1. Read `references/qa-checks.md` before designing verification.
2. Derive checks from the changed UX: architecture, component behavior, layout, forms, SSH sessions, markdown, motion, or logs.
3. Run automated checks first: `go test ./...`, existing lint/type/static checks, and targeted tests.
4. Run the app manually in a real terminal or pseudo-terminal at multiple sizes.
5. Capture screenshots or VHS recordings when layout, motion, onboarding, demos, or regression artifacts matter.
6. Report exact commands, dimensions, scenarios, and any unverified risk.

## Verification

Use both automated checks and manual terminal checks. Pick the depth based on the UI risk, but never claim terminal UX quality from build output alone.

## Automated Checks

- Unit-test Bubble Tea `Update` behavior with typed messages and keypresses.
- Test pure render helpers and stable view output where the project already supports golden/snapshot tests.
- Test validation functions for Huh forms.
- Test renderer configuration for Glamour.
- Test log configuration without asserting fragile color output unless the project already has terminal golden testing.

## Manual Checks

- Run at narrow, normal, and wide terminal sizes.
- Use keyboard-only navigation for every primary task.
- Exercise loading, empty, populated, error, success, cancel, back, quit, and resize states.
- Check long labels, wide Unicode, ANSI-colored content, pasted input, and no-color mode when supported.
- Confirm logs do not corrupt full-screen output.

## Recording

Use VHS when a feature needs a reproducible visual demo, screenshot, or motion proof. Keep tapes deterministic: fixed terminal size, seeded data, short sleeps, explicit waits, and no network dependency unless the product requires it.

## Avoid

- Avoid claiming visual quality from unit tests alone.
- Avoid recording a demo with private data, live tokens, or unstable timestamps.
- Avoid only testing the happy path.
- Avoid treating one terminal size as proof of responsiveness.
- Avoid leaving manual observations vague. Record dimensions and actions.

## Completion Evidence

A strong final verification note includes:

- Commands run and whether they passed.
- Manual scenarios and terminal dimensions.
- Screenshots/video paths when captured.
- Known gaps, if any.
- Any follow-up risk that remains outside the task scope.

## References

- `references/qa-checks.md`: test strategy, manual terminal checklist, VHS tape pattern, and evidence template.
