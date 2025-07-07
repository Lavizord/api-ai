# Go API Template

This is a minimal and extensible Go Web API built using Gorilla Mux. It includes a basic router, Swagger documentation, and AWS Secrets Manager integration for securely loading configuration values. The project also includes a set of commonly used middlewares:

- **CORS**: Cross-Origin Resource Sharing  
- **JWT**: JSON Web Token authentication  
- **Logging**: Request logging  
- **Recovery**: Panic recovery with stack trace  

In addition, the API provides a simple **health check endpoint** for uptime monitoring and a **login endpoint** for JWT authentication.

## Features

- Gorilla Mux router  
- Middleware support (CORS, JWT, Logging, Recovery)  
- Swagger (OpenAPI) integration  
- AWS Secrets Manager (simple implementation)  
- Health check and login endpoints  
- Other example endpoints. 


## Features

- Gorilla Mux router  
- Middleware support  
- Swagger (OpenAPI) integration  
- AWS Secrets Manager (simple implementation)  

## Getting Started

```bash
git clone https://github.com/Retromindgames/api-template.git
cd api-template
go run .
```

### Swagger

API has built in documentation and testing. 
Init swagger to access the api documentation:

```bash
swag init
```
Acess via: http://localhost:8080/swagger/index.html#/


## Project Structure


```bash
.
├── main.go              # Entry point
├── router/              # Route definitions
├── middleware/          # Custom middlewares
├── docs/                # Swagger docs (generated via swag)
├── handlers/            # Endpoint declaration.
├── secrets/             # AWS Secrets Manager logic
├── logger/              # For logging tools
├── auth/                # Authentication related code.
├── ent/                 # EntGo schema and generated code.
├── pkg/                 # For internal packages
├── go.mod
└── go.sum
```

## Entgo

ORM, use the bellow command to generate the code from the schema:

```bash
go run entgo.io/ent/cmd/ent generate ./ent/schema
```

