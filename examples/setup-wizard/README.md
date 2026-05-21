# Setup Wizard

This demo shows how the Charm Skills guide an agent from a prompt to a Huh-powered terminal form.

It demonstrates structured input, validation, accessible mode, safe defaults, and deterministic sample output that can run in CI.

## Run

Interactive mode:

```bash
go run ./cmd/setup-wizard
```

Deterministic sample output:

```bash
go run ./cmd/setup-wizard --print-sample
```

Accessible mode:

```bash
ACCESSIBLE=1 go run ./cmd/setup-wizard
```

Run tests:

```bash
go test ./...
```

## Interactions

The wizard collects a project name, runtime mode, feature set, and confirmation. Destructive or high-impact actions are not performed; the app prints a generated configuration preview.

## Prompt To Implementation

The prompt is form-shaped rather than dashboard-shaped, so `charm-huh-forms` maps the workflow to Huh fields and validation instead of a custom prompt loop.

`charm-lipgloss-layout` styles the generated preview without mixing formatting into validation logic.

`charm-tui-builder` keeps the domain configuration as a typed struct so interactive and deterministic paths share the same code.

`charm-tui-qa` drives the `--print-sample` path and pure tests for validation and generated config output.
