---
name: charm-huh-forms
description: Build polished terminal forms, prompts, setup wizards, confirmations, and accessible input flows with Charm Huh. Use when Codex needs standalone prompts, multi-page forms, select/multiselect/input/text/confirm fields, validation, dynamic fields, accessible screen-reader mode, or Huh embedded in a Bubble Tea application.
---

# Charm Huh Forms

## Overview

Use Huh when the user needs to collect structured input in the terminal. Prefer it over hand-built prompt loops unless the flow requires custom full-screen behavior that belongs in Bubble Tea.

## Workflow

1. Read `references/form-patterns.md` before implementing non-trivial forms.
2. Model the answers first, then choose fields: input, text, select, multiselect, confirm, or file picker if available in the current Huh version.
3. Group fields into pages by task and cognitive load.
4. Add validation and limits where invalid answers would cause downstream failure.
5. Provide an accessible mode controlled by config or environment.
6. Verify with keyboard-only entry, invalid input, cancel/error handling, and accessible mode.

## Choose Huh When

- The task is a setup wizard, config generator, survey, confirmation, selection, or one-off prompt.
- Users need validation, defaults, choices, and clear field-level errors.
- Accessibility matters and a plain prompt mode is acceptable.
- The form can run standalone or as a bounded subflow inside a larger app.

## Prefer Bubble Tea Instead When

- The UI needs custom navigation, live dashboards, split panes, background updates, or persistent multi-screen state.
- Form fields must interact with non-form components continuously.
- The user spends most of the session browsing, comparing, or manipulating data rather than answering fields.

## Implementation Rules

- Store field results in typed variables or a config struct.
- Use typed options rather than stringly typed sentinel values when possible.
- Validate near the field so the user sees the error before submission.
- Use dynamic field functions for dependent questions, but bind recomputation narrowly to the values that actually change.
- Offer an explicit accessible mode through environment or config, not as a hidden code path.
- Theme forms to match the surrounding app when a brand/theme exists; otherwise use a restrained default.

## Avoid

- Avoid asking too many questions on one screen.
- Avoid making destructive confirmations default to "yes".
- Avoid dynamic option functions that perform expensive work on every field update.
- Avoid accessibility as an afterthought. Wire it before final verification.
- Avoid using Huh for long-running progress or real-time dashboards.

## Verification

- Unit-test validation functions.
- Manually run the form through happy path, invalid input, cancel, empty/default values, and dynamic-field changes.
- Run with accessible mode enabled, for example an `ACCESSIBLE=1` project convention if present.
- Confirm output/config serialization is deterministic and does not include prompt-only labels where machine values are expected.

## References

- `references/form-patterns.md`: field selection, dynamic forms, accessible mode, and verification checklist.
