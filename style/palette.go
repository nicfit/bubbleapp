package style

import (
	"image/color"

	"github.com/charmbracelet/lipgloss/v2"
)

// Palette defines the color palette using Tailwind CSS color names and values.
type Palette struct {
	Red50  color.Color
	Red100 color.Color
	Red200 color.Color
	Red300 color.Color
	Red400 color.Color
	Red500 color.Color
	Red600 color.Color
	Red700 color.Color
	Red800 color.Color
	Red900 color.Color
	Red950 color.Color

	Orange50  color.Color
	Orange100 color.Color
	Orange200 color.Color
	Orange300 color.Color
	Orange400 color.Color
	Orange500 color.Color
	Orange600 color.Color
	Orange700 color.Color
	Orange800 color.Color
	Orange900 color.Color
	Orange950 color.Color

	Amber50  color.Color
	Amber100 color.Color
	Amber200 color.Color
	Amber300 color.Color
	Amber400 color.Color
	Amber500 color.Color
	Amber600 color.Color
	Amber700 color.Color
	Amber800 color.Color
	Amber900 color.Color
	Amber950 color.Color

	Yellow50  color.Color
	Yellow100 color.Color
	Yellow200 color.Color
	Yellow300 color.Color
	Yellow400 color.Color
	Yellow500 color.Color
	Yellow600 color.Color
	Yellow700 color.Color
	Yellow800 color.Color
	Yellow900 color.Color
	Yellow950 color.Color

	Lime50  color.Color
	Lime100 color.Color
	Lime200 color.Color
	Lime300 color.Color
	Lime400 color.Color
	Lime500 color.Color
	Lime600 color.Color
	Lime700 color.Color
	Lime800 color.Color
	Lime900 color.Color
	Lime950 color.Color

	Green50  color.Color
	Green100 color.Color
	Green200 color.Color
	Green300 color.Color
	Green400 color.Color
	Green500 color.Color
	Green600 color.Color
	Green700 color.Color
	Green800 color.Color
	Green900 color.Color
	Green950 color.Color

	Emerald50  color.Color
	Emerald100 color.Color
	Emerald200 color.Color
	Emerald300 color.Color
	Emerald400 color.Color
	Emerald500 color.Color
	Emerald600 color.Color
	Emerald700 color.Color
	Emerald800 color.Color
	Emerald900 color.Color
	Emerald950 color.Color

	Teal50  color.Color
	Teal100 color.Color
	Teal200 color.Color
	Teal300 color.Color
	Teal400 color.Color
	Teal500 color.Color
	Teal600 color.Color
	Teal700 color.Color
	Teal800 color.Color
	Teal900 color.Color
	Teal950 color.Color

	Cyan50  color.Color
	Cyan100 color.Color
	Cyan200 color.Color
	Cyan300 color.Color
	Cyan400 color.Color
	Cyan500 color.Color
	Cyan600 color.Color
	Cyan700 color.Color
	Cyan800 color.Color
	Cyan900 color.Color
	Cyan950 color.Color

	Sky50  color.Color
	Sky100 color.Color
	Sky200 color.Color
	Sky300 color.Color
	Sky400 color.Color
	Sky500 color.Color
	Sky600 color.Color
	Sky700 color.Color
	Sky800 color.Color
	Sky900 color.Color
	Sky950 color.Color

	Blue50  color.Color
	Blue100 color.Color
	Blue200 color.Color
	Blue300 color.Color
	Blue400 color.Color
	Blue500 color.Color
	Blue600 color.Color
	Blue700 color.Color
	Blue800 color.Color
	Blue900 color.Color
	Blue950 color.Color

	Indigo50  color.Color
	Indigo100 color.Color
	Indigo200 color.Color
	Indigo300 color.Color
	Indigo400 color.Color
	Indigo500 color.Color
	Indigo600 color.Color
	Indigo700 color.Color
	Indigo800 color.Color
	Indigo900 color.Color
	Indigo950 color.Color

	Violet50  color.Color
	Violet100 color.Color
	Violet200 color.Color
	Violet300 color.Color
	Violet400 color.Color
	Violet500 color.Color
	Violet600 color.Color
	Violet700 color.Color
	Violet800 color.Color
	Violet900 color.Color
	Violet950 color.Color

	Purple50  color.Color
	Purple100 color.Color
	Purple200 color.Color
	Purple300 color.Color
	Purple400 color.Color
	Purple500 color.Color
	Purple600 color.Color
	Purple700 color.Color
	Purple800 color.Color
	Purple900 color.Color
	Purple950 color.Color

	Fuchsia50  color.Color
	Fuchsia100 color.Color
	Fuchsia200 color.Color
	Fuchsia300 color.Color
	Fuchsia400 color.Color
	Fuchsia500 color.Color
	Fuchsia600 color.Color
	Fuchsia700 color.Color
	Fuchsia800 color.Color
	Fuchsia900 color.Color
	Fuchsia950 color.Color

	Pink50  color.Color
	Pink100 color.Color
	Pink200 color.Color
	Pink300 color.Color
	Pink400 color.Color
	Pink500 color.Color
	Pink600 color.Color
	Pink700 color.Color
	Pink800 color.Color
	Pink900 color.Color
	Pink950 color.Color

	Rose50  color.Color
	Rose100 color.Color
	Rose200 color.Color
	Rose300 color.Color
	Rose400 color.Color
	Rose500 color.Color
	Rose600 color.Color
	Rose700 color.Color
	Rose800 color.Color
	Rose900 color.Color
	Rose950 color.Color

	Slate50  color.Color
	Slate100 color.Color
	Slate200 color.Color
	Slate300 color.Color
	Slate400 color.Color
	Slate500 color.Color
	Slate600 color.Color
	Slate700 color.Color
	Slate800 color.Color
	Slate900 color.Color
	Slate950 color.Color

	Gray50  color.Color
	Gray100 color.Color
	Gray200 color.Color
	Gray300 color.Color
	Gray400 color.Color
	Gray500 color.Color
	Gray600 color.Color
	Gray700 color.Color
	Gray800 color.Color
	Gray900 color.Color
	Gray950 color.Color

	Zinc50  color.Color
	Zinc100 color.Color
	Zinc200 color.Color
	Zinc300 color.Color
	Zinc400 color.Color
	Zinc500 color.Color
	Zinc600 color.Color
	Zinc700 color.Color
	Zinc800 color.Color
	Zinc900 color.Color
	Zinc950 color.Color

	Neutral50  color.Color
	Neutral100 color.Color
	Neutral200 color.Color
	Neutral300 color.Color
	Neutral400 color.Color
	Neutral500 color.Color
	Neutral600 color.Color
	Neutral700 color.Color
	Neutral800 color.Color
	Neutral900 color.Color
	Neutral950 color.Color

	Stone50  color.Color
	Stone100 color.Color
	Stone200 color.Color
	Stone300 color.Color
	Stone400 color.Color
	Stone500 color.Color
	Stone600 color.Color
	Stone700 color.Color
	Stone800 color.Color
	Stone900 color.Color
	Stone950 color.Color

	Black color.Color
	White color.Color
}

// NewDefaultPalette returns a Palette initialized with default Tailwind CSS colors.
func NewDefaultPalette() Palette {
	return Palette{
		Red50:  lipgloss.Color("#fee2e2"),
		Red100: lipgloss.Color("#fecaca"),
		Red200: lipgloss.Color("#fca5a5"),
		Red300: lipgloss.Color("#f87171"),
		Red400: lipgloss.Color("#ef4444"),
		Red500: lipgloss.Color("#dc2626"),
		Red600: lipgloss.Color("#b91c1c"),
		Red700: lipgloss.Color("#991b1b"),
		Red800: lipgloss.Color("#7f1d1d"),
		Red900: lipgloss.Color("#681818"),
		Red950: lipgloss.Color("#450a0a"),

		Orange50:  lipgloss.Color("#fff7ed"),
		Orange100: lipgloss.Color("#ffedd5"),
		Orange200: lipgloss.Color("#fed7aa"),
		Orange300: lipgloss.Color("#fdba74"),
		Orange400: lipgloss.Color("#fb923c"),
		Orange500: lipgloss.Color("#f97316"),
		Orange600: lipgloss.Color("#ea580c"),
		Orange700: lipgloss.Color("#c2410c"),
		Orange800: lipgloss.Color("#9a3412"),
		Orange900: lipgloss.Color("#7c2d12"),
		Orange950: lipgloss.Color("#431407"),

		Amber50:  lipgloss.Color("#fffbeb"),
		Amber100: lipgloss.Color("#fef3c7"),
		Amber200: lipgloss.Color("#fde68a"),
		Amber300: lipgloss.Color("#fcd34d"),
		Amber400: lipgloss.Color("#fbbf24"),
		Amber500: lipgloss.Color("#f59e0b"),
		Amber600: lipgloss.Color("#d97706"),
		Amber700: lipgloss.Color("#b45309"),
		Amber800: lipgloss.Color("#92400e"),
		Amber900: lipgloss.Color("#78350f"),
		Amber950: lipgloss.Color("#451a03"),

		Yellow50:  lipgloss.Color("#fefce8"),
		Yellow100: lipgloss.Color("#fef9c3"),
		Yellow200: lipgloss.Color("#fef08a"),
		Yellow300: lipgloss.Color("#fde047"),
		Yellow400: lipgloss.Color("#facc15"),
		Yellow500: lipgloss.Color("#eab308"),
		Yellow600: lipgloss.Color("#ca8a04"),
		Yellow700: lipgloss.Color("#a16207"),
		Yellow800: lipgloss.Color("#854d0e"),
		Yellow900: lipgloss.Color("#713f12"),
		Yellow950: lipgloss.Color("#422006"),

		Lime50:  lipgloss.Color("#f7fee7"),
		Lime100: lipgloss.Color("#ecfccb"),
		Lime200: lipgloss.Color("#d9f99d"),
		Lime300: lipgloss.Color("#bef264"),
		Lime400: lipgloss.Color("#a3e635"),
		Lime500: lipgloss.Color("#84cc16"),
		Lime600: lipgloss.Color("#65a30d"),
		Lime700: lipgloss.Color("#4d7c0f"),
		Lime800: lipgloss.Color("#3f6212"),
		Lime900: lipgloss.Color("#365314"),
		Lime950: lipgloss.Color("#1a2e05"),

		Green50:  lipgloss.Color("#f0fdf4"),
		Green100: lipgloss.Color("#dcfce7"),
		Green200: lipgloss.Color("#bbf7d0"),
		Green300: lipgloss.Color("#86efac"),
		Green400: lipgloss.Color("#4ade80"),
		Green500: lipgloss.Color("#22c55e"),
		Green600: lipgloss.Color("#16a34a"),
		Green700: lipgloss.Color("#15803d"),
		Green800: lipgloss.Color("#166534"),
		Green900: lipgloss.Color("#14532d"),
		Green950: lipgloss.Color("#052e16"),

		Emerald50:  lipgloss.Color("#ecfdf5"),
		Emerald100: lipgloss.Color("#d1fae5"),
		Emerald200: lipgloss.Color("#a7f3d0"),
		Emerald300: lipgloss.Color("#6ee7b7"),
		Emerald400: lipgloss.Color("#34d399"),
		Emerald500: lipgloss.Color("#10b981"),
		Emerald600: lipgloss.Color("#059669"),
		Emerald700: lipgloss.Color("#047857"),
		Emerald800: lipgloss.Color("#065f46"),
		Emerald900: lipgloss.Color("#064e3b"),
		Emerald950: lipgloss.Color("#022c22"),

		Teal50:  lipgloss.Color("#f0fdfa"),
		Teal100: lipgloss.Color("#ccfbf1"),
		Teal200: lipgloss.Color("#99f6e4"),
		Teal300: lipgloss.Color("#5eead4"),
		Teal400: lipgloss.Color("#2dd4bf"),
		Teal500: lipgloss.Color("#14b8a6"),
		Teal600: lipgloss.Color("#0d9488"),
		Teal700: lipgloss.Color("#0f766e"),
		Teal800: lipgloss.Color("#115e59"),
		Teal900: lipgloss.Color("#134e4a"),
		Teal950: lipgloss.Color("#042f2e"),

		Cyan50:  lipgloss.Color("#ecfeff"),
		Cyan100: lipgloss.Color("#cffafe"),
		Cyan200: lipgloss.Color("#a5f3fc"),
		Cyan300: lipgloss.Color("#67e8f9"),
		Cyan400: lipgloss.Color("#22d3ee"),
		Cyan500: lipgloss.Color("#06b6d4"),
		Cyan600: lipgloss.Color("#0891b2"),
		Cyan700: lipgloss.Color("#0e7490"),
		Cyan800: lipgloss.Color("#155e75"),
		Cyan900: lipgloss.Color("#164e63"),
		Cyan950: lipgloss.Color("#083344"),

		Sky50:  lipgloss.Color("#f0f9ff"),
		Sky100: lipgloss.Color("#e0f2fe"),
		Sky200: lipgloss.Color("#bae6fd"),
		Sky300: lipgloss.Color("#7dd3fc"),
		Sky400: lipgloss.Color("#38bdf8"),
		Sky500: lipgloss.Color("#0ea5e9"),
		Sky600: lipgloss.Color("#0284c7"),
		Sky700: lipgloss.Color("#0369a1"),
		Sky800: lipgloss.Color("#075985"),
		Sky900: lipgloss.Color("#0c4a6e"),
		Sky950: lipgloss.Color("#082f49"),

		Blue50:  lipgloss.Color("#eff6ff"),
		Blue100: lipgloss.Color("#dbeafe"),
		Blue200: lipgloss.Color("#bfdbfe"),
		Blue300: lipgloss.Color("#93c5fd"),
		Blue400: lipgloss.Color("#60a5fa"),
		Blue500: lipgloss.Color("#3b82f6"),
		Blue600: lipgloss.Color("#2563eb"),
		Blue700: lipgloss.Color("#1d4ed8"),
		Blue800: lipgloss.Color("#1e40af"),
		Blue900: lipgloss.Color("#1e3a8a"),
		Blue950: lipgloss.Color("#172554"),

		Indigo50:  lipgloss.Color("#eef2ff"),
		Indigo100: lipgloss.Color("#e0e7ff"),
		Indigo200: lipgloss.Color("#c7d2fe"),
		Indigo300: lipgloss.Color("#a5b4fc"),
		Indigo400: lipgloss.Color("#818cf8"),
		Indigo500: lipgloss.Color("#6366f1"),
		Indigo600: lipgloss.Color("#4f46e5"),
		Indigo700: lipgloss.Color("#4338ca"),
		Indigo800: lipgloss.Color("#3730a3"),
		Indigo900: lipgloss.Color("#312e81"),
		Indigo950: lipgloss.Color("#1e1b4b"),

		Violet50:  lipgloss.Color("#f5f3ff"),
		Violet100: lipgloss.Color("#ede9fe"),
		Violet200: lipgloss.Color("#ddd6fe"),
		Violet300: lipgloss.Color("#c4b5fd"),
		Violet400: lipgloss.Color("#a78bfa"),
		Violet500: lipgloss.Color("#8b5cf6"),
		Violet600: lipgloss.Color("#7c3aed"),
		Violet700: lipgloss.Color("#6d28d9"),
		Violet800: lipgloss.Color("#5b21b6"),
		Violet900: lipgloss.Color("#4c1d95"),
		Violet950: lipgloss.Color("#2e1065"),

		Purple50:  lipgloss.Color("#faf5ff"),
		Purple100: lipgloss.Color("#f3e8ff"),
		Purple200: lipgloss.Color("#e9d5ff"),
		Purple300: lipgloss.Color("#d8b4fe"),
		Purple400: lipgloss.Color("#c084fc"),
		Purple500: lipgloss.Color("#a855f7"),
		Purple600: lipgloss.Color("#9333ea"),
		Purple700: lipgloss.Color("#7e22ce"),
		Purple800: lipgloss.Color("#6b21a8"),
		Purple900: lipgloss.Color("#581c87"),
		Purple950: lipgloss.Color("#3b0764"),

		Fuchsia50:  lipgloss.Color("#fdf4ff"),
		Fuchsia100: lipgloss.Color("#fae8ff"),
		Fuchsia200: lipgloss.Color("#f5d0fe"),
		Fuchsia300: lipgloss.Color("#f0abfc"),
		Fuchsia400: lipgloss.Color("#e879f9"),
		Fuchsia500: lipgloss.Color("#d946ef"),
		Fuchsia600: lipgloss.Color("#c026d3"),
		Fuchsia700: lipgloss.Color("#a21caf"),
		Fuchsia800: lipgloss.Color("#86198f"),
		Fuchsia900: lipgloss.Color("#701a75"),
		Fuchsia950: lipgloss.Color("#4a044e"),

		Pink50:  lipgloss.Color("#fdf2f8"),
		Pink100: lipgloss.Color("#fce7f3"),
		Pink200: lipgloss.Color("#fbcfe8"),
		Pink300: lipgloss.Color("#f9a8d4"),
		Pink400: lipgloss.Color("#f472b6"),
		Pink500: lipgloss.Color("#ec4899"),
		Pink600: lipgloss.Color("#db2777"),
		Pink700: lipgloss.Color("#be185d"),
		Pink800: lipgloss.Color("#9d174d"),
		Pink900: lipgloss.Color("#831843"),
		Pink950: lipgloss.Color("#500724"),

		Rose50:  lipgloss.Color("#fff1f2"),
		Rose100: lipgloss.Color("#ffe4e6"),
		Rose200: lipgloss.Color("#fecdd3"),
		Rose300: lipgloss.Color("#fda4af"),
		Rose400: lipgloss.Color("#fb7185"),
		Rose500: lipgloss.Color("#f43f5e"),
		Rose600: lipgloss.Color("#e11d48"),
		Rose700: lipgloss.Color("#be123c"),
		Rose800: lipgloss.Color("#9f1239"),
		Rose900: lipgloss.Color("#881337"),
		Rose950: lipgloss.Color("#4c0519"),

		Slate50:  lipgloss.Color("#f8fafc"),
		Slate100: lipgloss.Color("#f1f5f9"),
		Slate200: lipgloss.Color("#e2e8f0"),
		Slate300: lipgloss.Color("#cbd5e1"),
		Slate400: lipgloss.Color("#94a3b8"),
		Slate500: lipgloss.Color("#64748b"),
		Slate600: lipgloss.Color("#475569"),
		Slate700: lipgloss.Color("#334155"),
		Slate800: lipgloss.Color("#1e293b"),
		Slate900: lipgloss.Color("#0f172a"),
		Slate950: lipgloss.Color("#020617"),

		Gray50:  lipgloss.Color("#f9fafb"),
		Gray100: lipgloss.Color("#f3f4f6"),
		Gray200: lipgloss.Color("#e5e7eb"),
		Gray300: lipgloss.Color("#d1d5db"),
		Gray400: lipgloss.Color("#9ca3af"),
		Gray500: lipgloss.Color("#6b7280"),
		Gray600: lipgloss.Color("#4b5563"),
		Gray700: lipgloss.Color("#374151"),
		Gray800: lipgloss.Color("#1f2937"),
		Gray900: lipgloss.Color("#111827"),
		Gray950: lipgloss.Color("#030712"),

		Zinc50:  lipgloss.Color("#fafafa"),
		Zinc100: lipgloss.Color("#f4f4f5"),
		Zinc200: lipgloss.Color("#e4e4e7"),
		Zinc300: lipgloss.Color("#d4d4d8"),
		Zinc400: lipgloss.Color("#a1a1aa"),
		Zinc500: lipgloss.Color("#71717a"),
		Zinc600: lipgloss.Color("#52525b"),
		Zinc700: lipgloss.Color("#3f3f46"),
		Zinc800: lipgloss.Color("#27272a"),
		Zinc900: lipgloss.Color("#18181b"),
		Zinc950: lipgloss.Color("#09090b"),

		Neutral50:  lipgloss.Color("#fafafa"),
		Neutral100: lipgloss.Color("#f5f5f5"),
		Neutral200: lipgloss.Color("#e5e5e5"),
		Neutral300: lipgloss.Color("#d4d4d4"),
		Neutral400: lipgloss.Color("#a3a3a3"),
		Neutral500: lipgloss.Color("#737373"),
		Neutral600: lipgloss.Color("#525252"),
		Neutral700: lipgloss.Color("#404040"),
		Neutral800: lipgloss.Color("#262626"),
		Neutral900: lipgloss.Color("#171717"),
		Neutral950: lipgloss.Color("#0a0a0a"),

		Stone50:  lipgloss.Color("#fafaf9"),
		Stone100: lipgloss.Color("#f5f5f4"),
		Stone200: lipgloss.Color("#e7e5e4"),
		Stone300: lipgloss.Color("#d6d3d1"),
		Stone400: lipgloss.Color("#a8a29e"),
		Stone500: lipgloss.Color("#78716c"),
		Stone600: lipgloss.Color("#57534e"),
		Stone700: lipgloss.Color("#44403c"),
		Stone800: lipgloss.Color("#292524"),
		Stone900: lipgloss.Color("#1c1917"),
		Stone950: lipgloss.Color("#0c0a09"),

		Black: lipgloss.Color("#000000"),
		White: lipgloss.Color("#ffffff"),
	}
}
