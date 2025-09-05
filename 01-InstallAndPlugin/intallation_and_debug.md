# Tools

## Install Go Language 1.25.0

### Windows:

[Download here](https://go.dev/dl/go1.25.0.windows-amd64.msi) and then install

### Linux Method 1:

Run these commands in terminal:

```bash
sudo apt update
sudo apt search golang-go
sudo apt search gccgo-go
sudo apt install golang-go
```

**Note:** This method installs the version available in your distribution's repository, which may not be Go 1.25.0. For the latest version, use Method 2.

### Linux Method 2:

1. [Download file](https://go.dev/dl/go1.25.0.linux-amd64.tar.gz)
2. Extract file `tar -xf "go1.25.0.linux-amd64.tar.gz"`
3. Setup permissions `sudo chown -R root:root ./go`
4. Move go binaries to local folder `sudo mv -v go /usr/local`
5. For setup go path, open profile: `code ~/.bash_profile` or `code ~/.profile`
6. Append following lines to file:
   ```bash
   export GOPATH=$HOME/go
   export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
   ```
7. Source this file: `source ~/.bash_profile` or `source ~/.profile`
8. Run `go version` to check if go is installed

### macOS:

#### Method 1 (Installer):
[Download here](https://go.dev/dl/go1.25.0.darwin-amd64.pkg) for Intel Macs or [here](https://go.dev/dl/go1.25.0.darwin-arm64.pkg) for Apple Silicon Macs, then install

#### Method 2 (Manual):
1. [Download file](https://go.dev/dl/go1.25.0.darwin-amd64.tar.gz) for Intel Macs or [here](https://go.dev/dl/go1.25.0.darwin-arm64.tar.gz) for Apple Silicon Macs
2. Extract file `tar -xf "go1.25.0.darwin-amd64.tar.gz"` (or arm64 version)
3. Move go binaries `sudo mv go /usr/local`
4. Add to your shell profile (`~/.bash_profile`, `~/.zshrc`, etc.):
   ```bash
   export GOPATH=$HOME/go
   export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
   ```
5. Source your profile or restart terminal
6. Run `go version` to verify installation

## Install Go Debugger

Run this command: `go install github.com/go-delve/delve/cmd/dlv@latest`