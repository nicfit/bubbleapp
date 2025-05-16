package context

import (
	"fmt"
	"sync/atomic"

	"github.com/alexanderbh/bubbleapp/app"
)

var nextContextID uint64

// Context is a generic type representing a context object.
type Context[T any] struct {
	id           uint64
	initialValue T // This is only returned if no Provider is found in the ancestry.
}

// Create creates a new context object with an initial reference value.
// This initialValue is returned by UseContext if no Provider is found in the ancestry.
func Create[T any](initialValue T) *Context[T] {
	return &Context[T]{
		id:           atomic.AddUint64(&nextContextID, 1),
		initialValue: initialValue,
	}
}

// ProviderProps are the props for the ContextProvider component.
type ProviderProps[T any] struct {
	Context *Context[T]
	Value   T // This is the specific value this provider instance will make available
	Child   app.Child
	app.Layout
}

// NewProvider creates a new ContextProvider component.
// It takes the context object, the specific value to provide, and children.
func NewProvider[T any](c *app.Ctx, context *Context[T], valueToProvide T, child app.Child) app.C {
	if context == nil {
		panic("NewProvider called with nil Context object")
	}

	p := ProviderProps[T]{
		Context: context,
		Value:   valueToProvide, // Use the explicitly passed valueToProvide
		Child:   child,
		Layout: app.Layout{
			GrowX: true,
			GrowY: true,
		},
	}

	return c.RenderWithName(ContextProvider[T], p, "CtxProvider{"+fmt.Sprintf("%T", valueToProvide)+"}")
}

// ContextProvider is a component that makes a value available to all components
// in its subtree.
func ContextProvider[T any](c *app.Ctx, props app.Props) string {
	p, ok := props.(ProviderProps[T])
	if !ok {
		panic(`ContextProvider: Invalid props type.`)
	}

	if p.Context == nil {
		panic("ContextProvider: Context object is nil in ProviderProps")
	}

	c.PushContextValue(p.Context.id, p.Value)
	defer c.PopContextValue(p.Context.id)

	return p.Child(c).String()
}

// UseContext is a hook that allows components to subscribe to a context's value.
// It returns the current value for the given context, searching upwards through
// component ancestors for the nearest Provider. If no Provider is found,
// it returns the initialValue specified when the context was created.
func UseContext[T any](c *app.Ctx, context *Context[T]) T {
	if context == nil {
		panic("UseContext called with nil Context object")
	}

	value, found := c.GetContextValue(context.id)
	if found {
		if typedValue, ok := value.(T); ok {
			return typedValue
		}
		// This panic indicates a type mismatch between what a Provider stored
		// and what UseContext expected. This is a critical programming error.
		panic(fmt.Sprintf("Context value type mismatch for context ID %d. Expected type *%T, but found type %T in provider.", context.id, *new(T), value))
	}

	// No Provider found in the ancestry, return the initialValue for this context.
	return context.initialValue
}
