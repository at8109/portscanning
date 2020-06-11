package main

import (
	"fmt"

	"github.com/at8109/portscanning/port"
)

func main() {
	fmt.Println("Port Scanner in Go")
	open := port.ScanPort("tcp", "localhost", 4000)
	fmt.Printf("Port Open: %t\n", open)

	results := port.IntialScan("localhost")
	fmt.Println(results)
}
