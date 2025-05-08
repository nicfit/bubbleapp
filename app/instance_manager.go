package app

import (
	"reflect"
	"strings"
)

// instanceContext represents an instance of a functional component (FC).
// It holds the component's ID, focusable state, function reference, props,
// event handlers, and state management for both state and effects.
type instanceContext struct {
	id        string
	focusable bool
	fc        FC
	props     Props
	handlers  map[string]interface{}
	States    []any          // Added for UseState
	Effects   []effectRecord // New field for UseEffect states
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
			id:       id,
			handlers: make(map[string]interface{}),
			States:   make([]any, 0),
			Effects:  make([]effectRecord, 0),
		}
		c.ctxs[id] = instance
	}
	// Always update fc and props
	instance.fc = fc
	instance.props = props

	// Re-extract handlers
	instance.handlers = make(map[string]interface{}) // Clear old handlers
	if props != nil {
		propsValue := reflect.ValueOf(props)
		if propsValue.Kind() == reflect.Ptr {
			propsValue = propsValue.Elem()
		}
		if propsValue.Kind() == reflect.Struct {
			propsType := propsValue.Type()
			for i := 0; i < propsValue.NumField(); i++ {
				field := propsType.Field(i)
				fieldValue := propsValue.Field(i)
				if strings.HasPrefix(field.Name, "On") && fieldValue.Kind() == reflect.Func && !fieldValue.IsNil() {
					instance.handlers[field.Name] = fieldValue.Interface()
				}
			}
		}
	}
	return instance
}

func (c *fcInstanceContext) cleanupEffects(removedIDs []string) {
	for _, id := range removedIDs {
		if instance, ok := c.ctxs[id]; ok {
			for i := range instance.Effects {
				if instance.Effects[i].cleanupFn != nil {
					instance.Effects[i].cleanupFn()
					instance.Effects[i].cleanupFn = nil // Avoid double cleanup
				}
			}
		}
	}
}

func (c *fcInstanceContext) getAllIDs() []string {
	ids := make([]string, 0, len(c.ctxs))
	for id := range c.ctxs {
		ids = append(ids, id)
	}
	return ids
}
