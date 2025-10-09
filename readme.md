# URL Shortener using Go

## Pre-requisites

- Go (Golang)
- Docker

## Functionalities

🚀 Features

Shorten any long URL into a short one.

Optionally specify your own custom backhalf.

Prevent duplicate short links for the same original URL.

Automatic generation of unique 8-character backhalves.

In-memory storage (no database dependency).

Clean RESTful API design using Gin.

## Installation & Setup

git clone https://github.com/srbhgalinde/url-shortner.git
cd url-shortner

## Run app
1. go mod tidy
2. go run cmd/server/main.go

## API Endpoints
1. Shorten URL

POST /shorten
```
Request Body:

{
  "url": "https://www.example.com/very/long/url",
  "backhalf": "custom123" // optional
}

```

Response (201 Created):
```

{
  "shortUrl": "http://localhost:8080/custom123",
  "backhalf": "custom123"
}
```

If backhalf is omitted, the server will generate a random 8-character one.

Error Responses:
```
400 – Invalid request body.

409 – Custom backhalf already taken.
```

2. Redirect Short URL

GET /{backhalf}

Redirects the user to the original long URL.

Example:
```
GET http://localhost:8080/custom123
```
→ Redirects to https://www.example.com/very/long/url

🧠 Example Usage (with curl)
## Create a shortened URL
```
curl -X POST http://localhost:8080/shorten \
  -H "Content-Type: application/json" \
  -d '{"url": "https://longurl.org/guygu26v62828882"}'
```

## Example output:
```
 {
  "shortUrl": "http://localhost:8080/9b1d7e2a",
  "backhalf": "9b1d7e2a"
}
```

Test redirection
curl -v http://localhost:8080/9b1d7e2a

## Testing

You can use tools like Postman, curl, or HTTPie to test API requests.

# Future Enhancements

Add persistent database storage (e.g. Redis, PostgreSQL).

Add user authentication.

Add rate limiting and analytics (click tracking, expiration).

Deploy with Docker.
