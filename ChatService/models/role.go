package models

type Role int64

const (
	NormalUser Role = iota
	Administrator
)
