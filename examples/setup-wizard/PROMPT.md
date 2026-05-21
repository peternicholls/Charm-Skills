# Prompt

Build a terminal setup wizard that asks for a new TUI project's name, runtime mode, optional features, and confirmation before generating a small configuration preview. Use Charm Skills so the form is accessible, validated, visually clear, and testable without interactive input.

## Expected Skills

- `charm-huh-forms`
- `charm-lipgloss-layout`
- `charm-tui-builder`
- `charm-tui-qa`

## Acceptance Criteria

- Runs locally without network calls.
- Uses Huh for interactive form fields.
- Uses Lip Gloss for deterministic preview styling.
- Provides accessible mode through a flag or environment variable.
- Provides a deterministic `--print-sample` path.
- Validates project names and configuration values.
- Includes tests for validation, defaults, and generated output.
