## Go (Golang) Commands

### 1. Run Go Project
```bash
make run
```
- Runs the Go application using the `run` target defined in the `Makefile`.
- Executes the commands configured in the Makefile.

### 2. Manage Go Dependencies
```bash
go mod tidy
```
- Downloads missing dependencies required by the project.
- Adds missing direct dependencies to `go.mod`.
- Removes unused dependencies.
- Updates `go.sum` with the required dependency checksums.
- Keeps direct and indirect dependencies in sync.

**Note:** `go mod tidy` manages project dependencies; it does not install packages globally.




<!-- Go Run command
make run

iss command se import package install ho jata hai saath, me direct/indirect install package ko vi fix karta hai
go mod tidy -->