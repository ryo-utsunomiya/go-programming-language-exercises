package memo

import (
	"context"
)

type Func func(key string, ctx context.Context) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{}
}

type request struct {
	key      string
	response chan<- result
	ctx      context.Context
	cancel   context.CancelFunc
}

type Memo struct {
	requests chan request
	ctx      context.Context
	cancel   context.CancelFunc
}

func New(f Func) *Memo {
	ctx, cancel := context.WithCancel(context.Background())
	memo := &Memo{
		requests: make(chan request),
		ctx:      ctx,
		cancel:   cancel,
	}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string, ctx context.Context) (interface{}, error) {
	response := make(chan result)
	req := request{key: key, response: response, ctx: ctx}
	memo.requests <- req
	res := <-response

	select {
	case <-ctx.Done():
		memo.ctx = context.WithValue(memo.ctx, "key", key)
		memo.cancel()
	default:
		// noop
	}

	return res.value, res.err
}

func (memo *Memo) Close() { close(memo.requests) }

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for {
		select {
		case <-memo.ctx.Done():
			delete(cache, memo.ctx.Value("key").(string))
		case req := <-memo.requests:
			e := cache[req.key]
			if e == nil {
				e = &entry{ready: make(chan struct{})}
				cache[req.key] = e
				go e.call(f, req.key, memo.ctx)
			}
			go e.deliver(req.response)
		}
	}
}

func (e *entry) call(f Func, key string, ctx context.Context) {
	e.res.value, e.res.err = f(key, ctx)
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	<-e.ready
	response <- e.res
}
