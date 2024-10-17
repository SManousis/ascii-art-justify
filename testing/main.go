package main

import (
	"fmt"
	"os"
	"strings"

	"ascii/funcs"
)

func main() {
	// Check if the number of command-line arguments is correct
	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER] || EX: go run . --align=right something standard")
		return
	}

	var text string
	var option string
	banner := "standard"
	var alignment string
	hasAlign := false

	if len(os.Args) == 4 {
		option = os.Args[1]
		text = os.Args[2]
		banner = strings.ToLower(os.Args[3])
		hasAlign = true

		if !strings.HasPrefix(option, "--align") || !strings.HasSuffix(option, ".txt") {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER] || EX: go run . --align=right something standard")
			return
		}

		alignment = option[8:]
	}

	if len(os.Args) == 3 {
		if strings.HasPrefix(option, "--output=") && strings.HasSuffix(option, ".txt") {
			option = os.Args[1]
			text = os.Args[2]
			hasAlign = true
		} else if strings.HasPrefix(option, "--align=") && !strings.HasSuffix(option, ".txt") {

			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER] || EX: go run . --align=right something standard")
			return
		} else {

			text = os.Args[1]
			banner = strings.ToLower(os.Args[2])
		}
	}

	if len(os.Args) == 2 {
		text = os.Args[1]
	}

	// Split the input text by '\n' to handle newlines in the input
	sepText := strings.Split(text, "\\n")

	// Read the banner file (e.g., 'standard.txt', 'shadow.txt')
	file, err := os.ReadFile("banners/" + banner + ".txt")
	if err != nil {
		fmt.Println("Banner not found")
		return
	}

	// Convert the file content to a string and normalize line endings
	str := string(file)
	str = strings.Replace(str, "\r\n", "\n", -1) // Replace Windows line endings with Unix ones

	// Split the banner file content into lines
	lines := strings.Split(str, "\n")

	if !hasAlign {
		// If no output file specified, print the ASCII art to the console
		funcs.PrintAsciiArt(sepText, lines)
	}
	if hasAlign {
		alignment = strings.ToLower(option[8:])
		width, _, _ := funcs.TerminalSize()
		if alignment == "" {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER] || EX: go run . --align=right something standard")
			return
		} else {
			switch alignment {
			case "left", "right", "center", "justify":
			default:
				fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER] || EX: go run . --align=right something standard")
				return
			}
		}
		funcs.PrintAsciiArtWithAlign(alignment, width, sepText, lines)
	}
}
