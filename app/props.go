package app

import (
	"image/color"

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

type Padding struct {
	P  int
	PT int
	PB int
	PL int
	PR int
	PY int
	PX int
}

func ApplyPadding(s lipgloss.Style, options Padding) lipgloss.Style {
	paddingTop := 0
	paddingBottom := 0
	paddingLeft := 0
	paddingRight := 0

	if options.P > 0 {
		paddingTop = options.P
		paddingBottom = options.P
		paddingLeft = options.P
		paddingRight = options.P
	}
	if options.PY > 0 {
		paddingTop = options.PY
		paddingBottom = options.PY
	}
	if options.PX > 0 {
		paddingLeft = options.PX
		paddingRight = options.PX
	}
	// Individual settings override broader ones
	if options.PT > 0 {
		paddingTop = options.PT
	}
	if options.PB > 0 {
		paddingBottom = options.PB
	}
	if options.PL > 0 {
		paddingLeft = options.PL
	}
	if options.PR > 0 {
		paddingRight = options.PR
	}

	return s.Padding(paddingTop, paddingRight, paddingBottom, paddingLeft)
}

type Border struct {
	Border       lipgloss.Border
	BorderTop    bool
	BorderBottom bool
	BorderLeft   bool
	BorderRight  bool
	Color        color.Color
}

func ApplyBorder(s lipgloss.Style, options Border) lipgloss.Style {
	s = s.Border(options.Border).BorderBottom(false).BorderLeft(false).BorderRight(false).BorderTop(false)
	if options.BorderTop {
		s = s.BorderTop(options.BorderTop)
	}
	if options.BorderBottom {
		s = s.BorderBottom(options.BorderBottom)
	}
	if options.BorderLeft {
		s = s.BorderLeft(options.BorderLeft)
	}
	if options.BorderRight {
		s = s.BorderRight(options.BorderRight)
	}
	if options.Color != nil {
		s = s.BorderForeground(options.Color)
	}
	return s
}
