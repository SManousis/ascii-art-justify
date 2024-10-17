package funcs

import (
	"fmt"
	"syscall"
	"unsafe"
)

type winsize struct {
	Row    uint16 // Number of rows (height) in the terminal
	Col    uint16 // Number of columns (width) in the terminal
	Xpixel uint16 // Width of the terminal in pixels
	Ypixel uint16 // Height of the terminal in pixels
}

// TerminalSize retrieves the current size of the terminal window.
// It returns the width and height of the terminal in characters,
// along with an error if there is an issue obtaining the size.
func TerminalSize() (width int, height int, err error) {
	ws := &winsize{} // Create an instance of the winsize struct to hold terminal size information

	// Use the syscall package to make a system call to retrieve terminal size
	retCode, _, err := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdout),     // Standard output file descriptor
		uintptr(syscall.TIOCGWINSZ), // Command to get window size
		uintptr(unsafe.Pointer(ws)), // Pointer to winsize struct to populate with size data
	)

	// Check if the system call was successful
	if int(retCode) == -1 {
		// Print an error message if getting terminal size failed
		fmt.Println("Error getting terminal size:", err)
		return 0, 0, err // Return zero values for width and height
	}

	// Return the terminal size as width and height
	return int(ws.Col), int(ws.Row), nil
}
