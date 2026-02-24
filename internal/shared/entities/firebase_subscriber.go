package entities

type FirebaseSubscriber struct {
	Token  string `gorm:"primaryKey"`
	UserID uint
}

func (FirebaseSubscriber) TableName() string { return "firebase_subscribers" }
