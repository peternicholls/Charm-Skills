---
name: charm-wish-ssh-apps
description: Build, secure, test, and polish SSH-accessible terminal applications with Charm Wish. Use when Codex needs to serve Bubble Tea apps over SSH, compose Wish middleware, handle SSH key authentication, access control, PTY dimensions, per-session app state, logging middleware, local SSH testing, or deployment guidance for Go SSH apps.
---

# Charm Wish SSH Apps

## Overview

Use Wish when the terminal UI should be reachable over SSH without giving users a shell. Each SSH session should get isolated app state, terminal dimensions, and clean connection lifecycle handling.

## Workflow

1. Read `references/ssh-patterns.md` before adding Wish or changing SSH behavior.
2. Decide whether the app serves a Bubble Tea session, a command-oriented SSH API, Git behavior, or a narrow custom handler.
3. Compose middleware intentionally: access control, active terminal requirements, logging, Bubble Tea, then app-specific handlers.
4. Keep authentication, authorization, and session state explicit.
5. Test locally with `ssh localhost -p <port>` and throwaway host-key handling.
6. Verify resize, disconnect, auth failure, unsupported command, and concurrent sessions.

## Use Wish For

- Remote TUIs where users connect from any terminal.
- Per-user dashboards, admin tools, games, internal tools, or self-hosted terminal apps.
- SSH-key identity and terminal-native access without HTTPS certificates.
- Serving Bubble Tea apps through PTY-backed SSH sessions.

## Implementation Rules

- Never expose a default shell unless the product explicitly requires it and authorization is reviewed.
- Give each SSH session its own model, dependencies, logger, and cleanup path.
- Use access-control middleware for allowed commands or session types.
- Log connection metadata without leaking secrets or raw user input.
- Respect PTY dimensions and window resize messages.
- Handle disconnects and canceled contexts without goroutine leaks.

## Avoid

- Avoid global mutable model state shared across sessions.
- Avoid treating all SSH clients as interactive terminals.
- Avoid unauthenticated write actions unless the user explicitly requested public access.
- Avoid logging private key material, tokens, command payload secrets, or terminal content by default.
- Avoid deployment instructions that require replacing OpenSSH.

## Verification

- Run `go test ./...` and any integration tests in the project.
- Manually test local connection, denied connection, resize, disconnect, and unsupported commands.
- Test at least two concurrent sessions when state isolation matters.
- Confirm logs show connection lifecycle and errors without secrets.

## References

- `references/ssh-patterns.md`: Wish middleware ordering, Bubble Tea session pattern, auth/access control, and local test flow.
