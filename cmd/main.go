package main

import (
	"bufio"
	"fmt"
	"github.com/natefinch/npipe"
	"github.com/zepocas/komorebi-layout-manager/internal"
	"log"
	"net"
	"os"
	"os/signal"
)

func main() {
	shortPipeName, listener := internal.CreatePipe()
	defer func(listener *npipe.PipeListener) {
		err := listener.Close()
		if err != nil {
			log.Fatalf("Fatal while closing listener: %v", err)
		}
	}(listener)

	internal.SubscribePipe(shortPipeName)

	// Handle Ctrl+C gracefully
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		i := 1
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Printf("Error accepting connection: %v", err)
				if i == 10 {
					break
				}
				i++
				continue
			}
			log.Println("New connection accepted!")

			go handleConnection(conn)
		}
	}()

	// Wait for Ctrl+C
	<-c
	fmt.Println("\nShutting down...")
}

func handleConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Fatal connection error: %v", err)

		}
	}(conn)

	reader := bufio.NewReader(conn)
	tracker := internal.NewStateTracker()
	for {
		event, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Error reading from connection: %v", err)
			return
		}

		// we should parse this json into a event struct
		err = tracker.ProcessEvent(event)
		tracker.PrintSummary()
		if err != nil {
			log.Fatal(err)
		}

	}
}
