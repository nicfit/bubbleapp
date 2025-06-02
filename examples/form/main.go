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
	"github.com/alexanderbh/bubbleapp/style"

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
		huh.NewSelect[string]().Key("rememberme").
			Title("Remember me").
			Description("Log in automatically when using this SSH key").
			Options(huh.NewOptions("Yes", "No")...),
	),
).WithTheme(huh.ThemeFunc(huh.ThemeDracula)).WithShowHelp(false)

func NewRoot(c *app.Ctx) *app.C {
	formSubmit, setFormSubmit := app.UseState[*FormData](c, nil)

	return stack.New(c, func(c *app.Ctx) []*app.C {
		cs := make([]*app.C, 0)
		cs = append(cs, c.Render(loginLogo, nil))

		if formSubmit == nil {
			cs = append(cs, form.New(c, loginForm, func() {
				setFormSubmit(&FormData{
					email:    loginForm.GetString("email"),
					password: loginForm.GetString("password"),
					remember: loginForm.GetString("rememberme"),
				})
			}))
		} else if formSubmit != nil {
			cs = append(cs,
				text.New(c, "Email: "+formSubmit.email, nil),
				text.New(c, "Password 🙈: "+formSubmit.password, nil),
				text.New(c, "Remember me: "+formSubmit.remember, nil),
			)
		}

		return append(cs,
			box.NewEmpty(c),
			divider.New(c),
			button.New(c, "Quit", c.Quit, button.WithVariant(style.Danger)),
		)
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
	f := lipgloss.NewStyle().Foreground(c.Theme.Colors.Secondary)
	b := lipgloss.NewStyle().Foreground(c.Theme.Colors.Base700)
	return f.Render("██") + b.Render("╗      ") + f.Render("██████") + b.Render("╗  ") + f.Render("██████") + b.Render("╗ ") + f.Render("██") + b.Render("╗") + f.Render("███") + b.Render("╗   ") + f.Render("██") + b.Render("╗") + "\n" +
		f.Render("██") + b.Render("║     ") + f.Render("██") + b.Render("╔═══") + f.Render("██") + b.Render("╗") + f.Render("██") + b.Render("╔════╝ ") + f.Render("██") + b.Render("║") + f.Render("████") + b.Render("╗  ") + f.Render("██") + b.Render("║") + "\n" +
		f.Render("██") + b.Render("║     ") + f.Render("██") + b.Render("║   ") + f.Render("██") + b.Render("║") + f.Render("██") + b.Render("║  ") + f.Render("███") + b.Render("╗") + f.Render("██") + b.Render("║") + f.Render("██") + b.Render("╔") + f.Render("██") + b.Render("╗ ") + f.Render("██") + b.Render("║") + "\n" +
		f.Render("██") + b.Render("║     ") + f.Render("██") + b.Render("║   ") + f.Render("██") + b.Render("║") + f.Render("██") + b.Render("║   ") + f.Render("██") + b.Render("║") + f.Render("██") + b.Render("║") + f.Render("██") + b.Render("║╚") + f.Render("██") + b.Render("╗") + f.Render("██") + b.Render("║") + "\n" +
		f.Render("███████") + b.Render("╗╚") + f.Render("██████") + b.Render("╔╝╚") + f.Render("██████") + b.Render("╔╝") + f.Render("██") + b.Render("║") + f.Render("██") + b.Render("║ ╚") + f.Render("████") + b.Render("║") + "\n" +
		b.Render("╚══════╝ ╚═════╝  ╚═════╝ ╚═╝╚═╝  ╚═══╝")
}
