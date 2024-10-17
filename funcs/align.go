package funcs

import (
	"fmt"
)

// PrintAsciiArtWithAlign prints ASCII art text aligned according to the specified alignment.
// Parameters:
// - align: A string that specifies the alignment type ("left", "right", "center", "justify").
// - width: An integer that represents the total width for alignment.
// - sentences: A slice of strings containing the sentences to print.
// - textFile: A slice of strings representing the ASCII art lines for each character.
func PrintAsciiArtWithAlign(align string, width int, sentences []string, textFile []string) {
	// Slice to store the total length of each sentence in ASCII art representation
	var length []int

	// Loop through each sentence to calculate the total length of ASCII art for each word
	for _, word := range sentences {
		if word == "" {
			// If the sentence is empty (i.e., there was a '\n'), skip to the next iteration
			continue
		}
		n := 0
		// Calculate the length of the ASCII art for the current word
		for j := 0; j < len(word); j++ {
			for lineIndex, line := range textFile {
				// Each character's ASCII art is 9 lines tall, starting from lineIndex
				if lineIndex == (int(word[j])-32)*9+1 {
					n += len(line) // Accumulate the length of each character's line
				}
			}
		}
		length = append(length, n) // Store the total length for the current word
	}

	// Handle left alignment by directly printing the ASCII art
	if align == "left" {
		PrintAsciiArt(sentences, textFile) // Call function to print ASCII art
	}

	// Handle right alignment
	if align == "right" {
		for i, word := range sentences {
			if word == "" {
				// Print an empty line if the sentence is empty
				if i != 0 {
					fmt.Println()
				}
				continue
			}
			for h := 1; h < 9; h++ { // Loop through each line of ASCII art
				// Print leading spaces to right-align the text
				if width > length[i] {
					for k := 0; k < width-length[i]; k++ {
						fmt.Print(" ") // Print spaces for right alignment
					}
				}
				// Print the ASCII art for each character in the word
				for j := 0; j < len(word); j++ {
					for lineIndex, line := range textFile {
						if lineIndex == (int(word[j])-32)*9+h {
							fmt.Print(line) // Print the corresponding ASCII line
						}
					}
				}
				fmt.Println() // Move to the next line
			}
		}
	}

	// Handle center alignment
	if align == "center" {
		for i, word := range sentences {
			if word == "" {
				// Print an empty line if the sentence is empty
				if i != 0 {
					fmt.Println()
				}
				continue
			}
			for h := 1; h < 9; h++ { // Loop through each line of ASCII art
				// Print leading spaces to center the text
				if width > length[i] {
					for k := 0; k < (width-length[i])/2; k++ {
						fmt.Print(" ") // Print spaces for center alignment
					}
				}
				// Print the ASCII art for each character in the word
				for j := 0; j < len(word); j++ {
					for lineIndex, line := range textFile {
						if lineIndex == (int(word[j])-32)*9+h {
							fmt.Print(line) // Print the corresponding ASCII line
						}
					}
				}
				fmt.Println() // Move to the next line
			}
		}
	}

	// Handle justified alignment
	if align == "justify" {
		for i, word := range sentences {
			if word == "" {
				// Print an empty line if the sentence is empty
				if i != 0 {
					fmt.Println()
				}
				continue
			}
			count := 0
			// Count the number of spaces in the current word
			for _, ch := range word {
				if ch == ' ' {
					count++ // Increment count for each space
				}
			}
			for h := 1; h < 9; h++ { // Loop through each line of ASCII art
				for j := 0; j < len(word); j++ {
					if count == 0 || width == length[i] {
						// If there are no spaces or the length matches the width, print normally
						for lineIndex, line := range textFile {
							if lineIndex == (int(word[j])-32)*9+h {
								fmt.Print(line) // Print the corresponding ASCII line
							}
						}
					} else {
						// Calculate the space to distribute between words
						space := (width - length[i]) / count
						for lineIndex, line := range textFile {
							if lineIndex == (int(word[j])-32)*9+h {
								fmt.Print(line) // Print the corresponding ASCII line
								if word[j] == ' ' {
									// Print distributed spaces after each word
									for d := 0; d < space; d++ {
										fmt.Print(" ") // Print spaces between words
									}
								}
							}
						}
					}
				}
				fmt.Println() // Move to the next line
			}
		}
	}
}
