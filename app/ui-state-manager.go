package app

type uiStateContext struct {
	heights map[string]int
	widths  map[string]int
	Focused string
	Hovered string
}

func NewUIStateContext() *uiStateContext {
	return &uiStateContext{
		heights: make(map[string]int),
		widths:  make(map[string]int),
	}
}

func (s *uiStateContext) GetHeight(id string) int {
	v, ok := s.heights[id]
	if !ok {
		return 0
	}
	return v
}
func (s *uiStateContext) setHeight(id string, value int) {
	s.heights[id] = value
}
func (s *uiStateContext) GetWidth(id string) int {
	v, ok := s.widths[id]
	if !ok {
		return 0
	}
	return v
}
func (s *uiStateContext) setWidth(id string, value int) {
	s.widths[id] = value
}

func (s *uiStateContext) cleanup(existingIDs []string) {
	for id := range s.heights {
		if !contains(existingIDs, id) {
			delete(s.heights, id)
		}
	}
	for id := range s.widths {
		if !contains(existingIDs, id) {
			delete(s.widths, id)
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
