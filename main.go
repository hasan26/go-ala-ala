package main

import "github.com/kataras/iris"

func main() {
	iris.Get("/hi", func(ctx *iris.Context) {
		ctx.Writef("Hi %s", "Hasan")
	})
	iris.Listen(":8080")
}