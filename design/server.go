package design

import (
	"goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
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

	cors.Origin("*", func() {
		cors.Headers("*")
		cors.Methods("GET", "POST", "DELETE")
		cors.MaxAge(100)
		cors.Credentials()
	})
})
