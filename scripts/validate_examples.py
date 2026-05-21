#!/usr/bin/env python3
"""Validate demonstration example structure and optionally run Go tests."""

from __future__ import annotations

import argparse
import subprocess
import sys
from pathlib import Path


ROOT = Path(__file__).resolve().parents[1]
EXAMPLES = ROOT / "examples"
REQUIRED_FILES = ("PROMPT.md", "README.md", "go.mod")


def collect_examples() -> list[Path]:
    if not EXAMPLES.exists():
        return []
    return sorted(path for path in EXAMPLES.iterdir() if path.is_dir())


def has_go_source(example: Path) -> bool:
    return any(path.suffix == ".go" for path in example.rglob("*.go"))


def has_go_test(example: Path) -> bool:
    return any(path.name.endswith("_test.go") for path in example.rglob("*.go"))


def first_heading(path: Path) -> str:
    for line in path.read_text(encoding="utf-8").splitlines():
        if line.startswith("# "):
            return line[2:].strip()
    return ""


def validate_example(example: Path, errors: list[str]) -> None:
    for required in REQUIRED_FILES:
        if not (example / required).exists():
            errors.append(f"{example}: missing {required}")

    prompt = example / "PROMPT.md"
    if prompt.exists():
        text = prompt.read_text(encoding="utf-8")
        if not text.lstrip().startswith("# Prompt"):
            errors.append(f"{prompt}: should start with '# Prompt'")
        for expected in ("Expected Skills", "Acceptance Criteria"):
            if expected not in text:
                errors.append(f"{prompt}: missing {expected!r} section")

    readme = example / "README.md"
    if readme.exists():
        text = readme.read_text(encoding="utf-8")
        heading = first_heading(readme)
        if not heading:
            errors.append(f"{readme}: missing H1")
        for expected in ("Run", "Prompt To Implementation"):
            if expected not in text:
                errors.append(f"{readme}: missing {expected!r} section")

    if not has_go_source(example):
        errors.append(f"{example}: missing Go source files")
    if not has_go_test(example):
        errors.append(f"{example}: missing Go test files")


def run_go_tests(example: Path) -> int:
    print(f"==> go test ./... ({example.relative_to(ROOT)})")
    completed = subprocess.run(
        ["go", "test", "./..."],
        cwd=example,
        text=True,
        stdout=subprocess.PIPE,
        stderr=subprocess.STDOUT,
        check=False,
    )
    if completed.stdout:
        print(completed.stdout, end="" if completed.stdout.endswith("\n") else "\n")
    return completed.returncode


def main() -> int:
    parser = argparse.ArgumentParser()
    parser.add_argument("--test", action="store_true", help="also run go test ./... in each example")
    args = parser.parse_args()

    errors: list[str] = []
    examples = collect_examples()
    if not examples:
        errors.append("No example directories found")

    for example in examples:
        validate_example(example, errors)

    if errors:
        print("Example validation failed:\n", file=sys.stderr)
        for error in errors:
            print(f"- {error}", file=sys.stderr)
        return 1

    if args.test:
        failures = [example for example in examples if run_go_tests(example) != 0]
        if failures:
            print("Go tests failed for examples:", file=sys.stderr)
            for example in failures:
                print(f"- {example.relative_to(ROOT)}", file=sys.stderr)
            return 1

    print(f"Validated {len(examples)} examples.")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
