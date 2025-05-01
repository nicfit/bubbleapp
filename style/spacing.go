package style

import (
	"github.com/charmbracelet/lipgloss/v2"
)

type Margin struct {
	M  int
	MT int
	MB int
	ML int
	MR int
	MY int
	MX int
}

func ApplyMargin(s lipgloss.Style, options Margin) lipgloss.Style {
	marginTop := 0
	marginBottom := 0
	marginLeft := 0
	marginRight := 0

	if options.M > 0 {
		marginTop = options.M
		marginBottom = options.M
		marginLeft = options.M
		marginRight = options.M
	}
	if options.MY > 0 {
		marginTop = options.MY
		marginBottom = options.MY
	}
	if options.MX > 0 {
		marginLeft = options.MX
		marginRight = options.MX
	}
	// Individual settings override broader ones
	if options.MT > 0 {
		marginTop = options.MT
	}
	if options.MB > 0 {
		marginBottom = options.MB
	}
	if options.ML > 0 {
		marginLeft = options.ML
	}
	if options.MR > 0 {
		marginRight = options.MR
	}

	return s.Margin(marginTop, marginRight, marginBottom, marginLeft)
}
