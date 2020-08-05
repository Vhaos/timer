package cmd

import (
	"log"
	"errors"
	"strings"
	"time"
  "fmt"
	
	"github.com/gen2brain/beeep"
  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(setCmd)
}

var (
	timer time.Duration = 0 
	setCmd = &cobra.Command{
		Use:   "set <time>",
		Short: "Sets a new timer",
		Long:  `Set a timer to expire in a specified amount of time`,
		Example: "  timer set 5m15s\n  timer set 30s",
		Args: parseArgs,
		Run: setTimer,
	}
)

func parseArgs(cmd *cobra.Command, args []string) error {
	err := cobra.ExactArgs(1)(cmd, args)
	if err != nil {
		return err
	}

	value := strings.Join(args, " ")
	timer, err = time.ParseDuration(value)
	if err != nil {
		e := fmt.Sprintf("\n  couldn't parse the time duration you entered.")
		return errors.New(e)
	}

	return nil
}

func setTimer(cmd *cobra.Command, args []string) {
	title := "⏰ Timer CLI"
	body := fmt.Sprintf("A timer has been started for %s", timer)
	icon := "assets/information.png"
	
	err := beeep.Notify(title, body, icon)
	if err != nil {
		log.Fatal("error trying to create notification")
		panic(err)
	}
}