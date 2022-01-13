### Tribal exercise

Go project with docker environment

For execute this project you need only execute the next line
```
make build-dev
```

Or for execute the docker-compose command directly you can use 
```
docker-compose --env-file .env up --build --remove-orphans -d
```

And open your browser with the port 8080 and you can see the message **_Hello, Tribal! <3_**
```
http://localhost:8080
```

For make a request to validate credit lines you can use the next cURL
```
curl --location --request POST 'http://localhost:8080/credit/validate' \
--header 'Content-Type: application/json' \
--data-raw '{
    "foundingType": "SME",
    "cashBalance": 435.30,
    "monthlyRevenue": 4235.45,
    "requestedCreditLine": 100,
    "requestedDate": "2022-01-15T00:45:39.860Z"
}'
```

Also you can find the Postman collection for this on the root path of this project with the name of **collection.json**
