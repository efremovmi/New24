package models

type StandardResponses struct {
	Ok  string
	Err string
}

type AuthResponses struct {
	Ok    string
	Err   string
	Token string
}

type GetAllUsersResponses struct {
	Ok    string
	Err   string
	Users []*User
}

type GetUserResponses struct {
	Ok   string
	Err  string
	User *User
}
