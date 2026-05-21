# Prompt

Build a small terminal docs browser that renders Markdown beautifully and can also be served over SSH for remote readers. Use Charm Skills so the Markdown output is readable, the SSH app shape is safe, and the demo can be verified without starting a server.

## Expected Skills

- `charm-glamour-markdown`
- `charm-wish-ssh-apps`
- `charm-tui-builder`
- `charm-lipgloss-layout`
- `charm-tui-motion-observability`
- `charm-tui-qa`

## Acceptance Criteria

- Uses embedded sample Markdown and no network calls.
- Uses Glamour to render Markdown.
- Includes a documented Wish SSH server path.
- Provides a deterministic `--render-sample` path.
- Keeps session/app configuration explicit and safe.
- Includes tests for topic lookup and rendering fallback behavior.
