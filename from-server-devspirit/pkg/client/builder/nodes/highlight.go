package nodes

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/fatih/color"
	"github.com/zyedidia/highlight"
)

func out(inputString string) {

	// Load the go syntax file
	// Make sure that the syntax_files directory is in the current directory
	syntaxFile, err := ioutil.ReadFile("go.yml")
	if err != nil {
		panic(err)
	}

	// Parse it into a `*highlight.Def`
	syntaxDef, err := highlight.ParseDef(syntaxFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Make a new highlighter from the definition
	h := highlight.NewHighlighter(syntaxDef)
	// Highlight the string
	// Matches is an array of maps which point to groups
	// matches[lineNum][colNum] will give you the change in group at that line and column number
	// Note that there is only a group at a line and column number if the syntax highlighting changed at that position
	matches := h.HighlightString(inputString)

	// We split the string into a bunch of lines
	// Now we will print the string
	lines := strings.Split(inputString, "\n")
	for lineN, l := range lines {
		for colN, c := range l {
			// Check if the group changed at the current position
			if group, ok := matches[lineN][colN]; ok {
				// Check the group name and set the color accordingly (the colors chosen are arbitrary)
				if group == highlight.Groups["statement"] {
					color.Set(color.FgGreen)
				} else if group == highlight.Groups["preproc"] {
					color.Set(color.FgHiRed)
				} else if group == highlight.Groups["special"] {
					color.Set(color.FgHiYellow)
				} else if group == highlight.Groups["constant.string"] {
					color.Set(color.FgCyan)
				} else if group == highlight.Groups["constant.specialChar"] {
					color.Set(color.FgHiMagenta)
				} else if group == highlight.Groups["type"] {
					color.Set(color.FgHiCyan)
				} else if group == highlight.Groups["constant.number"] {
					color.Set(color.FgCyan)
				} else if group == highlight.Groups["comment"] {
					color.Set(color.FgHiGreen)
				} else {
					color.Unset()
				}
			}
			// Print the character
			fmt.Print(string(c))
		}
		// This is at a newline, but highlighting might have been turned off at the very end of the line so we should check that.
		if group, ok := matches[lineN][len(l)]; ok {
			if group == highlight.Groups["default"] || group == highlight.Groups[""] {
				color.Unset()
			}
		}

		fmt.Print("\n")
	}
}
