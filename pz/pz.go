package main

import (
	"github.com/spf13/cobra"
	"time"
)

func main() {

	var Verbose bool
	var hostname string

	var cmdPz = &cobra.Command{
		Use: "pz",
	}

	cmdPz.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	cmdPz.PersistentFlags().StringVarP(&hostname, "registry-host", "r", "http://localhost:8080", "registry host name")

	var cmdRegistry = &cobra.Command{
		Use:   "registry",
		Short: "Start the registry service",
		Long:  "Start the registry service",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Registry(hostname)
		},
	}

	var cmdGateway = &cobra.Command{
		Use:   "gateway",
		Short: "Start the gateway service",
		Long:  "Start the gateway service",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Gateway(hostname)
		},
	}

	var cmdDispatch = &cobra.Command{
		Use:   "dispatcher",
		Short: "Start the dispatcher service",
		Long:  "Start the dispatcher service",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Dispatcher(hostname)
		},
	}

	var sleepDuration time.Duration
	var cmdSleeper = &cobra.Command{
		Use:   "sleeper",
		Short: "Start the sleeper service",
		Long:  "Start the sleeper service",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Sleeper(hostname, sleepDuration)
		},
	}
	cmdSleeper.PersistentFlags().DurationVarP(&sleepDuration, "duration", "d", 5*time.Second, "duration to sleep")

	cmdPz.AddCommand(cmdRegistry, cmdDispatch, cmdSleeper, cmdGateway)
	cmdPz.Execute()
}
