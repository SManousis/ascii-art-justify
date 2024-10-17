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
		fmt.Println("ERROR! Usage: go run . [OPTION] [STRING] [BANNER] \nEX: go run . --output=<fileName.txt> \"something\" standard")
		return
	}

	var text string
	var option string
	banner := "standard"
	var alignment string
	var outputfile string
	hasAlign := false
	hasOutput := false

	if len(os.Args) > 2 {
		if len(os.Args) == 3 {
			if !strings.HasPrefix(os.Args[1], "--") {
				text = os.Args[1]
				banner = strings.ToLower(os.Args[2])
			} else {
				option = os.Args[1]
				text = os.Args[2]
			}
		} else {
			option = os.Args[1]
			text = os.Args[2]
			banner = strings.ToLower(os.Args[3])
		}
		if option != "" {
			if strings.HasPrefix(option, "--output=") && strings.HasSuffix(option, ".txt") {
				hasOutput = true
				outputfile = option[9:]
			} else if strings.HasPrefix(option, "--align=") {
				hasAlign = true
				alignment = strings.ToLower(option[8:])
			} else {
				fmt.Println("ERROR! Usage: go run . [OPTION] [STRING] [BANNER] \nEX: go run . --output=<fileName.txt> \"something\" standard")
				return
			}
		}
	} else {
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

	if option == "" {
		// If no output file specified, print the ASCII art to the console
		funcs.PrintAsciiArt(sepText, lines)
	} else if hasAlign {
		width, _, _ := funcs.TerminalSize()
		if alignment == "" {
			fmt.Println("ERROR! Usage: go run . [OPTION] [STRING] [BANNER] \nEX: go run . --output=<fileName.txt> \"something\" standard")
			return
		} else {
			switch alignment {
			case "left", "right", "center", "justify":
			default:
				fmt.Println("ERROR! Acceptable alignments are: left, right, center, justify")
				return
			}
		}
		funcs.PrintAsciiArtWithAlign(alignment, width, sepText, lines)
	} else if hasOutput {
		// If output file is specified, create or overwrite the file
		file, err := os.Create(outputfile)
		if err != nil {
			fmt.Println("ERROR!", err)
			return
		}
		// Generate the ASCII art and write it to the file
		funcs.PrintAsciiArtToFile(file, sepText, lines)
	}
}
