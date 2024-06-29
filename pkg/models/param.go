package models

type Param struct {
	ID    ParamName `gorm:"primaryKey;type:varchar(255);not null"`
	Value string    `gorm:"type:varchar(255);not null"`
}

func (Param) TableName() string {
	return "params"
}
