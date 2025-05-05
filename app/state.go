package app

type StateStore struct {
	data    map[string]any
	Focused string
	Hovered string
}

func NewStateStore() *StateStore {
	return &StateStore{data: make(map[string]any)}
}

func GetUIState[T any, UIState any](ctx *Context[T], id string) *UIState {
	v := ctx.UIState.get(id)
	typedV, ok := v.(*UIState)
	if !ok {
		return nil
	}
	return typedV
}

func SetUIState[T any, UIState any](ctx *Context[T], id string, value *UIState) {
	if value == nil {
		ctx.UIState.set(id, nil)
		return
	}
	ctx.UIState.set(id, value)
}

func (s *StateStore) get(id string) (v any) {
	v, ok := s.data[id]
	if !ok {
		return nil
	}
	return v
}

func (s *StateStore) set(id string, value any) {
	s.data[id] = value
}

func (s *StateStore) cleanup(existingIDs []string) {
	for id := range s.data {
		if !contains(existingIDs, id) {
			delete(s.data, id)
		}
	}

}

// contains checks if a slice of strings contains a specific string.
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
