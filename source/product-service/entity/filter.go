package entity

type Filter interface {
	GetValue() string
	GetOperator() string
	GetName() string
}
