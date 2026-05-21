# Versioning

Charm Skills uses Semantic Versioning for the Skill suite as a whole.

Version numbers are stored in [VERSION](../VERSION) and released as Git tags named `vMAJOR.MINOR.PATCH`.

## Version Rules

Increment `MAJOR` for:

- Removing or renaming a Skill.
- Changing Skill trigger intent in a way that breaks existing user prompts.
- Removing documented workflows or verification expectations.
- Reorganizing the repository layout in a way that breaks installation instructions.

Increment `MINOR` for:

- Adding a new Skill.
- Adding new library coverage, reference patterns, or verification workflows.
- Expanding a Skill's supported task surface in a backward-compatible way.

Increment `PATCH` for:

- Fixing inaccurate guidance.
- Clarifying wording without changing supported behavior.
- Updating metadata, examples, links, or validation rules.
- Small compatibility fixes for current Charm library behavior.

## Pre-Releases

Use pre-release versions for broad experimental work:

```text
0.2.0-alpha.1
0.2.0-rc.1
```

Do not tag pre-releases as stable releases.

## Compatibility

The suite is intended for Codex-compatible Skill consumers. A release should not be considered stable unless:

- Every Skill validates with `python3 scripts/validate_skills.py`.
- Installation instructions still work.
- `SKILL.md` frontmatter remains compatible with Codex Skill discovery.
- Breaking changes are documented in [CHANGELOG.md](../CHANGELOG.md).
