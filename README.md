# Go WebApp

- The first need [Go](https://golang.org/) installed
  

- create
  ```
  go
   |_ src
   |_ bin
   |_ pkg
  ```
  

- `$GOPATH` is set to `$HOME/go` for default. You can customize `$GOPATH`:
  
    `export GOPATH=$HOME/<YOUR_PATH>/go`
  
    Verify that it's ok
    
    `go env`
  

- create new project in `go/src`
  

- Open new project in GoLand
  Enable Go Modules Integration in Preferences... -> Go -> Go Modules 
  
  
- create layout for Go application project by following [guide](https://github.com/golang-standards/project-layout)


- Let's make the project directory the root of a module by using go mod init:
  
    `go mod init go-webapp`
  

- Use the below Go command to install package.
  
    `go get -u github.com/gin-gonic/gin`
  
    It inserts `require ( ... )` in go module and packages in `go/pkg/mod`


-  Create vendor folder which contains external dependency, by following command:
   
    `go mod vendor`
   
- Build project

  `go build`

- Run project

  `go run .`