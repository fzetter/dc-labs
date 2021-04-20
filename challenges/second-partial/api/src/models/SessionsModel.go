package models

import (
  "api/src/utils"
  "time"
  "errors"
  "github.com/dgrijalva/jwt-go"
)

// Login
func Login(body *utils.Authentication) (message utils.MessageStruct, err error) {

  found := false
  jwt_token := ""
  user := utils.UserStruct{}

  // ** CHECK IF USER IS REGISTERED
  for _, element := range utils.Users {
    if element.User == body.User {
      // Passwords Don't Match
      if element.Password != body.Password {
        return utils.MessageStruct{}, errors.New("Error: Incorrect Password")
      // Success
      } else {
        found = true
        user = element

        // Create Token
        claims := utils.ClaimsStruct{
            User: user.User,
            Email: user.Email,
            Role: user.Role,
            StandardClaims: jwt.StandardClaims{
                ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
                Issuer: "DC",
            },
        }

        token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
        signedToken, _ := token.SignedString(utils.JWTKey)
        jwt_token = signedToken
        utils.Tokens = append(utils.Tokens, signedToken)
      }
    }
	}

  // ** RESOLVE
  if found == true {

    res := utils.MessageStruct {
      Message: "Hi " + user.User + ", welcome to the DPIP System",
      Token: jwt_token,
    }

    return res, nil

  } else {
    return utils.MessageStruct{}, errors.New("Error: User Not Found")
  }

}
