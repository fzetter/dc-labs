package utils

import ()

var JWTKey = []byte("spH2x-3@KUwG8#HP")

var Tokens []string

var Users = []UserStruct {
    UserStruct {
      User: "admin",
      Email: "admin@admin.com",
      Role: "admin",
      Password: "admin123",
    },
    UserStruct {
      User: "client",
      Email: "client@client.com",
      Role: "client",
      Password: "client123",
    },
}
