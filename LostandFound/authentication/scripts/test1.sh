#!/bin/bash

count=1
while [ $count -le 1000 ]
do
    curl --location --request GET 'http://localhost:9090/validate' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcnlUaW1lIjoxNzI2NDUyOTE1LCJwaW5jb2RlIjo1NjAyMDEsInVzZXJuYW1lIjoiVXNlcjYifQ.G7nIC2goAp2xJ0KeOUk5baLYPo3_O-J9QH8pzPokyZ4' \
--header 'Content-Type: application/json' \
--data '{
    "username": "User6"
}'
count=$((count + 1))
sleep 1
done