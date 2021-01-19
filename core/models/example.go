package models

type Example struct {
	ID uint
}

func (Example) TableName() string {
	return "public.example"
}
