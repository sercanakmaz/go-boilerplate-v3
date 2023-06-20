package cmd

import (
	core_api "github.com/sercanakmaz/go-boilerplate-v3/contexts/oms/core-api"
	"github.com/spf13/cobra"
)

func init() {
	Root.AddCommand(&cobra.Command{
		Use:   "oms-core-api",
		Short: "OMS Core API",
		Long:  "OMS Core API",
		RunE:  core_api.Init,
	})
}
