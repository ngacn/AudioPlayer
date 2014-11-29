package main

import (
	"fmt"
	"gopkg.in/qml.v1"
    "log"
	"os"
    "os/exec"
)

func main() {
    cmd := exec.Command("warpten", "-d")
    err := cmd.Start()
    if err != nil {
        log.Fatal(err)
    }
    // log.Printf("Waiting for command to finish...")
    // err = cmd.Wait()
    // log.Printf("Command finished with error: %v", err)

	if err = qml.Run(run); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	engine := qml.NewEngine()

	controls, err := engine.LoadFile("main.qml")
	if err != nil {
		return err
	}

	window := controls.CreateWindow(nil)

	window.Show()
	window.Wait()
	return nil
}
