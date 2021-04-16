# 一个微信提醒机器人、可以用于群聊提醒、签到、打卡等

# 如何使用
1. 打开 mysql，执行 `remind.sql` 文件里的 sql 
2. 配置 app.ini 里的数据库连接
3. 运行程序 `WECHATY_PUPPET_SERVICE_TOKEN=xxx go run main.go`
4. 把机器人拉到群里，回复 `#开启打卡`, 机器人会按照 app.ini 的 `CronSpec` 设置的时间间隔定时提醒群里未打卡的成员。

# 支持命令
- $以后不要提醒我
- $关闭打卡
- $帮助
- $开启打卡
- $提醒我
- $打卡
- $外卖红包
- ......
