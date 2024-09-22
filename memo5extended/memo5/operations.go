package memo5


var	cancelKey chan string
var cache  = make(map[string]*entry) 

// New function.
func New(f Func) *Memo {
	memo := Memo{request: make(chan request)}
	go memo.server(f)
	return &memo
}

func (memo *Memo) server(f Func) {
	for {
		select {
		case req := <-memo.request:
			e := cache[req.key]
			if e == nil {
				e = &entry{ready: make(chan struct{})}
				cache[req.key] = e
				select {
				case <-cancelKey:
				delete(cache,req.key)
				default:
					go e.call(f, req.key, req.done)
				}

			}
			go e.deliver(req.response)
		default:
		}
	}

}

func (e *entry) call(f Func, key string, ch <-chan struct{}) {
	e.response.value, e.response.err = f(key, ch)
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	<-e.ready
	response <- e.response
}

// Get method
func (memo *Memo) Get(url string, ch <-chan struct{}) (interface{}, error) {
	response := make(chan result)
	memo.request <- request{url, response, ch}
	res := <-response
	select {
	case <-ch:
		cancelKey <- url
	default:
	}
	return res.value, res.err
}

