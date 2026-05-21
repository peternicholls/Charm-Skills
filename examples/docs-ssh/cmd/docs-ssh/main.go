package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"

	"charm.land/glamour/v2"
	"charm.land/lipgloss/v2"
	"charm.land/wish/v2"
	"github.com/charmbracelet/ssh"
)

type serverConfig struct {
	Addr       string
	DefaultDoc string
	AllowShell bool
}

var docs = map[string]string{
	"welcome": `# Charm Skills Docs

Welcome to the SSH docs demo.

- Rendered with Glamour
- Safe by default
- Backed by embedded Markdown

Use this as a miniature shape for remote terminal documentation.`,
	"qa": `# QA Checklist

1. Render with a narrow width.
2. Verify links, lists, and code fences.
3. Test SSH locally before publishing.

` + "```bash\npython3 scripts/validate_examples.py --test\n```",
}

func topics() []string {
	names := make([]string, 0, len(docs))
	for name := range docs {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

func markdownFor(topic string) (string, bool) {
	if topic == "" {
		topic = "welcome"
	}
	doc, ok := docs[topic]
	return doc, ok
}

func renderTopic(topic string, width int) string {
	source, ok := markdownFor(topic)
	if !ok {
		source = "# Missing topic\n\nAvailable topics: " + strings.Join(topics(), ", ")
	}
	rendered, err := glamour.Render(source, "dark")
	if err != nil {
		rendered = source
	}
	if width > 0 {
		rendered = lipgloss.NewStyle().Width(width).Render(rendered)
	}
	return titleStyle.Render("Docs SSH") + "\n\n" + rendered
}

func defaultServerConfig(addr string) serverConfig {
	return serverConfig{
		Addr:       addr,
		DefaultDoc: "welcome",
		AllowShell: false,
	}
}

func serve(cfg serverConfig) error {
	if cfg.AllowShell {
		return fmt.Errorf("shell access is intentionally disabled in this demo")
	}
	server, err := wish.NewServer(
		wish.WithAddress(cfg.Addr),
		wish.WithMiddleware(func(next ssh.Handler) ssh.Handler {
			return func(s ssh.Session) {
				_, _ = s.Write([]byte(renderTopic(cfg.DefaultDoc, 80)))
			}
		}),
	)
	if err != nil {
		return err
	}
	fmt.Printf("serving docs over ssh at %s\n", cfg.Addr)
	return server.ListenAndServe()
}

var titleStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("12"))

func main() {
	renderSample := flag.Bool("render-sample", false, "render sample Markdown and exit")
	listTopics := flag.Bool("list", false, "list embedded topics and exit")
	serveFlag := flag.Bool("serve", false, "start local SSH server")
	addr := flag.String("addr", "127.0.0.1:23234", "SSH listen address")
	topic := flag.String("topic", "welcome", "topic to render")
	flag.Parse()

	switch {
	case *listTopics:
		fmt.Println(strings.Join(topics(), "\n"))
	case *renderSample:
		fmt.Println(renderTopic(*topic, 80))
	case *serveFlag:
		if err := serve(defaultServerConfig(*addr)); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		fmt.Println(renderTopic(*topic, 80))
	}
}
