package main

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("a command must be provided!")
	}
	count := 1
	max := 100
	var err error = errors.New("nil")
	for count <= max && err != nil {
		log.Printf("%v/%v run starts", count, max)
		cmd := exec.Command(os.Args[1], os.Args[2:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		log.Printf("%v/%v run finishes", count, max)
		count++
	}
	if err == nil {
		log.Printf("%v/%v run finishes", count, max)
	} else {
		log.Fatalf("command \"%v\" failed", strings.Join(os.Args, " "))
	}
}
