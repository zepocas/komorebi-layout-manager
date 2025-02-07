package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "os/exec"
    "os/signal"
    "syscall"
    "time"

    "github.com/natefinch/npipe"
)

func main() {
    // 1. Start komorebi subscribe-pipe (SYNCHRONOUSLY and WAIT)
    shortPipeName := "komorebi-layout"
    pipeName := `\\.\pipe\` + shortPipeName

    cmd := exec.Command("komorebic", "subscribe-pipe", shortPipeName)
    err := cmd.Run() // Wait for the command to finish
    if err != nil {
        log.Fatalf("Error running komorebi subscribe-pipe: %v", err)
    }

    // 2. Connect to the Named Pipe (Windows-Specific) - Retry logic added
    var conn *npipe.PipeConn // Correct type: *npipe.PipeConn

    for i := 0; i < 5; i++ {
        log.Println("Connecting to pipe")
        conn, err = npipe.Dial(pipeName) // Dial returns a *npipe.PipeConn
        if err == nil {
            log.Println("Connection successful")
            break // Connection successful
        }
        log.Printf("Attempt %d: Error connecting to named pipe: %v", i+1, err)
        time.Sleep(500) // Wait before retrying
    }

    if err != nil {
        log.Fatalf("Failed to connect to named pipe after multiple attempts: %v", err)
    }
    defer conn.Close() // Close the connection

    fmt.Println("Connected to Komorebi named pipe. Listening for events...")

    // 3. Event Loop (using bufio.NewReader for io.Reader)
    reader := bufio.NewReader(conn) // Use bufio.NewReader with the connection

    go func() {
        for {
            line, err := reader.ReadString('\n') // Read until newline
            if err != nil {
                log.Printf("Error reading from pipe: %v", err)
                break // Or handle the error as needed
            }
            eventData := line[:len(line)-1] // Remove trailing newline
            fmt.Printf("Received event: %s\n", eventData)
            // TODO: Process the event data (unmarshal, act upon it)
        }
    }()

    // 4. Graceful Shutdown
    sigchan := make(chan os.Signal, 1)
    signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
    <-sigchan

    fmt.Println("Shutting down...")
}
