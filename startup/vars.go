package startup

import (
	"fmt"
	"github.com/dchaofei/wechat-remind-bot/bcmd"
	"github.com/dchaofei/wechat-remind-bot/bcron"
	"github.com/dchaofei/wechat-remind-bot/vars"
	"github.com/robfig/cron/v3"
	"github.com/wechaty/go-wechaty/wechaty"
	"github.com/wechaty/go-wechaty/wechaty-puppet/schemas"
	"github.com/wechaty/go-wechaty/wechaty/user"
	"gopkg.in/ini.v1"
	"log"
	"strings"
)

func SetupVars() {
	loadIni()
	vars.Bot = getBot()
	vars.CronInstance = getCron()
}

func getBot() *wechaty.Wechaty {
	bot := wechaty.NewWechaty()
	bot.OnScan(func(context *wechaty.Context, qrCode string, status schemas.ScanStatus, data string) {
		log.Printf("Scan QR Code to login: %v\nhttps://api.qrserver.com/v1/create-qr-code/?data=%s\n", status, qrCode)
	}).OnLogin(func(context *wechaty.Context, user *user.ContactSelf) {
		log.Printf("%s logined\n", user.Name())
	}).OnLogout(func(context *wechaty.Context, user *user.ContactSelf, reason string) {
		log.Printf("%s logout, reason: %s\n", user.Name(), reason)
	}).OnMessage(func(context *wechaty.Context, message *user.Message) {
		str := message.String()
		if message.Room() != nil {
			str = fmt.Sprintf("roomID: %s ; %s", message.Room().ID(), str)
		}
		log.Println(str)
		h := bcmd.GetHandler(strings.Replace(message.Text(), "＄", "$", 1))
		if h != nil {
			h.Handle(message)
		}
	}).OnStart(func(context *wechaty.Context) {
		log.Println("started")
	})

	var err = bot.Start()
	if err != nil {
		log.Fatalf("getBot Start: %v", err)
	}
	return bot
}

func getCron() *cron.Cron {
	c := cron.New()
	if _, err := c.AddFunc(vars.AppSetting.CronSpec, bcron.Remind); err != nil {
		log.Fatalf("getCrom c.AddFunc: %v", err)
	}
	if vars.AppSetting.EatRemindCronSpec != "" {
		if _, err := c.AddFunc(vars.AppSetting.EatRemindCronSpec, bcron.EatRemind); err != nil {
			log.Fatalf("getCrom c.AddFunc: %v", err)
		}
	}
	c.Start()
	return c
}

var cfg *ini.File

func loadIni() {
	var err error
	cfg, err = ini.Load("app.ini")
	if err != nil {
		log.Fatalf("loadIni, fail to parse 'app.ini': %v", err)
	}

	mapTo("app", vars.AppSetting)
	mapTo("database", vars.DataBaseSetting)

	vars.AppSetting.EatRemindRoomIds = cfg.Section("app").Key("EatRemindRoomIds").Strings(",")
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
