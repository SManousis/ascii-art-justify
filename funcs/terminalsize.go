package funcs

import (
	"fmt"
	"syscall"
	"unsafe"
)

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func TerminalSize() (width int, height int, err error) {
	ws := &winsize{}
	retCode, _, err := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdout),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)),
	)
	if int(retCode) == -1 {
		fmt.Println("Error getting terminal size:", err)
		return 0, 0, err
	}
	return int(ws.Col), int(ws.Row), nil
}
