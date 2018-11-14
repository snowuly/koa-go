[KOA](https://koajs.com/) Go Version
================================

```go

app := koa.NewApp()
app.Use(func(ctx context.Context, w http.ResponseWriter, r *http.Request, next func()) {
	fmt.Println(r.URL)
	w.Write([]byte("ok"))
})
app.Listen(":8080")

```
