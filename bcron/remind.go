package bcron

import (
	"github.com/dchaofei/wechat-remind-bot/models"
	"github.com/dchaofei/wechat-remind-bot/vars"
	"github.com/wechaty/go-wechaty/wechaty-puppet/helper"
	"github.com/wechaty/go-wechaty/wechaty-puppet/schemas"
	_interface "github.com/wechaty/go-wechaty/wechaty/interface"
	"log"
	"math"
	"time"
)

func Remind() {
	rooms, err := models.GetOpenStatusRooms()
	if err != nil {
		log.Println("Remind models.GetOpenStatusRooms() err: ", err)
		return
	}
	async := helper.NewAsync(0)
	for _, room := range rooms {
		room := room
		async.AddTask(func() (interface{}, error) {
			remind(room)
			return nil, nil
		})
	}
	async.Result()
}

func remind(roomModel *models.Room) {
	bot := vars.Bot
	room := bot.Room().Find(&schemas.RoomQueryFilter{
		Id: roomModel.WxRoomID,
	})
	if room == nil {
		log.Println("remind 没有找到 room: ", roomModel.WxRoomID)
		return
	}

	// 防止新成员没有进来
	if err := room.Sync(); err != nil {
		log.Println("room.Sync err: ", err.Error())
	}

	notRemindWxIds, err := models.GetNotRemindWxIDsBy(roomModel.ID)
	if err != nil {
		log.Println("remind GetNotRemindWxIDsBy err: ", err.Error())
		return
	}

	alreadyCheckinIds, err := models.GetAlreadyCheckinWxIdsBy(roomModel.ID, time.Now().Format("2006-01-02"))
	if err != nil {
		log.Println("remind GetAlreadyCheckinWxIdsBy err: ", err.Error())
		return
	}
	notRemindWxIds = append(notRemindWxIds, alreadyCheckinIds...)

	contacts, err := room.MemberAll(nil)
	if err != nil {
		log.Println("remind MemberAll err: ", err.Error())
		return
	}

	var remindContacts []_interface.IContact
	for _, contact := range contacts {
		if inArray(contact.ID(), notRemindWxIds) || contact.Self() {
			continue
		}
		remindContacts = append(remindContacts, contact)
		// 防止联系人昵称变更，导致 @ 失败
		err := contact.Sync()
		if err != nil {
			log.Println("contact.Sync err: ", err.Error())
		}
	}

	length := len(remindContacts)

	if length == 0 {
		return
	}

	// 分批@,每次只@100人
	start := 0
	max := 100
	for {
		if start >= length {
			return
		}
		min := math.Min(float64(start+max), float64(length))
		room.Say("\n\n不要忘记打卡哦!!!\n\n\n如果今天不想收到提醒请回复:＄打卡\n了解更多命令回复:＄帮助\n点外卖回复:＄红包", remindContacts[start:int(min)]...)
		start+=max
	}
}

func inArray(s string, ss []string) bool {
	for _, v := range ss {
		if s == v {
			return true
		}
	}
	return false
}
