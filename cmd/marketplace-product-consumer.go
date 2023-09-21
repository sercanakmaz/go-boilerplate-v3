package cmd

import (
	marketplace_product_consumer "github.com/sercanakmaz/go-boilerplate-v3/contexts/product/marketplace-product-consumer"
	"github.com/spf13/cobra"
)

func init() {
	Root.AddCommand(&cobra.Command{
		Use:   "marketplace-product-consumer",
		Short: "MP Product consumer",
		Long:  "MP Product consumer",
		RunE:  marketplace_product_consumer.Init,
	})
}
