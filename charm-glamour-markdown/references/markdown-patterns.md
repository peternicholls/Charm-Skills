# Glamour Markdown Patterns

## Import

Current docs show:

```go
import "charm.land/glamour/v2"
```

Verify the module path against the target project's dependencies before editing, because older Glamour releases used `github.com/charmbracelet/glamour`.

## Simple Rendering

```go
out, err := glamour.Render(markdown, "dark")
if err != nil {
    return err
}
fmt.Print(out)
```

Use this for one-shot CLI output when a default style is sufficient.

## Width-Aware Renderer

```go
r, err := glamour.NewTermRenderer(
    glamour.WithWordWrap(width),
)
if err != nil {
    return err
}

rendered, err := r.Render(markdown)
```

Use the available content width, not the whole terminal width, when the Markdown lives inside a panel or viewport.

## Bubble Tea Viewport Pattern

- Store raw Markdown in the model.
- Store rendered Markdown or regenerate it in a helper when width changes.
- Set viewport content after rendering.
- Clamp scroll position after content changes.

```go
func (m *model) renderMarkdown() tea.Cmd {
    width := max(20, m.viewport.Width)
    rendered, err := renderMarkdown(m.rawMarkdown, width)
    if err != nil {
        m.err = err
        return nil
    }
    m.viewport.SetContent(rendered)
    return nil
}
```

## Styles

Start with built-in styles such as `dark`, `light`, or project defaults. Use `GLAMOUR_STYLE` or environment-aware configuration when users expect customization.

Add custom styles when:

- The app already has a terminal theme.
- Headings, code blocks, and blockquotes need a product-specific treatment.
- The same rendered Markdown appears across many screens.

## Color Downsampling

Glamour rendering is pure and may not know terminal capabilities. If the app needs terminal-aware color downsampling, run final output through the surrounding terminal-aware path, such as Lip Gloss printing, or use the project's established output abstraction.

## QA Sample

Test with Markdown containing:

````markdown
# Heading

Paragraph with a [link](https://charm.land) and a long sentence that must wrap cleanly in a narrow viewport.

- One
- Two
  - Nested

```go
fmt.Println("code")
```

> Quote with emphasis.
````

Confirm heading hierarchy, links, lists, quotes, code blocks, and wrapping remain readable.
