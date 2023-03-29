package cmd

import (
	core_api "github.com/sercanakmaz/go-boilerplate-v3/contexts/product/core-api"
	"github.com/spf13/cobra"
)

func init() {
	Root.AddCommand(&cobra.Command{
		Use:   "product-core-api",
		Short: "Product Core API",
		Long:  "Product Core API",
		RunE:  core_api.Init,
	})
}
