package bcmd

import (
	"github.com/dchaofei/wechat-remind-bot/models"
	"github.com/wechaty/go-wechaty/wechaty/user"
)

const CloseCheckInCmdName = "#关闭签到"

func init() {
	registerHandle(CloseCheckInCmdName, new(closeCheckIn))
}

type closeCheckIn struct{}

func (o *closeCheckIn) Handle(message *user.Message) {
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
	if from.ID() != roomModel.AdminWxID {
		room.Say("只有当前群的 bot 管理员才能操作此功能", from)
		return
	}
	if roomModel.Status != models.OpenCheckinStatus {
		room.Say("签到已关闭,请不要重复关闭", from)
		return
	}
	err = models.UpdateRoomStatus(models.CloseCheckinStatus)
	if err != nil {
		room.Say("关闭签到失败: "+err.Error(), from)
		return
	}
	room.Say("签到已关闭", from)
}
