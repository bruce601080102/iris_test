package main

import (
	"github.com/kataras/iris/v12"
)

func main() {

	// ============================================================
	app := iris.New()
	app.AllowMethods(iris.MethodOptions)
	app.Use(iris.NoCache)
	app.HandleDir("/", "./static/js")

	app.Get("/", func(ctx iris.Context) {
		ctx.Writef("Hello World")
	})

	app.Run(iris.Addr("0.0.0.0:80"))
}
