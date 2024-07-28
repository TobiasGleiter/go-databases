# go-databases

Interface and Databases in Go(lang)

## SQL Implementation

- Add your development database credentials and make sure you have a mysql database running.
- Run `make dev service=sql ` to start the service using a mysql database.
- Send curl POST request: `curl -X POST http://localhost:4000/users/create`

## MongoDB Implementation

- Add your development database credentials and make sure you have a mongodb database running.
- Run `make dev service=mongodb ` to start the service using a mongodb database.
- Send curl POST request: `curl -X POST http://localhost:4000/users/create`
