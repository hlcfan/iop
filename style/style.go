package style

import "fmt"

type Color interface {
	toRGB() (RGB, error)
}

type Style struct {
	IdentifierColorHex Color
	NumberColorHex     Color
	StringColorHex     Color
	BoolColorHex       Color
	DelimiterColorHex  Color
}

func (s Style) PrintIdentifier(text string) string {
	return textWithColor(s.IdentifierColorHex, text)
}

func (s Style) PrintNumber(text string) string {
	return textWithColor(s.NumberColorHex, text)
}

func (s Style) PrintString(text string) string {
	return textWithColor(s.StringColorHex, text)
}

func (s Style) PrintBool(text string) string {
	return textWithColor(s.BoolColorHex, text)
}

func (s Style) PrintCharacter(text string) string {
	return textWithColor(s.DelimiterColorHex, text)
}

func textWithColor(colorHex Color, text string) string {
	rgb, err := colorHex.toRGB()
	if err != nil {
		return text
	}
	return fmt.Sprintf("\x1b[%sm%s\x1b[0m", rgb, text)
}
