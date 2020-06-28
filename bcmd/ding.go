package bcmd

import (
	"github.com/wechaty/go-wechaty/wechaty/user"
	"log"
)

const DingCmdName = "#ding"

func init() {
	registerHandle(DingCmdName, new(ding))
}

type ding struct{}

func (d *ding) Handle(message *user.Message) {
	_, err := message.Say("dong")
	if err != nil {
		log.Println("ding handler exception:", err)
	}
}
