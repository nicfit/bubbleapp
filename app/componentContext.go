package app

import (
	"reflect"
	"strings"
)

type fcInstance struct {
	id        string
	focusable bool
	fc        FC
	props     Props
	handlers  map[string]interface{}
	States    []any // Added for UseState
}

type fcInstanceContext struct {
	ctxs map[string]*fcInstance
}

func newFCInstanceContext() *fcInstanceContext {
	return &fcInstanceContext{
		ctxs: make(map[string]*fcInstance),
	}
}

func (c *fcInstanceContext) get(id string) (*fcInstance, bool) {
	instance, ok := c.ctxs[id]
	return instance, ok
}

func (c *fcInstanceContext) set(id string, fc FC, props Props) *fcInstance {
	instance, ok := c.ctxs[id]
	if !ok {
		instance = &fcInstance{
			id:       id,
			handlers: make(map[string]interface{}),
			States:   make([]any, 0), // Initialize States for a new instance
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
