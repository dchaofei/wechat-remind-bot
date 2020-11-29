package bcmd

import (
	"github.com/dchaofei/wechat-remind-bot/models"
	"github.com/wechaty/go-wechaty/wechaty/user"
)

const OpenCheckInCmdName = "$开启打卡"

func init() {
	registerHandle(OpenCheckInCmdName, new(openCheckIn))
}

type openCheckIn struct{}

func (o *openCheckIn) Handle(message *user.Message) {
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
		if err := o.create(message); err != nil {
			message.Say(err.Error())
			return
		}
		message.Say("开启打卡成功")
		return
	}
	from := message.From()
	if from.ID() != roomModel.AdminWxID {
		room.Say("只有当前群的 bot 管理员才能操作此功能", from)
		return
	}
	if roomModel.Status != models.CloseCheckinStatus {
		room.Say("打卡已开启,请不要重复开启", from)
		return
	}
	err = models.UpdateRoomStatus(roomModel.ID, models.OpenCheckinStatus)
	if err != nil {
		room.Say("开启打卡失败: "+err.Error(), from)
		return
	}
	room.Say("打卡已开启", from)
}

func (o *openCheckIn) create(message *user.Message) error {
	roomModel := &models.Room{
		ID:        0,
		WxRoomID:  message.Room().ID(),
		AdminWxID: message.From().ID(),
		Status:    models.OpenCheckinStatus,
	}
	return models.AddRoom(roomModel)
}
