# Todolist Application
## Description
This Golang-based todolist application, following Clean Architecture principles, leverages the Gin Framework for REST API development. Authentication and authorization are secured through JWT, while Postgresql serves as the database. Validator ensures request body validation, and Bcrypt handles password hashing, all within a clean and modular architectural design.

## Tech Stack
- [Gin Framework](https://github.com/gin-gonic/gin) 
- [JWT (JSON Web Token)](https://github.com/golang-jwt/jwt)
- [Postgresql](https://www.postgresql.org/)
- [Validator](https://github.com/go-playground/validator)
- [Bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)

## Running 
- Clone this repository
- Run `go run main.go` in the root directory
- The application will run on port 8080
- You can access the application on `http://localhost:8080`

## Endpoints
### User Group 
- Register a new user: `POST /auth/register`
- Login: `POST /auth/login`
### Scope Group
- Add new scope: `POST /api/scopes`
- Delete scope: `DELETE /api/scopes/:id`
- Assign scope to user: `GET /api/user/:user_id/scopes/:scope_id/assign`
### Todolist Group
- Add new todolist: `POST /api/todolist`
- Get all todolist: `GET /api/todolist`
- Get todolist by id: `GET /api/todolist/:id`
- Update todolist: `PUT /api/todolist/:id`
- Delete todolist: `DELETE /api/todolist/:id`
- Checklist todolist: `POST /api/todolist/:id/check`

## Soon to be implemented
- [ ] Routing group
- [ ] JWT refresh token