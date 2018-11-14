package koa

import "context"
import "net/http"

type key int

const (
	wrkey = key(iota)
)

type wr struct {
	w http.ResponseWriter
	r *http.Request
}

type App struct {
	queue Queue
}

func NewApp() *App {
	var q Queue
	return &App{q}
}

func (app *App) Use(fn func(context.Context, http.ResponseWriter, *http.Request, func())) {
	app.queue.Add(func(ctx context.Context, next func()) {
		wr := ctx.Value(wrkey).(*wr)
		fn(ctx, wr.w, wr.r, next)
	})
}

func (app *App) Listen(addr string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(context.Background(), wrkey, &wr{w, r})
		app.queue.Run(ctx)
	})
	http.ListenAndServe(addr, mux)
}
