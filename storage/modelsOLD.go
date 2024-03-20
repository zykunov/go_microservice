package storage

import (
	"time"

	"gorm.io/gorm"
)

type SummaryModel struct {
	gorm.Model
	OrderUid          string    `json:"order_uid"`
	TrackNumber       string    `json:"track_number"`
	Entry             string    `json:"entry"`
	Delivery          Delivery  `gorm:"foreignKey:CompanyRefer"`
	Payment           Payment   `gorm:"foreignKey:transaction"`
	Items             Items     `gorm:"foreignKey:track_number"`
	Locale            string    `json:"locale"`
	InternalSignature string    `json:"internal_signature"`
	CustomerId        string    `json:"customer_id"`
	DeliveryService   string    `json:"delivery_service"`
	Shardkey          string    `json:"shardkey"`
	Sm_id             int       `json:"sm_id"`
	DateCreated       time.Time `json:"date_created"`
	OofShard          string    `json:"oof_shard"`
}

type Delivery struct {
	gorm.Model
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

type Payment struct {
	gorm.Model
	Transaction   string `json:"transaction"`
	RequestId     string `json:"request_id"`
	Currency      string `json:"currency"`
	Provider      string `json:"provider"`
	Amount        int    `json:"amount"`
	PaymentDt     int    `json:"payment_dt"`
	Bank          string `json:"bank"`
	Delivery_cost int    `json:"delivery_cost"`
	GoodsTotal    int    `json:"goods_total"`
	CustomFee     int    `json:"custom_fee"`
}

type Items struct {
	gorm.Model
	ChrtId      int    `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price"`
	Rid         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int    `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  int    `json:"total_price"`
	NmId        int    `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      int    `json:"status"`
}
