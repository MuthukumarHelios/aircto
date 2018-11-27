package middlewares

import (	
	"net/http"
)

type JwtClaims struct{
	UserName string `json:"userName" bson:"userName"`
    	
}




func VerifyJwt() {

	
}