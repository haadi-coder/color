package color

import "fmt"

// Black returns the given text in black color.
func Black(text string) string {
	return Style(text, AttrFgBlack)
}

// Blackf formats the given string using fmt.Sprintf and returns it in black color.
func Blackf(format string, a ...interface{}) string {
	return Black(fmt.Sprintf(format, a...))
}

// BrightBlack returns the given text in bright black color.
func BrightBlack(text string) string {
	return Style(text, AttrFgBrightBlack)
}

// BgBlack returns the given text with black background.
func BgBlack(text string) string {
	return Style(text, AttrBgBlack)
}

// BgBrightBlack returns the given text with bright black background.
func BgBrightBlack(text string) string {
	return Style(text, AttrBgBrightBlack)
}

// Red returns the given text in red color.
func Red(text string) string {
	return Style(text, AttrFgRed)
}

// Redf formats the given string using fmt.Sprintf and returns it in red color.
func Redf(format string, a ...interface{}) string {
	return Red(fmt.Sprintf(format, a...))
}

// BrightRed returns the given text in bright red color.
func BrightRed(text string) string {
	return Style(text, AttrFgBrightRed)
}

// BgRed returns the given text with red background.
func BgRed(text string) string {
	return Style(text, AttrBgRed)
}

// BgBrightRed returns the given text with bright red background.
func BgBrightRed(text string) string {
	return Style(text, AttrBgBrightRed)
}

// Green returns the given text in green color.
func Green(text string) string {
	return Style(text, AttrFgGreen)
}

// Greenf formats the given string using fmt.Sprintf and returns it in green color.
func Greenf(format string, a ...interface{}) string {
	return Green(fmt.Sprintf(format, a...))
}

// BrightGreen returns the given text in bright green color.
func BrightGreen(text string) string {
	return Style(text, AttrFgBrightGreen)
}

// BgGreen returns the given text with green background.
func BgGreen(text string) string {
	return Style(text, AttrBgGreen)
}

// BgBrightGreen returns the given text with bright green background.
func BgBrightGreen(text string) string {
	return Style(text, AttrBgBrightGreen)
}

// Yellow returns the given text in yellow color.
func Yellow(text string) string {
	return Style(text, AttrFgYellow)
}

// Yellowf formats the given string using fmt.Sprintf and returns it in yellow color.
func Yellowf(format string, a ...interface{}) string {
	return Yellow(fmt.Sprintf(format, a...))
}

// BrightYellow returns the given text in bright yellow color.
func BrightYellow(text string) string {
	return Style(text, AttrFgBrightYellow)
}

// BgYellow returns the given text with yellow background.
func BgYellow(text string) string {
	return Style(text, AttrBgYellow)
}

// BgBrightYellow returns the given text with bright yellow background.
func BgBrightYellow(text string) string {
	return Style(text, AttrBgBrightYellow)
}

// Blue returns the given text in blue color.
func Blue(text string) string {
	return Style(text, AttrFgBlue)
}

// Bluef formats the given string using fmt.Sprintf and returns it in blue color.
func Bluef(format string, a ...interface{}) string {
	return Blue(fmt.Sprintf(format, a...))
}

// BrightBlue returns the given text in bright blue color.
func BrightBlue(text string) string {
	return Style(text, AttrFgBrightBlue)
}

// BgBlue returns the given text with blue background.
func BgBlue(text string) string {
	return Style(text, AttrBgBlue)
}

// BgBrightBlue returns the given text with bright blue background.
func BgBrightBlue(text string) string {
	return Style(text, AttrBgBrightBlue)
}

// Magenta returns the given text in magenta color.
func Magenta(text string) string {
	return Style(text, AttrFgMagenta)
}

// Magentaf formats the given string using fmt.Sprintf and returns it in magenta color.
func Magentaf(format string, a ...interface{}) string {
	return Magenta(fmt.Sprintf(format, a...))
}

// BrightMagenta returns the given text in bright magenta color.
func BrightMagenta(text string) string {
	return Style(text, AttrFgBrightMagenta)
}

// BgMagenta returns the given text with magenta background.
func BgMagenta(text string) string {
	return Style(text, AttrBgMagenta)
}

// BgBrightMagenta returns the given text with bright magenta background.
func BgBrightMagenta(text string) string {
	return Style(text, AttrBgBrightMagenta)
}

// Cyan returns the given text in cyan color.
func Cyan(text string) string {
	return Style(text, AttrFgCyan)
}

// Cyanf formats the given string using fmt.Sprintf and returns it in cyan color.
func Cyanf(format string, a ...interface{}) string {
	return Cyan(fmt.Sprintf(format, a...))
}

// BrightCyan returns the given text in bright cyan color.
func BrightCyan(text string) string {
	return Style(text, AttrFgBrightCyan)
}

// BgCyan returns the given text with cyan background.
func BgCyan(text string) string {
	return Style(text, AttrBgCyan)
}

// BgBrightCyan returns the given text with bright cyan background.
func BgBrightCyan(text string) string {
	return Style(text, AttrBgBrightCyan)
}

// White returns the given text in white color.
func White(text string) string {
	return Style(text, AttrFgWhite)
}

// Whitef formats the given string using fmt.Sprintf and returns it in white color.
func Whitef(format string, a ...interface{}) string {
	return White(fmt.Sprintf(format, a...))
}

// BrightWhite returns the given text in bright white color.
func BrightWhite(text string) string {
	return Style(text, AttrFgBrightWhite)
}

// BgWhite returns the given text with white background.
func BgWhite(text string) string {
	return Style(text, AttrBgWhite)
}

// BgBrightWhite returns the given text with bright white background.
func BgBrightWhite(text string) string {
	return Style(text, AttrBgBrightWhite)
}

// Bold returns the given text in bold style.
func Bold(text string) string {
	return Style(text, AttrBold)
}

// Italic returns the given text in italic style.
func Italic(text string) string {
	return Style(text, AttrItalic)
}

// Underline returns the given text underlined.
func Underline(text string) string {
	return Style(text, AttrUnderline)
}

// Dim returns the given text in dimmed style.
func Dim(text string) string {
	return Style(text, AttrDim)
}
