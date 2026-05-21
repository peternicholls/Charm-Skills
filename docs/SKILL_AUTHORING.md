# Skill Authoring Guide

Use this guide when adding or changing Skills in this repository.

## Shape

Every Skill should have:

```text
skill-name/
├── SKILL.md
├── agents/
│   └── openai.yaml
└── references/
    └── focused-topic.md
```

Only add `scripts/` or `assets/` when they directly support repeated execution.

## SKILL.md

Keep `SKILL.md` short and procedural:

- What the Skill enables.
- Workflow.
- Decision rules.
- Implementation rules.
- Mistakes to avoid.
- Verification.
- Pointers to references.

Do not use `SKILL.md` as general documentation. It is loaded into an agent context and should be efficient.

## Description

The frontmatter `description` is the trigger surface. Include:

- What the Skill does.
- When to use it.
- Specific task and library names that should trigger it.

## References

References should contain deeper patterns, code examples, checklists, and edge cases.

Use one level of references from `SKILL.md`; avoid reference chains where a reference points to another reference that must also be loaded.

## Quality Bar

A Skill is ready when a fresh agent can answer:

- When should I use this?
- Which Charm tool should I choose?
- What implementation pattern should I start with?
- What should I avoid?
- How do I verify the work?
