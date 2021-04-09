# Web application startup: Goland
![GitHub repo size](https://img.shields.io/github/repo-size/NicoMincuzzi/go-webapp-startup)

## Prerequisites

- The first need [Go](https://golang.org/) installed
  

- Create directory for `GOPATH` setting, making a `/go` folder and relative subfolders `/src`, `/bin` and `/pkg`, as follow:
  ```
  /go
   |_ /src
   |_ /bin
   |_ /pkg
  ```
  

- For `GOPATH` envirionment variable, add the following lines to the `~/.bashrc`. The `GOPATH` envirionment variable is set to `$HOME/go` for default. However, you can customize it by:
  
    ```sh
    export GOPATH=$HOME/<YOUR_PATH>/go
    export PATH=$GOPATH/bin:$PATH 
    ```
  
  By command `go env` you can list all set env variables about Go.
  
## Create a new project

- Create your new project in `go/src` folder
  

- Create layout for Go application project by following [guide](https://github.com/golang-standards/project-layout)


- Let's make the project directory the root of a module by using `go mod init`:
  
    `go mod init go-webapp`

## IDEs

Can open your project with your favorite IDE.

### Goland

- One of the most famous IDE is GoLand IDE, in which you need to flag `Enable Go Modules Integration` in `Preferences... -> Go -> Go Modules `

  ![Schermata 2021-03-30 alle 00 24 04](https://user-images.githubusercontent.com/48289901/112907559-5fdf0680-90ee-11eb-8129-5ea7ef657e4b.png)

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
  
`go get -u github.com/<DEPENDENCY>`
    
 For example, 
    
 `go get -u github.com/gin-gonic/gin`
  
 It inserts `require ( ... )` in go module (`go.mod` file) and packages in `go/pkg/mod`

### `/vendor`

Application dependencies (managed manually or by your favorite dependency management tool like the new built-in Go Modules feature). The `go mod vendor` command will create the `/vendor` directory for you. 
   
    `go mod vendor`

## Build and Run  

- Build project

  `go build`

- Run project

  `go run .`
