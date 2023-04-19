package dto

type LoginDto struct {
	Username string
	Password string
}

type RegisterDto struct {
	Username string
	Password string
	ConPass  string
}
