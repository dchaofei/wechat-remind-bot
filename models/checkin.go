package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Checkin struct {
	ID        int        `gorm:"primary_key" json:"id"`
	WxID      string     `json:"wx_id"`
	RoomID    int64      `json:"room_id"`
	Date      string     `json:"date"`
	CreatedOn *time.Time `json:"created_on"`
}

func ExistCheckinBy(wxID, roomID, date interface{}) (bool, error) {
	var checkin Checkin
	err := db.Select("id").Where("wx_id = ? AND room_id = ? AND date = ?", wxID, roomID, date).First(&checkin).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if checkin.ID > 0 {
		return true, nil
	}
	return false, nil
}

func GetAlreadyCheckinWxIdsBy(roomID, date interface{}) ([]string, error) {
	var ids []string
	return ids, db.Model(&Checkin{}).Where("room_id = ? and date = ?", roomID, date).Pluck("wx_id", &ids).Error
}

func AddCheckIn(checkin *Checkin) error {
	return db.Create(checkin).Error
}
