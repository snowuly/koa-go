// +build ignore

package main

import (
	"context"
	"fmt"
	"koa"
	"net/http"
)

func main() {
	app := koa.NewApp()
	app.Use(func(ctx context.Context, w http.ResponseWriter, r *http.Request, next func(context.Context)) {
		fmt.Println(r.URL)
		w.Write([]byte("ok\n"))
		next(ctx)
		w.Write([]byte("haha\n"))
	})
	app.Use(func(ctx context.Context, w http.ResponseWriter, r *http.Request, next func(context.Context)) {
		w.Write([]byte("middel\n"))
	})
	app.Listen(":8080")
}
