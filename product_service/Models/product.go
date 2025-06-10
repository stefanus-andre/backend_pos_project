package Models

type Product struct {
	ID    uint   `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Image string `json:"image"`
}
