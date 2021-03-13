package models

import (
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* Claim struct */
type Claim struct {
	Email string             `bson:"email"  json:"email"`
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	jwt.StandardClaims
}
