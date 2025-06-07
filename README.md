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