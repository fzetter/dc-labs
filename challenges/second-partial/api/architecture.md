# Architecture Document

###### FILE STRUCTURE

      api
      ├── public
      │   ├── footer.html
      │   ├── header.html
      │   ├── index.html
      │   └── menu.html
      ├── src
      │   ├── controllers
      │   │    ├── AboutController.go
      │   │    ├── ImagesController.go
      │   │    └── SessionsController.go
      │   ├── models
      │   │    └── SessionsModel.go
      │   ├── routes
      │   │    └── routes.go
      │   └── utils
      │        ├── db.go
      │        ├── structures.go
      │        └── utils.go
      ├── app.go
      ├── go.sum
      ├── go.mod
      ├── architecture.md
      └── user-guide.md

###### INFORMATION

  + The server framework is done with Gin Web Framework.
  + JSON Web Token (JWT) are used for user authentication.

###### ENDPOINTS

  + GET /about
  + POST /login
  + DELETE /logout - Bearer Token
  + GET /status - Bearer Token
  + POST /upload - Bearer Token





  
