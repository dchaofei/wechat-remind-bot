package bcmd

import (
	"github.com/wechaty/go-wechaty/wechaty/user"
	"sort"
)

type Handler interface {
	Handle(message *user.Message)
}

var handlers = map[string]Handler{}

func registerHandle(name string, handler Handler) {
	if _, ok := handlers[name]; ok {
		panic(name + " handler 已经存在")
	}
	handlers[name] = handler
}

func GetHandler(name string) Handler {
	return handlers[name]
}

func GetHandlerNames() []string {
	var names []string
	for k, _ := range handlers {
		names = append(names, k)
	}
	sort.Strings(names)
	return names
}
