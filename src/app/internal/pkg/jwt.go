package pkg

import "github.com/golang-jwt/jwt/v4"

type jwtParameter struct {
	Username string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Email    string `form:"username" json:"username" binding:"required"`
	jwt.RegisteredClaims
}
