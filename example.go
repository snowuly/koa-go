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
	app.Use(func(ctx context.Context, w http.ResponseWriter, r *http.Request, next func()) {
		fmt.Println(r.URL)
		w.Write([]byte("ok"))
	})
	app.Listen(":8080")
}
