package design

import "goa.design/goa/v3/dsl"

var _ = dsl.Service("user", func() {
	dsl.Description("The user service performs operations on users")

	dsl.HTTP(func() {
		dsl.Path("/users")
	})

	dsl.Method("create", func() {
		dsl.Meta("swagger:summary", "ユーザ作成")

		dsl.Payload(func() {
			dsl.Attribute("name", dsl.String, func() {
				dsl.Description("ユーザ名")
				dsl.Example("yyh-gl")
			})
			dsl.Attribute("age", dsl.Int, func() {
				dsl.Description("年齢")
				dsl.Example(18)
			})
			dsl.Required("name", "age")
		})

		dsl.Result(func() {
			dsl.Attribute("user", user, "作成されたユーザ情報")
			dsl.Required("user")
		})

		dsl.HTTP(func() {
			dsl.POST("")
			dsl.Response(dsl.StatusCreated)
		})
	})

	dsl.Method("index", func() {
		dsl.Meta("swagger:summary", "ユーザ一覧取得")

		dsl.Payload(dsl.Empty)

		dsl.Result(func() {
			dsl.Attribute("users", dsl.ArrayOf(user), "ユーザ一覧")
			dsl.Required("users")
		})

		dsl.HTTP(func() {
			dsl.GET("")
			dsl.Response(dsl.StatusOK)
		})
	})
})

var user = dsl.Type("User", func() {
	dsl.Attribute("name", dsl.String, "ユーザ名", func() {
		dsl.Example("yyh-gl")
	})
	dsl.Attribute("age", dsl.Int, "年齢", func() {
		dsl.Example(18)
	})
	dsl.Required("name", "age")
})
