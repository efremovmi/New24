package models

type StandardResponses struct {
	Status string `json:"status,omitempty"`
	Err    string `json:"err,omitempty"`
}

type AuthResponses struct {
	Err   string `json:"err,omitempty"`
	Token string `json:"token,omitempty"`
}

type GetAllUsersResponses struct {
	Status string  `json:"status,omitempty"`
	Err    string  `json:"err,omitempty"`
	Users  []*User `json:"users,omitempty"`
}

type GetUserResponses struct {
	Status string `json:"status,omitempty"`
	Err    string `json:"err,omitempty"`
	User   *User  `json:"user,omitempty"`
}
