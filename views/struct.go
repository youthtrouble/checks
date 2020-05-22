package views

// Credentials is a struct that models the structure of a user, both in the request body, and in the DB
type Credentials struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type Manny struct{

}