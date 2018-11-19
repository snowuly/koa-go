package koa

import "context"
import "net/http"

type key int

const (
	wrkey = key(iota)
)

type Handler func(context.Context, http.ResponseWriter, *http.Request, func(context.Context))

type wr struct {
	w http.ResponseWriter
	r *http.Request
}

func (app *App) Use(fn Handler) {
	app.Add(func(ctx context.Context, next func(context.Context)) {
		wr := ctx.Value(wrkey).(*wr)
		fn(ctx, wr.w, wr.r, next)
	})
}

type App struct {
	Queue
}

func NewApp() *App {
	var queue Queue
	return &App{queue}
}

func (app *App) Listen(addr string, ctx context.Context) {
	if ctx == nil {
		ctx = context.Background()
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(ctx, wrkey, &wr{w, r})
		app.Run(ctx)
	})
	http.ListenAndServe(addr, mux)
}
