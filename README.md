# Go API Template

This is a minimal and extensible Go Web API built using Gorilla Mux. It includes a basic router, Swagger documentation, and AWS Secrets Manager integration for securely loading configuration values. The project also includes a set of commonly used middlewares:

- **CORS**: Cross-Origin Resource Sharing  
- **JWT**: JSON Web Token authentication  
- **Logging**: Request logging  
- **Recovery**: Panic recovery with stack trace  

In addition, the API provides a simple **health check endpoint** for uptime monitoring and a **login endpoint** for JWT authentication.

## Features

- Gorilla Mux router.
- Middleware support (CORS, JWT, Logging, Recovery).
- Swagger (OpenAPI) integration.
- AWS Secrets Manager (simple implementation). 
- Health check and login endpoints.
- Other example endpoints. 
- Cloudflare R2 bucket upload.
- Entgo for ORM, using postgress (docker-compose).

## Getting Started

```bash
git clone https://github.com/Retromindgames/api-template.git
cd api-template
go run .
```

## Env variables

The project requires a few env variable defines in a .env file that is used by docker-compose.
If they are not defined, there will be console errors.
Just create a .env file and add the proper credentials / configurations.

Here is a list:

```bash
# Database config
DATABASE_URL=

# R2 storage credentials
R2_ACCOUNT_ID=
R2_ACCESS_KEY_ID=
R2_ACCESS_KEY_SECRET=
R2_BUCKET_NAME=
```

### Swagger

API has built in documentation and testing. 
Init swagger to access the api documentation:

```bash
swag init
```
Acess via: http://localhost:8080/swagger/index.html#/

The command needs to be run if there are any changes to the swagger annotation.

NOTE: Currently there is an issue with swagger and the UploadPDF endpoint.

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

The human writen schema files are in ./ent/schema/, the other files in ./ent/ are generated files from the command above.

The command needs to be used when there is any schema change.
