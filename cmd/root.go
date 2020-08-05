package cmd

import (
	"os"
	"log"

	"github.com/spf13/cobra"
)

// rootCmd represents the root command <timer> before subcommands are invoked
var (
	rootCmd = &cobra.Command{
		Version: "0.1.0",
		Use:   "timer <command> <value>",
		Short: "A simple CLI command for setting timers",
		Long: `Timer ⏰ is a simple CLI tool, written in Go, for quickly creating simple timers 
that display a tiny native notification when they expire. Timer is built with 
Go, Cobra (CLI framework) and Beeep (cross-platform notifactions and alerts)`,
		Example: `  timer set 1m30s - sets a timer of 1 minute and 30 seconds.`,
	}
)

// Execute adds all child commands to the root command and sets flags appropriately
// called by main.main(). 
func Execute() {
  if err := rootCmd.Execute(); err != nil {
    log.Fatal(err)
		os.Exit(1)
  }
}