package cmd

import (
	"github.com/sercanakmaz/go-boilerplate-v3/contexts/user/created-log-consumer"
	"github.com/spf13/cobra"
)

func init() {
	Root.AddCommand(&cobra.Command{
		Use:   "user-created-log-consumer",
		Short: "User Created Log Consumer",
		Long:  "User Created Log Consumer",
		RunE:  created_log_consumer.Init,
	})
}
