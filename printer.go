package printer

import (
	"fmt"
	"os"
	"strings"
	"syscall"
	"unsafe"
)

// Enables verbose printing
var Verbose = false

const line = "━"
const bar = "┃ "
const cornerTop = "┏"
const cornerBottom = "┗"

const leftPad = 3

func Success(text string, a ...interface{}) {
	fmt.Printf(Green("✔ ")+text+"\n", a...)
}

func Fail(text string, a ...interface{}) {
	fmt.Printf(Red("✗ ")+text+"\n", a...)
}

func Error(text string, a ...interface{}) {
	fmt.Printf(Red("✗ ")+text+"\n", a...)
}

func Info(text string, a ...interface{}) {
	fmt.Printf(Blue("🐧  ")+text+"\n", a...)
}

func Warning(text string, a ...interface{}) {
	fmt.Printf(Yellow("⚠ ")+text+"\n", a...)
}

func SuccessBar(text string, a ...interface{}) {
	fmt.Printf(Green(bar)+text+"\n", a...)
}

func ErrorBar(text string, a ...interface{}) {
	fmt.Printf(Red(bar)+text+"\n", a...)
}

func InfoBar(text string, a ...interface{}) {
	fmt.Printf(Blue(bar)+text+"\n", a...)
}

func WarningBar(text string, a ...interface{}) {
	fmt.Printf(Yellow(bar)+text+"\n", a...)
}

func SuccessBarIcon(text string, a ...interface{}) {
	fmt.Printf(Greenf("%s✔ ", bar)+text+"\n", a...)
}

func ErrorBarIcon(text string, a ...interface{}) {
	fmt.Printf(Redf("%s✗ ", bar)+text+"\n", a...)
}

func InfoBarIcon(text string, a ...interface{}) {
	fmt.Printf(Bluef("%s🐧 ", bar)+text+"\n", a...)
}

func WarningBarIcon(text string, a ...interface{}) {
	fmt.Printf(Yellowf("%s⚠ ", bar)+text+"\n", a...)
}

func SuccessLine() {
	width, _, _ := getSize(int(os.Stdout.Fd()))
	fmt.Printf(Green(strings.Repeat(line, width)))
}

func ErrorLine() {
	width, _, _ := getSize(int(os.Stdout.Fd()))
	fmt.Printf(Red(strings.Repeat(line, width)))
}

func InfoLine() {
	width, _, _ := getSize(int(os.Stdout.Fd()))
	fmt.Printf(Blue(strings.Repeat(line, width)))
}

func WarningLine() {
	width, _, _ := getSize(int(os.Stdout.Fd()))
	fmt.Printf(Yellow(strings.Repeat(line, width)))
}

func SuccessLineText(text string, a ...interface{}) {
	// text = fmt.Sprintf(text, a...)
	// width, _, _ := getSize(int(os.Stdout.Fd()))

	// fmt.Printf(Green(strings.Repeat(line, 3)) + Nc strings.Repeat(line, width) + Nc
}

func ErrorLineText(text string, a ...interface{}) {
	width, _, _ := getSize(int(os.Stdout.Fd()))
	fmt.Printf(Red("%s"), strings.Repeat(line, width))
}

func InfoLineText(text string, a ...interface{}) {
	// width, _, _ := getSize(int(os.Stdout.Fd()))
	// fmt.Printf(Blue + strings.Repeat(line, width) + Nc)
}

func WarningLineText(text string, a ...interface{}) {
	// width, _, _ := getSize(int(os.Stdout.Fd()))
	// fmt.Printf(Yellow + strings.Repeat(line, width) + Nc)
}

func SuccessLineTop() {
	width, _, _ := getSize(int(os.Stdout.Fd()))
	fmt.Printf(Bold(Green(cornerTop + strings.Repeat(line, width-1))))
}

func ErrorLineTop() {
	width, _, _ := getSize(int(os.Stdout.Fd()))
	fmt.Printf(Bold(Red(cornerTop + strings.Repeat(line, width-1))))
}

func InfoLineTop() {
	width, _, _ := getSize(int(os.Stdout.Fd()))
	fmt.Printf(Bold(Blue(cornerTop + strings.Repeat(line, width-1))))
}

func WarningLineTop() {
	width, _, _ := getSize(int(os.Stdout.Fd()))
	fmt.Printf(Bold(Yellow(cornerTop + strings.Repeat(line, width-1))))
}

func SuccessLineTextTop(text string, a ...interface{}) {
	width, _, _ := getSize(int(os.Stdout.Fd()))
	fmt.Printf(Bold(Green(cornerTop + strings.Repeat(line, width-1))))
}

func ErrorLineTextTop(text string, a ...interface{}) {
	width, _, _ := getSize(int(os.Stdout.Fd()))
	fmt.Printf(Bold(Red(cornerTop + strings.Repeat(line, width-1))))
}

func InfoLineTextTop(text string, a ...interface{}) {
	width, _, _ := getSize(int(os.Stdout.Fd()))
	fmt.Printf(Bold(Blue(cornerTop + strings.Repeat(line, width-1))))
}

func WarningLineTextTop(text string, a ...interface{}) {
	width, _, _ := getSize(int(os.Stdout.Fd()))
	fmt.Printf(Bold(Yellow(cornerTop + strings.Repeat(line, width-1))))
}

func SuccessLineBottom() {
	width, _, _ := getSize(int(os.Stdout.Fd()))
	fmt.Printf(Bold(Green(cornerBottom + strings.Repeat(line, width-1))))
}

func ErrorLineBottom() {
	width, _, _ := getSize(int(os.Stdout.Fd()))
	fmt.Printf(Bold(Red(cornerBottom + strings.Repeat(line, width-1))))
}

func InfoLineBottom() {
	width, _, _ := getSize(int(os.Stdout.Fd()))
	fmt.Printf(Bold(Blue(cornerBottom + strings.Repeat(line, width-1))))
}

func WarningLineBottom() {
	width, _, _ := getSize(int(os.Stdout.Fd()))
	fmt.Printf(Bold(Yellow(cornerBottom + strings.Repeat(line, width-1))))
}

func VerboseSuccess(text string, a ...interface{}) {
	if Verbose {
		Success(text, a...)
	}
}

func VerboseFail(text string, a ...interface{}) {
	if Verbose {
		Fail(text, a...)
	}
}

func VerboseError(text string, a ...interface{}) {
	if Verbose {
		Error(text, a...)
	}
}

func VerboseInfo(text string, a ...interface{}) {
	if Verbose {
		Info(text, a...)
	}
}

func VerboseWarning(text string, a ...interface{}) {
	if Verbose {
		Warning(text, a...)
	}
}

func VerboseSuccessBar(text string, a ...interface{}) {
	if Verbose {
		SuccessBar(text, a...)
	}
}

func VerboseErrorBar(text string, a ...interface{}) {
	if Verbose {
		ErrorBar(text, a...)
	}
}

func VerboseInfoBar(text string, a ...interface{}) {
	if Verbose {
		InfoBar(text, a...)
	}
}

func VerboseWarningBar(text string, a ...interface{}) {
	if Verbose {
		WarningBar(text, a...)
	}
}

// GetSize returns the dimensions of the given terminal.
// https://github.com/golang/crypto/blob/master/ssh/terminal/util.go#L80
func getSize(fd int) (width, height int, err error) {
	var dimensions [4]uint16

	if _, _, err := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(fd), uintptr(syscall.TIOCGWINSZ), uintptr(unsafe.Pointer(&dimensions)), 0, 0, 0); err != 0 {
		return -1, -1, err
	}
	return int(dimensions[1]), int(dimensions[0]), nil
}
