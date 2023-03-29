package cmd

import (
	core_api "github.com/sercanakmaz/go-boilerplate-v3/contexts/user/core-api"
	"github.com/spf13/cobra"
)

func init() {
	Root.AddCommand(&cobra.Command{
		Use:   "user-core-api",
		Short: "User Core API",
		Long:  "User Core API",
		RunE:  core_api.Init,
	})
}
