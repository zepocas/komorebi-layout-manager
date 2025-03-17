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
	} else {
		log.Println("No error on subscribe-pipe")
	}

}

func CreatePipe() (string, *npipe.PipeListener) {
	shortPipeName := "komorebi-manager"
	pipeName := `\\.\pipe\` + shortPipeName

	listener, err := npipe.Listen(pipeName)
	if err != nil {
		log.Fatalf("Error creating pipe: %v", err)
	}

	//defer func(listener *npipe.PipeListener) {
	//	err := listener.Close()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}(listener)

	fmt.Printf("Created pipe: %s\n", pipeName)
	return shortPipeName, listener
}
