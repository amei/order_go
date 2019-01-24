package model

type OrderStatus int
const (
	Status_summit  = 0
	Status_finish  = 1
)
type OrderItem struct {
	OrderId string
	OrderUser string
	OrderType string
	CreateTime string
	Status int
}