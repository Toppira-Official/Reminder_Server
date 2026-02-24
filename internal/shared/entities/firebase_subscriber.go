package entities

type FirebaseSubscriber struct {
	Token  string `gorm:"primaryKey"`
	UserID uint
}

func (FirebaseSubscriber) TableName() string { return "Firebase_Subscribers" }
