package style

import (
	"image/color"
)

type Colors struct {
	Palette Palette

	Base50  color.Color
	Base100 color.Color
	Base200 color.Color
	Base300 color.Color
	Base400 color.Color
	Base500 color.Color
	Base600 color.Color
	Base700 color.Color
	Base800 color.Color
	Base900 color.Color
	Base950 color.Color

	BaseAlt50  color.Color
	BaseAlt100 color.Color
	BaseAlt200 color.Color
	BaseAlt300 color.Color
	BaseAlt400 color.Color
	BaseAlt500 color.Color
	BaseAlt600 color.Color
	BaseAlt700 color.Color
	BaseAlt800 color.Color
	BaseAlt900 color.Color
	BaseAlt950 color.Color

	Primary        color.Color
	PrimaryLight   color.Color
	PrimaryLighter color.Color
	PrimaryDark    color.Color
	PrimaryDarker  color.Color
	PrimaryBg      color.Color
	PrimaryFg      color.Color

	Secondary        color.Color
	SecondaryLight   color.Color
	SecondaryLighter color.Color
	SecondaryDark    color.Color
	SecondaryDarker  color.Color
	SecondaryBg      color.Color
	SecondaryFg      color.Color

	Tertiary        color.Color
	TertiaryLight   color.Color
	TertiaryLighter color.Color
	TertiaryDark    color.Color
	TertiaryDarker  color.Color
	TertiaryBg      color.Color
	TertiaryFg      color.Color

	Info        color.Color
	InfoLight   color.Color
	InfoLighter color.Color
	InfoDark    color.Color
	InfoDarker  color.Color
	InfoBg      color.Color
	InfoFg      color.Color

	Danger        color.Color
	DangerLight   color.Color
	DangerLighter color.Color
	DangerDark    color.Color
	DangerDarker  color.Color
	DangerBg      color.Color
	DangerFg      color.Color

	Success        color.Color
	SuccessLight   color.Color
	SuccessLighter color.Color
	SuccessDark    color.Color
	SuccessDarker  color.Color
	SuccessBg      color.Color
	SuccessFg      color.Color

	Warning        color.Color
	WarningLight   color.Color
	WarningLighter color.Color
	WarningDark    color.Color
	WarningDarker  color.Color
	WarningBg      color.Color
	WarningFg      color.Color
}

func NewDefaultColors() Colors {
	return NewColors(NewDefaultPalette())
}

func NewColors(palette Palette) Colors {
	return Colors{
		Palette: palette,

		Base50:  palette.Neutral50,
		Base100: palette.Neutral100,
		Base200: palette.Neutral200,
		Base300: palette.Neutral300,
		Base400: palette.Neutral400,
		Base500: palette.Neutral500,
		Base600: palette.Neutral600,
		Base700: palette.Neutral700,
		Base800: palette.Neutral800,
		Base900: palette.Neutral900,
		Base950: palette.Neutral950,

		BaseAlt50:  palette.Slate50,
		BaseAlt100: palette.Slate100,
		BaseAlt200: palette.Slate200,
		BaseAlt300: palette.Slate300,
		BaseAlt400: palette.Slate400,
		BaseAlt500: palette.Slate500,
		BaseAlt600: palette.Slate600,
		BaseAlt700: palette.Slate700,
		BaseAlt800: palette.Slate800,
		BaseAlt900: palette.Slate900,
		BaseAlt950: palette.Slate950,

		Primary:        palette.Indigo500,
		PrimaryLight:   palette.Indigo400,
		PrimaryLighter: palette.Indigo300,
		PrimaryDark:    palette.Indigo600,
		PrimaryDarker:  palette.Indigo700,
		PrimaryBg:      palette.Indigo800,
		PrimaryFg:      palette.Indigo300,

		Secondary:        palette.Emerald500,
		SecondaryLight:   palette.Emerald400,
		SecondaryLighter: palette.Emerald300,
		SecondaryDark:    palette.Emerald600,
		SecondaryDarker:  palette.Emerald700,
		SecondaryBg:      palette.Emerald800,
		SecondaryFg:      palette.Emerald300,

		Tertiary:        palette.Orange500,
		TertiaryLight:   palette.Orange400,
		TertiaryLighter: palette.Orange300,
		TertiaryDark:    palette.Orange600,
		TertiaryDarker:  palette.Orange700,
		TertiaryBg:      palette.Orange800,
		TertiaryFg:      palette.Orange300,

		Info:        palette.Sky500,
		InfoLight:   palette.Sky400,
		InfoLighter: palette.Sky300,
		InfoDark:    palette.Sky600,
		InfoDarker:  palette.Sky700,
		InfoBg:      palette.Sky800,
		InfoFg:      palette.Sky300,

		Danger:        palette.Red500,
		DangerLight:   palette.Red400,
		DangerLighter: palette.Red300,
		DangerDark:    palette.Red600,
		DangerDarker:  palette.Red700,
		DangerBg:      palette.Red800,
		DangerFg:      palette.Red300,

		Success:        palette.Green500,
		SuccessLight:   palette.Green500,
		SuccessLighter: palette.Green300,
		SuccessDark:    palette.Green600,
		SuccessDarker:  palette.Green700,
		SuccessBg:      palette.Green800,
		SuccessFg:      palette.Green300,

		Warning:        palette.Amber500,
		WarningLight:   palette.Amber400,
		WarningLighter: palette.Amber300,
		WarningDark:    palette.Amber600,
		WarningDarker:  palette.Amber700,
		WarningBg:      palette.Amber800,
		WarningFg:      palette.Amber300,
	}
}
