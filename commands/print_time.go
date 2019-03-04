package commands

import (
	"time"

	"github.com/spf13/cobra"
)

// PrintTime will print the current time in a pretty way
func PrintTime() *cobra.Command {
	return &cobra.Command{
		Use: "curtime",
		RunE: func(cmd *cobra.Command, args []string) error {
			now := time.Now()
			prettyTime := now.Format(time.RubyDate)
			cmd.Println("Hey Gophers! The current time is", prettyTime)
			return nil
		},
	}
}
