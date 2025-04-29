package app

import (
	"strings"
	"time"

	"slices"

	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/google/uuid"
)

type UIModel interface {
	View() string
	Update(msg tea.Msg) (tea.Model, tea.Cmd)
	Init() tea.Cmd
	Base() *Base
}

type Base struct {
	Ctx      *Context
	ID       string
	Focused  bool
	Hovered  bool
	Children []UIModel
	Width    int
	Height   int
	Opts     BaseOptions
}

type BaseOptions struct {
	TickFPS   time.Duration
	GrowX     bool
	GrowY     bool
	Focusable bool
	IsRoot    bool
}

type Option func(*BaseOptions)

// Make sure to only have 1 base model with Tick FPS
func WithTick(fps time.Duration) Option {
	return func(o *BaseOptions) {
		o.TickFPS = fps
	}
}
func WithGrowX(grow bool) Option {
	return func(o *BaseOptions) {
		o.GrowX = grow
	}
}
func WithGrowY(grow bool) Option {
	return func(o *BaseOptions) {
		o.GrowY = grow
	}
}

func WithGrow(grow bool) Option {
	return func(o *BaseOptions) {
		o.GrowX = grow
		o.GrowY = grow
	}
}

func WithFocusable(focusable bool) Option {
	return func(o *BaseOptions) {
		o.Focusable = focusable
	}
}

// Is this really the best way? Is the root really special?
func AsRoot() Option {
	return func(o *BaseOptions) {
		o.IsRoot = true
		o.TickFPS = time.Second / 12
	}
}

func New(ctx *Context, opts ...Option) *Base {
	options := BaseOptions{
		GrowX:     false,
		GrowY:     false,
		Focusable: false,
	}

	for _, opt := range opts {
		opt(&options)
	}

	return &Base{
		Ctx:      ctx,
		ID:       uuid.New().String(),
		Focused:  false,
		Opts:     options,
		Children: []UIModel{},
	}
}

func (m Base) Init() tea.Cmd {
	cmds := []tea.Cmd{}
	if m.GetChildren() != nil && len(m.GetChildren()) > 0 {
		for _, child := range m.GetChildren() {
			cmds = append(cmds, child.Init())
		}
	}
	if m.Opts.TickFPS > 0 {
		cmds = append(cmds, m.tick())
	}
	if m.Opts.IsRoot {
		cmds = append(cmds, m.Ctx.FocusManager.FocusFirstCmd(m.GetChildren()[0]))
	}
	return tea.Batch(cmds...)
}
func (m *Base) Update(msg tea.Msg) tea.Cmd {
	var (
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.Opts.IsRoot {
			switch msg.String() {
			case "tab":
				return m.Ctx.FocusManager.FocusNextCmd(m.GetChildren()[0])
			case "shift+tab":
				return m.Ctx.FocusManager.FocusPrevCmd(m.GetChildren()[0])
			}
		}
	case FocusComponentMsg:
		targetIsSelf := m.ID == msg.TargetID
		if targetIsSelf {

			if m.Opts.Focusable {
				m.Focused = true // Target should be focused
			} else {
				if m.GetChildren() != nil && len(m.GetChildren()) > 0 {
					cmds = append(cmds, sendFocusMsg(m.GetChildren()[0].Base().ID))
				}
			}

		} else {
			m.Focused = false
		}

	case BlurAllMsg:
		m.Focused = false

	case tea.WindowSizeMsg:
		m.Height = msg.Height
		m.Width = msg.Width
		if m.Opts.IsRoot {
			m.Ctx.Width = msg.Width
			m.Ctx.Height = msg.Height
		}

	case tea.MouseMsg:
		switch msg := msg.(type) {
		case tea.MouseMotionMsg:
			if m.Ctx.Zone.Get(m.ID).InBounds(msg) {
				m.Hovered = true
			} else {
				m.Hovered = false
			}
		}

	case TickMsg:
		if m.Opts.TickFPS > 0 {
			cmds = append(cmds, m.tick())
		}
	}

	// For each child, update and collect commands
	if m.GetChildren() != nil && len(m.GetChildren()) > 0 {
		newChildren := make([]UIModel, len(m.GetChildren()))
		for i, child := range m.GetChildren() {
			updatedChild, cmd := child.Update(msg)
			child = updatedChild.(UIModel)
			cmds = append(cmds, cmd)
			newChildren[i] = child
		}
		m.Children = newChildren
	}

	return tea.Batch(cmds...)
}

func (base *Base) View() string {
	children := []string{}
	for _, child := range base.Children {
		children = append(children, child.View())
	}
	result := strings.Join(children, "\n")

	if base.Opts.IsRoot {
		return base.Ctx.Zone.Scan(result)
	}
	return result
}

func (base *Base) AddChild(child UIModel) {
	base.Children = append(base.Children, child)
}

func (base *Base) AddChildren(children ...UIModel) {
	base.Children = append(base.Children, children...)
}

func (base *Base) RemoveChild(ID string) bool {
	for i, c := range base.Children {
		if c.Base().ID == ID {
			base.Children = slices.Delete(base.Children, i, i+1)
			return true
		}
	}
	for _, c := range base.Children {
		rec := c.Base().RemoveChild(ID)
		if rec {
			return true
		}
	}
	return false
}
func (base *Base) ReplaceChild(ID string, new UIModel) {
	for i, c := range base.Children {
		if c.Base().ID == ID {
			base.Children[i] = new
			break
		}
	}
}

func (base *Base) GetChild(id string) UIModel {
	for _, child := range base.Children {
		if child.Base().ID == id {
			return child
		}
	}
	for _, child := range base.Children {
		rec := child.Base().GetChild(id)
		if rec != nil {
			return rec
		}
	}
	return nil
}

func (fc *Base) GetChildren() []UIModel {
	return fc.Children
}
