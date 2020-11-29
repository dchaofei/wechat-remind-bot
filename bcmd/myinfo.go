package bcmd

import (
	"fmt"
	"github.com/wechaty/go-wechaty/wechaty/user"
)

const MyInfoCmdName = "$我的信息"

func init() {
	registerHandle(MyInfoCmdName, new(myInfo))
}

type myInfo struct{}

func (m *myInfo) Handle(message *user.Message) {
	room := message.Room()
	if room == nil {
		message.Say("该功能仅支持群聊")
		return
	}
	from := message.From()
	name := from.Name()
	id := from.ID()
	room.Say(fmt.Sprintf("\nwx_id:%s\n昵称:%s", id, name), from)
}
