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
}

type fcInstanceContext struct {
	ctxs map[string]*fcInstance
}

func newFCInstanceContext() *fcInstanceContext {
	return &fcInstanceContext{
		ctxs: make(map[string]*fcInstance),
	}
}

func (c *fcInstanceContext) get(id string) *fcInstance {
	if ctx, ok := c.ctxs[id]; ok {
		return ctx
	}
	return &fcInstance{}
}

func (c *fcInstanceContext) set(id string, fc FC, props Props) {
	instance := &fcInstance{
		id:       id,
		fc:       fc,
		props:    props,
		handlers: make(map[string]interface{}),
	}

	// --- Extract handlers using reflection based on convention ---
	if props != nil {
		propsValue := reflect.ValueOf(props)
		// Ensure props is a struct or a pointer to a struct
		if propsValue.Kind() == reflect.Ptr {
			propsValue = propsValue.Elem()
		}

		if propsValue.Kind() == reflect.Struct {
			propsType := propsValue.Type()
			for i := 0; i < propsValue.NumField(); i++ {
				field := propsType.Field(i)
				fieldValue := propsValue.Field(i)

				// Convention: Starts with "On" and is a function
				if strings.HasPrefix(field.Name, "On") && fieldValue.Kind() == reflect.Func && !fieldValue.IsNil() {
					instance.handlers[field.Name] = fieldValue.Interface() // Store the function
				}
			}
		}
	}
	c.ctxs[id] = instance

}
