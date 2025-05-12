package style

import (
	"image/color"

	"github.com/charmbracelet/lipgloss/v2"
)

type Colors struct {
	Primary      color.Color
	PrimaryLight color.Color
	PrimaryDark  color.Color

	Secondary      color.Color
	SecondaryLight color.Color
	SecondaryDark  color.Color

	Tertiary      color.Color
	TertiaryLight color.Color
	TertiaryDark  color.Color

	Info      color.Color
	InfoLight color.Color
	InfoDark  color.Color

	Danger      color.Color
	DangerLight color.Color
	DangerDark  color.Color

	Success      color.Color
	SuccessLight color.Color
	SuccessDark  color.Color

	Warning      color.Color
	WarningLight color.Color
	WarningDark  color.Color

	Ghost      color.Color
	GhostLight color.Color
	GhostDark  color.Color

	Background          color.Color
	SecondaryBackground color.Color
	TertiaryBackground  color.Color
	UIPanelBackground   color.Color
	HighlightBackground color.Color
	CodeBackground      color.Color

	White color.Color
	Black color.Color
}

func defaultColors() Colors {
	return Colors{
		Primary:      lipgloss.Color("#BA99FF"),
		PrimaryLight: lipgloss.Color("#D4B3FF"),
		PrimaryDark:  lipgloss.Color("#A07FE5"),

		Secondary:      lipgloss.Color("#539FFF"),
		SecondaryLight: lipgloss.Color("#6DB9FF"),
		SecondaryDark:  lipgloss.Color("#3985E5"),

		Tertiary:      lipgloss.Color("#1AE9D7"),
		TertiaryLight: lipgloss.Color("#34FFF1"),
		TertiaryDark:  lipgloss.Color("#00CFBD"),

		Info:      lipgloss.Color("#8AC3F5"),
		InfoLight: lipgloss.Color("#A4DDFF"),
		InfoDark:  lipgloss.Color("#70A9DB"),

		Danger:      lipgloss.Color("#F56E7B"),
		DangerLight: lipgloss.Color("#FF8895"),
		DangerDark:  lipgloss.Color("#DB5461"),

		Success:      lipgloss.Color("#7BEAAF"),
		SuccessLight: lipgloss.Color("#95FFC9"),
		SuccessDark:  lipgloss.Color("#61D095"),

		Warning:      lipgloss.Color("#FFDC43"),
		WarningLight: lipgloss.Color("#FFF65D"),
		WarningDark:  lipgloss.Color("#E5C229"),

		Ghost:      lipgloss.Color("#A4A7BD"),
		GhostLight: lipgloss.Color("#BEC1D7"),
		GhostDark:  lipgloss.Color("#8A8DA3"),

		Background:          lipgloss.Color("#24242F"),
		SecondaryBackground: lipgloss.Color("#2C2C3D"),
		TertiaryBackground:  lipgloss.Color("#34344A"),
		UIPanelBackground:   lipgloss.Color("#2E2E42"),
		HighlightBackground: lipgloss.Color("#3F3F5C"),
		CodeBackground:      lipgloss.Color("#272734"),

		White: lipgloss.Color("#FFFFFF"),
		Black: lipgloss.Color("#000000"),
	}
}
