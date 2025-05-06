package main

import (
	"os"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/divider"
	"github.com/alexanderbh/bubbleapp/component/form"
	"github.com/alexanderbh/bubbleapp/component/stack"
	"github.com/alexanderbh/bubbleapp/component/text"

	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/huh/v2"
	"github.com/charmbracelet/lipgloss/v2"
)

type CustomData struct {
	email    string
	password string
	remember string
}

var loginForm = huh.NewForm(
	huh.NewGroup(
		huh.NewInput().Key("email").Title("Email"),
		huh.NewInput().Key("password").Title("Password").EchoMode(huh.EchoModePassword),
		huh.NewSelect[string]().Key("rememberme").Title("Remember me").Description("Log in automatically when using this SSH key").Options(huh.NewOptions("Yes", "No")...),
	),
).WithTheme(huh.ThemeFunc(huh.ThemeDracula)).WithShowHelp(false)

func NewRoot(ctx *app.Context[CustomData]) app.Fc[CustomData] {

	return stack.New(ctx, func(ctx *app.Context[CustomData]) []app.Fc[CustomData] {
		view := []app.Fc[CustomData]{
			text.New(ctx, loginLogo(ctx), nil),
			divider.New(ctx),
			form.New(ctx, loginForm, func(ctx *app.Context[CustomData]) {
				ctx.Data.email = loginForm.GetString("email")
				ctx.Data.password = loginForm.GetString("password")
				ctx.Data.remember = loginForm.GetString("rememberme")
				ctx.Update()
			}, nil),
		}

		if ctx.Data.email != "" {
			view = append(view, text.New(ctx, "Email: "+ctx.Data.email, nil))
			view = append(view, text.New(ctx, "Password 🙈: "+ctx.Data.password, nil))
			view = append(view, text.New(ctx, "Remember me: "+ctx.Data.remember, nil))
		}

		return view
	}, nil)
}

func main() {
	ctx := app.NewContext(&CustomData{})

	app := app.NewApp(ctx, NewRoot)
	p := tea.NewProgram(app, tea.WithAltScreen(), tea.WithMouseAllMotion())
	app.SetTeaProgram(p)
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}

var loginLogo = func(ctx *app.Context[CustomData]) string {
	f := lipgloss.NewStyle().Foreground(ctx.Styles.Colors.Primary)
	b := lipgloss.NewStyle().Foreground(ctx.Styles.Colors.GhostDark)
	return f.Render("██") + b.Render("╗      ") + f.Render("██████") + b.Render("╗  ") + f.Render("██████") + b.Render("╗ ") + f.Render("██") + b.Render("╗") + f.Render("███") + b.Render("╗   ") + f.Render("██") + b.Render("╗") + "\n" +
		f.Render("██") + b.Render("║     ") + f.Render("██") + b.Render("╔═══") + f.Render("██") + b.Render("╗") + f.Render("██") + b.Render("╔════╝ ") + f.Render("██") + b.Render("║") + f.Render("████") + b.Render("╗  ") + f.Render("██") + b.Render("║") + "\n" +
		f.Render("██") + b.Render("║     ") + f.Render("██") + b.Render("║   ") + f.Render("██") + b.Render("║") + f.Render("██") + b.Render("║  ") + f.Render("███") + b.Render("╗") + f.Render("██") + b.Render("║") + f.Render("██") + b.Render("╔") + f.Render("██") + b.Render("╗ ") + f.Render("██") + b.Render("║") + "\n" +
		f.Render("██") + b.Render("║     ") + f.Render("██") + b.Render("║   ") + f.Render("██") + b.Render("║") + f.Render("██") + b.Render("║   ") + f.Render("██") + b.Render("║") + f.Render("██") + b.Render("║") + f.Render("██") + b.Render("║╚") + f.Render("██") + b.Render("╗") + f.Render("██") + b.Render("║") + "\n" +
		f.Render("███████") + b.Render("╗╚") + f.Render("██████") + b.Render("╔╝╚") + f.Render("██████") + b.Render("╔╝") + f.Render("██") + b.Render("║") + f.Render("██") + b.Render("║ ╚") + f.Render("████") + b.Render("║") + "\n" +
		b.Render("╚══════╝ ╚═════╝  ╚═════╝ ╚═╝╚═╝  ╚═══╝")
}
