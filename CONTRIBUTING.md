# Contributing

Contributions are welcome when they make the Skills more accurate, practical, concise, and useful for real agent workflows.

## Principles

- Keep Skills actionable. Prefer instructions an agent can execute over general background.
- Keep `SKILL.md` concise. Move deeper examples and patterns into `references/`.
- Prefer existing Charm idioms and official library behavior over invented abstractions.
- Add verification guidance for every meaningful behavior change.
- Avoid adding dependencies or generated assets unless there is a clear repeated need.

## Development Workflow

1. Create a branch from `main`.
2. Make a small, focused change.
3. Run validation:

   ```bash
   python3 scripts/validate_skills.py
   ```

4. Update [CHANGELOG.md](CHANGELOG.md) under `Unreleased`.
5. Open a pull request using the template.

## Skill Requirements

Each Skill directory must contain:

- `SKILL.md`
- `agents/openai.yaml`
- Optional `references/`, `scripts/`, or `assets/` only when they directly support the Skill.

`SKILL.md` frontmatter must include only:

```yaml
---
name: skill-name
description: What the skill does and when to use it.
---
```

The folder name and frontmatter `name` must match.

## Reference Files

Use references for detailed patterns, examples, APIs, and checklists that are too large for `SKILL.md`.

If a reference file is longer than 100 lines, include a short `## Contents` section near the top.

## Commit Messages

Use clear commit messages that explain why the change exists. For larger or decision-heavy changes, include trailers such as:

```text
Constraint: <external constraint>
Rejected: <alternative> | <reason>
Confidence: <low|medium|high>
Scope-risk: <narrow|moderate|broad>
Tested: <what was verified>
Not-tested: <known gaps>
```

## Pull Request Review

Reviewers should check:

- Does the change make the Skill more usable by an agent?
- Is the guidance current with Charm library behavior?
- Is the Skill concise enough to load into context?
- Are detailed examples moved to references?
- Is verification guidance concrete?
- Does validation pass?
