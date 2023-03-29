package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var Root = &cobra.Command{
	Use:   "GoBoilerplate",
	Short: "Go Boilerplate Application",
	Long:  "Command Line Interface for Oms Projection Applications",
}

func Execute() {
	if err := Root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func init() {
	viper.AutomaticEnv()
}
