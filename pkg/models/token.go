package models

type UserClaims struct {
	UserID        string `json:"userId" form:"userId"`
	Email         string `json:"email" form:"email"`
	EmailVerified bool   `json:"emailVerified" form:"emailVerified"`
}
