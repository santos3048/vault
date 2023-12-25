package color

import (
	"fmt"
	"os"
	"golang.org/x/sys/windows"
	) 

func Init(){
	var outMode uint32
	out := windows.Handle(os.Stdout.Fd())
	if err := windows.GetConsoleMode(out, &outMode); err != nil {
		return
	}
	outMode |= windows.ENABLE_PROCESSED_OUTPUT | windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING
	_ = windows.SetConsoleMode(out, outMode)
}

func Set(cl string){
	switch cl{
	case "red":
		fmt.Print("\x1b[31m")
	case "bold":
		fmt.Print("\x1b[1m")
	case "blue":
		fmt.Print("\x1b[34m")
	case "cyan":
		fmt.Print("\x1b[36m")
	case "green":
		fmt.Print("\x1b[32m")
	default:
		fmt.Print("\x1b[0m")
	}
}

func Unset() {
	fmt.Print("\x1b[0m")
}