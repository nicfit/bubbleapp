package style

import (
	"github.com/charmbracelet/lipgloss/v2"
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

	ButtonBordered ButtonStyles
	Button         ButtonStyles
}

type ButtonStyles struct {
	Primary   lipgloss.Style
	Secondary lipgloss.Style
	Tertiary  lipgloss.Style
	Success   lipgloss.Style
	Danger    lipgloss.Style
	Info      lipgloss.Style
	Warning   lipgloss.Style

	PrimaryFocused   lipgloss.Style
	SecondaryFocused lipgloss.Style
	TertiaryFocused  lipgloss.Style
	SuccessFocused   lipgloss.Style
	DangerFocused    lipgloss.Style
	InfoFocused      lipgloss.Style
	WarningFocused   lipgloss.Style

	PrimaryHovered   lipgloss.Style
	SecondaryHovered lipgloss.Style
	TertiaryHovered  lipgloss.Style
	SuccessHovered   lipgloss.Style
	DangerHovered    lipgloss.Style
	InfoHovered      lipgloss.Style
	WarningHovered   lipgloss.Style
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

	buttonBaseNormal := lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder())
	s.ButtonBordered = ButtonStyles{
		Primary:   buttonBaseNormal.BorderForeground(s.Colors.Primary).Foreground(s.Colors.Primary),
		Secondary: buttonBaseNormal.BorderForeground(s.Colors.Secondary).Foreground(s.Colors.Secondary),
		Tertiary:  buttonBaseNormal.BorderForeground(s.Colors.Tertiary).Foreground(s.Colors.Tertiary),
		Success:   buttonBaseNormal.BorderForeground(s.Colors.Success).Foreground(s.Colors.Success),
		Danger:    buttonBaseNormal.BorderForeground(s.Colors.Danger).Foreground(s.Colors.Danger),
		Info:      buttonBaseNormal.BorderForeground(s.Colors.Info).Foreground(s.Colors.Info),
		Warning:   buttonBaseNormal.BorderForeground(s.Colors.Warning).Foreground(s.Colors.Warning),

		PrimaryFocused:   buttonBaseNormal.BorderForeground(s.Colors.PrimaryLight).Foreground(s.Colors.Black).Background(s.Colors.PrimaryDark),
		SecondaryFocused: buttonBaseNormal.BorderForeground(s.Colors.SecondaryLight).Foreground(s.Colors.Black).Background(s.Colors.SecondaryDark),
		TertiaryFocused:  buttonBaseNormal.BorderForeground(s.Colors.TertiaryLight).Foreground(s.Colors.Black).Background(s.Colors.TertiaryDark),
		SuccessFocused:   buttonBaseNormal.BorderForeground(s.Colors.SuccessLight).Foreground(s.Colors.Black).Background(s.Colors.SuccessDark),
		DangerFocused:    buttonBaseNormal.BorderForeground(s.Colors.DangerLight).Foreground(s.Colors.Black).Background(s.Colors.DangerDark),
		WarningFocused:   buttonBaseNormal.BorderForeground(s.Colors.WarningLight).Foreground(s.Colors.Black).Background(s.Colors.WarningDark),
		InfoFocused:      buttonBaseNormal.BorderForeground(s.Colors.InfoLight).Foreground(s.Colors.Black).Background(s.Colors.InfoDark),

		PrimaryHovered:   buttonBaseNormal.BorderForeground(s.Colors.PrimaryDark).Foreground(s.Colors.Black).Background(s.Colors.PrimaryLight),
		SecondaryHovered: buttonBaseNormal.BorderForeground(s.Colors.SecondaryDark).Foreground(s.Colors.Black).Background(s.Colors.SecondaryLight),
		TertiaryHovered:  buttonBaseNormal.BorderForeground(s.Colors.TertiaryDark).Foreground(s.Colors.Black).Background(s.Colors.TertiaryLight),
		SuccessHovered:   buttonBaseNormal.BorderForeground(s.Colors.SuccessDark).Foreground(s.Colors.Black).Background(s.Colors.SuccessLight),
		DangerHovered:    buttonBaseNormal.BorderForeground(s.Colors.DangerDark).Foreground(s.Colors.Black).Background(s.Colors.DangerLight),
		WarningHovered:   buttonBaseNormal.BorderForeground(s.Colors.WarningDark).Foreground(s.Colors.Black).Background(s.Colors.WarningLight),
		InfoHovered:      buttonBaseNormal.BorderForeground(s.Colors.InfoDark).Foreground(s.Colors.Black).Background(s.Colors.InfoLight),
	}

	s.Button = ButtonStyles{
		Primary:   lipgloss.NewStyle().Background(s.Colors.Primary).Foreground(s.Colors.Black),
		Secondary: lipgloss.NewStyle().Background(s.Colors.Secondary).Foreground(s.Colors.Black),
		Tertiary:  lipgloss.NewStyle().Background(s.Colors.Tertiary).Foreground(s.Colors.Black),
		Success:   lipgloss.NewStyle().Background(s.Colors.Success).Foreground(s.Colors.Black),
		Danger:    lipgloss.NewStyle().Background(s.Colors.Danger).Foreground(s.Colors.Black),
		Warning:   lipgloss.NewStyle().Background(s.Colors.Warning).Foreground(s.Colors.Black),
		Info:      lipgloss.NewStyle().Background(s.Colors.Info).Foreground(s.Colors.Black),

		PrimaryFocused:   lipgloss.NewStyle().Background(s.Colors.PrimaryLight).Foreground(s.Colors.Black),
		SecondaryFocused: lipgloss.NewStyle().Background(s.Colors.SecondaryLight).Foreground(s.Colors.Black),
		TertiaryFocused:  lipgloss.NewStyle().Background(s.Colors.TertiaryLight).Foreground(s.Colors.Black),
		SuccessFocused:   lipgloss.NewStyle().Background(s.Colors.SuccessLight).Foreground(s.Colors.Black),
		DangerFocused:    lipgloss.NewStyle().Background(s.Colors.DangerLight).Foreground(s.Colors.Black),
		WarningFocused:   lipgloss.NewStyle().Background(s.Colors.WarningLight).Foreground(s.Colors.Black),
		InfoFocused:      lipgloss.NewStyle().Background(s.Colors.InfoLight).Foreground(s.Colors.Black),

		PrimaryHovered:   lipgloss.NewStyle().Background(s.Colors.PrimaryDark).Foreground(s.Colors.Black),
		SecondaryHovered: lipgloss.NewStyle().Background(s.Colors.SecondaryDark).Foreground(s.Colors.Black),
		TertiaryHovered:  lipgloss.NewStyle().Background(s.Colors.TertiaryDark).Foreground(s.Colors.Black),
		SuccessHovered:   lipgloss.NewStyle().Background(s.Colors.SuccessDark).Foreground(s.Colors.Black),
		DangerHovered:    lipgloss.NewStyle().Background(s.Colors.DangerDark).Foreground(s.Colors.Black),
		WarningHovered:   lipgloss.NewStyle().Background(s.Colors.WarningDark).Foreground(s.Colors.Black),
		InfoHovered:      lipgloss.NewStyle().Background(s.Colors.InfoDark).Foreground(s.Colors.Black),
	}

	return s
}
