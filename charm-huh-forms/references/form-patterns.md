# Huh Form Patterns

## Import

Use:

```go
import "charm.land/huh/v2"
```

## Basic Shape

```go
type answers struct {
    name     string
    region   string
    features []string
    confirm  bool
}

var a answers

form := huh.NewForm(
    huh.NewGroup(
        huh.NewInput().
            Title("Project name").
            Value(&a.name).
            Validate(validateProjectName),
        huh.NewSelect[string]().
            Title("Region").
            Options(
                huh.NewOption("US East", "us-east"),
                huh.NewOption("EU West", "eu-west"),
            ).
            Value(&a.region),
    ),
)

if err := form.Run(); err != nil {
    return err
}
```

Keep labels user-facing and option values machine-facing.

## Field Selection

- Use `Input` for short strings.
- Use `Text` for multiline content.
- Use `Select` for one choice.
- Use `MultiSelect` for several choices, with `Limit` when the domain has a real cap.
- Use `Confirm` for yes/no decisions, especially before irreversible work.
- Use standalone `Run` on a single field for quick prompts.

## Dynamic Forms

Use dynamic field functions when later choices depend on earlier answers. Bind recomputation to the specific value that drives the change.

Avoid network calls inside dynamic functions unless cached or already available. Preload options with a Bubble Tea or CLI step when the source is slow or fallible.

## Accessible Mode

Expose accessible mode through config or environment:

```go
accessibleMode := os.Getenv("ACCESSIBLE") != ""
form.WithAccessible(accessibleMode)
```

Use simple titles and validation messages that make sense without spatial context. Do not depend on color, cursor position, or visual grouping alone.

## Validation

Validation should be fast, deterministic, and local when possible. For validation that needs I/O, prefer validating after submission with a clear recovery path.

Good validation messages:

- Name the field.
- Explain the constraint.
- Say what to do next.

## UX Checklist

- Defaults are useful and not dangerous.
- Destructive confirmations default to safe answers.
- Field order matches user intent, not data model order.
- Multi-select limits are explained in the title or description.
- Dynamic choices update predictably.
- Accessible mode has been run manually.
