package helpers

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// ErrCheck checks if err is preset. If so, the program panics.
func ErrCheck(msg string, err error) {
	if err != nil {
		fmt.Printf("%s: %+v", msg, err)
		panic(err)
	}
}

// PrettyPrint prints the contents of data with spaces for better distinction.
func PrettyPrint(data interface{}) {
	var p []byte
	p, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s \n", p)
}

// Run makes a program run indefinitely until the user hits CTRL-C.
func Run() {
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGTERM)
	<-sc
}
