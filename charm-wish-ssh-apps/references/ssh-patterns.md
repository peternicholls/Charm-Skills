# Wish SSH App Patterns

## Import

Use:

```go
import "charm.land/wish/v2"
```

Wish also provides middleware packages. Verify exact package paths in the target project and current docs before editing imports.

## Middleware Model

Wish middlewares behave like HTTP middleware: each one wraps the next. Be deliberate about ordering because the last middleware may execute first depending on composition.

Common concerns:

- Access control: reject unsupported commands or non-interactive sessions.
- Active terminal: require PTY when rendering a TUI.
- Logging: record connection lifecycle and dimensions.
- Bubble Tea: attach SSH session input/output to a `tea.Program`.

## Bubble Tea Session Pattern

Each SSH connection should construct a fresh model:

```go
func newSessionModel(s ssh.Session) tea.Model {
    return model{
        user:   s.User(),
        width:  80,
        height: 24,
    }
}
```

Do not reuse one `tea.Program` across sessions.

## Authentication And Authorization

Keep separate:

- Authentication: who connected.
- Authorization: what that identity may do.
- Session capability: whether the client has a PTY, command, or subsystem.

Default to deny for write operations. For public read-only apps, still constrain commands and resources.

## Local Testing

Use a high local port and avoid polluting known hosts during iteration. A local SSH config can set:

```sshconfig
Host localhost
    UserKnownHostsFile /dev/null
    StrictHostKeyChecking no
```

Then test:

```bash
ssh localhost -p 23234
```

Only use relaxed host-key settings for local development.

## Concurrent Session Checklist

- Each session has isolated model state.
- Disconnect cancels any session work.
- Logs include session id or user when useful.
- Resize events affect only the connected session.
- Shared services are concurrency-safe.

## Deployment Notes

Wish implements its own SSH server and does not require OpenSSH. For system services, use a dedicated user, constrained working directory, explicit host key path, restart policy, and structured logs.
