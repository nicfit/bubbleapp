package text

import "image/color"

// WithFg sets the foreground color.
func WithFg(fg color.Color) prop {
	return func(props *Props) {
		props.Foreground = fg
	}
}

// WithBg sets the background color.
func WithBg(bg color.Color) prop {
	return func(props *Props) {
		props.Background = bg
	}
}

// WithBold enables or disables bold text.
func WithBold(bold bool) prop {
	return func(props *Props) {
		props.Bold = bold
	}
}

// WithM sets uniform margin for all sides.
func WithM(m int) prop {
	return func(props *Props) {
		props.Margin.M = m
	}
}

// WithMargin sets individual margins.
func WithMargin(top, right, bottom, left int) prop {
	return func(props *Props) {
		props.Margin.MT = top
		props.Margin.MR = right
		props.Margin.MB = bottom
		props.Margin.ML = left
	}
}

// WithMT sets the top margin.
func WithMT(m int) prop {
	return func(props *Props) {
		props.Margin.MT = m
	}
}

// WithMR sets the right margin.
func WithMR(m int) prop {
	return func(props *Props) {
		props.Margin.MR = m
	}
}

// WithMB sets the bottom margin.
func WithMB(m int) prop {
	return func(props *Props) {
		props.Margin.MB = m
	}
}

// WithML sets the left margin.
func WithML(m int) prop {
	return func(props *Props) {
		props.Margin.ML = m
	}
}

// WithMX sets horizontal (left and right) margins.
func WithMX(m int) prop {
	return func(props *Props) {
		props.Margin.MX = m
	}
}

// WithMY sets vertical (top and bottom) margins.
func WithMY(m int) prop {
	return func(props *Props) {
		props.Margin.MY = m
	}
}

// WithP sets uniform padding for all sides.
func WithP(p int) prop {
	return func(props *Props) {
		props.Padding.P = p
	}
}

// WithPadding sets individual paddings.
func WithPadding(top, right, bottom, left int) prop {
	return func(props *Props) {
		props.Padding.PT = top
		props.Padding.PR = right
		props.Padding.PB = bottom
		props.Padding.PL = left
	}
}

// WithPT sets the top padding.
func WithPT(p int) prop {
	return func(props *Props) {
		props.Padding.PT = p
	}
}

// WithPR sets the right padding.
func WithPR(p int) prop {
	return func(props *Props) {
		props.Padding.PR = p
	}
}

// WithPB sets the bottom padding.
func WithPB(p int) prop {
	return func(props *Props) {
		props.Padding.PB = p
	}
}

// WithPL sets the left padding.
func WithPL(p int) prop {
	return func(props *Props) {
		props.Padding.PL = p
	}
}

// WithPX sets horizontal (left and right) paddings.
func WithPX(p int) prop {
	return func(props *Props) {
		props.Padding.PX = p
	}
}

// WithPY sets vertical (top and bottom) paddings.
func WithPY(p int) prop {
	return func(props *Props) {
		props.Padding.PY = p
	}
}
