package ledger

import (
	"github.com/fatih/color"
)

func prettyGreen(prettyText interface{}) string {
	green := color.New(color.FgGreen).SprintFunc()
	return green(prettyText)
}

func prettyGreenBold(prettyText interface{}) string {
	green := color.New(color.FgGreen).Add(color.Bold).SprintFunc()
	return green(prettyText)
}

func prettyBlue(prettyText interface{}) string {
	green := color.New(color.FgBlue).SprintFunc()
	return green(prettyText)
}

func prettyBlueBold(prettyText interface{}) string {
	green := color.New(color.FgBlue).Add(color.Bold).SprintFunc()
	return green(prettyText)
}

func prettyYellow(prettyText interface{}) string {
	green := color.New(color.FgYellow).SprintFunc()
	return green(prettyText)
}

func prettyYellowBold(prettyText interface{}) string {
	green := color.New(color.FgYellow).Add(color.Bold).SprintFunc()
	return green(prettyText)
}

func prettyRed(prettyText interface{}) string {
	green := color.New(color.FgRed).SprintFunc()
	return green(prettyText)
}

func prettyRedBold(prettyText interface{}) string {
	green := color.New(color.FgRed).Add(color.Bold).SprintFunc()
	return green(prettyText)
}

func prettyMagenta(prettyText interface{}) string {
	green := color.New(color.FgMagenta).SprintFunc()
	return green(prettyText)
}

func prettyMagentaBold(prettyText interface{}) string {
	green := color.New(color.FgMagenta).Add(color.Bold).SprintFunc()
	return green(prettyText)
}

func prettyCyan(prettyText interface{}) string {
	green := color.New(color.FgCyan).SprintFunc()
	return green(prettyText)
}

func prettyCyanBold(prettyText interface{}) string {
	green := color.New(color.FgCyan).Add(color.Bold).SprintFunc()
	return green(prettyText)
}
