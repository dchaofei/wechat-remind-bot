package bcmd

import (
	"github.com/wechaty/go-wechaty/wechaty-puppet/schemas"
	"github.com/wechaty/go-wechaty/wechaty/user"
	"log"
)

const MoneyCmdName = "$外卖红包"

func init() {
    m :=new(money)
	registerHandle(MoneyCmdName, m, 98)
	registerHandle("$红包", m, 99)
}

type money struct{}

var MiniProgram = user.NewMiniProgram(&schemas.MiniProgramPayload{
	Appid:       "wx68b30d5e22041892",
	Description: "",
	PagePath:    "pages/index/index.html",
	ThumbUrl:    "3062020100045630540201000204996066a702032f802902042049110e02045fbf9ad3042f6175706170706d73675f613033383638643431626632356535635f313630363339323533313431395f3836343431340204010800030201000405004c56f900",
	Title:       "这里的外卖最便宜,因为有券",
	Username:    "gh_915306feedc7@app",
	ThumbKey:    "4c3c5c93a3ff093ce9d8f740767801ff",
})

func (h *money) Handle(message *user.Message) {
	_, err := message.Say(MiniProgram)
	if err != nil {
		log.Println("money handler exception:", err)
	}
}
