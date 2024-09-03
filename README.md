# Backend
Built with:
- golang
- graphql
- gql-gen
- postgresSQL

## Docs
- Please Check The Wiki for contribution information! 

### How to run this code
```bash
# clone this repository
git clone https://github.com/Moral-Authority/backend.git
cd backend

# run go mod tidy
go mod tidy

# setting your environment in .env file

# serve your server
go run .
```

### How to run it via docker-compose

```bash
# run this command in the project root directory
docker-compose up
```

### Environment Variables
      - ENVIRONMENT: running environment(default=dev)
      - SERVER_PORT: port of the server(default=8080)
      - SERVER_NAME: server name
      - DATABASE_NAME: db name
      - DATABASE_USERNAME: db username
      - DATABASE_PASSWORD: db password
      - DATABASE_CONNECTION_URL: db host url
      - DATABASE_CONNECTION_PORT: db port

Note : you can check to my documentation CRUD but you better use [Insomnia](https://insomnia.rest/), it was in this root directory with json format


### Seed DB

`heroku run go run cmd/seed/seed.go --app moral-authority-backend `

OR 

```bash
heroku run bash --app moral-authority-backend
go run cmd/seed/seed.go
```

### Dump DB