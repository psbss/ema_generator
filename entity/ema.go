package entity

import (
	"time"
)

type Ema struct {
	Id uint
	Content string `gorm:"size:50"`
	Author string `gorm:"size:25"`
	createAt time.Time
}