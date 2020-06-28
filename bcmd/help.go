package bcmd

import "github.com/wechaty/go-wechaty/wechaty/user"

const HelpCmdName = "#帮助"

func init() {
	registerHandle(HelpCmdName, new(help))
}

type help struct{}

func (h *help) Handle(message *user.Message) {
	s := ""
	for _, name := range GetHandlerNames() {
		s += name + "\n"
	}
	s = "支持的命令:\n\n" + s
	message.Say(s)
}
