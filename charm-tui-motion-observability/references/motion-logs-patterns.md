# Motion And Logging Patterns

## Contents

- Imports
- Harmonica Spring Pattern
- Animation UX
- Charm Log Pattern
- Wish Session Logging
- Observability Checklist

## Imports

Use:

```go
import "github.com/charmbracelet/harmonica"
import log "charm.land/log/v2"
```

Older Log examples may use `github.com/charmbracelet/log`; verify against the target project's module path before editing imports.

## Harmonica Spring Pattern

Harmonica is framework-agnostic. Store position and velocity in the model:

```go
type model struct {
    x        float64
    velocity float64
    target   float64
    spring   harmonica.Spring
    animating bool
}
```

Initialize with the same frame rate you will actually tick:

```go
m.spring = harmonica.NewSpring(harmonica.FPS(60), 6.0, 0.8)
```

Advance on tick messages:

```go
type tickMsg time.Time

func tick() tea.Cmd {
    return tea.Tick(time.Second/60, func(t time.Time) tea.Msg {
        return tickMsg(t)
    })
}
```

Stop ticking when the value is close enough to target and velocity is negligible.

## Animation UX

Good uses:

- Progress bar easing when progress updates are discrete.
- Sliding focus indicator.
- Expanding/collapsing a detail region.

Bad uses:

- Continuous decorative movement.
- Animating content that users need to read.
- Large layout shifts in narrow terminals.

## Charm Log Pattern

Create a logger once and inject it:

```go
logger := log.NewWithOptions(os.Stderr, log.Options{
    ReportTimestamp: true,
    Level:           log.InfoLevel,
})

logger.Info("starting tui", "screen", "projects")
logger.Error("load failed", "err", err)
```

For full-screen Bubble Tea apps, use a file or disabled logger for debug output:

```go
f, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o600)
if err != nil {
    return err
}
logger := log.New(f)
```

## Wish Session Logging

Include safe connection metadata:

- user
- remote address
- command or subsystem
- TERM value
- initial dimensions
- duration

Do not log raw screen content or secrets.

## Observability Checklist

- The user can still see a clean TUI while debug logging is enabled.
- Logs have enough fields to diagnose failed actions.
- Error logs include the actual error value.
- Sensitive inputs are omitted or redacted.
- `Fatal` is only used at process boundaries where immediate exit is correct.
- Animation loops stop after completion or screen changes.
