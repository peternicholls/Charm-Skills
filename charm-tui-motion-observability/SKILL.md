---
name: charm-tui-motion-observability
description: Add tasteful terminal motion and useful observability to Charm applications with Harmonica and Charm Log. Use when Codex needs spring animations, progress motion, animated state transitions, spinner/progress tuning, structured human-readable logs, slog/logfmt/JSON output, debug logging that does not corrupt a TUI, or instrumentation for Bubble Tea and Wish apps.
---

# Charm TUI Motion And Observability

## Overview

Use this when the terminal UI needs to feel responsive and diagnosable. Motion should clarify state changes; logging should help developers and operators without corrupting the rendered terminal surface.

## Workflow

1. Read `references/motion-logs-patterns.md` before adding animation loops or logging.
2. Decide whether motion is necessary. Prefer Bubbles spinner/progress for common loading and progress states.
3. Use Harmonica for spring movement only when natural motion improves comprehension.
4. Decide where logs go before adding them. Full-screen TUIs should usually log to a file, debug pane, or disabled-by-default logger.
5. Add structured fields to logs for user action, screen, session, operation, and error.
6. Verify animation stops, logs avoid secrets, and the TUI output remains clean.

## Motion Rules

- Use motion for progress, focus transition, reveal/hide, or spatial continuity.
- Keep animations short and interruptible.
- Store animation state in the model and advance it through tick messages or component commands.
- Disable or simplify motion in tests and non-interactive output.
- Do not animate critical information that must be read immediately.

## Logging Rules

- Use Charm Log for human-readable structured logs.
- Prefer per-session or per-component loggers over global logging in reusable packages.
- Write debug logs away from stdout/stderr while Bubble Tea owns the terminal.
- Use levels consistently: debug for internal state, info for lifecycle, warn for recoverable anomalies, error for failed operations.
- Include errors as structured fields.

## Avoid

- Avoid unbounded tick loops after a screen is closed.
- Avoid motion that changes layout dimensions every frame and causes text jitter.
- Avoid logging keystrokes, tokens, passwords, private keys, or raw terminal contents.
- Avoid `Fatal` in library code or recoverable TUI paths because it exits the process.
- Avoid using logs as user-facing error messages.

## Verification

- Run tests for tick/update transitions and log configuration helpers.
- Manually verify animation start, completion, interruption, resize, and screen switch.
- Confirm logs are present at the configured level and absent below it.
- Confirm full-screen rendering is not corrupted by log output.
- Confirm logs do not contain secrets or sensitive user input.

## References

- `references/motion-logs-patterns.md`: Harmonica spring pattern, Bubble Tea tick integration, Charm Log configuration, and observability checklist.
