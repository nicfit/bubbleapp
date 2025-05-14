package form

import (
	"strings"
	"unicode"

	"github.com/alexanderbh/bubbleapp/app"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/huh/v2"
)

type Props struct {
	HuhForm  *huh.Form
	OnSubmit func() // Can this somehow pass the form data? Or should values be set from parent?
	app.Layout
}

type prop func(*Props)

func New(c *app.Ctx, huhForm *huh.Form, onSubmit func(), opts ...prop) string {
	appliedProps := Props{
		HuhForm:  huhForm,
		OnSubmit: onSubmit,
	}

	for _, opt := range opts {
		if opt != nil {
			opt(&appliedProps)
		}
	}

	return c.Render(form, appliedProps)
}

func form(c *app.Ctx, props app.Props) string {
	p, ok := props.(Props)
	if !ok {
		panic("Form: incorrect props type")
	}

	_ = app.UseIsFocused(c)

	activeForm, setActiveForm := app.UseState(c, p.HuhForm)

	app.UseOnFocused(c, func(isReverse bool) {
		if activeForm == nil {
			return
		}
		var cmds []tea.Cmd
		cmds = append(cmds, activeForm.Init())

		// This will handle the one case where user is tabbing forward
		// and the form is on the last field. User continues to tab forward
		// and wraps around to the form again. This will move the focus
		// to the first field.

		// Unfortunately is it not possible with the current API to do the same
		// for the reverse case. If I move to NextField it will end up Submitting
		// the form. And it is not possible to know from the outside where the focus
		// is in the form.
		if !isReverse {
			for range 20 {
				cmds = append(cmds, activeForm.PrevField())
			}
		}
		c.ExecuteCmd(tea.Batch(cmds...))
	})

	app.UseEffect(c, func() {
		if p.HuhForm != activeForm {
			setActiveForm(p.HuhForm)
		}
	}, []any{p.HuhForm})

	app.UseEffect(c, func() {
		if activeForm == nil {
			return
		}

		activeForm.SubmitCmd = func() tea.Msg {
			if p.OnSubmit != nil {
				p.OnSubmit()
			}
			c.Update()
			return nil
		}

		initCmd := activeForm.Init()
		if initCmd != nil {
			c.ExecuteCmd(initCmd)
		}

		currentInstanceID := app.UseID(c)
		c.UIState.Focused = currentInstanceID

	}, []any{activeForm, p.OnSubmit})

	app.UseMsgHandler(c, func(msg tea.Msg) tea.Cmd {
		if activeForm == nil {
			return nil
		}

		if updatedFormModel, cmd := activeForm.Update(msg); cmd != nil {
			if updatedTypedForm, ok := updatedFormModel.(*huh.Form); ok {
				setActiveForm(updatedTypedForm)
			}
			return cmd
		}
		return nil
	})

	app.UseKeyHandler(c, func(msg tea.KeyMsg) bool {
		if activeForm == nil {
			return false
		}
		switch msg.String() {
		case "ctrl+c": // Do not exit on ctrl+c (part of default keymap for huh.Form for some reason)
			return false
		case "shift+tab":
			beforeTab := activeForm.GetFocusedField().GetKey()
			activeForm.PrevField()
			afterTab := activeForm.GetFocusedField().GetKey()
			if beforeTab == afterTab {
				return false
			}
			activeForm.NextField()
		case "tab":
			beforeTab := activeForm.GetFocusedField().GetKey()
			activeForm.NextField()
			afterTab := activeForm.GetFocusedField().GetKey()
			if beforeTab == afterTab {
				return false
			}
			activeForm.PrevField()
		}

		updatedFormModel, cmd := activeForm.Update(msg)

		if updatedTypedForm, ok := updatedFormModel.(*huh.Form); ok {
			setActiveForm(updatedTypedForm)
		}

		if cmd != nil {
			c.ExecuteCmd(cmd)
		}
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
	})

	if activeForm == nil {
		return ""
	}
	view := activeForm.View()

	// Remove the last newline from the form output. huhForm has a new line at the end unfortunately.
	lastNewlineIdx := strings.LastIndexByte(view, '\n')
	if lastNewlineIdx == -1 {
		return view
	}
	preNewlinePart := view[:lastNewlineIdx]
	postNewlinePart := view[lastNewlineIdx+1:]
	trimmedPreNewlinePart := strings.TrimRightFunc(preNewlinePart, func(r rune) bool {
		return r != '\n' && unicode.IsSpace(r)
	})
	return trimmedPreNewlinePart + postNewlinePart
}
