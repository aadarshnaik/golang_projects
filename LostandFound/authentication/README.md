
# go-Authentication

This project is an authentication service written in golang. 
The purpose of creating this projects is purely educational. As authentication systems are widely used across services this service can be used as a microservice in other educational projects that I am going to build.
This project uses mysql as the database. 
I am using gorm as the ORM here. Also using gorm Automigrate to create database from given struct. 

#### Project Flow
User registers using the /register api with necessary details, we check in the database if an user exists, If not we create a salt, concat it to password, encrypt it using bcrypt package in golang and save the encrypted hash in db. 

User Login - When user tries to login we check if the credetials match with the user credetials (comparing password hash as we store the salt in db). If credentials is a match we issue a JWT token with an expiry of 1 hr.

Validate - We get the token from the user in request header and validate it using the secret string we used to sign the token.


## API Reference

#### Get all items

Register API
```http
  POST /register
```
```
curl --location 'http://localhost:9090/register' \
--header 'Content-Type: application/json' \
--data '
    {
        "username": "test1",
        "passwordhash": "passtest1",
        "pincode": 123456
    }
'
```
Login API
```http
  POST /login
```
```
curl --location 'http://localhost:9090/login' \
--header 'Content-Type: application/json' \
--data '{
    "username": "test1",
    "passwordhash": "passtest1"
}'
```
Validate API
```http
  POST /validate
```
```
curl --location --request GET 'http://localhost:9090/validate' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcnlUaW1lIjoxNzI2NDUyOTE1LCJwaW5jb2RlIjo1NjAyMDEsInVzZXJuYW1lIjoiVXNlcjYifQ.G7nIC2goAp2xJ0KeOUk5baLYPo3_O-J9QH8pzPokyZ4' \
--header 'Content-Type: application/json' \
--data '{
    "username": "test1"
}'
```

## RUN the project

There is a K8s directory but haven't deployed on K8s yet. 
If you have docker in your machine you can do 

To pull a docker image from docker hub and run it 
```
make mysql
```
To run this project 
```
make run
```

If you don't have docker on the machine you can always download go
and build the file using go build . and run the executable file. just make sure that you have mysql running and active on your machine with port 3306 open. 

Please reach out to me for any suggestions, concerns. 



## Authors

- [@aadarshnaik](https://www.github.com/aadarshnaik)
- email: naikaadarsh23@gmail.com 

