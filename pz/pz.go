package main

import (
	"github.com/spf13/cobra"
	"time"
)

func main() {

	var verbose bool
	var registryHost, serviceHost string

	var cmdPz = &cobra.Command{
		Use: "pz",
	}
	cmdPz.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

	var cmdRegistry = &cobra.Command{
		Use:   "registry",
		Short: "Start the registry service",
		Long:  "Start the registry service",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Registry(serviceHost)
		},
	}
	cmdRegistry.Flags().StringVar(&serviceHost, "host", "localhost:12300", "service host name")

	var cmdGateway = &cobra.Command{
		Use:   "gateway",
		Short: "Start the gateway service",
		Long:  "Start the gateway service",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Gateway(serviceHost, registryHost)
		},
	}
	cmdGateway.Flags().StringVar(&registryHost, "registry", "localhost:12300", "registry host name")
	cmdGateway.Flags().StringVar(&serviceHost, "host", "localhost:12301", "service host name")

	var cmdDispatch = &cobra.Command{
		Use:   "dispatcher",
		Short: "Start the dispatcher service",
		Long:  "Start the dispatcher service",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Dispatcher(serviceHost, registryHost)
		},
	}
	cmdDispatch.Flags().StringVar(&registryHost, "registry", "localhost:12300", "registry host name")
	cmdDispatch.Flags().StringVar(&serviceHost, "host", "localhost:12302", "service host name")

	var sleepDuration time.Duration
	var cmdSleeper = &cobra.Command{
		Use:   "sleeper",
		Short: "Start the sleeper service",
		Long:  "Start the sleeper service",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Sleeper(serviceHost, registryHost, sleepDuration)
		},
	}
	cmdSleeper.Flags().DurationVar(&sleepDuration, "seconds", 5*time.Second, "duration to sleep (seconds)")
	cmdSleeper.Flags().StringVar(&registryHost, "registry", "localhost:12300", "registry host name")
	cmdSleeper.Flags().StringVar(&serviceHost, "host", "localhost:12303", "service host name")

	cmdPz.AddCommand(cmdRegistry, cmdDispatch, cmdSleeper, cmdGateway)
	cmdPz.Execute()
}
