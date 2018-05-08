package sharedmap

const (
	set    = "set"
	get    = "get"
	remove = "remove"
	count  = "count"
)

// SharedMap a thread safe map(type: map[interface{}]interface{}).
// This using channel, not mutex.
type SharedMap struct {
	m map[interface{}]interface{}
	c chan command
}

type command struct {
	action string
	key    interface{}
	value  interface{}
	result chan<- interface{}
}

// Set sets value in SharedMap.m
// to associated with the given key.
// It update if the key already has been registered.
func (sm SharedMap) Set(k interface{}, v interface{}) {
	sm.c <- command{action: set, key: k, value: v}
}

// Get returns the first value associated with the given key
// and whether or not it exists.
// It returns nil if the key has not been registered.
func (sm SharedMap) Get(k interface{}) (value interface{}, exist bool) {
	callback := make(chan interface{})
	sm.c <- command{action: get, key: k, result: callback}
	result := (<-callback).([2]interface{})
	value = result[0]
	exist = result[1].(bool)
	return
}

// Remove removes value from SharedMap.m
// to associated with the given key.
func (sm SharedMap) Remove(k interface{}) {
	sm.c <- command{action: remove, key: k}
}

// Count returns the number of SharedMap.m length.
func (sm SharedMap) Count() int {
	callback := make(chan interface{})
	sm.c <- command{action: count, result: callback}
	return (<-callback).(int)
}

func (sm SharedMap) run() {
	for cmd := range sm.c {
		switch cmd.action {
		case set:
			sm.m[cmd.key] = cmd.value
		case get:
			v, ok := sm.m[cmd.key]
			cmd.result <- [2]interface{}{v, ok}
		case remove:
			delete(sm.m, cmd.key)
		case count:
			cmd.result <- len(sm.m)
		}
	}
}

// New is Create a new SharedMap and returns.
func New() SharedMap {
	sm := SharedMap{
		m: make(map[interface{}]interface{}),
		c: make(chan command),
	}
	go sm.run()
	return sm
}
