package form

import (
	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/style"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/huh/v2"
	"github.com/charmbracelet/lipgloss/v2"
)

type Options struct {
	style.Margin
}

type form[T any] struct {
	base     *app.Base
	opts     *Options
	style    lipgloss.Style
	onSubmit func(ctx *app.Context[T])
}

type formSubmitMsg struct {
	id string
}
type uiState struct {
	form *huh.Form
}

// Note: This needs to be in the UI tree after submissions
func New[T any](ctx *app.Context[T], huhForm *huh.Form, onSubmit func(ctx *app.Context[T]), options *Options, baseOptions ...app.BaseOption) *form[T] {
	if options == nil {
		options = &Options{}
	}
	if baseOptions == nil {
		baseOptions = []app.BaseOption{}
	}

	s := lipgloss.NewStyle()
	s = style.ApplyMargin(s, options.Margin)

	base, cleanup := app.NewBase(ctx, "form", append([]app.BaseOption{app.WithFocusable(true)}, baseOptions...)...)
	defer cleanup()

	state := app.GetUIState[T, uiState](ctx, base.ID)
	if state == nil || state.form == nil || state.form != huhForm {
		app.SetUIState(ctx, base.ID, &uiState{form: huhForm})
		huhForm.SubmitCmd = func() tea.Msg {
			return formSubmitMsg{id: base.ID}
		}

		ctx.AddCmd(huhForm.Init())
		ctx.UIState.Focused = base.ID
	}

	return &form[T]{
		base:     base,
		style:    s,
		opts:     options,
		onSubmit: onSubmit,
	}
}

func (m *form[T]) Render(ctx *app.Context[T]) string {
	return m.getState(ctx).form.View()
}

func (m *form[T]) Update(ctx *app.Context[T], msg tea.Msg) bool {

	switch msg := msg.(type) {
	case formSubmitMsg:
		if msg.id == m.base.ID {
			m.onSubmit(ctx)
			return true
		}
	}

	state := m.getState(ctx)
	newForm, cmd := state.form.Update(msg)
	newFormTyped := newForm.(*huh.Form)
	state.form = newFormTyped
	ctx.AddCmd(cmd)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "shift+tab":
			return true
		case "tab":
			return true
		}
	}
	return false
}

func (m *form[T]) Children(ctx *app.Context[T]) []app.Fc[T] {
	return nil
}
func (m *form[T]) Base() *app.Base {
	return m.base
}

func (m *form[T]) getState(ctx *app.Context[T]) *uiState {
	state := app.GetUIState[T, uiState](ctx, m.base.ID)
	return state
}
