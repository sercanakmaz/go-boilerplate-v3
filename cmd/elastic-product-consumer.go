package cmd

import (
	elastic_product_consumer "github.com/sercanakmaz/go-boilerplate-v3/contexts/product/elastic-product-consumer"
	"github.com/spf13/cobra"
)

func init() {
	Root.AddCommand(&cobra.Command{
		Use:   "elastic-product-consumer",
		Short: "Elastic consumer",
		Long:  "Elastic consumer",
		RunE:  elastic_product_consumer.Init,
	})
}
