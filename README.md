# backend-boilerplate

Simple go boilerplate using hexagonal architecture for microservices. Already equipped with basic docker-compose file and docker file for future deployments.

## Packages:
 - [Gin Framework](https://github.com/gin-gonic/gin): Base framework for the project and handling http requests [documentation](https://gin-gonic.com/docs/)
 - [SQLBoiler](https://github.com/volatiletech/sqlboiler): Generates the ORM for the project based on the database schema
 - [Goose](https://github.com/pressly/goose): Migration tool capable of working with different database drivers and migrations in sql and golang
 - [Go Swag](https://github.com/swaggo/swag): Generates swagger documentation for the api based on annotations
 - [Gin Swagger](https://github.com/swaggo/gin-swagger): Integrates the swagger documentation with gin
 - [SQL Mock](https://github.com/DATA-DOG/go-sqlmock): Mocks database for better unit tests

There are 2 makefiles to facilitate the use of cli commands, one in the root folder and other in the migration folder. If environment variables are changed these makefiles may need to be updated.
