package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type NotRemind struct {
	ID        int        `gorm:"primary_key" json:"id"`
	WxID      string     `json:"wx_id"`
	RoomID    int64      `json:"room_id"`
	CreatedOn *time.Time `json:"created_on"`
}

func GetNotRemindWxIDsBy(roomID interface{}) ([]string, error) {
	var ids []string
	return ids, db.Model(&NotRemind{}).Where("room_id = ?", roomID).Pluck("wx_id", &ids).Error
}

func AddNotRemind(remind *NotRemind) error {
	return db.Create(remind).Error
}

func ExistNotRemindBy(roomID, wxID interface{}) (bool, error) {
	var notRemind NotRemind
	err := db.Select("id").Where("wx_id = ? AND room_id = ?", wxID, roomID).First(&notRemind).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if notRemind.ID > 0 {
		return true, nil
	}
	return false, nil
}

func DeleteBy(roomID, wxID interface{}) error {
	return db.Where("wx_id = ? AND room_id = ?", wxID, roomID).Delete(&NotRemind{}).Error
}
