package main

import (
	"AutoGenProtocol/src/autogen"
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("---------------------------auto gen protocol start---------------------------")
	autogen.AutoGenProtocol()
	fmt.Println("---------------------------auto gen protocol end-----------------------------")
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\r')
}
