module samples

go 1.15

replace bigcommerce.com/apis/clients/storeinfo => ../clients/storeinfo

replace bigcommerce.com/apis/clients/themes => ../clients/themes

replace bigcommerce.com/apis/clients/widgets => ../clients/widgets

require (
	bigcommerce.com/apis/clients/storeinfo v0.0.0-00010101000000-000000000000
	bigcommerce.com/apis/clients/themes v0.0.0-00010101000000-000000000000
	bigcommerce.com/apis/clients/widgets v0.0.0-00010101000000-000000000000
	github.com/go-openapi/runtime v0.19.24
)
