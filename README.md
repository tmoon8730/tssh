# tssh
A simple Go CLI utility for storing and recalling SSH commands

## Usage
**Add a new template**: `tssh template add <name> <ssh_arguments>`

_Note: the ssh arguments should include everything except the ssh for the command you want to save._
_Ex. If_ `ssh 192.168.0.8` _was the expected command then enter_ `tssh template add test 192.168.0.8`

**Remove saved template**: `tssh template remove <name>`

**List saved templates**: `tssh template list`

**Execute saved ssh command**: `tssh <name>`

## Development

### Prerequisites
* [Go CLI](https://golang.org/doc/install) v1.13 or higher
* Any Linux Distro

### Download Dependencies
From the project root run `go get` to install the needed dependency packages

### Run locally
Run `go run main.go <arugments>` from the project root to execute the program locally. In place of `<arguments>` use any of the listed tssh commands from above.

### Build
Run `go build -o bin/tssh main.go` to compile into an executable binary. Optionally run `./build.sh` to build, test, and copy the binary to `/usr/local/bin/` for executing globally.

### Run Unit Tests
Change to the **sshtemplate** directory and run `go test -v` to execute the unit tests for the package.