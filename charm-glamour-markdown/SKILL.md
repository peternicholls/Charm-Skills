---
name: charm-glamour-markdown
description: Render, style, wrap, and test Markdown output in terminal applications with Charm Glamour. Use when Codex needs to display Markdown in a CLI or TUI, choose Glamour styles, build custom Markdown renderers, integrate Markdown with Bubble Tea viewports, handle terminal word wrapping, or avoid hand-written Markdown-to-ANSI rendering.
---

# Charm Glamour Markdown

## Overview

Use Glamour for Markdown. Do not manually parse Markdown or fake headings/lists with ad hoc string formatting when the source content is Markdown or Markdown-like documentation.

## Workflow

1. Read `references/markdown-patterns.md` before implementing custom rendering or Bubble Tea integration.
2. Identify whether the output is one-shot CLI text, scrollable TUI content, or a themed document pane.
3. Choose a built-in style first. Add a custom style only when the product already has a terminal theme or branded output.
4. Set word wrap from terminal width or viewport width.
5. Use Lip Gloss printing or downstream rendering where color downsampling matters.
6. Verify with headings, links, code fences, tables, lists, long lines, and narrow widths.

## Use Glamour For

- README/help rendering in a CLI.
- Markdown detail panes in a Bubble Tea viewport.
- Release notes, task descriptions, issue bodies, and generated reports.
- User-authored Markdown previews in terminal apps.
- Style-driven markdown output that should remain consistent across screens.

## Implementation Rules

- Keep source Markdown separate from rendered ANSI output.
- Re-render when width changes if wrapping depends on terminal size.
- Sanitize or constrain untrusted Markdown according to the target app's security posture.
- Use a viewport for long rendered content inside a full-screen TUI.
- Preserve plain-text or no-color output modes when the surrounding CLI supports them.

## Avoid

- Avoid regex-based Markdown rendering.
- Avoid fixed 80-column wrapping inside resizable TUIs.
- Avoid rendering Markdown into active prompts if the ANSI output will break cursor movement.
- Avoid custom styles that reduce code block, link, heading, or list readability.

## Verification

- Add tests for renderer configuration and width-sensitive output where stable.
- Manually inspect at narrow and normal widths.
- Verify dark/light or no-color modes when applicable.
- Include at least one code block, nested list, link, and long paragraph in manual samples.

## References

- `references/markdown-patterns.md`: renderer choices, viewport integration, wrapping, styles, and QA samples.
