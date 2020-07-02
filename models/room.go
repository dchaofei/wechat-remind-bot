package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type RoomStatus uint8

const OpenCheckinStatus RoomStatus = 1
const CloseCheckinStatus RoomStatus = 0

type Room struct {
	ID         int64      `gorm:"primary_key" json:"id"`
	WxRoomID   string     `json:"wx_room_id"`
	AdminWxID  string     `json:"admin_wx_id"`
	Status     RoomStatus `json:"status"`
	CreatedOn  *time.Time `json:"created_on"`
	ModifiedOn *time.Time `json:"modified_on"`
}

func GetRoom(wxRoomID string) (*Room, error) {
	var room Room
	err := db.Where("wx_room_id = ?", wxRoomID).First(&room).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func GetOpenStatusRooms() ([]*Room, error) {
	var rooms []*Room
	return rooms, db.Where("status = ?", OpenCheckinStatus).Find(&rooms).Error
}

func AddRoom(room *Room) error {
	return db.Create(room).Error
}

func UpdateRoomStatus(roomID int64, status RoomStatus) error {
	return db.Model(&Room{}).Where("id = ?", roomID).Update("status", status).Error
}
