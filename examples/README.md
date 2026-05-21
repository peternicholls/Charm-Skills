# Examples

These examples show how the Charm Skills can guide an agent from prompt to working proof of concept.

Each example starts with `PROMPT.md`, then includes the project that an agent built from that prompt. The examples are deliberately small, deterministic, and local-only so users can inspect the prompt, run the code, and see how the Skills affect implementation decisions.

## Projects

| Example | Demonstrates | Primary Skills |
| --- | --- | --- |
| `deploy-dashboard` | Bubble Tea dashboard with component focus, responsive layout, logs, progress, and QA hooks. | `charm-tui-builder`, `charm-lipgloss-layout`, `charm-bubbletea-components`, `charm-tui-motion-observability`, `charm-tui-qa` |
| `setup-wizard` | Huh form workflow with validation, defaults, accessible mode, and deterministic sample output. | `charm-huh-forms`, `charm-lipgloss-layout`, `charm-tui-builder`, `charm-tui-qa` |
| `docs-ssh` | Markdown docs rendering and SSH-accessible terminal app shape. | `charm-glamour-markdown`, `charm-wish-ssh-apps`, `charm-tui-builder`, `charm-lipgloss-layout`, `charm-tui-qa` |
| `task-board` | Kanban-style board with custom domain state, responsive columns, and card movement. | `charm-tui-builder`, `charm-lipgloss-layout`, `charm-bubbletea-components`, `charm-tui-qa` |
| `log-inspector` | Filterable support log triage app with text input, table rows, severity counts, and detail preview. | `charm-tui-builder`, `charm-bubbletea-components`, `charm-lipgloss-layout`, `charm-tui-motion-observability`, `charm-tui-qa` |
| `palette-studio` | Terminal theme preview studio with swatches, state samples, cards, and theme switching. | `charm-tui-builder`, `charm-lipgloss-layout`, `charm-tui-qa` |

## How To Read An Example

1. Read `PROMPT.md` first.
2. Check the expected Skills and acceptance criteria.
3. Read `README.md` for the agent's implementation decisions.
4. Run the deterministic command before trying the interactive mode.
5. Run the example tests when Go dependencies are available.

## Validate Example Structure

From the repository root:

```bash
python3 scripts/validate_examples.py
```

To run Go tests in every example as well:

```bash
python3 scripts/validate_examples.py --test
```

The `--test` mode may download Go modules.
