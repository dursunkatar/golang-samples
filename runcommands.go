package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {

	// Print Go Version
	cmdOutput, err := exec.Command("tasklist").Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", cmdOutput)
}
