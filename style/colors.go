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
		Primary:      lipgloss.Color("#A07FFF"),
		PrimaryLight: lipgloss.Color("#C0A0FF"),
		PrimaryDark:  lipgloss.Color("#8050DD"),

		Secondary:      lipgloss.Color("#3985FF"),
		SecondaryLight: lipgloss.Color("#6BABFF"),
		SecondaryDark:  lipgloss.Color("#2260C4"),

		Tertiary:      lipgloss.Color("#00CFBD"),
		TertiaryLight: lipgloss.Color("#40E8D9"),
		TertiaryDark:  lipgloss.Color("#00A297"),

		Info:      lipgloss.Color("#70A9DB"),
		InfoLight: lipgloss.Color("#90BFE9"),
		InfoDark:  lipgloss.Color("#4F7FAD"),

		Danger:      lipgloss.Color("#DB5461"),
		DangerLight: lipgloss.Color("#E87983"),
		DangerDark:  lipgloss.Color("#B33845"),

		Success:      lipgloss.Color("#61D095"),
		SuccessLight: lipgloss.Color("#83E0AE"),
		SuccessDark:  lipgloss.Color("#3FAF72"),

		Warning:      lipgloss.Color("#E6C229"),
		WarningLight: lipgloss.Color("#F0D35E"),
		WarningDark:  lipgloss.Color("#C2A017"),

		Ghost:      lipgloss.Color("#8A8DA3"),
		GhostLight: lipgloss.Color("#A6A9BC"),
		GhostDark:  lipgloss.Color("#565869"),

		Background:          lipgloss.Color("#0A0A15"),
		SecondaryBackground: lipgloss.Color("#121223"),
		TertiaryBackground:  lipgloss.Color("#1A1A30"),
		UIPanelBackground:   lipgloss.Color("#141428"),
		HighlightBackground: lipgloss.Color("#252542"),
		CodeBackground:      lipgloss.Color("#0D0D1A"),

		White: lipgloss.Color("#FFFFFF"),
		Black: lipgloss.Color("#000000"),
	}
}
