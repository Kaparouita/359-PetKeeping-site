package domain

import "time"

type Order struct {
	Id        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	BookRefer int  `json:"Book_id"`
	UserRefer int  `json:"user_id"`
	User      User `gorm:"foreignKey:UserRefer"`
}
