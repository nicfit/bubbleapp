package app

import (
	"strings"
	"time"

	"slices"

	"github.com/alexanderbh/bubbleapp/shader"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
	"github.com/google/uuid"
)

type UIModel[T any] interface {
	View() string
	Update(msg tea.Msg) (tea.Model, tea.Cmd)
	Init() tea.Cmd
	Base() *Base[T]
}

type Base[T any] struct {
	Ctx      *Context[T]
	ID       string
	Focused  bool
	Hovered  bool
	Shader   shader.Shader
	Model    UIModel[T]
	Children []*Base[T]
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
	Shader    shader.Shader
}

type Option func(*BaseOptions)

// TODO: REWRITE THIS TO USE AN STRUCT INSTEAD OF THESE WITH FUNCTIONS. LIKE OTHER COMPONENTS
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
func WithShader(shader shader.Shader) Option {
	return func(o *BaseOptions) {
		o.Shader = shader
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

func New[T any](ctx *Context[T], opts ...Option) *Base[T] {
	options := BaseOptions{
		GrowX:     false,
		GrowY:     false,
		Focusable: false,
	}

	for _, opt := range opts {
		opt(&options)
	}

	b := &Base[T]{
		Ctx:      ctx,
		ID:       uuid.New().String(),
		Focused:  false,
		Opts:     options,
		Children: []*Base[T]{},
		Shader:   options.Shader,
	}

	return b
}

func (m *Base[T]) Init() tea.Cmd {
	cmds := []tea.Cmd{}
	if len(m.Children) > 0 {
		for _, child := range m.Children {
			cmds = append(cmds, child.Model.Init())
		}
	}
	if m.Opts.TickFPS > 0 {
		cmds = append(cmds, m.tick())
	}
	if m.Opts.IsRoot {
		// This could be replaced with a RequestWindowSize when that works
		cmds = append(cmds, func() tea.Msg {
			return tea.WindowSizeMsg{
				Width:  m.Ctx.Width,
				Height: m.Ctx.Height,
			}
		}, tea.RequestBackgroundColor, tea.RequestForegroundColor)
		cmds = append(cmds, m.Ctx.FocusFirstCmd(m.Children[0]))
	}
	return tea.Batch(cmds...)
}
func (m *Base[T]) Update(msg tea.Msg) tea.Cmd {
	var (
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.Opts.IsRoot {
			switch msg.String() {
			case "tab":
				return m.Ctx.FocusNextCmd(m.Children[0])
			case "shift+tab":
				return m.Ctx.FocusPrevCmd(m.Children[0])
			}
		}
	case FocusComponentMsg:
		targetIsSelf := m.ID == msg.TargetID
		if targetIsSelf {

			if m.Opts.Focusable {
				m.Focused = true // Target should be focused
			} else {
				if len(m.Children) > 0 {
					cmds = append(cmds, sendFocusMsg(m.Children[0].ID))
				}
			}

		} else {
			m.Focused = false
		}

	case BlurAllMsg:
		m.Focused = false
	case tea.BackgroundColorMsg:
		// Not working yet it seems
		if m.Opts.IsRoot {
			m.Ctx.BackgroundColor = msg.Color
		}
		return nil
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
		if ds, ok := m.Shader.(shader.DynamicShader); ok {
			ds.Tick()
		}
		if m.Opts.TickFPS > 0 {
			cmds = append(cmds, m.tick())
		}
	}

	// For each child, update and collect commands
	if len(m.Children) > 0 {
		newChildren := make([]*Base[T], len(m.Children))
		for i, child := range m.Children {
			updatedChild, cmd := child.Model.Update(msg)
			typedChild := updatedChild.(UIModel[T])
			typedChild.Base().Model = typedChild // Update the UIModel on the *Base
			child = updatedChild.(UIModel[T]).Base()
			cmds = append(cmds, cmd)
			newChildren[i] = child
		}
		m.Children = newChildren
	}

	return tea.Batch(cmds...)
}

func (base *Base[T]) Render() string {
	elements := []string{}

	for _, child := range base.Children {
		elements = append(elements, child.Model.View())
	}
	result := strings.Join(elements, "\n")

	if base.Opts.IsRoot {
		return base.Ctx.Zone.Scan(result)
	}
	return result
}

func (base *Base[T]) ApplyShader(input string) string {
	if base.Shader != nil {
		return base.Shader.Render(input, nil)
	}
	return input
}
func (base *Base[T]) ApplyShaderWithStyle(input string, style lipgloss.Style) string {
	if base.Shader != nil {
		return base.Shader.Render(input, &style)
	}
	return input
}

func (base *Base[T]) AddChild(child *Base[T]) {
	base.Children = append(base.Children, child)
}

func (base *Base[T]) AddChildren(children ...*Base[T]) {
	base.Children = append(base.Children, children...)
}

func (base *Base[T]) RemoveChild(ID string) bool {
	for i, c := range base.Children {
		if c.ID == ID {
			base.Children = slices.Delete(base.Children, i, i+1)
			return true
		}
	}
	for _, c := range base.Children {
		rec := c.RemoveChild(ID)
		if rec {
			return true
		}
	}
	return false
}
func (base *Base[T]) ReplaceChild(ID string, new *Base[T]) {
	for i, c := range base.Children {
		if c.ID == ID {
			base.Children[i] = new
			break
		}
	}
}

func (base *Base[T]) GetChild(id string) *Base[T] {
	for _, child := range base.Children {
		if child.ID == id {
			return child
		}
	}
	for _, child := range base.Children {
		rec := child.GetChild(id)
		if rec != nil {
			return rec
		}
	}
	return nil
}
