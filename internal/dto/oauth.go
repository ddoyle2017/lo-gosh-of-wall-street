package dto

type AuthToken struct {
	AccessToken string
	TokenType   string
	ExpiresIn   uint64 // Maybe some sort of timestamp/date type instead?
	Sub         string
}
