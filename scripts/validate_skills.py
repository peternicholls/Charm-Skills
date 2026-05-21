#!/usr/bin/env python3
"""Validate the Charm Skills repository without external dependencies."""

from __future__ import annotations

import re
import sys
from pathlib import Path


ROOT = Path(__file__).resolve().parents[1]
SKILL_NAME_RE = re.compile(r"^[a-z0-9]+(?:-[a-z0-9]+)*$")
REFERENCE_RE = re.compile(r"`?(references/[A-Za-z0-9_.\-/]+)`?")
FORBIDDEN_TEXT = (
    "[TODO",
    "TODO:",
    "Structuring This Skill",
    "Resources (optional)",
    "placeholder script",
    "Example Asset File",
)


def fail(errors: list[str], message: str) -> None:
    errors.append(message)


def parse_frontmatter(path: Path, errors: list[str]) -> tuple[dict[str, str], str]:
    text = path.read_text(encoding="utf-8")
    if not text.startswith("---\n"):
        fail(errors, f"{path}: missing YAML frontmatter")
        return {}, text

    end = text.find("\n---\n", 4)
    if end == -1:
        fail(errors, f"{path}: frontmatter is not closed")
        return {}, text

    raw = text[4:end]
    body = text[end + 5 :]
    data: dict[str, str] = {}
    for line_no, line in enumerate(raw.splitlines(), start=2):
        if not line.strip():
            continue
        if ":" not in line:
            fail(errors, f"{path}:{line_no}: invalid frontmatter line")
            continue
        key, value = line.split(":", 1)
        key = key.strip()
        value = value.strip()
        if key in data:
            fail(errors, f"{path}:{line_no}: duplicate frontmatter key {key!r}")
        data[key] = value.strip("\"'")
    return data, body


def validate_openai_yaml(skill_dir: Path, errors: list[str]) -> None:
    path = skill_dir / "agents" / "openai.yaml"
    if not path.exists():
        fail(errors, f"{skill_dir}: missing agents/openai.yaml")
        return

    text = path.read_text(encoding="utf-8")
    for key in ("display_name", "short_description", "default_prompt"):
        if not re.search(rf"^\s*{key}:\s*.+", text, re.MULTILINE):
            fail(errors, f"{path}: missing interface.{key}")


def validate_references(skill_dir: Path, body: str, errors: list[str]) -> None:
    refs = sorted(set(REFERENCE_RE.findall(body)))
    for ref in refs:
        ref_path = skill_dir / ref
        if not ref_path.exists():
            fail(errors, f"{skill_dir / 'SKILL.md'}: referenced file does not exist: {ref}")

    references_dir = skill_dir / "references"
    if references_dir.exists():
        for ref_path in sorted(references_dir.glob("*.md")):
            lines = ref_path.read_text(encoding="utf-8").splitlines()
            if len(lines) > 100 and not any(line.strip() == "## Contents" for line in lines[:20]):
                fail(errors, f"{ref_path}: references over 100 lines need a ## Contents section near the top")


def validate_skill(skill_dir: Path, errors: list[str]) -> None:
    name = skill_dir.name
    if not SKILL_NAME_RE.fullmatch(name):
        fail(errors, f"{skill_dir}: skill folder must use lowercase hyphen-case")
    if len(name) > 64:
        fail(errors, f"{skill_dir}: skill folder name exceeds 64 characters")

    skill_md = skill_dir / "SKILL.md"
    if not skill_md.exists():
        fail(errors, f"{skill_dir}: missing SKILL.md")
        return

    frontmatter, body = parse_frontmatter(skill_md, errors)
    allowed = {"name", "description"}
    keys = set(frontmatter)
    if keys != allowed:
        fail(errors, f"{skill_md}: frontmatter keys must be exactly name and description; found {sorted(keys)}")
    if frontmatter.get("name") != name:
        fail(errors, f"{skill_md}: frontmatter name must match folder name {name!r}")
    if len(frontmatter.get("description", "")) < 60:
        fail(errors, f"{skill_md}: description is too short to be a useful trigger")
    if not body.lstrip().startswith("# "):
        fail(errors, f"{skill_md}: body should start with an H1 title after frontmatter")

    full_text = skill_md.read_text(encoding="utf-8")
    for forbidden in FORBIDDEN_TEXT:
        if forbidden in full_text:
            fail(errors, f"{skill_md}: remove template or placeholder text: {forbidden}")

    validate_openai_yaml(skill_dir, errors)
    validate_references(skill_dir, body, errors)


def discover_skill_dirs() -> list[Path]:
    return sorted(
        path
        for path in ROOT.iterdir()
        if path.is_dir()
        and not path.name.startswith(".")
        and (path / "SKILL.md").exists()
    )


def main() -> int:
    errors: list[str] = []
    skill_dirs = discover_skill_dirs()
    if not skill_dirs:
        fail(errors, "No skill directories found")

    for skill_dir in skill_dirs:
        validate_skill(skill_dir, errors)

    if errors:
        print("Skill validation failed:\n", file=sys.stderr)
        for error in errors:
            print(f"- {error}", file=sys.stderr)
        return 1

    print(f"Validated {len(skill_dirs)} skills.")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
