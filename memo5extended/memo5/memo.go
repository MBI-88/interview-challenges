package memo5


type request struct {
	key string
	response chan<- result
	done <-chan struct{}
}

type result struct {
	value interface{}
	err error
}

// Memo struct 
type Memo struct {
	request chan request
}

type entry struct {
	response result
	ready chan struct{}
}

// Func func
type Func func(key string,ch <-chan struct{}) (value interface{},er error)