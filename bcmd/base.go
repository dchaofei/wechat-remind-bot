package bcmd

import (
	"github.com/wechaty/go-wechaty/wechaty/user"
	"sort"
)

type Handler interface {
	Handle(message *user.Message)
}

type handlerName struct {
	name string
	sort int
}

type handlerNames []handlerName

func (h handlerNames) Len() int {
	return len(h)
}

func (h handlerNames) Less(i, j int) bool {
	return h[i].sort < h[j].sort
}

func (h handlerNames) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

var handlers = map[string]Handler{}

var sortedHandlerNames handlerNames

var outputNames []string

func registerHandle(name string, handler Handler, sortI ...int) {
	if _, ok := handlers[name]; ok {
		panic(name + " handler 已经存在")
	}
	handlers[name] = handler
	sortField := 0
	if len(sortI) > 0 {
		sortField = sortI[0]
	}
	sortedHandlerNames = append(sortedHandlerNames, handlerName{
		name: name,
		sort: sortField,
	})
	sort.Sort(sortedHandlerNames)
	setHandlerNames()
}

func GetHandler(name string) Handler {
	return handlers[name]
}

func setHandlerNames() {
	var names []string
	var unSortNames []string
	for _, v := range sortedHandlerNames {
		if v.sort == 0 {
			unSortNames = append(unSortNames, v.name)
			continue
		}
		names = append(names, v.name)
	}
	sort.Strings(unSortNames)

	outputNames = append(unSortNames, names...)
}

func GetHandlerNames() []string {
	return outputNames
}
