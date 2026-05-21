# Terminal UI QA Checks

## Test Strategy By Change

Architecture or model changes:

- Unit-test `Update` transitions for key messages.
- Test async command result messages by constructing typed messages directly.
- Test quit/cancel/back behavior.

Layout/style changes:

- Test pure render helpers when output is deterministic.
- Manually inspect at multiple terminal sizes.
- Capture screenshots when visual quality is the main deliverable.

Forms:

- Unit-test validation functions.
- Manually run happy path, invalid values, cancel, defaults, dynamic choices, and accessible mode.

SSH apps:

- Test local connect, denied connect, unsupported command, resize, disconnect, and concurrent sessions.

Markdown:

- Render sample content with headings, links, lists, code blocks, quotes, and long paragraphs.

Motion/logging:

- Test tick lifecycle and stop conditions.
- Confirm logs go to the intended destination and do not include secrets.

## Manual Terminal Matrix

Use project-appropriate equivalents when exact dimensions are awkward:

- Narrow: `60x20`
- Standard: `80x24`
- Wide: `120x32`

For each size, check:

- First screen loads cleanly.
- Primary action is visible.
- Focus is visible without color alone.
- Footer/help does not hide important content.
- Text wraps or truncates intentionally.
- Resize while interacting does not break state.

## VHS Pattern

VHS writes terminal GIFs, videos, screenshots, and frame sequences from tape files.

Basic tape:

```tape
Output demo.gif
Set Width 1200
Set Height 700
Set FontSize 18

Type "./my-tui --demo"
Enter
Wait /Ready/
Sleep 1s
Screenshot
```

Prefer:

- `Require` commands at the top for required binaries.
- `Wait` instead of long sleeps.
- Fixed dimensions and deterministic fixture data.
- Redacted or fake data.

Do not install VHS only for trivial code changes. Use it when the UI surface, motion, or demo artifact is part of the acceptance criteria.

## Evidence Template

```text
Verified:
- go test ./...: pass
- Manual 60x20, 80x24, 120x32: primary flow, resize, cancel/back, error state
- VHS: artifacts/demo.gif captures setup flow

Not verified:
- SSH deployment under systemd
```

## Red Flags

- The app only works in the developer's terminal size.
- Help text lists keys that do not work.
- Error messages disappear on resize.
- Form validation blocks progress without explaining recovery.
- Logs print over the TUI.
- Screen reader or accessible mode is mentioned but not run.
