package helper

import (
	"fmt"
	"os"
	"os/exec"
)

func ClearScreen() {
	cmd:= exec.Command("cmd", "/c", "/cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// back to menu
func BackHandler()  {
	fmt.Print("Click enter, for back to Menu")
	var back int
	fmt.Scanln(&back)
}