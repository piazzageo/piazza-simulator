package main

import (
	//"log"
	"github.com/spf13/cobra"
	"time"
	//"strings"

	//"github.com/mpgerlek/piazza-simulator/piazza"
)

func main() {

	var Verbose bool
	var registryPort int

	var cmdRegistry = &cobra.Command{
		Use:   "registry",
		Short: "Start the registry service",
		Long:  "Start the registry service",
		Run: func(cmd *cobra.Command, args []string) {
			Registry(registryPort)
		},
	}

	var cmdDispatch = &cobra.Command{
		Use:   "dispatcher",
		Short: "Start the dispatcher service",
		Long:  "Start the dispatcher service",
		Run: func(cmd *cobra.Command, args []string) {
			Dispatcher(registryPort)
		},
	}

	var sleepDuration time.Duration
	var cmdSleeper = &cobra.Command{
		Use:   "sleeper",
		Short: "Start the sleeper service",
		Long:  "Start the sleeper service",
		Run: func(cmd *cobra.Command, args []string) {
			Sleeper(registryPort, sleepDuration)
		},
	}
	cmdSleeper.PersistentFlags().DurationVarP(&sleepDuration, "duration", "d", 5*time.Second, "duration to sleep")

	var cmdPz = &cobra.Command{Use: "pz"}

	cmdPz.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")

	cmdRegistry.PersistentFlags().IntVarP(&registryPort, "port", "p", 8080, "port number")

	cmdPz.AddCommand(cmdRegistry, cmdDispatch, cmdSleeper)
	cmdPz.Execute()
}
