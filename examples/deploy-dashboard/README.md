# Deploy Dashboard

This demo shows how the Charm Skills guide an agent from a prompt to a small Bubble Tea deployment dashboard.

It demonstrates a component-driven TUI with a deployment list, detail/log viewport, progress math, key help, responsive layout, and deterministic sample output.

## Run

```bash
go run ./cmd/deploy-dashboard
```

For deterministic non-interactive output:

```bash
go run ./cmd/deploy-dashboard --print-sample
```

Run tests:

```bash
go test ./...
```

## Interactions

- `j` / `down`: move to the next deployment.
- `k` / `up`: move to the previous deployment.
- `tab`: switch focus between the deployment list and log pane.
- `?`: toggle full help.
- `q` / `ctrl+c`: quit.

## Prompt To Implementation

The prompt asks for a release-operator dashboard, so `charm-tui-builder` shapes the app as a Bubble Tea model with explicit state, key routing, and deterministic sample data.

`charm-bubbletea-components` maps the browsing workflow to Bubbles `list`, `viewport`, `key`, and `help` components instead of custom widgets.

`charm-lipgloss-layout` keeps visual hierarchy in styles and render helpers, with narrow-width behavior that stacks content rather than assuming one terminal size.

`charm-tui-motion-observability` is represented by progress/status modeling and a clean boundary for logs: log lines are shown in the viewport, not printed over the full-screen UI.

`charm-tui-qa` drives the deterministic `--print-sample` path and tests for progress and rendering behavior.
