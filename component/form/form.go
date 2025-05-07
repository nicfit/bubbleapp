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

// !Work in progress!
//
// huh.form assumes it is the only thing running so it does not play well with other things.
// The plan is to have a composable form, but for now to use huh.form it works best when it
// is the only thing running in the UI tree.
//
// Note: This needs to be in the UI tree after submission to access the state/values
// Note: Only supports default KeyMap for now. To support custom keymaps the keymap
// needs to be passed in as a parameter to the form. This is because the keymap is private in the form.
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
		// Assume the form is the main thing and force focus on it
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
	state := m.getState(ctx)

	switch msg := msg.(type) {
	case formSubmitMsg:
		if msg.id == m.base.ID {
			m.onSubmit(ctx)
			return true
		}
	case tea.KeyMsg:
		// Hardcoded to DefaultKeymap from huh for now
		switch msg.String() {
		case "ctrl+c": // Do not exit on ctrl+c (part of default keymap for huh.Form for some reason)
			return false
		case "shift+tab":
			beforeTab := state.form.GetFocusedField().GetKey()
			state.form.PrevField()
			afterTab := state.form.GetFocusedField().GetKey()
			if beforeTab == afterTab {
				//state.form.NextField()
				//state.form.PrevField()
				return false
			}
			state.form.NextField()
		case "tab":
			beforeTab := state.form.GetFocusedField().GetKey()
			state.form.NextField()
			afterTab := state.form.GetFocusedField().GetKey()
			if beforeTab == afterTab {
				//state.form.PrevField()
				//state.form.NextField()
				return false
			}
			state.form.PrevField()
		}
	}

	newForm, cmd := state.form.Update(msg)
	newFormTyped := newForm.(*huh.Form)
	state.form = newFormTyped
	ctx.AddCmd(cmd)

	// Return if key was handled or not
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
