# Description
build using tutorial from http://thenewstack.io/make-a-restful-json-api-go/

# operations
## list all
`curl -v http://localhost:8080/todos`

## list one
`curl -v http://localhost:8080/todos/1`

## add todo
`curl -v -H "Content-Type: application/json" -d "{\"name\":\"bla\"}" http://localhost:8080/todos`
