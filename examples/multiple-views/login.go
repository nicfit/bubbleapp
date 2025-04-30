package main

import (
	"time"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/button"
	"github.com/alexanderbh/bubbleapp/component/loader"
	"github.com/alexanderbh/bubbleapp/component/stack"
	"github.com/alexanderbh/bubbleapp/component/text"
	"github.com/alexanderbh/bubbleapp/style"

	zone "github.com/alexanderbh/bubblezone/v2"
	tea "github.com/charmbracelet/bubbletea/v2"
)

type CustomData struct {
	UserID string
}

func NewLogin() model[CustomData] {
	ctx := &app.Context[CustomData]{
		Styles: style.DefaultStyles(),
		Zone:   zone.New(),
		Data:   &CustomData{},
	}

	loginButton := button.New(ctx, "Log in",
		button.WithVariant(button.Primary),
	)

	failButton := button.New(ctx, "Fail log in",
		button.WithVariant(button.Warning),
	)

	quitButton := button.New(ctx, "Quit App",
		button.WithVariant(button.Danger),
	)

	stackView := stack.New(ctx, stack.Options[CustomData]{
		Children: []*app.Base[CustomData]{
			text.New(ctx, "██       ██████   ██████  ██ ███    ██\n██      ██    ██ ██       ██ ████   ██\n██      ██    ██ ██   ███ ██ ██ ██  ██\n██      ██    ██ ██    ██ ██ ██  ██ ██\n███████  ██████   ██████  ██ ██   ████\n\n").Base(),
			text.New(ctx, "Log in or fail! Up to you!").Base(),
			// Put a horizontal stack here once we have it perhaps
			loginButton.Base(),
			failButton.Base(),
			quitButton.Base(),
		}},
	)

	base := app.New(ctx, app.AsRoot())
	base.AddChild(stackView.Base())

	loggingInView := stack.New(ctx, stack.Options[CustomData]{
		Children: []*app.Base[CustomData]{
			text.New(ctx, "Please wait...").Base(),
			loader.New(ctx, loader.Meter, loader.WithText("Logging in...")).Base(),
		}},
	)

	return model[CustomData]{
		base:          base,
		loggingInView: loggingInView.Base(),
		inputView:     stackView.Base(),
		failButtonID:  failButton.Base().ID,
		loginButtonID: loginButton.Base().ID,
		quitButtonID:  quitButton.Base().ID,
	}
}

type model[T CustomData] struct {
	base *app.Base[CustomData]

	loggingInView *app.Base[CustomData]
	inputView     *app.Base[CustomData]

	errorTextID   string
	failButtonID  string
	loginButtonID string
	quitButtonID  string
}

type LoginSuccessMsg struct{}

type LoginFailedMsg struct {
	Error string
}

func LoginCmd(data *CustomData, fail bool) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(2 * time.Second)
		if fail {
			return LoginFailedMsg{
				Error: "\nLogin failed. Ouch!",
			}
		}

		// Setting global state here. Could be from DB or something else.
		data.UserID = "1234abc"
		return LoginSuccessMsg{}
	}
}

func (m model[T]) Init() tea.Cmd {
	return m.base.Init()
}

func (m model[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}

	case button.ButtonPressMsg:
		switch msg.ID {
		case m.quitButtonID:
			return m, tea.Quit
		case m.failButtonID:
			if m.errorTextID != "" {
				m.base.RemoveChild(m.errorTextID) // This could be done on tea.KeyMsg as well
				m.errorTextID = ""
			}
			m.base.ReplaceChild(
				m.inputView.ID,
				m.loggingInView,
			)
			return m, LoginCmd(m.base.Ctx.Data, true)
		case m.loginButtonID:
			if m.errorTextID != "" {
				m.base.RemoveChild(m.errorTextID) // This could be done on tea.KeyMsg as well
				m.errorTextID = ""
			}
			m.base.ReplaceChild(
				m.inputView.ID,
				m.loggingInView,
			)
			return m, LoginCmd(m.base.Ctx.Data, false)
		}
	case LoginSuccessMsg:
		return NewAuthModel(m.base.Ctx), nil // Context is thrown away here. Is that what we want?
	case LoginFailedMsg:
		m.base.ReplaceChild(
			m.loggingInView.ID,
			m.inputView,
		)

		errorText := text.New(m.base.Ctx, msg.Error, text.WithColor(m.base.Ctx.Styles.Colors.Danger)) // Add variant to text for Error text
		m.errorTextID = errorText.Base().ID
		m.base.GetChildren()[0].AddChild(
			errorText.Base(),
		)
	}

	cmd = m.base.Update(msg)

	return m, cmd

}

func (m model[T]) View() string {
	return m.base.Render()
}
