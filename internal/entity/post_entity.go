package entity

type Posts struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id" gorm:"column:user_id;foreignKey:UserID"`
	Title       string `json:"title"`
	Description string `json:"description"`
	User        Users  `json:"user"` 
}