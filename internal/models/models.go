package models

type calculateobj struct {
	id  int `gorm:"primaryKey"`
	a   int
	b   int
	ops string
}

type latestentry struct {
	id       int `gorm:"primaryKey"`
	latestid int
}
