# Docs SSH

This demo shows how the Charm Skills guide an agent from a prompt to a Markdown docs browser that can be inspected locally or served over SSH.

It demonstrates Glamour Markdown rendering, explicit topic routing, safe SSH server configuration, and deterministic rendering for CI.

## Run

Render a deterministic sample without starting a server:

```bash
go run ./cmd/docs-ssh --render-sample
```

List topics:

```bash
go run ./cmd/docs-ssh --list
```

Start the local SSH server:

```bash
go run ./cmd/docs-ssh --serve --addr 127.0.0.1:23234
```

Then connect from another terminal:

```bash
ssh localhost -p 23234
```

Run tests:

```bash
go test ./...
```

## Prompt To Implementation

The prompt asks for Markdown-first documentation, so `charm-glamour-markdown` chooses Glamour instead of hand-written ANSI formatting and adds a deterministic `--render-sample` path.

`charm-wish-ssh-apps` shapes the server as a local Wish SSH service with explicit address configuration and no shell exposure.

`charm-tui-builder` keeps topic selection and rendering as pure functions that can be tested outside the SSH server.

`charm-lipgloss-layout` adds a small title treatment around rendered documents.

`charm-tui-motion-observability` is represented by explicit startup messages and clean separation between rendered docs and operational output.

`charm-tui-qa` drives tests for topic lookup and rendering fallback, plus manual SSH commands in this README.
