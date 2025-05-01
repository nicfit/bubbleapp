package style

var smallCapsMap = map[rune]rune{
	'A': 'ᴀ', // U+1D00
	'B': 'ʙ', // U+0299
	'C': 'ᴄ', // U+1D04
	'D': 'ᴅ', // U+1D05
	'E': 'ᴇ', // U+1D07
	'F': 'ꜰ', // U+A730
	'G': 'ɢ', // U+0262
	'H': 'ʜ', // U+029C
	'I': 'ɪ', // U+026A
	'J': 'ᴊ', // U+1D0A
	'K': 'ᴋ', // U+1D0B
	'L': 'ʟ', // U+029F
	'M': 'ᴍ', // U+1D0D
	'N': 'ɴ', // U+0274
	'O': 'ᴏ', // U+1D0F
	'P': 'ᴘ', // U+1D18
	'Q': 'ꞯ', // U+A7AF
	'R': 'ʀ', // U+0280
	'S': 'ꜱ', // U+A731
	'T': 'ᴛ', // U+1D1B
	'U': 'ᴜ', // U+1D1C
	'V': 'ᴠ', // U+1D20
	'W': 'ᴡ', // U+1D21
	'X': 'x', // U+0078 (Note: Standard lowercase x)
	'Y': 'ʏ', // U+028F
	'Z': 'ᴢ', // U+1D22
}

func ConvertToSmallCaps(input string) string {
	var result []rune
	for _, r := range input {
		if smallCap, ok := smallCapsMap[r]; ok {
			result = append(result, smallCap)
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}
