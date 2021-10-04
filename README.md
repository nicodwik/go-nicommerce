# Go-Nicommerce
### _Mini Ecommerce API project built with Golang Language_

## System Design
This project made with Model View Controller (MVC)

## Entity Relational Diagram
![ERD](https://i.ibb.co/DV1b4zJ/ERD-ECOMMERCE-NICO-drawio.png)

## Features
- Authentication Middleware using [JWT](https://github.com/dgrijalva/jwt-go)
- Object Relational Mapping using [GORM](https://gorm.io)
- Routing using [Echo](https://https://echo.labstack.com)
- Database using [GORM-Mysql](https://gorm.io/docs/connecting_to_the_database.html)
- Auto Reload using [Fresh](https://github.com/gravityblast/fresh)

## API Documentation
[Link](https://documenter.getpostman.com/view/11174314/UUxzCTtP)

[Postman Collection](https://www.getpostman.com/collections/ff3b73d96108950d4d18)

## Installation
- Fork this repository
- Clone from forked repository
- Change directory
```sh
cd go-nicommerce
```
- Run this to download all packages
```sh
go mod download
```
- customize constant variables in `constans/constanst.go`
- create new MYSQL database
- run `main.go` file
```sh
go run main.go
```
- or enable auto-reload development with `fresh` command
```sh
fresh
```
