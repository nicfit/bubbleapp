package style

import (
	"image/color"

	"github.com/charmbracelet/lipgloss/v2"
)

type ComponentState int

const (
	Normal ComponentState = iota
	Hover
	Focus
	Disabled
)

type Variant int

const (
	Primary Variant = iota
	Secondary
	Tertiary
	Success
	Danger
	Info
	Warning
	Base
)

type AppTheme struct {
	Colors Colors

	BackgroundColor color.Color
	ForegroundColor color.Color

	Button map[Variant]map[ComponentState]lipgloss.Style
	Text   map[Variant]map[ComponentState]lipgloss.Style
}

func NewDefaultAppTheme() *AppTheme {
	return NewAppTheme(NewDefaultColors())
}

func NewAppTheme(colors Colors) *AppTheme {

	// Define a base padding for buttons
	buttonPadding := lipgloss.NewStyle()

	return &AppTheme{
		Colors:          colors,
		BackgroundColor: colors.Base900,
		ForegroundColor: colors.Base50,

		Button: map[Variant]map[ComponentState]lipgloss.Style{
			Primary: {
				Normal:   buttonPadding.Background(colors.Primary).Foreground(colors.Base50),
				Hover:    buttonPadding.Background(colors.PrimaryLighter).Foreground(colors.Base950),
				Focus:    buttonPadding.Background(colors.PrimaryLight).Foreground(colors.Base950),
				Disabled: buttonPadding.Background(colors.PrimaryLight).Foreground(colors.PrimaryDark),
			},
			Secondary: {
				Normal:   buttonPadding.Background(colors.Secondary).Foreground(colors.Base50),
				Hover:    buttonPadding.Background(colors.SecondaryLighter).Foreground(colors.Base950),
				Focus:    buttonPadding.Background(colors.SecondaryLight).Foreground(colors.Base950),
				Disabled: buttonPadding.Background(colors.SecondaryLight).Foreground(colors.SecondaryDark),
			},
			Tertiary: {
				Normal:   buttonPadding.Background(colors.Tertiary).Foreground(colors.Base50),
				Hover:    buttonPadding.Background(colors.TertiaryLighter).Foreground(colors.Base950),
				Focus:    buttonPadding.Background(colors.TertiaryLight).Foreground(colors.Base950),
				Disabled: buttonPadding.Background(colors.TertiaryLight).Foreground(colors.TertiaryDark),
			},
			Success: {
				Normal:   buttonPadding.Background(colors.Success).Foreground(colors.Base50),
				Hover:    buttonPadding.Background(colors.SuccessLighter).Foreground(colors.Base950),
				Focus:    buttonPadding.Background(colors.SuccessLight).Foreground(colors.Base950),
				Disabled: buttonPadding.Background(colors.SuccessLight).Foreground(colors.SuccessDark),
			},
			Danger: {
				Normal:   buttonPadding.Background(colors.Danger).Foreground(colors.Base50),
				Hover:    buttonPadding.Background(colors.DangerLighter).Foreground(colors.Base950),
				Focus:    buttonPadding.Background(colors.DangerLight).Foreground(colors.Base950),
				Disabled: buttonPadding.Background(colors.DangerLight).Foreground(colors.DangerDark),
			},
			Info: {
				Normal:   buttonPadding.Background(colors.Info).Foreground(colors.Base50),
				Hover:    buttonPadding.Background(colors.InfoLighter).Foreground(colors.Base950),
				Focus:    buttonPadding.Background(colors.InfoLight).Foreground(colors.Base950),
				Disabled: buttonPadding.Background(colors.InfoLight).Foreground(colors.InfoDark),
			},
			Warning: {
				Normal:   buttonPadding.Background(colors.WarningDark).Foreground(colors.Base50),
				Hover:    buttonPadding.Background(colors.WarningLighter).Foreground(colors.Base950),
				Focus:    buttonPadding.Background(colors.WarningLight).Foreground(colors.Base950),
				Disabled: buttonPadding.Background(colors.WarningLight).Foreground(colors.WarningDark),
			},
			Base: {
				Normal:   buttonPadding.Background(colors.Base700).Foreground(colors.Base50),
				Hover:    buttonPadding.Background(colors.Base800).Foreground(colors.Base50),
				Focus:    buttonPadding.Background(colors.Base900).Foreground(colors.Base50),
				Disabled: buttonPadding.Background(colors.Base900).Foreground(colors.Base700),
			},
		},
		Text: map[Variant]map[ComponentState]lipgloss.Style{
			Primary: {
				Normal:   lipgloss.NewStyle().Foreground(colors.PrimaryFg),
				Hover:    lipgloss.NewStyle().Foreground(colors.PrimaryLight),
				Focus:    lipgloss.NewStyle().Foreground(colors.PrimaryDark).Underline(true),
				Disabled: lipgloss.NewStyle().Foreground(colors.Base400),
			},
			Secondary: {
				Normal:   lipgloss.NewStyle().Foreground(colors.SecondaryFg),
				Hover:    lipgloss.NewStyle().Foreground(colors.SecondaryLight),
				Focus:    lipgloss.NewStyle().Foreground(colors.SecondaryDark).Underline(true),
				Disabled: lipgloss.NewStyle().Foreground(colors.Base400),
			},
			Tertiary: {
				Normal:   lipgloss.NewStyle().Foreground(colors.TertiaryFg),
				Hover:    lipgloss.NewStyle().Foreground(colors.TertiaryLight),
				Focus:    lipgloss.NewStyle().Foreground(colors.TertiaryDark).Underline(true),
				Disabled: lipgloss.NewStyle().Foreground(colors.Base400),
			},
			Success: {
				Normal:   lipgloss.NewStyle().Foreground(colors.SuccessFg),
				Hover:    lipgloss.NewStyle().Foreground(colors.SuccessLight),
				Focus:    lipgloss.NewStyle().Foreground(colors.SuccessDark).Underline(true),
				Disabled: lipgloss.NewStyle().Foreground(colors.Base400),
			},
			Danger: {
				Normal:   lipgloss.NewStyle().Foreground(colors.DangerFg),
				Hover:    lipgloss.NewStyle().Foreground(colors.DangerLight),
				Focus:    lipgloss.NewStyle().Foreground(colors.DangerDark).Underline(true),
				Disabled: lipgloss.NewStyle().Foreground(colors.Base400),
			},
			Info: {
				Normal:   lipgloss.NewStyle().Foreground(colors.InfoFg),
				Hover:    lipgloss.NewStyle().Foreground(colors.InfoLight),
				Focus:    lipgloss.NewStyle().Foreground(colors.InfoDark).Underline(true),
				Disabled: lipgloss.NewStyle().Foreground(colors.Base400),
			},
			Warning: {
				Normal:   lipgloss.NewStyle().Foreground(colors.WarningFg),
				Hover:    lipgloss.NewStyle().Foreground(colors.WarningLight),
				Focus:    lipgloss.NewStyle().Foreground(colors.WarningDark).Underline(true),
				Disabled: lipgloss.NewStyle().Foreground(colors.Base400),
			},
			Base: {
				Normal:   lipgloss.NewStyle().Foreground(colors.Base50),
				Hover:    lipgloss.NewStyle().Foreground(colors.Base50),
				Focus:    lipgloss.NewStyle().Foreground(colors.Base50).Underline(true),
				Disabled: lipgloss.NewStyle().Foreground(colors.Base400),
			},
		},
	}
}
