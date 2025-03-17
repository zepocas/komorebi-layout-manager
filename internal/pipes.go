package internal

import (
	"fmt"
	"github.com/natefinch/npipe"
	"log"
	"os"
	"os/exec"
)

func SubscribePipe(shortPipeName string) {
	subscribeCmd := exec.Command("komorebic", "subscribe-pipe", shortPipeName)

	subscribeCmd.Stdout = os.Stdout
	subscribeCmd.Stderr = os.Stderr

	if err := subscribeCmd.Start(); err != nil {
		log.Fatalf("Error starting subscribe process: %v", err)
	}
}

func CreatePipe() (string, *npipe.PipeListener) {
	shortPipeName := "komorebi-manager"
	pipeName := `\\.\pipe\` + shortPipeName

	listener, err := npipe.Listen(pipeName)
	if err != nil {
		log.Fatalf("Error creating pipe: %v", err)
	}

	fmt.Printf("Created pipe: %s\n", pipeName)
	return shortPipeName, listener
}
