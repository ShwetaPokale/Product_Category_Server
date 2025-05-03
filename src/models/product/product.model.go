package product

type Product struct {
	ProductID          uint    `gorm:"primaryKey;autoIncrement" json:"product_id"`
	ProductName        string  `gorm:"type:varchar(100);not null" json:"product_name"`
	ProductDescription string  `gorm:"type:text" json:"product_description"`
	ProductCategory    string  `gorm:"type:varchar(50)" json:"product_category"`
	ProductPrice       float64 `gorm:"type:numeric(10,2);check:product_price >= 0" json:"product_price"`
	QuantityAvailable  int     `gorm:"check:quantity_available >= 0" json:"quantity_available"`
}
