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

func NewLogin() model {
	ctx := &app.Context{
		Styles:       style.DefaultStyles(),
		FocusManager: app.NewFocusManager(),
		Zone:         zone.New(),
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

	stackView := stack.New(ctx)
	stackView.AddChildren(
		text.New(ctx, `██       ██████   ██████  ██ ███    ██ 
██      ██    ██ ██       ██ ████   ██ 
██      ██    ██ ██   ███ ██ ██ ██  ██ 
██      ██    ██ ██    ██ ██ ██  ██ ██ 
███████  ██████   ██████  ██ ██   ████`+"\n\n"),
		text.New(ctx, "Log in or fail! Up to you!"),
		// Put a horizontal stack here once we have it perhaps
		loginButton,
		failButton,
		quitButton,
	)

	base := app.New(ctx, app.AsRoot())
	base.AddChild(stackView)

	loggingInView := stack.New(ctx)
	loggingInView.AddChildren(
		text.New(ctx, "Please wait..."),
		loader.New(ctx, loader.Meter, loader.WithText("Logging in...")),
	)

	return model{
		base:          base,
		loggingInView: loggingInView,
		inputView:     stackView,
		failButtonID:  failButton.Base().ID,
		loginButtonID: loginButton.Base().ID,
		quitButtonID:  quitButton.Base().ID,
	}
}

type model struct {
	base *app.Base

	loggingInView app.UIModel
	inputView     app.UIModel

	errorTextID   string
	failButtonID  string
	loginButtonID string
	quitButtonID  string
}

type LoginSuccessMsg struct {
	UserID string
}

type LoginFailedMsg struct {
	Error string
}

func LoginCmd(m model, fail bool) tea.Cmd {
	return func() tea.Msg {
		if m.errorTextID != "" {
			m.base.RemoveChild(m.errorTextID) // This could be done on tea.KeyMsg as well
			m.errorTextID = ""
		}
		m.base.ReplaceChild(
			m.inputView.Base().ID,
			m.loggingInView,
		)

		time.Sleep(2 * time.Second)
		if fail {
			return LoginFailedMsg{
				Error: "Login failed by chance",
			}
		}

		return LoginSuccessMsg{
			UserID: "1234abc",
		}
	}
}

func (m model) Init() tea.Cmd {
	return m.base.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			return m, LoginCmd(m, true)
		case m.loginButtonID:
			return m, LoginCmd(m, false)
		}
	case LoginSuccessMsg:
		return NewAuthModel(msg.UserID), nil // Context is thrown away here. Is that what we want?
	case LoginFailedMsg:
		m.base.ReplaceChild(
			m.loggingInView.Base().ID,
			m.inputView,
		)

		errorText := text.New(m.base.Ctx, msg.Error) // Add variant to text for Error text
		m.errorTextID = errorText.Base().ID
		m.base.GetChildren()[0].Base().AddChild(
			errorText,
		)
	}

	cmd = m.base.Update(msg)

	return m, cmd

}

func (m model) View() string {
	return m.base.View()
}
