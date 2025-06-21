# gRPC up and running

<hr>

###### 🛠️ *References*

- GitHub: [https://github.com/grpc-up-and-running/samples](https://github.com/grpc-up-and-running/samples)

<hr>

# Installation necessary tools:
- Download and install the latest protocol buffer version 3 compiler from the [GitHub release page](https://github.com/protocolbuffers/protobuf/releases). Then run this command:
  ```bash
  unzip protoc-*-linux-x86_64.zip -d $HOME/.local
  echo 'export PATH="$PATH:$HOME/.local/bin"' >> $HOME/.zshrc
  source $HOME/.zshrc
  ```
  
- Install the gRPC library using the following command:
  ```bash
  go get -u google.golang.org/grpc
  ```
  
- Install the protoc plug-in for Go using the following command:
  ```bash
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  echo 'export PATH="$PATH:$(go env GOPATH)/bin"' >> $HOME/.zshrc
  source $HOME/.zshrc

  # Verfication
  protoc-gen-go --version
  ```

# Completed chaps
- [x] chap01
- [x] [chap02](./chap02/README.md)