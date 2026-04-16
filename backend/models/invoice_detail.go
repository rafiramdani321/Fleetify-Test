package models

type InvoiceDetail struct {
	ID        uint  `gorm:"primaryKey;autoIncrement" json:"id"`
	InvoiceID uint  `gorm:"not null" json:"invoice_id"`
	ItemID    uint  `gorm:"not null" json:"item_id"`
	Quantity  int   `gorm:"not null" json:"quantity"`
	Price     int64 `gorm:"not null" json:"price"`
	SubTotal  int64 `gorm:"not null" json:"subtotal"`

	Item Item `gorm:"foreignKey:ItemID" json:"item"`
}