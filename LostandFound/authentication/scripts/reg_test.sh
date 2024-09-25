#!/bin/bash

#!/bin/bash

count=1

while [ $count -lt 1000 ]
do
    curl --location 'http://localhost:9090/register' \
    --header 'Content-Type: application/json' \
    --data "
    {
        \"username\": \"User$count\",
        \"passwordhash\": \"Userpass$count\",
        \"pincode\": 567890
    }
    "
    count=$((count + 1))
    sleep 1
done
