package vars

import (
	"github.com/robfig/cron/v3"
	"github.com/wechaty/go-wechaty/wechaty"
)

type App struct {
	CronSpec string
}

type Database struct {
	User     string
	Password string
	Host     string
	Name     string
}

var (
	AppSetting      = &App{}
	DataBaseSetting = &Database{}
	Bot             *wechaty.Wechaty
	CronInstance    *cron.Cron
)
