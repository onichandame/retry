package retry

import (
	"errors"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"strings"
)

var count = 0
var max = 100

var rootCmd = &cobra.Command{
	Use:   "retry",
	Args:  cobra.MinimumNArgs(1),
	Short: "retry a command on failure",
	Long:  "re-run a command until it is successful or 100 tries all failed",
	Run: func(cmd *cobra.Command, args []string) {
		var err error = errors.New("nil")
		for count < max && err != nil {
			count++
			log.Printf("%v/%v run starts", count, max)
			cmd := exec.Command(os.Args[1], os.Args[2:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err = cmd.Run()
			log.Printf("%v/%v run finishes", count, max)
		}
		if err == nil {
			log.Printf("completed %v/%v", count, max)
		} else {
			log.Fatalf("command \"%v\" failed", strings.Join(os.Args, " "))
		}
	},
}

func Exec() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
