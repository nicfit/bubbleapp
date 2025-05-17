package app

import (
	tea "github.com/charmbracelet/bubbletea/v2"
)

// KeyHandler defines the signature for component-internal key handlers.
// It returns true if the key press was handled, false otherwise.
type KeyHandler func(keyMsg tea.KeyMsg) bool

// MouseHandler defines the signature for component-internal mouse handlers.
type MouseHandler func(msg tea.MouseMsg, childID string) bool

// MsgHandler is for receiving raw tea.Msg messages.
type MsgHandler func(msg tea.Msg) tea.Cmd

// instanceContext represents an instance of a functional component (FC).
// It holds the component's ID, focusable state, function reference, props,
// event handlers, and state management for both state and effects.
type instanceContext struct {
	id                string
	focusable         bool
	states            []any
	effects           []effectRecord
	keyHandlers       []KeyHandler
	globalKeyHandlers []KeyHandler
	mouseHandlers     []MouseHandler
	messageHandlers   []MsgHandler
	onFocused         func(isReverse bool)
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

func (c *fcInstanceContext) set(id string) *instanceContext {
	instance, ok := c.ctxs[id]
	if !ok {
		instance = &instanceContext{
			id:              id,
			states:          make([]any, 0),
			effects:         make([]effectRecord, 0),
			mouseHandlers:   make([]MouseHandler, 0),
			keyHandlers:     make([]KeyHandler, 0),
			messageHandlers: make([]MsgHandler, 0),
		}
		c.ctxs[id] = instance
	}

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

func (c *fcInstanceContext) getAllGlobalKeyHandlers() []KeyHandler {
	var handlers []KeyHandler
	ids := make([]string, 0, len(c.ctxs))
	for id := range c.ctxs {
		ids = append(ids, id)
	}

	for i := len(ids) - 1; i >= 0; i-- {
		instance := c.ctxs[ids[i]]
		handlers = append(handlers, instance.globalKeyHandlers...)
	}
	return handlers
}
