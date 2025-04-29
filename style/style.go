package style

import (
	"github.com/charmbracelet/lipgloss/v2"
)

const (
	TopHeight = 13
)

type Styles struct {
	Colors Colors

	TextPrimary   lipgloss.Style
	TextSecondary lipgloss.Style
	TextTertiary  lipgloss.Style
	TextDanger    lipgloss.Style
	TextSuccess   lipgloss.Style
	TextWarning   lipgloss.Style
	TextInfo      lipgloss.Style
	TextGhost     lipgloss.Style
}

// TODO: Add a way to define Themes.
// Maybe use bubbletint if possible? Perhaps we need more colors though.
func DefaultStyles() *Styles {

	s := new(Styles)

	s.Colors = defaultColors()

	s.TextPrimary = lipgloss.NewStyle().Foreground(s.Colors.Primary)
	s.TextSecondary = lipgloss.NewStyle().Foreground(s.Colors.Secondary)
	s.TextTertiary = lipgloss.NewStyle().Foreground(s.Colors.Tertiary)
	s.TextDanger = lipgloss.NewStyle().Foreground(s.Colors.Danger)
	s.TextSuccess = lipgloss.NewStyle().Foreground(s.Colors.Success)
	s.TextWarning = lipgloss.NewStyle().Foreground(s.Colors.Warning)
	s.TextInfo = lipgloss.NewStyle().Foreground(s.Colors.Info)
	s.TextGhost = lipgloss.NewStyle().Foreground(s.Colors.Ghost)

	return s
}
