package cmd

import (
	"sync"
	"errors"
	"strings"
	"time"
  "fmt"
	
	"github.com/gen2brain/beeep"
  "github.com/spf13/cobra"
)

// initialise amd add command to rootCmd
func init() {
  rootCmd.AddCommand(setCmd)
}

// initialise command vars and timer duration
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

// parses to ensure there is only 1 argument passed in
// and that it is a valid time duration 
func parseArgs(cmd *cobra.Command, args []string) error {
	// only 1 arg
	err := cobra.ExactArgs(1)(cmd, args)
	if err != nil {
		return err
	}

	// arg is valid time duration
	value := strings.Join(args, " ")
	timer, err = time.ParseDuration(value)
	if err != nil {
		e := fmt.Sprintf("\n  couldn't parse the time duration you entered.")
		return errors.New(e)
	}

	return nil
}

// setTimer starts a timer for a set amount of time as a goroutine
// notifications are sent at the start and end of the timer
func setTimer(cmd *cobra.Command, args []string) {
	title := "⏰ Timer CLI - Countdown Started"
	body := fmt.Sprintf("A timer has been started for %s", timer)
	icon := "assets/icon.png"
	beeep.Notify(title, body, icon)
	
	// start timer in a waitgroup in case multiple timers started simultaneously 
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		// start a countdown
		done := time.Until(time.Now().Add(timer))
		<-time.After(done)

		// update title and body and notify
		title = "⏰ Timer CLI - Countdown finished"
		body = fmt.Sprintf("Your timer of %v is over.", args[0])
		beeep.Notify(title, body, icon)
	}()

	// wait for all the goroutines launched here to finish.
	wg.Wait()
}
