package funcs

import (
	"fmt"
)

func PrintAsciiArtWithAlign(align string, width int, sentences []string, textFile []string) {
	// Loop through each sentence (split by '\n')
	var length []int

	for _, word := range sentences {
		if word == "" {
			// If the sentence is empty (i.e., there was a '\n'), print an empty line

			continue
		}
		n := 0
		for j := 0; j < len(word); j++ {
			for lineIndex, line := range textFile {
				if lineIndex == (int(word[j])-32)*9+1 {
					n += len(line)
				}
			}
		}
		length = append(length, n)
	}

	if align == "left" {
		PrintAsciiArt(sentences, textFile)
	}
	if align == "right" {
		for i, word := range sentences {
			if word == "" {
				if i != 0 {
					fmt.Println()
				}
				continue
			}
			for h := 1; h < 9; h++ {
				if width > length[i] {
					for k := 0; k < width-length[i]; k++ {
						fmt.Print(" ")
					}
				}
				for j := 0; j < len(word); j++ {
					for lineIndex, line := range textFile {
						if lineIndex == (int(word[j])-32)*9+h {
							fmt.Print(line)
						}
					}
				}
				fmt.Println()
			}
		}
	}
	if align == "center" {
		for i, word := range sentences {
			if word == "" {
				if i != 0 {
					fmt.Println()
				}
				continue
			}
			for h := 1; h < 9; h++ {
				if width > length[i] {
					for k := 0; k < (width-length[i])/2; k++ {
						fmt.Print(" ")
					}
				}
				for j := 0; j < len(word); j++ {
					for lineIndex, line := range textFile {
						if lineIndex == (int(word[j])-32)*9+h {
							fmt.Print(line)
						}
					}
				}
				fmt.Println()
			}
		}
	}
	if align == "justify" {
		for i, word := range sentences {
			if word == "" {
				if i != 0 {
					fmt.Println()
				}
				continue
			}
			count := 0
			for _, ch := range word {
				if ch == ' ' {
					count++
				}
			}
			for h := 1; h < 9; h++ {
				for j := 0; j < len(word); j++ {
					if count == 0 || width == length[i] {
						for lineIndex, line := range textFile {
							if lineIndex == (int(word[j])-32)*9+h {
								fmt.Print(line)
							}
						}
					} else {
						space := (width - length[i]) / count
						for lineIndex, line := range textFile {
							if lineIndex == (int(word[j])-32)*9+h {
								fmt.Print(line)
								if word[j] == ' ' {
									for d := 0; d < space; d++ {
										fmt.Print(" ")
									}
								}
							}
						}
					}
				}
				fmt.Println()
			}
		}
	}
}
