package models

type Battle struct {
	Round int
}

func NewBattle() *Battle {
	return &Battle{Round: 0}
}
