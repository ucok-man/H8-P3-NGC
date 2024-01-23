package entity

import "time"

type Message struct {
	MessageID  int `gorm:"primaryKey"`
	SenderID   int
	ReceiverID int
	Subject    string
	Body       string
	CreatedAt  time.Time `gorm:"autoCreateTime"`

	Sender   *User `gorm:"foreignKey:SenderID;References:UserID"`
	Receiver *User `gorm:"foreignKey:ReceiverID;References:UserID"`
}
