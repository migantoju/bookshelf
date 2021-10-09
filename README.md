# bookshelf
Learning Cloud Native with Go by creating a *Bookshelf* Restful API.

## Caching third-party libraries
We can cache go modules by copying `go.mod` and `go.sum` files to Docker
first and runnig `go mod download` before copying all the other files to
Docker.