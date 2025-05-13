package app

import (
	tea "github.com/charmbracelet/bubbletea/v2"
)

// KeyHandler defines the signature for component-internal key handlers.
// It returns true if the key press was handled, false otherwise.
type KeyHandler func(keyMsg tea.KeyMsg) bool

// MouseHandler defines the signature for component-internal mouse handlers.
type MouseHandler func(msg tea.MouseMsg) bool

// instanceContext represents an instance of a functional component (FC).
// It holds the component's ID, focusable state, function reference, props,
// event handlers, and state management for both state and effects.
type instanceContext struct {
	id            string
	focusable     bool
	fc            FC
	props         Props
	states        []any
	effects       []effectRecord
	keyHandlers   []KeyHandler
	mouseHandlers []MouseHandler
}

type fcInstanceContext struct {
	ctxs map[string]*instanceContext
}

func newInstanceContext() *fcInstanceContext {
	return &fcInstanceContext{
		ctxs: make(map[string]*instanceContext),
	}
}

func (c *fcInstanceContext) get(id string) (*instanceContext, bool) {
	instance, ok := c.ctxs[id]
	return instance, ok
}

func (c *fcInstanceContext) set(id string, fc FC, props Props) *instanceContext {
	instance, ok := c.ctxs[id]
	if !ok {
		instance = &instanceContext{
			id:            id,
			states:        make([]any, 0),
			effects:       make([]effectRecord, 0),
			mouseHandlers: make([]MouseHandler, 0),
			keyHandlers:   make([]KeyHandler, 0),
		}
		c.ctxs[id] = instance
	}
	instance.fc = fc
	instance.props = props

	return instance
}

func (c *fcInstanceContext) cleanupEffects(removedIDs []string) {
	for _, id := range removedIDs {
		if instance, ok := c.ctxs[id]; ok {
			for i := range instance.effects {
				if instance.effects[i].cleanupFn != nil {
					instance.effects[i].cleanupFn()
					instance.effects[i].cleanupFn = nil
				}
			}
		}
		delete(c.ctxs, id)
	}
}

func (c *fcInstanceContext) getAllIDs() []string {
	ids := make([]string, 0, len(c.ctxs))
	for id := range c.ctxs {
		ids = append(ids, id)
	}
	return ids
}
