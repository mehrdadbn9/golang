**Tools**:

### Install Go Language 1.18

#### Windows:

[Download here](https://go.dev/dl/go1.18.1.windows-amd64.msi) and then install

#### Linux Method 1:

Run this commands in terminal:

    sudo apt update
    sudo apt search golang-go
    sudo apt search gccgo-go
    sudo apt install golang-go

#### Linux Method 2:

1. [Download file](https://go.dev/dl/go1.18.1.linux-amd64.tar.gz)
2. Extract file `tar -xf "go1.18.1.linux-amd64.tar.gz"`
3. Setup permissions `sudo chown -R root:root ./go`
4. Move go binaries to local folder `sudo mv -v go /usr/local`
5. for Setup go path, open profile: `code ~/.bash_profile` or `code ~/.profile`
6. Append following to lines to file:
   ```
   export GOPATH=$HOME/go
   export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
   ```
7. source this file: `source ~/.bash_profile` or `source ~/.profile`
8. Run `go version` to check if go is installed

### Install Go Debugger

Run this command: `go install github.com/go-delve/delve/cmd/dlv@latest`
