## Test practice from:
- https://medium.com/@carloslopez_19744/adroid-todo-app-gotodo-part-1-creating-an-api-restful-with-docker-go-and-postgresql-beceddcb4364

## About

Backend project that creates a RESTful API in Golang for the todo project. I have created a front-end app using the flutter. 

## Resources used for this project:
- Golang
- Web Framework Gin
- PostgreSql for the database
- Docker to run Postgres
- Gin JWT for Auth service

## Summary
- CRUD operation on the todo list that are stored in the database through RESTful API
- Authentication using JWT token 

## In Details 
To build an app where you can organize all your tasks once you register an account. The API connects to a PostgreSQL database, and both are running on his respective Docker container instances.

Moreover, the database’s design is very simple, it consists of only two tables: User and Task; the first table stores the user’s information, and the second table stores all users’ tasks.

## Steps
- Create a [Golang Module Project](https://blog.golang.org/using-go-modules) which allow Go to manages all the dependencies of the project without using any other tool.