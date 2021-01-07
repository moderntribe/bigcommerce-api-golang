module samples

go 1.15

replace bigcommerce.com/apis/clients/storeinfo => ../clients/storeinfo

replace bigcommerce.com/apis/clients/themes => ../clients/themes

replace bigcommerce.com/apis/clients/widgets => ../clients/widgets

replace bigcommerce.com/apis/clients/subscribers => ../clients/subscribers

replace bigcommerce.com/apis/clients/scripts => ../clients/scripts

replace bigcommerce.com/apis/clients/pricelists => ../clients/pricelists

replace bigcommerce.com/apis/clients/ordersv2 => ../clients/ordersv2

replace bigcommerce.com/apis/clients/ordersv3 => ../clients/ordersv3

replace bigcommerce.com/apis/clients/channels => ../clients/channels

replace bigcommerce.com/apis/clients/carts => ../clients/carts

replace bigcommerce.com/apis/clients/wishlists => ../clients/wishlists

require (
	bigcommerce.com/apis/clients/carts v0.0.0-00010101000000-000000000000
	bigcommerce.com/apis/clients/channels v0.0.0-00010101000000-000000000000
	bigcommerce.com/apis/clients/ordersv2 v0.0.0-00010101000000-000000000000
	bigcommerce.com/apis/clients/ordersv3 v0.0.0-00010101000000-000000000000
	bigcommerce.com/apis/clients/pricelists v0.0.0-00010101000000-000000000000
	bigcommerce.com/apis/clients/scripts v0.0.0-00010101000000-000000000000
	bigcommerce.com/apis/clients/storeinfo v0.0.0-00010101000000-000000000000
	bigcommerce.com/apis/clients/subscribers v0.0.0-00010101000000-000000000000
	bigcommerce.com/apis/clients/themes v0.0.0-00010101000000-000000000000
	bigcommerce.com/apis/clients/widgets v0.0.0-00010101000000-000000000000
	bigcommerce.com/apis/clients/wishlists v0.0.0-00010101000000-000000000000
	github.com/go-openapi/runtime v0.19.24
)
