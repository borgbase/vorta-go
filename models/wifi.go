package models

import "time"

type KnownWifi struct {
	ID        int       `gorm:"column:id;not null;primary_key"`
	SSID       string    `gorm:"column:ssid;type:varchar(255);not null"`
	LastConnected time.Time       `gorm:"column:last_connected"`
	Allowed   bool `gorm:"column:allowed;not null"`
	ProfileID	int	`gorm:"column:profile_id"`
}

func (KnownWifi) TableName() string {
	return "wifisettingmodel"
}
