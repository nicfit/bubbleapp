package context

import (
	"fmt"
	"strings"
	"sync/atomic"

	"github.com/alexanderbh/bubbleapp/app"
)

// Not sure if this code belongs in App or not. It depends on Context methods on Ctx.

var nextContextID uint64

// Context is a generic type representing a context object.
type Context[T any] struct {
	id           uint64
	defaultValue *T
}

// Create creates a new context object with an optional default value.
// The default value is returned by UseContext if no Provider is found in the ancestry.
func Create[T any](defaultValue *T) *Context[T] {
	return &Context[T]{
		id:           atomic.AddUint64(&nextContextID, 1),
		defaultValue: defaultValue,
	}
}

// ProviderProps are the props for the ContextProvider component.
// The Children field is a function that will render the child components.
// This function will be called by the UseChildren hook within the ContextProvider.
type ProviderProps[T any] struct {
	Context  *Context[T]
	Value    *T
	Children app.Children // Children is of type func(ctx *Ctx)
}

func NewProvider[T any](c *app.Ctx, context *Context[T], children app.Children) string {
	if context == nil {
		panic("NewContextProvider called with nil Context object")
	}

	p := ProviderProps[T]{
		Context:  context,
		Value:    context.defaultValue,
		Children: children,
	}

	return ContextProvider[T](c, p)
}

// ContextProvider is a component that makes a value available to all components
// in its subtree.
func ContextProvider[T any](c *app.Ctx, props app.Props) string {
	p, ok := props.(ProviderProps[T])
	if !ok {
		// This case should ideally be prevented by correct Go typing,
		// but as Props is 'any', a runtime check is good.
		// Consider logging this error appropriately in a real application.
		return "[Error: Invalid props type for ContextProvider]"
	}

	if p.Context == nil {
		return "[Error: Context object is nil in Provider]"
	}

	// Push the context value onto the stack for this specific context ID.
	// This makes the value available to UseContext calls in descendant components.
	c.PushContextValue(p.Context.id, p.Value)
	// Ensure the context value is popped when this provider's rendering is complete,
	// restoring the context state for sibling or parent components.
	defer c.PopContextValue(p.Context.id)

	// Render children. The UseChildren hook will execute p.Children(c),
	// and any output from components rendered within p.Children will be collected.
	childrenOutputs := app.UseChildren(c, p.Children)

	// Concatenate the string outputs from all child components.
	var builder strings.Builder
	for _, childStr := range childrenOutputs {
		builder.WriteString(childStr)
	}
	return builder.String()
}

// UseContext is a hook that allows components to subscribe to a context's value.
// It returns the current value for the given context, searching upwards through
// component ancestors for the nearest Provider. If no Provider is found,
// it returns the default value specified when the context was created.
func UseContext[T any](c *app.Ctx, context *Context[T]) *T {
	if context == nil {
		// This is a programming error: UseContext called with a nil context object.
		panic("UseContext called with nil Context object")
	}

	value, found := c.GetContextValue(context.id)
	if found {
		// Value found from a Provider. Attempt to cast it to the expected type T.
		if typedValue, ok := value.(*T); ok {
			return typedValue
		}
		// Type mismatch: the value stored in the context by a Provider
		// is not of the type expected by this consumer. This is a programming error.
		panic(fmt.Sprintf("Context value type mismatch for context ID %d. Expected type %T (based on default value), but found type %T in provider.", context.id, context.defaultValue, value))
	}

	// No Provider found in the ancestry, return the default value for this context.
	return context.defaultValue
}
