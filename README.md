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
Link : `https://documenter.getpostman.com/view/11174314/2sA3QmEaZg`

Postman Collection: `https://www.getpostman.com/collections/ff3b73d96108950d4d18`

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

## Quick Usage
- Create 2 new account via Register API, both can be used as a seller & buyer

#### SELLER
- Login as seller account
  ```sh
    curl --location 'http://go-nicommerce.nicodk.my.id/api/v1/login' \
  --form 'email="nicopenjual@gmail.com"' \
  --form 'password="password"'
  ```
  
- Activate store
  ```sh
   curl --location 'http://go-nicommerce.nicodk.my.id/api/v1/member/activate-store' \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MTYzNDIzOTcsInVzZXJJZCI6MX0.36bwvrrJthXFZar9Wzi7r5XZn1HOZ3bGl2pUA7Hn6kQ' \
  --form 'name="Nico Store"' \
  --form 'description="Menyediakan perlengkapan Ternak Lele modern"' \
  --form 'avatar="images.jpg"' \
  --form 'province_id="35"' \
  --form 'city_id="3573"' \
  --form 'address="Jln. Merjosari, Kota Malang"'
  ```
  
- Create new Product Category
  ```sh
    curl --location 'http://go-nicommerce.nicodk.my.id/api/v1/member/category' \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MTYzNDIzOTcsInVzZXJJZCI6MX0.36bwvrrJthXFZar9Wzi7r5XZn1HOZ3bGl2pUA7Hn6kQ' \
  --form 'name="Foods"' \
  --form 'avatar="images.jpg"'
  ```
  
- Create new Product
   ```sh
    curl --location 'http://go-nicommerce.nicodk.my.id/api/v1/member/product/category/1' \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MTYzNDIzOTcsInVzZXJJZCI6MX0.36bwvrrJthXFZar9Wzi7r5XZn1HOZ3bGl2pUA7Hn6kQ' \
  --form 'name="Pakan Lele Super (Grade A++)"' \
  --form 'description="Ini pakan lele super kualitas terbaik"' \
  --form 'stock="100"' \
  --form 'weight="100"' \
  --form 'base_price="99000"' \
  --form 'price_cut="0"'
  ```
   
- Create new Shipment Option
   ```sh
  curl --location 'http://localhost:8000/api/v1/member/shipment-option' \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MTYzNDIzOTcsInVzZXJJZCI6MX0.36bwvrrJthXFZar9Wzi7r5XZn1HOZ3bGl2pUA7Hn6kQ' \
  --form 'name="JNE"' \
  --form 'avatar="images.jpg"'
  ```
   
#### BUYER
- Login as buyer account
   ```sh
  curl --location 'http://go-nicommerce.nicodk.my.id/api/v1/login' \
  --form 'email="nicopembeli@gmail.com"' \
  --form 'password="password"'
   ```
   
- Create new Address Option
   ```sh
  curl --location 'http://go-nicommerce.nicodk.my.id/api/v1/member/address-option' \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MTYzNDI1NjAsInVzZXJJZCI6Mn0.4zOH1vCZ4oND0jZffj_nLVBuFS29WoYkvhU1POx4VMU' \
  --form 'province_id="35"' \
  --form 'city_id="3573"' \
  --form 'address="Jln. Joyosari A, Kelurahan Merjosari"'
   ```
- Insert product from seller account to buyer's cart
   ```sh
   curl --location 'http://go-nicommerce.nicodk.my.id/api/v1/member/cart' \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MTYyODU5NDYsInVzZXJJZCI6Mn0.IhRNfWvtAEB4ArcuhPh-Y4GDJKps2GYJp0OM2pMVWZw' \
  --form 'product_id="1"' \
  --form 'qty="5"' \
  --form 'shipping_price="10000"'
   ```
- Process cart to transaction
  ```sh
  curl --location 'http://go-nicommerce.nicodk.my.id/api/v1/member/transaction' \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MTYzNDExMDgsInVzZXJJZCI6Mn0.xsBj1Ne4oCtIC3eQ1WggkkrkU2su1eln-a7SOjshy3o' \
  --form 'store_id="1"' \
  --form 'shipping_id="1"' \
  --form 'address_id="1"'
   ```
- Open the `redirect_url` from response to processing payment

- Choose Gopay / QRIS
  
  ![image](https://github.com/nicodwik/go-nicommerce/assets/55322279/e21c06f7-01f0-4206-8cc7-24087c189774)
  
- Go to simulate Gopay / QRIS, copy QR code address & paste to simulator field
  
  `https://simulator.sandbox.midtrans.com/qris/index`
  
  ![image](https://github.com/nicodwik/go-nicommerce/assets/55322279/e8eea32b-ff96-4400-8d51-043ad391f481)

- Payment Gateway will send notification through `callback` endpoint, check the database if payment was success / fail

## Next Improvement
- Implement Redis cache to the most requested endpoint (products, categories, user detail, cart)
