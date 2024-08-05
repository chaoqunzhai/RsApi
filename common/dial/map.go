package dial

import "sync"

type Map struct {
	sync.RWMutex
	M map[string]*MapEvent
}

type MapEvent struct {
	Idc int64 `json:"idc"`
}

var (
	MapCnf = &Map{M: make(map[string]*MapEvent)}
)

func (e *Map) Get(key string) (*MapEvent, bool) {
	e.RLock()
	defer e.RUnlock()
	event, exists := e.M[key]
	return event, exists
}
func (e *Map) Set(key string, event *MapEvent) {
	e.Lock()
	defer e.Unlock()
	e.M[key] = event
}

func (e *Map) Delete(key string) {
	e.Lock()
	defer e.Unlock()

	delete(e.M, key)
}
