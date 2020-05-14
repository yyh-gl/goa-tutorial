package design

import (
	"goa.design/goa/v3/dsl"
)

var _ = dsl.API("goa-tutorial", func() {
	dsl.Title("Goa Tutorial App")
	dsl.Description("Goa introduction")
	dsl.Server("goa-tutorial", func() {
		dsl.Host("localhost", func() { dsl.URI("http://localhost:8088") })
	})
	dsl.HTTP(func() {
		dsl.Path("/api/v1")
	})
})
