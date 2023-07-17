package entity

type Users struct {
	ID string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
	Image string `json:"image"`
	Bio string `json:"bio"`
}