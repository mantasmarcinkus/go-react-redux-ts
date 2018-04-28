package handler

import "gopkg.in/macaron.v1"

func IndexHandler(ctx *macaron.Context) {
	ctx.HTML(200, "index")
}