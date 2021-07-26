# Referral System

## Overall Design

<p align="center">
    <img src="https://github.com/hasbiasshidiq/Referral-System/blob/master/images/Referral-System.png" width="400" />
</p>

Overall, the system consists of referral service and auth service. Referral service jobs include almost all referral and contribution process. While auth service jobs is related to token generation and token introspection (validation). The system communicates with user via HTTP REST protocol which goes directly to the referral service. While the referral service interacts with the auth service via gRPC protocol.


## Database Scheme
The next figure is database schema of referral service.

<p align="center">
    <img src="https://github.com/hasbiasshidiq/Referral-System/blob/master/images/DB.png" width="600" />
</p>

## Tech Stack
List of main tech stacks
* Database  : postgres:10.5
* API       : golang:1.16.5

Here is architecture reference for our each service implementation. Though the Auth Service has not been implemented as in the reference below due to time limit. 

Architecture Reference -> [link](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

Reference Architecture Golang (Implemented) -> [link](https://eltonminetto.dev/en/post/2020-07-06-clean-architecture-2years-later/)

## How to run code
First, make sure no running program on port 5432 (database), 8080 (referral service) and 50059 (auth service)
```
# build images and run database, referral service and auth service in separate container
$ docker-compose up

# stop running container
$ docker-compose down
```

## Endpoint
### Register Generator

This endpoint is accessed for the purpose of registering generator and returns a referral link that can be accessed by contributor to make contribution. This link will be valid for 7 days. The payload data that must be provided are generator_id, generator_name, email, and password.

Example request :
```
curl --location --request POST 'http://localhost:8080/generator/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "generator_id":"ronaldo",
    "generator_name":"Ronaldo",
    "email":"r9@gmail.com",
    "password":"pass"
}'
```
Success Response (Example) :
```
- HTTP Response Status = 201 Created
- Body
{
    "generated_link": "http://localhost:8080/referral/vK9qxgXuuBBdziAuAMEh0QnESVgbaWMSzpLrNT9O"
}
```

### Referral
The link used on this endpoint is generated on registration process. It returns an acccess token to gives contribution permission. This access token will be valid until the associated referral link expire.

Example request :
```
curl --location --request GET 'http://localhost:8080/referral/vK9qxgXuuBBdziAuAMEh0QnESVgbaWMSzpLrNT9O'
```
Success Response (Example) :
```
- HTTP Response Status = 200 OK
- Body
{
    "access_token": <Access Token for Contributor>
}
```
The obtained access token will be attached in Authorization Header for each contribution request.

### Contribute
This endpoint is accessed by contributors by providing an email in the payload and a valid access token in the Authorization Header. This access token is obtained from previous referral enpoint.

Example request :
```
curl --location --request POST 'http://localhost:8080/contributor/contribute' \
--header 'Authorization: Bearer <Access Token for Contributor>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "davidb@gmail.com"
}'
```
Success Response :
```
- HTTP Response Status = 200 OK
- Body
{
    "status": "Contribution is counted"
}
```

### Login Generator
This endpoint is accessed for generator authoritative purpose. There are two parameters required for this login process, that is `generator_id` and `password`. When the login process is successful, this endpoint will return an access token with `generator`role being embedded in claims for authorization purpose. This access token will be valid for one day.

Example Request :
```
curl --location --request POST 'http://localhost:8080/generator/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "generator_id":"ronaldo",
    "password":"pass"
}'
```
Success Response :
```
- HTTP Response Status = 200 OK
- Body
{
    "access_token": <Access Token for Generator>
}
```

### Extend Referral Link Time
This endpoint is accessed if generator wants to extend the lifetime of associated referral link. Access token is provided in Authorization Header. Lifetime will be updated up to one week from this renewal process. As explained in previous, the access token is obtained from the `login generator` endpoint.

Example Request :
```
curl --location --request PUT 'http://localhost:8080/generator/extend-time' \
--header 'Authorization: Bearer <Access Token for Generator>'
```
Success Response (Example):
```
{
    "status": "Success Extending Link till 2021-08-02 09:12:35"
}
```

### List Contributor
This endpoint is accessed when the generator wants to see list of emails and its number of contributions. An access token is provided in Authorization Header. This access token is obtained from the `login generator` endpoint.

Example Request :
```
curl --location --request GET 'http://localhost:8080/contributor' \
--header 'Authorization: Bearer <Access Token for Generator>'
```
Success Response (Example):
```
[
    {
        "email": "davidb@gmail.com",
        "referral_link": "http://localhost:8080/referral/vK9qxgXuuBBdziAuAMEh0QnESVgbaWMSzpLrNT9O",
        "contribution": "2"
    },
    {
        "email": "davidbeckham@gmail.com",
        "referral_link": "http://localhost:8080/referral/vK9qxgXuuBBdziAuAMEh0QnESVgbaWMSzpLrNT9O",
        "contribution": "5"
    }
]
```