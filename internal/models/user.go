package models

type User struct {
	ID       interface{} `json:"id"`
	UserName interface{} `json:"user_name"`
	Password interface{} `json:"-"`
	Role     interface{} `json:"Role"`
}
