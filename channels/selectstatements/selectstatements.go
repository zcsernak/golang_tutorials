package main

import (
	"fmt"
	"time"
)

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{}) // signal only channel with empty struct

func main() {
	go logger()
	logCh <- logEntry{time.Now(), logInfo, "App is starting"}

	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
	time.Sleep(100 * time.Millisecond)
	// when the main function is done, then the logger gorutine is just shut down, and not closed gracefully

	// the deffered anonymous function can be used to close it properly
	/*defer func() {
		close(logCh)
	}()*/

	doneCh <- struct{}{}
}

func logger() {
	/*for entry := range logCh {
		fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
	}*/
L:
	for {
		// blocks until a message is received from one of its channels it is listenning for
		select {
		case entry := <-logCh:
			fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
		case <-doneCh:
			break L
			/*default:*/ // makes the select statement not blocking
		}
	}
}
