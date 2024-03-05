# Coding challenge
This project aims to code a solution for [problem](./PROBLEM.md)

## Areas of improvement (for time saving)
- I am not a FE developer (very little typescript/react experience) so frontend is basic/functional.
- Test more error handling and edge cases, increase unit test coverage.
- Write GODOG integration tests.
- Write a swagger doc for backend, also open-api based tests.

## Important
- Dont have a personal cloud subscription to deploy this to a public area (please run via docker on local).
- Used in-memory database for simplicity and for benifit of saving time.
- Used moq to generate mocks for interfaces.
- Table based unit tests, added some to showcase knowledge (test coverage is low)
- Used a layered approach for backend (api/handler/database)
- Dependency injection, easy to mock and test.

## ALGORITHM
The heart of the algo is in this [file](./backend/handlers/packHandler.go)
`calculateBestPackCombination` (have been commented AND unit tests added for all mentioned use-cases).
Went for a recursive approach that checks for minimum wastage/excess and then number of packs.

Simple/functional [frontend](./frontend/src/App.js)

## Makefile tagets
- Unit tests (```make test```)
- Linting (```make lint```)


# Execution steps
- bring up docker environment: this will spin up both fe/be in containers
```
docker-compose up
``` 
- open up localhost:8888 (for dozzle, viewing logs from all containers)
- open up localhost:8880 (application front-end)
- backend runs on port 8881
- application comes up with initial pack sizes of [250, 500, 1000, 2000, 5000]
- can run queries both from frontend OR using curl
```
curl --location 'http://localhost:8881/packs?items=251'
```
- For non container testing
```
in frontend folder
> npm start
```
```
in backend folder
> go run cmd/main.go
```