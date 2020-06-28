package startup

import (
	"github.com/dchaofei/wechat-remind-bot/bcmd"
	"github.com/dchaofei/wechat-remind-bot/bcron"
	"github.com/dchaofei/wechat-remind-bot/vars"
	"github.com/robfig/cron/v3"
	"github.com/wechaty/go-wechaty/wechaty"
	"github.com/wechaty/go-wechaty/wechaty-puppet/schemas"
	"github.com/wechaty/go-wechaty/wechaty/user"
	"gopkg.in/ini.v1"
	"log"
)

func SetupVars() {
	loadIni()
	vars.Bot = getBot()
	vars.CronInstance = getCron()
}

func getBot() *wechaty.Wechaty {
	bot := wechaty.NewWechaty()
	bot.OnScan(func(qrCode string, status schemas.ScanStatus, data string) {
		log.Printf("Scan QR Code to login: %v\nhttps://api.qrserver.com/v1/create-qr-code/?data=%s\n", status, qrCode)
	}).OnLogin(func(user *user.ContactSelf) {
		log.Printf("%s logined\n", user.Name())
	}).OnLogout(func(user *user.ContactSelf, reason string) {
		log.Printf("%s logout, reason: %s\n", user.Name(), reason)
	}).OnMessage(func(message *user.Message) {
		log.Println(message)
		h := bcmd.GetHandler(message.Text())
		if h != nil {
			h.Handle(message)
		}
	}).OnStart(func() {
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
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
