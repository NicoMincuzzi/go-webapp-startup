# Web application startup: Goland
[![CI](https://github.com/NicoMincuzzi/go-webapp-startup/actions/workflows/ci.yml/badge.svg)](https://github.com/NicoMincuzzi/go-webapp-startup/actions/workflows/ci.yml)
![Golang version](https://img.shields.io/badge/golang-1.16-9cf)
![GitHub repo size](https://img.shields.io/github/repo-size/NicoMincuzzi/go-webapp-startup)

## Prerequisites

- The first need [Go](https://golang.org/) installed
  

- Create directory for `GOPATH` setting, making a `/go` folder and relative subfolders `/src`, `/bin` and `/pkg`, as follow:
  ```
  .
  `-- go\
      |-- src\
      |-- bin\
      `-- pkg\
  ```
  

- For `GOPATH` environment variable, add the following lines to the `~/.bashrc`. The `GOPATH` envirionment variable is set to `$HOME/go` for default. However, you can customize it by:
  
    ```groovy
    export GOPATH=$HOME/<YOUR_PATH>/go
    export PATH=$GOPATH/bin:$PATH 
    ```
  
  By command `go env` you can list all set env variables about Go.
  
## Create a new project

- To create our new project, we need to create a directory for it in `go/src`
  

- Create layout for Go application project by following [guide](https://github.com/golang-standards/project-layout)


- Let's make the project directory the root of a module by using `go mod init`:

    ```sh
    $ go mod init nicomincuzzi/go-webapp
    ```

  It's common practice to use your github username to keep your project names globally unique, and avoid name conflicts with any of your project dependencies, but you can use any name you like.

- To add module requirements and sums:
  
  ```sh
  $ go mod tidy
  ```

## IDEs

Can open your project with your favorite IDE.

### Goland

One of the most famous IDE is GoLand IDE, in which you need to flag `Enable Go Modules Integration` in `Preferences... -> Go -> Go Modules `

  ![image](https://user-images.githubusercontent.com/48289901/112907559-5fdf0680-90ee-11eb-8129-5ea7ef657e4b.png)

### VSCode

After the installation, Launch VSCode.

Open the Extensions Marketplace (Cmd+Shift+X), search Go and install it.

![image](https://user-images.githubusercontent.com/48289901/114190603-c4e6f780-994b-11eb-81f6-296e8469a5cb.png)

open the Command Palette (Cmd+Shift+P) and run the Go: Install/Update Tools command.

![image](https://user-images.githubusercontent.com/48289901/114191135-5eaea480-994c-11eb-8613-e9d30620945c.png)

Install all the Go extensions listed there.

![image](https://user-images.githubusercontent.com/48289901/114191203-7423ce80-994c-11eb-9e97-75abfc5a94b5.png)


## Dependency

Use the below Go command to install package.
 
 ```sh 
 $ go get -u github.com/<PACKAGE_PATH>
 ```
    
For example, install gin package:
 
 ```sh
 $ go get -u github.com/gin-gonic/gin
 ```
  
It inserts package in go module (`go.mod` file) and packages in `go/pkg/mod`

If you look at the `go.mod` file, you'll see it now contains this:

  ```go.mod
  module nicomincuzzi/go-webapp
  
  go 1.16
  
  require (
   github.com/gin-gonic/gin v1.6.3 // indirect
  )
  ```

You will also see a `go.sum` file now. This is a text file containing references to the specific versions of all the package dependencies, and their dependencies, along with a cryptographic hash of the contents of that version of the relevant module.

The `go.sum` file serves a similar function to `package-lock.json` for a Javascript project and you should always check it into version control along with your source code.
### `/vendor`

Application dependencies (managed manually or by your favorite dependency management tool like the new built-in Go Modules feature). The `go mod vendor` command will create the `/vendor` directory for you. 
   
   ```sh
   $ go mod vendor
   ```

## Build, Testing and Run  

- Build project

  ```sh
  $ go build
  ```

  or specifying the target folder:

  ```sh
  $ go build -o build/hello hello.go
  ```

- Testing
  
  From your project’s root directory, run your first test:
  
  ```sh
  $ go test
  ```

  You will receive the following output:
  
  ```sh
  Output
  PASS
  ok      ./math 0.988s
  ```

  The `go test` subcommand only looks for files with the `_test.go` suffix. `go test` then scans those file(s) for special functions including `func TestXxx` and several others that we will cover in later steps. `go test` then generates a temporary main package that calls these functions in the proper way, builds and runs them, reports the results, and finally cleans everything up.

  `go test` is probably sufficient for our little program, but there will be times when you’ll wish to see what tests are running and how long each takes. Adding the `-v` flag increases verbosity. Rerun your test with the new flag:
  
  ```sh
  $ go test -v
  ```

  You will see the following output:
  
  ```sh
  Output
  === RUN   TestAdd
  --- PASS: TestAdd (0.00s)
  PASS
  ok      ./math 1.410s
  ```
  
  Go provides cover tool (out of many built-in tools) to analyze coverage information of a test. We use this tool to accept our coverage profile and outputs an HTML file which contains the human-readable information about the test in a very interactive format.
  
  ```sh
  $ go test -cover
  ```
  
  ```sh
  $ go test -cover -html=cover.txt -o cover.html
  ```
  
- Run project

  ```sh
  $ go run .
  ```

## Creat a web application

We're going to use [gin](https://github.com/gin-gonic/gin) for our web application, which is a lightweight web framework.

Let’s take a quick look at how a request is processed in Gin. The control flow for a typical web application, API server or a microservice looks as follows:

```
Request -> Route Parser -> [Optional Middleware] -> Route Handler -> [Optional Middleware] -> Response
```

When a request comes in, Gin first parses the route. If a matching route definition is found, Gin invokes the route handler and zero or more middleware in an order defined by the route definition.

Create a file called `main.go` containing this code:

```go
package main

import "github.com/gin-gonic/gin"

func main() {
r := gin.Default()

    r.GET("/status", func(c *gin.Context) {
        c.String(200, "Healthy!")
    })

    r.Run(":3030")
}
```
Let's break this down a little:
```go
r := gin.Default()
```
This creates a router object, r, using the built-in defaults that come with gin.

Then, we assign a handler function to be called for any HTTP GET requests to the path /hello, and to return the string "Hello, World!" and a 200 (HTTP OK) status code:
```go
r.GET("/status", func(c *gin.Context) {
c.String(200, "Healthy!")
})
```
Finally, we start our webserver and tell it to listen on port 3000:
```go
r.Run(":3030")
```
To run this code, execute:
```sh
$ go run main.go
```
You should see output like this:
```sh
go: finding module for package github.com/gin-gonic/gin
go: found github.com/gin-gonic/gin in github.com/gin-gonic/gin v1.6.3
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
- using env:   export GIN_MODE=release
- using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /hello                    --> main.main.func1 (3 handlers)
[GIN-debug] Listening and serving HTTP on :3030
```
Now if you visit `http://localhost:3030/status` in your web browser, you should see the message `"Healthy!"`

## How to Contribute

Everyone is very welcome to contribute to this project. You can contribute just by submitting bugs or suggesting improvements by
[opening an issue on GitHub](https://github.com/NicoMincuzzi/go-webapp-startup/issues).

## License
Distributed under Apache-2.0 License, please see license file within the code for more details.
