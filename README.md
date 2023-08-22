# Golang template

This is a project for [Go](https://golang.org/) API use :

- [Fiber](https://gofiber.io/) as HTTP handle
- [ent](https://entgo.io/) as ORM
- [wire](https://github.com/google/wire) as dependency injection

## How to use

### 1. Clone the repository

```sh
$ git clone https://github.com/FournyP/go-api-template
$ cd go-api-template
$ git remote set-url origin <YOUR REPOSITORY URL>
```

### 2. Rename module name

- Open `go.mod` file
- Replace `go-api-template` by your module name
- Correct import in files (You can use VS Code to do it)

### 3. Create a `.env` file based on `.env.example` and fill it

```sh
$ cp .env.example .env
```

### 4. Uncomment database driver in `database.go` file and install dependencies

- Open database.go file
- Uncomment the driver you want to use
- Run `go mod tidy` to download the driver and automatically install dependencies

### 5. Run

```sh
$ go run .
```

### 6. Build

```sh
$ go build .
```
