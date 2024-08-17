Go protocol buffers plugin

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

comando
protoc --version
protoc src/proto/user.proto --go_out=./

Instalar dependencia no projeto
go get -v google.golang.org/protobuf

