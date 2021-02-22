# My Progress - Section 3

- Learn about Software Dependency from this great [article](https://research.swtch.com/deps) from [Russ Cox](https://swtch.com/~rsc/).
- Run into issues with old version of Go.
- Uninstall and Reinstall newest version of Go.
- Delete all Go related files in /usr.
- Run into issue with $GOROOT env as it turns out that with newer version of Go, you don't need to set GOROOT env. Solve the issue by unset GOROOT and remove the env from PATH in my bashrc file.
- Install Go on VSCode and learn some Go commands on VSCode.
- Write a simple Hello World programm
- Learn go commands (go run, build, install)
- Set GOBIN
- Go Modules overview
- Create a Go Module (go mod init, go test)
- Adding a dependency to go mod (go test, go list -m all)
- Upgrading dependencies
- go mod tidy command cleans up these unused dependencies

### [Go modules](https://blog.golang.org/using-go-modules)

- go mod init creates a new module, initializing the go.mod file that describes it.
- go build, go test, and other package-building commands add new dependencies to go.mod as needed.
- go list -m all prints the current module’s dependencies.
- go get changes the required version of a dependency (or adds a new dependency).
- go mod tidy removes unused dependencies.
