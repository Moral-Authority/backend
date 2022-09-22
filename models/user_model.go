package models

type NewUser struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

type User struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
	Age       int    `json:"age"`
	Address   string `json:"address"`
}
