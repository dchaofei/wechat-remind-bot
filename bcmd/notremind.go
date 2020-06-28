package bcmd

import (
	"github.com/dchaofei/wechat-remind-bot/models"
	"github.com/wechaty/go-wechaty/wechaty/user"
)

const NotRemindCmdName = "#以后不要提醒我"

func init() {
	registerHandle(NotRemindCmdName, new(notRemind))
}

type notRemind struct{}

func (n *notRemind) Handle(message *user.Message) {
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
	exist, err := models.ExistNotRemindBy(roomModel.ID, from.ID())
	if err != nil {
		room.Say(err.Error(), from)
        return
	}
    if exist {
        return
    }
	err = models.AddNotRemind(&models.NotRemind{
		WxID:   from.ID(),
		RoomID: roomModel.ID,
	})
	if err != nil {
		room.Say("操作失败，请稍后重试", from)
		return
	}
	room.Say("以后将不在提醒你，如果需要继续提醒，请对我说:#提醒我", from)
}
