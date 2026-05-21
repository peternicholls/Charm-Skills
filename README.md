# Charm Skills

Agent Skills for building excellent terminal user interfaces with the Charm ecosystem.

This repository packages a focused suite of Codex-compatible Skills that teach agents how to choose and use Charm libraries such as Bubble Tea, Lip Gloss, Bubbles, Huh, Glamour, Wish, Log, and Harmonica. The Skills are practical operating guides: each one tells an agent when to use it, what decisions to make, what implementation patterns to prefer, what mistakes to avoid, and how to verify the result.

This project is independent and is not affiliated with Charmbracelet.

## Skills

| Skill | Purpose |
| --- | --- |
| `charm-tui-builder` | Entry point for Charm TUI architecture, library selection, Bubble Tea state flow, and suite routing. |
| `charm-lipgloss-layout` | Visual hierarchy, responsive terminal layout, ANSI-aware measurement, and Lip Gloss styling. |
| `charm-bubbletea-components` | Bubbles component choice, focus routing, key maps, sizing, and update/view composition. |
| `charm-huh-forms` | Huh prompts, multi-step forms, validation, dynamic fields, and accessible input flows. |
| `charm-glamour-markdown` | Glamour Markdown rendering, wrapping, styles, and Bubble Tea viewport integration. |
| `charm-wish-ssh-apps` | Wish SSH apps, middleware, authentication, PTY behavior, and per-session state. |
| `charm-tui-motion-observability` | Harmonica motion, Charm Log instrumentation, and TUI-safe logging. |
| `charm-tui-qa` | Automated and manual terminal UX checks, screenshots, videos, VHS, and completion evidence. |

## Install

Clone the repository:

```bash
git clone <repo-url> charm-skills
```

Copy all Skills into your Codex skills directory:

```bash
mkdir -p "${CODEX_HOME:-$HOME/.codex}/skills"
cp -R charm-skills/charm-* "${CODEX_HOME:-$HOME/.codex}/skills/"
```

Or copy only the Skills you need:

```bash
cp -R charm-skills/charm-tui-builder "${CODEX_HOME:-$HOME/.codex}/skills/"
cp -R charm-skills/charm-tui-qa "${CODEX_HOME:-$HOME/.codex}/skills/"
```

Restart or reload your agent environment if it caches the skill list.

## Use

Ask Codex or a compatible agent to use the relevant Skill by name:

```text
Use charm-tui-builder to design a Bubble Tea app for browsing deploy logs.
```

For broad Charm TUI work, start with `charm-tui-builder`. It routes to the specialized Skills when layout, forms, components, Markdown, SSH, motion/logging, or QA need deeper guidance.

## Examples

Runnable prompt-to-project demonstrations live in [examples](examples/README.md). Each example starts with a `PROMPT.md`, names the Skills an agent should use, and includes the generated project plus deterministic verification.

## Repository Layout

Each Skill follows the standard Codex Skill shape:

```text
skill-name/
├── SKILL.md
├── agents/
│   └── openai.yaml
└── references/
    └── topic-specific-patterns.md
```

The `SKILL.md` files stay concise so agents can load them quickly. Deeper examples and patterns live in `references/` and are loaded only when needed.

## Validate

Run the repository validator before opening a pull request:

```bash
python3 scripts/validate_skills.py
python3 scripts/validate_examples.py
```

The validators check Skill frontmatter, folder names, UI metadata, reference links, stale template text, reference structure, and example project structure.

## Versioning

This repository uses SemVer for the Skill suite. See [docs/VERSIONING.md](docs/VERSIONING.md).

## Releases

Releases are tagged as `vMAJOR.MINOR.PATCH` and documented in [CHANGELOG.md](CHANGELOG.md). See [docs/RELEASE_PROCESS.md](docs/RELEASE_PROCESS.md).

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md). Please keep Skills concise, practical, and focused on agent execution rather than general documentation.

## License

MIT. See [LICENSE](LICENSE).
