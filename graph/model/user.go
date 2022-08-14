package model

type User struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
	Age       int    `json:"age"`
	Address   string `json:"address"`
}
