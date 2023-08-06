# nesesary tools for proto generation
go-tools-install:
	go install github.com/bufbuild/buf/cmd/buf@latest && \
	go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest && \
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
	go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest && \
	go install github.com/volatiletech/sqlboiler/v4@latest && \
	go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest

set-path:
	echo "export PATH=$PATH:$(go env GOPATH)/bin" >> ~/.zshrc

models-gen:
	sqlboiler mysql -o ./pkg/models --no-tests --wipe