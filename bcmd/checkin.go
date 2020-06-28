package bcmd

import (
	"github.com/dchaofei/wechat-remind-bot/models"
	"github.com/wechaty/go-wechaty/wechaty/user"
	"time"
)

const CheckInCmdName = "#签到"

func init() {
	registerHandle(CheckInCmdName, new(checkIn))
}

type checkIn struct{}

func (s *checkIn) Handle(message *user.Message) {
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
	if roomModel.Status != models.OpenCheckinStatus {
		room.Say("签到功能未开启", from)
		return
	}
	date := time.Now().Format("2006-01-02")
	exist, err := models.ExistCheckinBy(from.ID(), roomModel.ID, date)
	if err != nil {
		message.Say(err.Error())
		return
	}
	if exist {
		room.Say("今天已经签到，请不要重复签到", from)
		return
	}
	if err := models.AddCheckIn(&models.Checkin{
		WxID:   from.ID(),
		RoomID: roomModel.ID,
		Date:   date,
	}); err != nil {
		room.Say(err.Error(), from)
		return
	}
	room.Say("签到成功"+date, from)
}
