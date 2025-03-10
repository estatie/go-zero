package cmd

import (
	"fmt"
	"runtime"

	"github.com/estatie/go-zero/tools/goctl/vars"
	"github.com/logrusorgru/aurora"
)

var colorRender = []func(v any) string{
	func(v any) string {
		return aurora.BrightRed(v).String()
	},
	func(v any) string {
		return aurora.BrightGreen(v).String()
	},
	func(v any) string {
		return aurora.BrightYellow(v).String()
	},
	func(v any) string {
		return aurora.BrightBlue(v).String()
	},
	func(v any) string {
		return aurora.BrightMagenta(v).String()
	},
	func(v any) string {
		return aurora.BrightCyan(v).String()
	},
}

func blue(s string) string {
	if runtime.GOOS == vars.OsWindows {
		return s
	}

	return aurora.BrightBlue(s).String()
}

func green(s string) string {
	if runtime.GOOS == vars.OsWindows {
		return s
	}

	return aurora.BrightGreen(s).String()
}

func rainbow(s string) string {
	if runtime.GOOS == vars.OsWindows {
		return s
	}
	s0 := s[0]
	return colorRender[int(s0)%(len(colorRender)-1)](s)
}

// rpadx adds padding to the right of a string.
func rpadx(s string, padding int) string {
	template := fmt.Sprintf("%%-%ds", padding)
	return rainbow(fmt.Sprintf(template, s))
}
