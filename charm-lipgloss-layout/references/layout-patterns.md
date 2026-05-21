# Lip Gloss Layout Patterns

## Import

Use:

```go
import "charm.land/lipgloss/v2"
```

Use subpackages when appropriate:

```go
import "charm.land/lipgloss/v2/table"
import "charm.land/lipgloss/v2/list"
import "charm.land/lipgloss/v2/tree"
```

## Semantic Styles

Prefer a small semantic style set:

```go
var (
    titleStyle = lipgloss.NewStyle().Bold(true)
    mutedStyle = lipgloss.NewStyle().Faint(true)
    errorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Bold(true)
    activeStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("12")).Bold(true)
)
```

Do not encode component names into styles when state names are clearer. `activeStyle` survives refactors better than `selectedListRowPurpleStyle`.

## Measuring And Fitting

Use Lip Gloss measurement helpers:

```go
w := lipgloss.Width(rendered)
h := lipgloss.Height(rendered)
```

Use explicit constraints:

```go
panel := lipgloss.NewStyle().
    Width(max(20, availableWidth)).
    MaxWidth(availableWidth).
    Render(content)
```

When composing views, compute available width after margins, borders, and side panels. Do not hard-code a table width and hope wrapping works.

## Responsive Composition

Use a breakpoint-like rule:

- Under 70 columns: stack vertically, hide decorative sidebars, use short help.
- 70-110 columns: show primary and secondary regions, truncate tertiary metadata.
- Over 110 columns: show sidebars or richer metadata if they improve scanning.

Example:

```go
if m.width < 70 {
    return lipgloss.JoinVertical(lipgloss.Left, header, main, footer)
}
return lipgloss.JoinHorizontal(lipgloss.Top, sidebar, main)
```

## Color Discipline

Use ANSI colors for broad compatibility unless the project already uses hex colors and tests color downsampling. Always ensure errors, warnings, success, and selection remain distinguishable in monochrome.

When runtime background detection is acceptable:

```go
hasDarkBG := lipgloss.HasDarkBackground(os.Stdin, os.Stdout)
choose := lipgloss.LightDark(hasDarkBG)
accent := choose(lipgloss.Color("#005F87"), lipgloss.Color("#7DD3FC"))
```

Keep I/O out of pure render helpers. Pass computed colors or theme into them.

## Tables, Lists, Trees

Use `table` for static/tabular output and Bubbles `table` for interactive tables.

Use `list` and `tree` for presentational output. Use Bubbles `list` for interactive browsing, filtering, pagination, and help.

## Visual QA Checklist

- Active focus can be identified without color.
- No text overlaps, disappears, or wraps into controls at narrow width.
- Help text truncates or switches to short mode instead of pushing core content away.
- Error and success states are visually distinct and include recovery/next action.
- Borders align around wide Unicode and ANSI-styled text.
