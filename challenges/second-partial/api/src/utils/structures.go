package utils

import (
  "github.com/dgrijalva/jwt-go"
)

// JWT CLAIMS

type ClaimsStruct struct {
    User string `json:"username" binding:"required"`
    Email string `json:"email" binding:"required"`
    Role string `json:"role" binding:"required"`
    jwt.StandardClaims
}

// SESSIONS

type Authentication struct {
    User string `json:"user" binding:"required"`
    Password string `json:"password" binding:"required"`
}

type MessageStruct struct {
    Message string `json:"message" binding:"required"`
    Token string `json:"token"`
    Time string `json:"time"`
}

type UserStruct struct {
    User string `json:"user" binding:"required"`
    Email string `json:"email" binding:"required"`
    Role string `json:"role" binding:"required"`
    Password string `json:"password" binding:"required"`
    Token string `json:"token" binding:"required"`
}
