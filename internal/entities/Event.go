package entitites

type Event struct {
	ID        uint   `gorm:"primaryKey"`
	EventType string `gorm:"size:50;not null"`
	Data      string `gorm:"type:jsonb;not null"`
	Timestamp string `gorm:"type:timestamp;default:current_timestamp"`
}
