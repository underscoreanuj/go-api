# go-api

- This is a golang based REST API with a diverse set storage capabilities, following best design practices and a loosley coupled model.


## Setup

- simply fork the project and run `go get github.com/<your_github_handle>/go-api`

## Running Locally

- currently by default the project uses a local sqlite3 file as a database.

- we can use a firestore instance instead of an sqlite3 file, but that will require an os environment variable `GOOGLE_APPLICATION_CREDENTIALS` pointing to your firebase project's secret json file path for access to firestore.

- run `go build .` in the root diretory

- run `export GO_API_PORT=8000`, you can select any other open port if you wish, this port will be used by the API server

- run `go run .`

### Caching

- the project currently relies on a redis cache instance running in the background, you can specify the host and port in the server.go main file.

### Testing

- you can run `go test ./...` to recursively execute all test cases in all directories, or explicitly specify their path to the go test command.

### switching routers

- currently the project supports chi-router and mux-router, you can add your own implementation or an existing library implmenting the router interface.

## Endpoints

- /posts : [GET] : returns an array of post objects `[{"id": _, "title": _, "text": _}]`
- /posts : [POST]: expects a json payload of a post object `{"title": _, "text": _}`, returns the created post object with a unique id `{"id": _, "title": _, "text": _}`
- /posts/{id}: [GET]: returns the post object associated with the given id, if the id is not found then an error message is returned `{"id": _, "title": _, "text": _}`
