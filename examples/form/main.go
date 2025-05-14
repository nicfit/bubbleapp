package main

import (
	"os"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/box"
	"github.com/alexanderbh/bubbleapp/component/button"
	"github.com/alexanderbh/bubbleapp/component/divider"
	"github.com/alexanderbh/bubbleapp/component/form"
	"github.com/alexanderbh/bubbleapp/component/stack"
	"github.com/alexanderbh/bubbleapp/component/text"

	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/huh/v2"
	"github.com/charmbracelet/lipgloss/v2"
)

type FormData struct {
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

func NewRoot(c *app.Ctx, _ app.Props) string {
	formSubmit, setFormSubmit := app.UseState[*FormData](c, nil)

	return stack.New(c, func(c *app.Ctx) {
		c.Render(loginLogo, nil)

		if formSubmit == nil {
			form.New(c, loginForm, func() {
				setFormSubmit(&FormData{
					email:    loginForm.GetString("email"),
					password: loginForm.GetString("password"),
					remember: loginForm.GetString("rememberme"),
				})
			})
		}

		if formSubmit != nil {
			text.New(c, "Email: "+formSubmit.email, nil)
			text.New(c, "Password 🙈: "+formSubmit.password, nil)
			text.New(c, "Remember me: "+formSubmit.remember, nil)
		}

		box.NewEmpty(c)
		divider.New(c)
		button.New(c, "Quit", c.Quit, button.WithVariant(button.Danger))
	})
}

func main() {
	c := app.NewCtx()

	bubbleApp := app.New(c, NewRoot)
	p := tea.NewProgram(bubbleApp, tea.WithAltScreen(), tea.WithMouseAllMotion())
	bubbleApp.SetTeaProgram(p)
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}

func loginLogo(c *app.Ctx, _ app.Props) string {
	f := lipgloss.NewStyle().Foreground(c.Styles.Colors.Secondary)
	b := lipgloss.NewStyle().Foreground(c.Styles.Colors.GhostDark)
	return f.Render("██") + b.Render("╗      ") + f.Render("██████") + b.Render("╗  ") + f.Render("██████") + b.Render("╗ ") + f.Render("██") + b.Render("╗") + f.Render("███") + b.Render("╗   ") + f.Render("██") + b.Render("╗") + "\n" +
		f.Render("██") + b.Render("║     ") + f.Render("██") + b.Render("╔═══") + f.Render("██") + b.Render("╗") + f.Render("██") + b.Render("╔════╝ ") + f.Render("██") + b.Render("║") + f.Render("████") + b.Render("╗  ") + f.Render("██") + b.Render("║") + "\n" +
		f.Render("██") + b.Render("║     ") + f.Render("██") + b.Render("║   ") + f.Render("██") + b.Render("║") + f.Render("██") + b.Render("║  ") + f.Render("███") + b.Render("╗") + f.Render("██") + b.Render("║") + f.Render("██") + b.Render("╔") + f.Render("██") + b.Render("╗ ") + f.Render("██") + b.Render("║") + "\n" +
		f.Render("██") + b.Render("║     ") + f.Render("██") + b.Render("║   ") + f.Render("██") + b.Render("║") + f.Render("██") + b.Render("║   ") + f.Render("██") + b.Render("║") + f.Render("██") + b.Render("║") + f.Render("██") + b.Render("║╚") + f.Render("██") + b.Render("╗") + f.Render("██") + b.Render("║") + "\n" +
		f.Render("███████") + b.Render("╗╚") + f.Render("██████") + b.Render("╔╝╚") + f.Render("██████") + b.Render("╔╝") + f.Render("██") + b.Render("║") + f.Render("██") + b.Render("║ ╚") + f.Render("████") + b.Render("║") + "\n" +
		b.Render("╚══════╝ ╚═════╝  ╚═════╝ ╚═╝╚═╝  ╚═══╝")
}
