package koa

import "context"

type Queue struct {
	List []func(context.Context, func())
}

func (q *Queue) Add(fn func(context.Context, func())) {
	q.List = append(q.List, fn)
}

func (q *Queue) genNext(ctx context.Context, index int) func() {
	if index >= len(q.List) {
		return empty
	}
	return func() {
		q.List[index](ctx, q.genNext(ctx, index+1))
	}
}

func (q *Queue) Run(ctx context.Context) {
	if len(q.List) > 0 {
		q.List[0](ctx, q.genNext(ctx, 1))
	}

}

func empty() {}
