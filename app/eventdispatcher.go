package app

import (
	tea "github.com/charmbracelet/bubbletea/v2"
)

// Should tea be exposed like that here? Or wrapped?
type MouseEvent struct {
	X, Y   int
	Button tea.MouseButton
	Mod    tea.KeyMod
}

type KeyEvent struct {
	Key tea.Key
}

const (
	semanticActionPrimary = "OnAction" // e.g., Click, Enter on button
)

// dispatchToHandler is a generic function to dispatch events to the appropriate handler.
// It checks if the handler exists for the given semantic intent and calls it with the provided event data.
// It supports different types of event data, such as MouseEvent, KeyEvent, and others.
// If the handler is not found or the event data type does not match, it falls back to a parameter-less function call.
// The function returns true if the handler was found and called, false otherwise.
func (m app) dispatchToHandler(instance *instanceContext, semanticIntent string, handlerName string, eventData ...interface{}) bool {

	// Try the semantic intent first
	if primaryActionHandler, ok := instance.handlers[semanticIntent].(func()); ok {
		primaryActionHandler()
		return true
	}

	handlerFuncInterface, ok := instance.handlers[handlerName]
	if !ok {
		return false
	}

	switch data := eventData[0].(type) {
	case MouseEvent:
		if fn, typeOK := handlerFuncInterface.(func(MouseEvent)); typeOK {
			fn(data)
			return true
		} else if fn, typeOK := handlerFuncInterface.(func()); typeOK { // Fallback: func()
			fn()
			return true
		} else {
			// log.Printf("Handler '%s' for %s is not func(MouseEventData) or func(). Type: %T", handlerName, instance.id, handlerFuncInterface)
		}

	// case KeyEventData:
	//     if fn, typeOK := handlerFuncInterface.(func(KeyEventData) bool); typeOK { // e.g., returns bool if handled
	//         // fmt.Printf("Calling handler '%s' for %s with KeyEventData\n", handlerName, instance.id)
	//         if fn(data) {
	//              reRenderNeeded = true
	//         }
	//     } else if fn, typeOK := handlerFuncInterface.(func(key rune, mods Modifiers) bool); typeOK { // Another common pattern
	//         // fmt.Printf("Calling handler '%s' for %s with rune, Modifiers (event was KeyEventData)\n", handlerName, instance.id)
	//         if fn(data.Key, data.Mods) {
	//              reRenderNeeded = true
	//         }
	//     } else if fn, typeOK := handlerFuncInterface.(func()); typeOK { // Fallback: func()
	//         // fmt.Printf("Calling parameter-less handler '%s' for %s (event was KeyEventData)\n", handlerName, instance.id)
	//         fn()
	//         reRenderNeeded = true
	//     } else {
	//         // log.Printf("Handler '%s' for %s is not func(KeyEventData)bool, func(rune,Mods)bool, or func(). Type: %T", handlerName, instance.id, handlerFuncInterface)
	//     }

	// case TickEventData:
	//     if fn, typeOK := handlerFuncInterface.(func(TickEventData)); typeOK {
	//         fn(data)
	//         reRenderNeeded = true
	//     } else if fn, typeOK := handlerFuncInterface.(func(float64)); typeOK { // If handler just wants deltaTime
	//         fn(data.DeltaTime)
	//         reRenderNeeded = true
	//     } else if fn, typeOK := handlerFuncInterface.(func()); typeOK { // Fallback
	//         fn()
	//         reRenderNeeded = true
	//     } else {
	//         // log.Printf("Handler '%s' for %s is not func(TickEventData), func(float64), or func(). Type: %T", handlerName, instance.id, handlerFuncInterface)
	//     }

	// // Add more event data types as needed

	default:
		// If eventData is empty or of an unknown type, try calling as func() as a last resort
		// (This part might overlap with the fallbacks above, structure carefully)
		if len(eventData) == 0 || eventData[0] == nil { // Or a specific signal for no-arg call
			if fn, typeOK := handlerFuncInterface.(func()); typeOK {
				// fmt.Printf("Calling parameter-less handler '%s' for %s (no specific event data provided)\n", handlerName, instance.id)
				fn()
				return true
			} else if handlerFuncInterface != nil { // It exists but isn't func()
				// log.Printf("Handler '%s' for %s exists but is not func() and no matching event data provided. Type: %T", handlerName, instance.id, handlerFuncInterface)
			}
		} else {
			// log.Printf("Handler '%s' for %s: Unhandled event data type %T or no matching signature.", handlerName, instance.id, eventData[0])
		}
	}

	return false
}
