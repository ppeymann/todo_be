Todo Backend API

A RESTful API for managing todos, built using Golang, Gin, and PostgreSQL.

Features

✅ User authentication (signup, signin)
✅ Create, read, update, and delete todos
✅ Change todo status (in-progress, complete, cancel)
✅ Secure endpoints with authentication
✅ Swagger documentation

Tech Stack

Language: Golang

Framework: Gin

Database: PostgreSQL

Authentication: Paseto

Documentation: Swagger

Installation

1. Clone the Repository

git clone https://github.com/ppeymann/todo_be.git
cd todo_be

2. Install Dependencies

go mod tidy

5. Start the Server

go run main.go

The server runs on http://localhost:8080 by default.

API Endpoints

Authentication

POST /signup - Register a new user

POST /signin - Login and get a token

PATCH /change_pass - Change password

Todo Management

GET /todos - Fetch all todos

GET /todos/{id} - Get a specific todo

POST /todos - Create a new todo

PATCH /todos/status/{id}/{status} - Update todo status

PUT /todos/{id} - Update a todo

DELETE /todos/{id} - Delete a todo

Swagger API Docs

The API is documented with Swagger.

After starting the server, open:➡️ http://localhost:8080/swagger/index.html
