package bcmd

import (
	"github.com/dchaofei/wechat-remind-bot/models"
	"github.com/wechaty/go-wechaty/wechaty/user"
)

const RemindMeCmdName = "#提醒我"

func init() {
	registerHandle(RemindMeCmdName, new(remindMe))
}

type remindMe struct{}

func (n *remindMe) Handle(message *user.Message) {
	room := message.Room()
	if room == nil {
		message.Say("该功能仅支持群聊")
		return
	}
	roomModel, err := models.GetRoom(room.ID())
	if err != nil {
		message.Say(err.Error())
		return
	}
	if roomModel == nil {
		return
	}
	from := message.From()
	err = models.DeleteBy(roomModel.ID, from.ID())
	if err != nil {
		room.Say("操作失败: "+err.Error(), from)
		return
	}
	room.Say("操作成功", from)
}
