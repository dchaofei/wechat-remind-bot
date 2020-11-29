package bcron

import (
    "fmt"
    "github.com/dchaofei/wechat-remind-bot/bcmd"
    "github.com/dchaofei/wechat-remind-bot/vars"
	"log"
)

func EatRemind() {
	for _, roomId := range vars.AppSetting.EatRemindRoomIds {
		go func(roomId string) {
			defer func() {
				if err := recover(); err != nil {
					log.Printf("发送 eatRemind {%s} panic: %v", roomId, err)
				}
			}()
            fmt.Println("执行中")
			room := vars.Bot.Room().Load(roomId)
			_, err := room.Say("兄弟姐们儿，点外卖的抓紧时间啦! #小程序:外卖券领取plus\n\n,仅在午饭晚饭的时候会有此提醒")
			if err != nil {
				log.Printf("发送消息失败: {%s} err: %s", roomId, err)
				return
			}
			if _, err := room.Say(bcmd.MiniProgram); err != nil {
                log.Printf("发送小程序失败: {%s} err: %s", roomId, err)
            }
		}(roomId)
	}
}
